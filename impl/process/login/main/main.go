package main

import (
	"log"
	"runtime"
	"time"

	"github.com/75912001/xr/impl/service/login"
	"github.com/75912001/xr/impl/service/login/handle_event"
)

func main() {
	var err error

	err = login.GServer.Init(handle_event.OnEventConnServer, handle_event.OnEventDisConnServer, handle_event.OnEventPacketServer,
		handle_event.OnParseProtoHeadServer, handle_event.OnEventAddrMulticast)
	if err != nil {
		log.Fatalf("server init err:%v", err)
		return
	}

	defer func() {
		login.GServer.Stop()
	}()

	for {
		time.Sleep(time.Second)
		log.Printf("goroutine cnt:%v", runtime.NumGoroutine())
	}
}
