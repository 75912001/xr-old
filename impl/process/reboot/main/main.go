package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/75912001/xr/lib/util"
	"log"
	"runtime"
	"time"

	"github.com/75912001/xr/impl/service/common/proto_head"
	"github.com/75912001/xr/impl/service/reboot/handle_event"
	"github.com/75912001/xr/lib/tcp"

	"github.com/75912001/xr/impl/service/reboot"
)

func main() {
	if !util.IsLittleEndian() {
		log.Fatalf("system is bigEndian!")
	}

	err := reboot.GRebootMgr.Init()
	if err != nil {
		log.Fatalf("rebootMgr init err:%v", err)
		return
	}

	for i := uint32(0); i < reboot.GBench.Json.Reboot.Cnt; i++ {
		go func() {
			var c tcp.Client
			j := &reboot.GRebootMgr.BenchMgr.Json
			address := fmt.Sprintf("%v:%v", reboot.GBench.Json.Reboot.ServerIP, reboot.GBench.Json.Reboot.ServerPort)
			err := c.Connect(address, j.Base.PacketLengthMax, reboot.GRebootMgr.EventChan,
				handle_event.OnEventDisConnClient, handle_event.OnEventPacketClient, handle_event.OnParseProtoHeadClient, j.Base.SendChanCapacity)
			if err != nil {
				log.Printf("connect server err:%v", err)
				return
			}
			for {
				time.Sleep(time.Second * 1)

				var ph proto_head.ProtoHead
				ph.PacketLength = 24
				ph.MessageID = 1
				ph.SessionID = 2
				ph.ResultID = 3
				ph.UserID = 4

				buf := new(bytes.Buffer)
				binary.Write(buf, binary.LittleEndian, ph.PacketLength)
				binary.Write(buf, binary.LittleEndian, ph.MessageID)
				binary.Write(buf, binary.LittleEndian, ph.SessionID)
				binary.Write(buf, binary.LittleEndian, ph.ResultID)
				binary.Write(buf, binary.LittleEndian, ph.UserID)

				//c.Remote.Send(buf.Bytes())
			}
		}()
	}

	for {
		time.Sleep(time.Second * 5)
		log.Printf("goroutine cnt:%v", runtime.NumGoroutine())
	}
}
