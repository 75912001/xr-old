package addr_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/75912001/xr/lib/addr"
	"github.com/75912001/xr/lib/log"
)

var eventChan = make(chan interface{}, 10000)
var GLog *log.Log

func init() {
	GLog = new(log.Log)
	GLog.Init("test_log")
}

func OnAddrMulticastType(name string, id uint32, ip string, port uint16, data string) int {
	GLog.Trace(fmt.Sprintf("name:%v, id:%v, ip:%v, port:%v, data:%v", name, id, ip, port, data))
	return 0
}

func handleEvent() {
	for v := range eventChan {
		switch v.(type) {
		//addr
		case *addr.AddrEvent:
			vv, ok := v.(*addr.AddrEvent)
			if ok {
				vv.Addr.OnAddr(vv.AddrJson.Name, vv.AddrJson.ID, vv.AddrJson.IP, vv.AddrJson.Port, vv.AddrJson.Data)
			}
		default:
			GLog.Crit(fmt.Sprintf("non-existent event:%v", v))
		}
	}
}

func TestAddr(t *testing.T) {
	defer func() {
		GLog.Exit()
	}()
	var a addr.Addr

	err := a.Start(GLog, eventChan, OnAddrMulticastType, "239.0.0.8", 8890, "Wi-Fi",
		"testService", 1, "127.0.0.1", 8899, "this is data.")
	if err != nil {
		t.Fatalf("addr multicase err:%v", err)
		return
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				GLog.Warn(fmt.Sprintf("handleEvent goroutine panic:%v", err))
			}
			GLog.Trace("handleEvent goroutine done.")
		}()
		handleEvent()
	}()

	for {
		time.Sleep(time.Second)
	}
}
