package tcp

import (
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/75912001/xr/lib/log"
)

//断开链接事件channel 基于自身是client
type CloseConnectEventChanClient struct {
	Client *Client
}

//收到数据事件channel 基于自身是client
type RecvEventChanClient struct {
	Buf    []byte
	Client *Client
}

type OnParseProtoHead func(buf []byte, length int) int //解析协议包头 返回长度:完整包总长度  返回0:不是完整包 返回-1:包错误
type OnCloseConnect func(client *Client) int           //远端链接关闭
type OnPacket func(client *Client, buf []byte) int     //远端包

//Client 己方作为客户端
type Client struct {
	OnPacket       OnPacket
	onCloseConnect OnCloseConnect
	server         remote //远端
}

//Connect 连接
//recvBufMax:接受数据的最大长度
//eventChan:外部传递的事件处理
func (p *Client) Connect(ip string, port uint16, recvBufMax uint32, eventChan chan interface{}, onParseProtoHead OnParseProtoHead, onCloseConnect OnCloseConnect, onPacket OnPacket) (err error) {
	p.onCloseConnect = onCloseConnect
	p.OnPacket = onPacket
	var addr = ip + ":" + strconv.Itoa(int(port))
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if nil != err {
		log.GLog.Crit(fmt.Sprintf("net.ResolveTCPAddr, err:%v, addr:%v", err, addr))
		return
	}
	p.server.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if nil != err {
		log.GLog.Crit(fmt.Sprintf("net.Dial, err:%v, addr:%v", err, addr))
		return
	}

	p.server.sendEventChanCnt = 1024
	p.server.sendChan = make(chan interface{}, p.server.sendEventChanCnt)
	go onEventSend(p.server.sendChan)

	go p.recv(p.server.conn, eventChan, recvBufMax, onParseProtoHead)
	return
}

func (p *Client) DisConnect() {
	p.server.lock.Lock()
	defer p.server.lock.Unlock()

	if p.server.sendChan != nil {
		close(p.server.sendChan)
		p.server.sendChan = nil
	}

	if p.server.isConnect() {
		p.onCloseConnect(p)
		p.server.conn.Close()
		p.server.conn = nil
	}
}

//发送数据(必须在处理EventChan事件中调用)
func (p *Client) Send(buf []byte) (err error) {
	if !p.server.isConnect() {
		log.GLog.Warn("link disconnect.")
		return errors.New("[ERROR]link disconnect.")
	}
	var c sendEventChan
	c.buf = buf
	c.dst = &p.server
	p.server.sendChan <- &c
	return
}

func (p *Client) recv(conn *net.TCPConn, eventChan chan interface{}, recvBufMax uint32, onParseProtoHead OnParseProtoHead) {
	log.GLog.Trace("recv goroutine start.")
	defer func() { //断开链接
		log.GLog.Trace("recv goroutine done.")
		var c CloseConnectEventChanClient
		c.Client = p
		eventChan <- &c
	}()

	buf := make([]byte, recvBufMax)
	var readIndex int
	for {
	LoopRead:
		readNum, err := conn.Read(buf[readIndex:])
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
				log.GLog.Crit(fmt.Sprintf("packetLength:%v, readIndex:%v, Buf:%v", packetLength, readIndex, buf))
				return
			}

			{ //接受数据
				var c RecvEventChanClient
				c.Buf = make([]byte, packetLength)
				c.Client = p
				copy(c.Buf, buf[:packetLength])
				eventChan <- &c
			}
			copy(buf, buf[packetLength:readIndex])
			readIndex -= packetLength

			if 0 == readIndex {
				goto LoopRead
			}
		}
	}
}
