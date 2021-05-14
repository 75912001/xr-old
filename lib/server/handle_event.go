package server

import (
	"fmt"

	"github.com/75912001/xr/lib/addr"

	"github.com/75912001/xr/lib/tcp"
	"github.com/75912001/xr/lib/timer"
)

func (p *Server) handleEvent() {
	for v := range p.eventChan {
		switch v.(type) {
		//server
		case *tcp.EventConnServer:
			vv, ok := v.(*tcp.EventConnServer)
			if ok {
				vv.Server.OnConn(vv.Remote)
			}
		case *tcp.EventDisConnServer:
			vv, ok := v.(*tcp.EventDisConnServer)
			if ok {
				vv.Server.EventDisconn(vv.Remote)
			}
		case *tcp.EventPacketServer:
			vv, ok := v.(*tcp.EventPacketServer)
			if ok {
				if !vv.Remote.IsConn() {
					continue
				}
				vv.Server.OnPacket(vv.Remote, vv.Data)
			}
			//client
		case *tcp.EventDisConnClient:
			vv, ok := v.(*tcp.EventDisConnClient)
			if ok {
				vv.Client.EventDisConn()
			}
		case *tcp.EventPacketClient:
			vv, ok := v.(*tcp.EventPacketClient)
			if ok {
				if !vv.Client.IsConn() {
					continue
				}
				vv.Client.OnEventPacket(vv.Client, vv.Data)
			}
			//timer
		case *timer.Second:
			v1, ok1 := v.(*timer.Second)
			if ok1 {
				if v1.IsValid() {
					v1.Function(v1.Arg)
				}
			}
		case *timer.Millisecond:
			v1, ok1 := v.(*timer.Millisecond)
			if ok1 {
				if v1.IsValid() {
					v1.Function(v1.Arg)
				}
			}
		//addrMulticase
		case *addr.EventAddrMulticast:
			vv, ok := v.(*addr.EventAddrMulticast)
			if ok {
				vv.Addr.OnEventAddrMulticast(vv.AddrJson.Name, vv.AddrJson.ID, vv.AddrJson.IP, vv.AddrJson.Port, vv.AddrJson.Data)
			}
		default:
			p.Log.Crit(fmt.Sprintf("non-existent event:%v", v))
		}
	}
}
