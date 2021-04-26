package tcp

import (
	"errors"
	"fmt"
	"net"

	"github.com/75912001/xr/lib/log"
)

//己方作为客户端
type Client struct {
	OnPacket  OnPacketClient
	onDisconn OnDisconnClient
	server    remote
	eventChan chan interface{}
}

//连接
//每个连接有一个 发送协程, 一个 接收协程
//事件放入 参数 eventChan 中, 以供外部处理
//address:127.0.0.1:8787
//recvPacketMaxLen:接受数据包的最大长度(包头+包体)
//eventChan:外部传递的事件处理管道.连接的事件会放入该管道,以供外部处理
//sendChanCapacity:发送管道容量
func (p *Client) Connect(address string, recvPacketMaxLen uint32, eventChan chan interface{},
	onParseProtoHead OnParseProtoHead, onDisconn OnDisconnClient, onPacket OnPacketClient,
	sendChanCapacity uint32) (err error) {
	log.GLog.Trace(fmt.Sprintf("address:%v, recvPacketMaxLen:%v, eventChan:%v, onParseProtoHead:%v, onDisconn:%v, onPacket:%v, sendChanCapacity:%v",
		address, recvPacketMaxLen, eventChan, onParseProtoHead, onDisconn, onPacket, sendChanCapacity))
	p.onDisconn = onDisconn
	p.OnPacket = onPacket
	p.eventChan = eventChan
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if nil != err {
		log.GLog.Crit(fmt.Sprintf("net.ResolveTCPAddr, err:%v, address:%v", err, address))
		return
	}

	p.server.connLock.Lock()
	defer p.server.connLock.Unlock()
	p.server.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if nil != err {
		log.GLog.Crit(fmt.Sprintf("net.Dial, err:%v address:%v", err, address))
		return
	}

	p.server.sendChan = make(chan interface{}, sendChanCapacity)
	go p.server.onSendEvent()

	go p.onRecvEvent(recvPacketMaxLen, onParseProtoHead)
	return
}

//主动断开连接
func (p *Client) DisConn() (err error) {
	if !p.server.isConn() {
		log.GLog.Warn("link disconnect.")
		return errors.New("[ERROR]link disconnect.")
	}
	p.eventChan <- &DisconnEventClient{
		Client: p,
	}
	return
}

//处理断开链接
func (p *Client) OnDisConn() {
	p.server.connLock.Lock()
	defer p.server.connLock.Unlock()

	if p.server.sendChan != nil {
		close(p.server.sendChan)
		p.server.sendChan = nil
	}

	if p.server.isConn() {
		p.onDisconn(p)
		p.server.conn.Close()
		p.server.conn = nil
	}
}

//发送数据(必须在处理EventChan事件中调用)
func (p *Client) Send(data []byte) (err error) {
	if !p.server.isConn() {
		log.GLog.Warn("link disconnect.")
		return errors.New("[ERROR]link disconnect.")
	}
	p.server.sendChan <- &sendEvent{
		data: data,
		dst:  &p.server,
	}
	return
}

//接收数据
func (p *Client) onRecvEvent(recvPacketMaxLen uint32, onParseProtoHead OnParseProtoHead) {
	log.GLog.Trace("goroutine start.")
	defer func() { //断开链接
		log.GLog.Trace("goroutine done.")
		p.eventChan <- &DisconnEventClient{
			Client: p,
		}
	}()

	buf := make([]byte, recvPacketMaxLen)
	var readIndex int
	for {
	LoopRead:
		readNum, err := p.server.conn.Read(buf[readIndex:])
		if nil != err {
			log.GLog.Error(fmt.Sprintf("Conn.Read, read num:%v, err:%v", readNum, err))
			return
		}
		readIndex += readNum
		for {
			packetLength := onParseProtoHead(buf, readIndex)
			if 0 == packetLength {
				goto LoopRead
			}

			if -1 == packetLength {
				log.GLog.Crit(fmt.Sprintf("packetLength:%v, readIndex:%v, Data:%v", packetLength, readIndex, buf))
				return
			}

			{ //接受数据
				var c PacketEventClient
				c.Data = make([]byte, packetLength)
				c.Client = p
				copy(c.Data, buf[:packetLength])
				p.eventChan <- &c
			}
			copy(buf, buf[packetLength:readIndex])
			readIndex -= packetLength

			if 0 == readIndex {
				goto LoopRead
			}
		}
	}
}
