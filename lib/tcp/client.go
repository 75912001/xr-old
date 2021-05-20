package tcp

import (
	"errors"
	"net"
)

//己方作为客户端
type Client struct {
	OnEventPacket  OnEventPacketClientFunc
	OnEventDisConn OnEventDisConnClientFunc
	Remote         Remote
	tcpChan        chan<- interface{}
}

//连接
//每个连接有一个 发送协程, 一个 接收协程
//事件放入 参数 eventChan 中, 以供外部处理
//address:127.0.0.1:8787
//recvPacketMaxLen:接受数据包的最大长度(包头+包体)
//tcpChan:外部传递的事件处理管道.连接的事件会放入该管道,以供外部处理
//sendChanCapacity:发送管道容量
func (p *Client) Connect(address string, recvPacketMaxLen uint32, tcpChan chan<- interface{},
	OnEventDisConn OnEventDisConnClientFunc,
	OnEventPacket OnEventPacketClientFunc,
	onParseProtoHead OnParseProtoHeadFunc,
	sendChanCapacity uint32) (err error) {
	p.OnEventPacket = OnEventPacket
	p.OnEventDisConn = OnEventDisConn
	p.tcpChan = tcpChan
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	if nil != err {
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if nil != err {
		return
	}

	p.Remote.addEventDisConn2EventChan = func() {
		p.tcpChan <- &EventDisConnClient{
			Client: p,
		}
	}

	p.Remote.addEventPacket2EventChan = func(data []byte, packetLength int) {
		pes := &EventPacketClient{
			Client: p,
			Data:   make([]byte, packetLength),
		}
		copy(pes.Data, data[:packetLength])
		p.tcpChan <- pes
	}

	p.Remote.start(conn, sendChanCapacity, recvPacketMaxLen, onParseProtoHead)
	return
}

func (p *Client) IsConn() bool {
	return p.Remote.IsConn()
}

//主动断开连接
func (p *Client) DisConn() (err error) {
	if !p.IsConn() {
		return errors.New("[ERROR]link disconnect.")
	}
	p.tcpChan <- &EventDisConnClient{
		Client: p,
	}
	return
}

//事件中处理时调用
func (p *Client) EventDisConn() {
	p.OnEventDisConn(p)
	if p.IsConn() {
		p.Remote.stop()
	}
}
