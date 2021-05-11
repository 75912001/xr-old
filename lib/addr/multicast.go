package addr

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/ipv4"
)

//multicast 组播
type multicast struct {
	conn                   *net.UDPConn
	mcaddr                 *net.UDPAddr
	cancelFunc             context.CancelFunc
	waitGroupGoroutineDone sync.WaitGroup
}

// 运行
func (p *multicast) start(ctx context.Context, ip string, port uint16, netName string, addr *Addr) (err error) {
	var strAddr = ip + ":" + strconv.Itoa(int(port))
	p.mcaddr, err = net.ResolveUDPAddr("udp4", strAddr)
	if err != nil {
		log.Printf("net.ResolveUDPAddr err:%v", err)
		return err
	}

	p.conn, err = net.ListenUDP("udp4", p.mcaddr)
	if err != nil {
		log.Printf("ListenUDP err:%v", err)
		return err
	}

	pc := ipv4.NewPacketConn(p.conn)

	iface, err := net.InterfaceByName(netName)
	if err != nil {
		log.Printf("can't find specified interface err:%v", err)
		return err
	}

	network, _ := net.ResolveIPAddr("ip4", ip)
	err = pc.JoinGroup(iface, network)
	if nil != err {
		log.Printf("err:%v, address::%v", err, network)
		return err
	}

	if loop, err := pc.MulticastLoopback(); err == nil {
		log.Printf("MulticastLoopback status:%v", loop)
		if !loop {
			if err := pc.SetMulticastLoopback(true); err != nil {
				log.Printf("SetMulticastLoopback err:%v", err)
			}
		}
	}

	ctxWithCancel, cancelFunc := context.WithCancel(ctx)
	p.cancelFunc = cancelFunc

	p.waitGroupGoroutineDone.Add(2)
	//读
	go func(addr *Addr) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("ReadFromUDP goroutine panic:%v", err)
			}
			p.waitGroupGoroutineDone.Done()
		}()

		//数据包大小
		const packetMax int = 1024
		for {
			recvBuf := make([]byte, packetMax)
			length, _, err := p.conn.ReadFromUDP(recvBuf)
			if nil != err {
				log.Printf("ReadFromUDP err:%v", err)
				break
			}
			recvBuf = recvBuf[0:length]
			err = addr.handleAddrMulticast(recvBuf)
			if err != nil {
				log.Printf("handleAddrMulticast err:%v", err)
			}
		}
	}(addr)

	//10-20sec 同步一次
	go func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("doAddrSYN goroutine panic:%v", err)
			}
			p.waitGroupGoroutineDone.Done()
		}()
		p.doAddrSYN([]byte(addr.addrFirstBuffer))
		for {
			select {
			case <-ctx.Done():
				log.Printf("doAddrSYN goroutine ctx done.")
				return
			case <-time.After(time.Duration(rand.Intn(10)+10) * time.Second):
				p.doAddrSYN([]byte(addr.addrBuffer))
			}
		}
	}(ctxWithCancel)
	return
}

func (p *multicast) exit() {
	//触发ReadFromUDP goroutine 退出
	if p.conn != nil {
		p.conn.Close()
	}

	//触发doAddrSYN goroutine 退出
	if p.cancelFunc != nil {
		p.cancelFunc()
		p.waitGroupGoroutineDone.Wait()
	}

	p.conn = nil
}

func (p *multicast) doAddrSYN(data []byte) {
	_, err := p.conn.WriteToUDP(data, p.mcaddr)

	if nil != err {
		log.Printf("doAddrSYN err:%v, %v", err, data)
	}
}
