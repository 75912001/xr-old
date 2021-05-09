package handleTcp

import (
	"fmt"

	"github.com/75912001/xr/lib/tcp"
)

func OnDisConnClient(client *tcp.Client) int {
	fmt.Println("OnDisconnClient")
	if !client.Remote.IsConn() {
		//GLog.Warn("duplicate shutdowns")
		return 0
	}
	return 0
}

func OnPacketClient(client *tcp.Client, data []byte) int {
	fmt.Println("OnPacketClient")
	return 0
}

func OnParseProtoHeadClient(data []byte, length int) int {
	//解析协议包头 返回长度:完整包总长度  返回0:不是完整包 返回-1:包错误
	fmt.Println("OnParseProtoHeadClient")
	return len(data)
}
