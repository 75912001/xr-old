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
		case *tcp.ConnEventServer:
			vv, ok := v.(*tcp.ConnEventServer)
			if ok {
				p.GLog.Debug(fmt.Sprintf("ConnEventServer, remote:%v", vv.Remote))
				vv.Server.OnConn(vv.Remote)

			}
		case *tcp.DisConnEventServer:
			vv, ok := v.(*tcp.DisConnEventServer)
			if ok {
				p.GLog.Debug(fmt.Sprintf("DisConnEventServer, remote:%v", vv.Remote))
				vv.Server.OnDisConn(vv.Remote)
				if vv.Remote.IsConn() {
					vv.Remote.Close()
				}
			}
		case *tcp.PacketEventServer:
			vv, ok := v.(*tcp.PacketEventServer)
			if ok {
				if !vv.Remote.IsConn() {
					continue
				}
				vv.Server.OnPacket(vv.Remote, vv.Data)
			}
			//client
		case *tcp.DisConnEventClient:
			vv, ok := v.(*tcp.DisConnEventClient)
			if ok {
				p.GLog.Debug(fmt.Sprintf("DisconnEventClient, remote:%v", vv.Client.Remote))
				vv.Client.OnDisConn(vv.Client)
				if vv.Client.Remote.IsConn() {
					vv.Client.Remote.Close()
				}
			}
		case *tcp.PacketEventClient:
			vv, ok := v.(*tcp.PacketEventClient)
			if ok {
				if !vv.Client.Remote.IsConn() {
					continue
				}
				vv.Client.OnPacket(vv.Client, vv.Data)
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
