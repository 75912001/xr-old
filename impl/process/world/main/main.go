package main

import (
	"fmt"
	"log"
	"path"
	"runtime"
	"time"

	"github.com/75912001/xr/lib/util"

	"github.com/75912001/xr/impl/service/world"

	"github.com/75912001/xr/impl/service/world/handle_event"
)

func main() {
	if !util.IsLittleEndian() {
		log.Panicf("system is bigEndian!")
	}

	var err error

	err = world.GServer.Init(handle_event.OnEventConnServer, handle_event.OnEventDisConnServer, handle_event.OnEventPacketServer,
		handle_event.OnParseProtoHeadServer, handle_event.OnEventAddrMulticast, handle_event.OnEventDefault)
	if err != nil {
		log.Fatalf("server init err:%v", err)
		return
	}
	{ //加载bench.json文件中world自有field
		currentPath, err := util.GetCurrentPath()
		if err != nil {
			world.GServer.Log.Crit("GetCurrentPath fatal:", err)
			return
		}
		{
			err = util.UnmarshalJsonFile(path.Join(currentPath, "bench.json"), &world.GBench.Json)
			if err != nil {
				world.GServer.Log.Crit("parse bench.json err:", err)
				return
			}
			world.GServer.Log.Info(fmt.Sprintf("bench json:%+v", world.GBench.Json))
		}
		{ //check
			if len(world.GBench.Json.DB.IP) == 0 || world.GBench.Json.DB.Port == 0 {
				world.GServer.Log.Crit("db address is error.")
				return
			}
		}
	}

	//err = world.GMongodbMgr.Connect(world.GServer.BenchMgr.Json.DB.IP, world.GServer.BenchMgr.Json.DB.Port)
	//if err != nil {
	//	log.Fatalf("mongodb connect err:%v", err)
	//	return
	//}

	defer func() {
		world.GServer.Stop()
	}()
	//go func() {
	//	fmt.Println("pprof start...")
	//	fmt.Println(http.ListenAndServe(":9876", nil))
	//	for {
	//		time.Sleep(time.Second)
	//	}
	//}()
	for {
		time.Sleep(time.Second)
		log.Printf("goroutine cnt:%v", runtime.NumGoroutine())
	}
}

/*


[server]
platform=20

[http_server]
ip=10.100.6.209
port=22501

[pay]
mysql_ip=10.100.6.185
mysql_port=3306
mysql_user=root
mysql_pwd=111111
mysql_db_name=S200001_USER

url_callback_pattern=/pay;/zfb_pay;/wx_pay;/dangbei_pay
http_ip=10.100.6.185
http_port=22502

[login]
wx_appid=wx95343f63f344167e
wx_secret=390658a8b5014c6a7f18968586575ca4
#有0、46、64、96、132数值可选，0代表640*640正方形头像
wx_head_size=96

#end
*/
