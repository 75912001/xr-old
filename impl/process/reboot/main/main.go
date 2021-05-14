package main

import (
	"log"
	"time"

	"github.com/75912001/xr/lib/tcp"

	"github.com/75912001/xr/impl/service/reboot"
	"github.com/75912001/xr/impl/service/reboot/handle_event"
)

func main() {
	err := reboot.GRebootMgr.Init()
	if err != nil {
		log.Fatalf("rebootMgr init err:%v", err)
		return
	}

	var c tcp.Client
	j := &reboot.GRebootMgr.BenchMgr.Json
	c.Connect(j.Server.Address, j.Base.PacketLengthMax, j.Base.PacketLengthMax, reboot.GRebootMgr.EventChan,
		handle_event.OnEventDisConnClient, handle_event.OnEventPacketClient, handle_event.OnParseProtoHeadClient, j.Base.SendChanCapacity)
	for {
		time.Sleep(time.Second * 10)
		c.Remote.Send([]byte("123"))
	}
}
