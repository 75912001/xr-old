package tcp

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

type Server struct {
	OnConn    OnEventConnServerFunc
	OnPacket  OnEventPacketServerFunc
	OnDisconn OnEventDisConnServerFunc
	tcpChan   chan<- interface{}
	listener  *net.TCPListener
	//+1 atomic.AddUint32(&recvChanCnt, 1)
	//-1 atomic.AddUint32(&recvChanCnt, ^uint32(0))
	//recvChanCnt uint32
	//sendChanCnt uint32
}

//运行服务
//address:127.0.0.1:8787
//rwBuffLen:tcp recv/send 缓冲大小
//recvPacketMaxLen:最大包长(包头+包体)
//eventChan:外部传递的事件处理
func (p *Server) Strat(address string, recvPacketMaxLen uint32, eventChan chan<- interface{},
	onConn OnEventConnServerFunc,
	onDisconn OnEventDisConnServerFunc,
	onPacket OnEventPacketServerFunc,
	onParseProtoHead OnParseProtoHeadFunc,
	sendChanCapacity uint32) (err error) {
	p.OnConn = onConn
	p.OnPacket = onPacket
	p.OnDisconn = onDisconn
	p.tcpChan = eventChan
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if nil != err {
		return
	}

	//TODO improvement [设置地址复用]
	//TODO improvement [设置监听的缓冲数量]

	p.listener, err = net.ListenTCP("tcp", tcpAddr)
	if nil != err {
		return
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Server Strat accept goroutine panic:%v", err)
			}
		}()
		var tempDelay time.Duration
		for {
			conn, err := p.listener.AcceptTCP()
			if nil != err {
				if ne, ok := err.(net.Error); ok && ne.Temporary() {
					if tempDelay == 0 {
						tempDelay = 5 * time.Millisecond
					} else {
						tempDelay *= 2
					}
					if max := 1 * time.Second; tempDelay > max {
						tempDelay = max
					}
					log.Printf("listen.AcceptTCP, tempDelay:%v, err:%v", tempDelay, err)
					time.Sleep(tempDelay)
					continue
				}
				log.Printf(fmt.Sprintf("listen.AcceptTCP, err:%v", err))
				return
			}
			tempDelay = 0

			go p.handleConn(conn, recvPacketMaxLen, onParseProtoHead, sendChanCapacity)
		}
	}()

	return
}

//停止 AcceptTCP
func (p *Server) Stop() {
	if p.listener != nil {
		p.listener.Close()
		p.listener = nil
	}
}

//主动断开连接
func (p *Server) DisConn(remote *Remote) (err error) {
	if !remote.IsConn() {
		return errors.New("[ERROR]link disconnect.")
	}
	p.tcpChan <- &EventDisConnServer{
		Server: p,
		Remote: remote,
	}
	return
}

//处理事件时调用
func (p *Server) EventDisconn(remote *Remote) {
	p.OnDisconn(remote)
	if remote.IsConn() {
		remote.stop()
	}
}

func (p *Server) handleConn(conn *net.TCPConn, recvPacketMaxLen uint32, onParseProtoHead OnParseProtoHeadFunc, sendChanCapacity uint32) {
	//log.Printf("connection from:%v", conn.RemoteAddr().String())
	remote := &Remote{}

	remote.addEventDisConn2EventChan = func() {
		p.tcpChan <- &EventDisConnServer{
			Server: p,
			Remote: remote,
		}
	}
	remote.addEventPacket2EventChan = func(data []byte, packetLength int) {
		pes := &EventPacketServer{
			Server: p,
			Data:   make([]byte, packetLength),
			Remote: remote,
		}

		copy(pes.Data, data[:packetLength])
		p.tcpChan <- pes
	}
	remote.start(conn, sendChanCapacity, recvPacketMaxLen, onParseProtoHead)

	//链接上
	p.tcpChan <- &EventConnServer{
		Server: p,
		Remote: remote,
	}
}
