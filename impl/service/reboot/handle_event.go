package reboot

import (
	"fmt"

	"github.com/75912001/xr/lib/tcp"
	"github.com/75912001/xr/lib/timer"
)

func (p *RebootMgr) handleEvent() {
	for v := range p.EventChan {
		switch v.(type) {
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
		default:
			p.Log.Crit(fmt.Sprintf("non-existent event:%v", v))
		}
	}
}
