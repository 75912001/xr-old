package tcp

import (
	"errors"
	"fmt"
	"net"
	"sync/atomic"
	"time"

	"github.com/75912001/xr/lib/log"

	"github.com/75912001/xr/lib/util"
)

type Server struct {
	OnConn      OnConnServerType
	OnPacket    OnPacketServerType
	OnDisConn   OnDisConnServerType
	tcpChan     chan<- interface{}
	listener    *net.TCPListener
	recvChanCnt uint32
	sendChanCnt uint32
	log         *log.Log
}

//运行服务
//address:127.0.0.1:8787
//rwBuffLen:tcp recv/send 缓冲大小
//recvPacketMaxLen:最大包长(包头+包体)
//eventChan:外部传递的事件处理
func (p *Server) Strat(address string, log *log.Log, recvPacketMaxLen int, eventChan chan<- interface{},
	onConn OnConnServerType, onDisconn OnDisConnServerType, onPacket OnPacketServerType, onParseProtoHead OnParseProtoHeadType,
	sendChanCapacity int) (err error) {
	p.log = log
	p.OnConn = onConn
	p.OnPacket = onPacket
	p.OnDisConn = onDisconn
	p.tcpChan = eventChan
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if nil != err {
		p.log.Crit(fmt.Sprintf("net.ResolveTCPAddr, err:%v, address:%v", err, address))
		return
	}

	//TODO improvement [设置地址复用]
	//TODO improvement [设置监听的缓冲数量]

	p.listener, err = net.ListenTCP("tcp", tcpAddr)
	if nil != err {
		p.log.Crit(fmt.Sprintf("net.ListenTCP, tcpAddr:%v, err:%v", tcpAddr, err))
		return
	}

	go func() {
		p.log.Trace("AcceptTCP goroutine start.")
		defer func() {
			if err := recover(); err != nil {
				p.log.Crit(fmt.Sprintf("%v accept goroutine panic:%v", util.GetCurrentFuncName(), err))
			}
			p.log.Trace("AcceptTCP goroutine exit.")
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
					p.log.Warn(fmt.Sprintf("listen.AcceptTCP, tempDelay:%v, err:%v", tempDelay, err))
					time.Sleep(tempDelay)
					continue
				}
				p.log.Crit(fmt.Sprintf("listen.AcceptTCP, err:%v", err))
				return
			}
			tempDelay = 0
			//TODO 去掉里面的go read
			go p.handleConn(conn, recvPacketMaxLen, onParseProtoHead, sendChanCapacity)
		}
	}()

	return
}

//退出 AcceptTCP
func (p *Server) Exit() {
	if p.listener != nil {
		p.listener.Close()
		p.listener = nil
	}
}

//主动断开连接
func (p *Server) DisConn(remote *Remote) (err error) {
	if !remote.IsConn() {
		p.log.Warn("link disconnect.")
		return errors.New("[ERROR]link disconnect.")
	}
	p.tcpChan <- &DisConnEventServer{
		Server: p,
		Remote: remote,
	}
	return
}

func (p *Server) Info() (recvChanCnt, sendChanCnt uint32) {
	return p.recvChanCnt, p.sendChanCnt
}

func (p *Server) handleConn(conn *net.TCPConn, recvPacketMaxLen int, onParseProtoHead OnParseProtoHeadType, sendChanCapacity int) {
	p.log.Debug(fmt.Sprintf("connection from:%v", conn.RemoteAddr().String()))

	conn.SetNoDelay(true)
	conn.SetReadBuffer(recvPacketMaxLen)
	conn.SetWriteBuffer(recvPacketMaxLen)

	remote := &Remote{
		conn:     conn,
		sendChan: make(chan interface{}, sendChanCapacity),
	}
	go func() {
		atomic.AddUint32(&p.sendChanCnt, 1)
		defer func() {
			if err := recover(); err != nil {
				p.log.Crit(fmt.Sprintf("onSendEvent goroutine panic:%v\n", err))
			}
			atomic.AddUint32(&p.sendChanCnt, ^uint32(0))
		}()
		remote.onSendEvent(p.log)
	}()

	//链接上
	p.tcpChan <- &ConnEventServer{
		Server: p,
		Remote: remote,
	}

	go func() {
		atomic.AddUint32(&p.recvChanCnt, 1)
		defer func() {
			if err := recover(); err != nil {
				p.log.Crit(fmt.Sprintf("onRecvEventChan goroutine panic:%v\n", err))
			}
			atomic.AddUint32(&p.recvChanCnt, ^uint32(0))
		}()
		p.onRecvEventChan(remote, recvPacketMaxLen, onParseProtoHead)
	}()
}
