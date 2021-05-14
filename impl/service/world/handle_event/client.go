package handle_event

import (
	"github.com/75912001/xr/lib/tcp"
)

func OnEventDisConnClient(client *tcp.Client) int {
	if !client.IsConn() {
		return 0
	}
	//TODO 业务逻辑
	return 0
}

func OnEventPacketClient(client *tcp.Client, data []byte) int {
	//TODO 业务逻辑
	return 0
}

func OnParseProtoHeadClient(data []byte, length int) int {
	//解析协议包头 返回长度:完整包总长度  返回0:不是完整包 返回-1:包错误
	//TODO 业务逻辑
	return length
}
