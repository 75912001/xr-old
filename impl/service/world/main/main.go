package main

import (
	"log"
	"time"

	"github.com/75912001/xr/impl/service/world/handleTcp"

	"github.com/75912001/xr/impl/service/common/server"
)

func t(d2 []byte) {
	//d2 = append(d2, '1')
	d2[0] = 'a'
	d2[1] = 'b'
	d2[2] = 'c'
	d2 = d2[1:]
	log.Print("d2:", d2)
}
func main() {
	{
		var data []byte
		data = make([]byte, 8)
		s := append(data, '1')
		log.Printf("s:%v\n", s)
		data[0] = '1'
		data[1] = '2'
		data[2] = '3'
		t(data)
		log.Print("data:", data)
	}
	var err error
	var server server.Server
	err = server.Init(handleTcp.OnConnServer, handleTcp.OnDisConnServer, handleTcp.OnPacketServer, handleTcp.OnParseProtoHeadServer)
	if err != nil {
		log.Fatalf("server init err:%v", err)
		return
	}
	defer func() {
		server.Exit()
	}()

	time.Sleep(time.Second * 3)
	return
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


[addr_multicast]
#mcast_ip=239.0.0.1
#mcast_port=5001
#interface on which arriving multicast datagrams will be received
#mcast_incoming_if=enp0s3
#data=0

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
