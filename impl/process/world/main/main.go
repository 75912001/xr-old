package main

import (
	"log"
	"time"

	"github.com/75912001/xr/impl/service/world"

	"github.com/75912001/xr/impl/service/world/handle_event"
)

func main() {
	var err error

	err = world.GServer.Init(handle_event.OnEventConnServer, handle_event.OnEventDisConnServer, handle_event.OnEventPacketServer,
		handle_event.OnParseProtoHeadServer, handle_event.OnEventAddrMulticast)
	if err != nil {
		log.Fatalf("server init err:%v", err)
		return
	}
	defer func() {
		world.GServer.Stop()
	}()

	for {
		time.Sleep(time.Second)
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
