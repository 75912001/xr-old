package handle_event

import (
	"github.com/75912001/xr/lib/tcp"
)

func OnEventConnServer(remote *tcp.Remote) int {
	//TODO 业务逻辑
	return 0
}

func OnEventDisConnServer(remote *tcp.Remote) int {
	if !remote.IsConn() {
		return 0
	}
	//TODO 业务逻辑
	return 0
}

func OnEventPacketServer(remote *tcp.Remote, data []byte) int {
	//TODO 业务逻辑
	return 0
}

func OnParseProtoHeadServer(data []byte, length int) int {
	//解析协议包头 返回长度:完整包总长度  返回0:不是完整包 返回-1:包错误
	//TODO 业务逻辑
	return length
}
