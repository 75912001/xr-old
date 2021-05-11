package addr_test

import (
	"log"
	"testing"
	"time"

	"github.com/75912001/xr/lib/addr"
)

var eventChan = make(chan interface{}, 10000)

func OnAddrMulticastType(name string, id uint32, ip string, port uint16, data string) int {
	log.Printf("name:%v, id:%v, ip:%v, port:%v, data:%v", name, id, ip, port, data)
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
			log.Fatalf("non-existent event:%v", v)
		}
	}
}

func TestAddr(t *testing.T) {

	var a addr.Addr

	err := a.Start(eventChan, OnAddrMulticastType, "239.0.0.8", 8890, "Wi-Fi",
		"testService", 1, "127.0.0.1", 8899, "this is data.")
	if err != nil {
		t.Fatalf("addr multicase err:%v", err)
		return
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("handleEvent goroutine panic:%v", err)
			}
			t.Logf("handleEvent goroutine done.")
		}()
		handleEvent()
	}()
	time.Sleep(time.Second * 5)
	a.Exit()
}
