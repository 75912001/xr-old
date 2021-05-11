package tcp

import (
	"errors"
	"fmt"
	"net"

	"github.com/75912001/xr/lib/log"
)

//己方作为客户端
type Client struct {
	OnPacket  OnPacketClientFunc
	OnDisConn OnDisConnClientFunc
	Remote    Remote
	tcpChan   chan<- interface{}
	log       *log.Log
}

//连接
//每个连接有一个 发送协程, 一个 接收协程
//事件放入 参数 eventChan 中, 以供外部处理
//address:127.0.0.1:8787
//rwBuffLen:tcp recv/send 缓冲大小
//recvPacketMaxLen:接受数据包的最大长度(包头+包体)
//eventChan:外部传递的事件处理管道.连接的事件会放入该管道,以供外部处理
//sendChanCapacity:发送管道容量
func (p *Client) Connect(address string, log *log.Log, rwBuffLen int, recvPacketMaxLen uint32, eventChan chan<- interface{},
	onDisConn OnDisConnClientFunc, onPacket OnPacketClientFunc, onParseProtoHead OnParseProtoHeadFunc,
	sendChanCapacity uint32) (err error) {
	log.Trace(fmt.Sprintf("address:%v, recvPacketMaxLen:%v, eventChan:%v,  onDisConn:%v, onPacket:%v, onParseProtoHead:%v, sendChanCapacity:%v",
		address, recvPacketMaxLen, eventChan, onDisConn, onPacket, onParseProtoHead, sendChanCapacity))
	p.OnDisConn = onDisConn
	p.OnPacket = onPacket
	p.tcpChan = eventChan
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if nil != err {
		log.Crit(fmt.Sprintf("net.ResolveTCPAddr, err:%v, address:%v", err, address))
		return
	}

	p.Remote.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if nil != err {
		log.Crit(fmt.Sprintf("net.Dial, err:%v address:%v", err, address))
		return
	}

	p.Remote.conn.SetNoDelay(true)
	p.Remote.conn.SetReadBuffer(rwBuffLen)
	p.Remote.conn.SetWriteBuffer(rwBuffLen)

	p.Remote.sendChan = make(chan interface{}, sendChanCapacity)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Crit(fmt.Sprintf("onSendEvent goroutine panic:%v\n", err))
			}
		}()
		p.Remote.onSendEvent(log)
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Crit(fmt.Sprintf("onRecvEventChan goroutine panic:%v\n", err))
			}
		}()
		p.onRecvEvent(recvPacketMaxLen, onParseProtoHead)
	}()
	return
}

//主动断开连接
func (p *Client) DisConn() (err error) {
	if !p.Remote.IsConn() {
		p.log.Warn("link disconnect.")
		return errors.New("[ERROR]link disconnect.")
	}
	p.tcpChan <- &DisConnEventClient{
		Client: p,
	}
	return
}

//接收数据
func (p *Client) onRecvEvent(recvPacketMaxLen uint32, onParseProtoHead OnParseProtoHeadFunc) {
	p.log.Trace("goroutine start.")

	defer func() { //断开链接
		if err := recover(); err != nil {
			p.log.Warn(fmt.Sprintf("goroutine panic:%v", err))
		} else { //断开链接
			p.tcpChan <- &DisConnEventClient{
				Client: p,
			}
		}
		p.log.Trace("goroutine done.")
	}()

	//TODO [improvement] 环形缓冲
	buf := make([]byte, recvPacketMaxLen)

	var readIndex int
	for {
	LoopRead:
		readNum, err := p.Remote.conn.Read(buf[readIndex:])
		if nil != err {
			p.log.Error(fmt.Sprintf("Conn.Read, read num:%v, err:%v", readNum, err))
			return
		}
		readIndex += readNum
		for {
			packetLength := onParseProtoHead(buf, readIndex)
			if 0 == packetLength {
				goto LoopRead
			}

			if -1 == packetLength {
				p.log.Crit(fmt.Sprintf("packetLength:%v, readIndex:%v, Data:%v", packetLength, readIndex, buf))
				return
			}

			//接受数据
			pes := &PacketEventClient{
				Client: p,
				//Data:    make([]byte, packetLength),
				Data:   buf[:packetLength],//TODO 测试1 ...
			}
			//TODO 测试1 ... copy(pes.Data, buf[:packetLength])
			p.tcpChan <- pes

			copy(buf, buf[packetLength:readIndex])
			readIndex -= packetLength

			if 0 == readIndex {
				goto LoopRead
			}
		}
	}
}
