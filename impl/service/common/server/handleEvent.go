package server

import (
	"fmt"

	"github.com/75912001/xr/lib/timer"

	"github.com/75912001/xr/lib/tcp"
)

func (p *Server) handleEvent() {
	for v := range p.eventChan {
		switch v.(type) {
		//server
		case *tcp.EventConnServer:
			vv, ok := v.(*tcp.EventConnServer)
			if ok {
				p.GLog.Debug(fmt.Sprintf("EventConnServer, remote:%v", vv.Remote))
				vv.Server.OnConn(vv.Remote)

			}
		case *tcp.EventDisConnServer:
			vv, ok := v.(*tcp.EventDisConnServer)
			if ok {
				p.GLog.Debug(fmt.Sprintf("EventDisConnServer, remote:%v", vv.Remote))
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
				p.GLog.Debug(fmt.Sprintf("DisconnEventClient, remote:%v", vv.Client.Remote))
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
		//case *addrmulticase.AddrMulticast:

		default:
			p.GLog.Crit(fmt.Sprintf("non-existent event:%v", v))
		}
	}
}
