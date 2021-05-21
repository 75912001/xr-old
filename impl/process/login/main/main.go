package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"runtime"
	"time"

	"github.com/75912001/xr/impl/service/login/handle_http"

	"github.com/75912001/xr/lib/util"

	"github.com/75912001/xr/impl/service/login"
	"github.com/75912001/xr/impl/service/login/handle_event"
)

func main() {
	if !util.IsLittleEndian() {
		log.Panicf("system is bigEndian!")
	}

	var err error

	err = login.GServer.Init(handle_event.OnEventConnServer, handle_event.OnEventDisConnServer, handle_event.OnEventPacketServer,
		handle_event.OnParseProtoHeadServer, handle_event.OnEventAddrMulticast, handle_event.OnEventDefault)
	if err != nil {
		log.Fatalf("server init err:%v", err)
		return
	}

	{ //加载bench.json文件中login自有field
		currentPath, err := util.GetCurrentPath()
		if err != nil {
			login.GServer.Log.Crit("GetCurrentPath fatal:", err)
			return
		}
		{
			err = util.UnmarshalJsonFile(path.Join(currentPath, "bench.json"), &login.GBench.Json)
			if err != nil {
				login.GServer.Log.Crit("parse bench.json err:", err)
				return
			}
			login.GServer.Log.Info(fmt.Sprintf("bench json:%+v", login.GBench.Json))
		}
		{ //check
			if len(login.GBench.Json.LoginHttp.Pattern) == 0 || len(login.GBench.Json.LoginHttp.IP) == 0 || login.GBench.Json.LoginHttp.Port == 0 {
				login.GServer.Log.Crit("http address is error.")
				return
			}
		}
	}

	//HTTP 登录服务
	{
		http.HandleFunc(login.GBench.Json.LoginHttp.Pattern, handle_http.LoginHttpHandler)
		//http 服务开始
		httpAddr := fmt.Sprintf("%v:%v", login.GBench.Json.LoginHttp.IP, login.GBench.Json.LoginHttp.Port)
		err = http.ListenAndServe(httpAddr, nil)
		if nil != err {
			login.GServer.Log.Crit("ListenAndServe err: ", err, httpAddr)
			return
		}
	}

	defer func() {
		login.GServer.Stop()
	}()

	for {
		time.Sleep(time.Second)
		log.Printf("goroutine cnt:%v", runtime.NumGoroutine())
	}
}
