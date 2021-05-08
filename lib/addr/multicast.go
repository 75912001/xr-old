package addr

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/75912001/xr/lib/log"
	"github.com/75912001/xr/lib/util"
	"golang.org/x/net/ipv4"
)

//数据包大小
const packetMax int = 1024

//multicast 组播
type multicast struct {
	conn   *net.UDPConn
	mcaddr *net.UDPAddr
	log    *log.Log
}

// 运行
func (p *multicast) start(ip string, port uint16, netName string, log *log.Log, addr *Addr) (err error) {
	p.log = log

	var strAddr = ip + ":" + strconv.Itoa(int(port))
	p.mcaddr, err = net.ResolveUDPAddr("udp4", strAddr)
	if err != nil {
		p.log.Crit("net.ResolveUDPAddr err:", err)
		return err
	}

	p.conn, err = net.ListenUDP("udp4", p.mcaddr)
	if err != nil {
		p.log.Crit("ListenUDP err:", err)
		return err
	}

	pc := ipv4.NewPacketConn(p.conn)

	iface, err := net.InterfaceByName(netName)
	if err != nil {
		p.log.Crit("can't find specified interface err:", err)
		return err
	}

	strAddrIpv4, _ := net.ResolveIPAddr("ip4", ip)
	err = pc.JoinGroup(iface, strAddrIpv4)
	if nil != err {
		p.log.Crit("err:", err, strAddrIpv4)
		return err
	}

	if loop, err := pc.MulticastLoopback(); err == nil {
		p.log.Trace("MulticastLoopback status:", loop)
		if !loop {
			if err := pc.SetMulticastLoopback(true); err != nil {
				p.log.Crit("SetMulticastLoopback err:", err)
			}
		}
	}

	//读
	go func(addr *Addr) {
		defer func() {
			p.conn.Close()
			if err := recover(); err != nil {
				p.log.Crit(fmt.Sprintf("%v handleRecv goroutine panic:%v", util.GetFuncName(), err))
			}
		}()

		for {
			recvBuf := make([]byte, packetMax)
			length, _, err := p.conn.ReadFromUDP(recvBuf)
			if nil != err {
				p.log.Crit("handleRecv err:", err)
				break
			}
			recvBuf = recvBuf[0:length]
			err = addr.parse(recvBuf)
			if err != nil {
				p.log.Crit("parse err:", err)
			}
		}
	}(addr)

	//10-20sec 同步一次
	go func(addr *Addr) {
		defer func() {
			if err := recover(); err != nil {
				p.log.Crit(fmt.Sprintf("%v doAddrSYN goroutine panic:%v", util.GetFuncName(), err))
			}
		}()
		var bFirst bool = true
		for {
			if bFirst {
				p.doAddrSYN([]byte(addr.addrFirstBuffer))
				bFirst = false
			} else {
				p.doAddrSYN([]byte(addr.addrBuffer))
			}

			time.Sleep(time.Duration(rand.Intn(10)+10) * time.Second)
		}
	}(addr)
	return err
}

func (p *multicast) doAddrSYN(data []byte) {
	_, err := p.conn.WriteToUDP(data, p.mcaddr)

	if nil != err {
		p.log.Error("doAddrSYN err:", err)
	}
}
