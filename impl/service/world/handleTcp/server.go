package handleTcp

import (
	"fmt"

	"github.com/75912001/xr/lib/tcp"
)

func OnConnServer(remote *tcp.Remote) int {
	//	fmt.Println("OnConnServer")
	return 0
}

func OnDisConnServer(remote *tcp.Remote) int {
	fmt.Println("OnDisconnServer")
	if !remote.IsConn() {
		//		GLog.Warn("duplicate shutdowns")
		return 0
	}
	return 0
}

func OnPacketServer(remote *tcp.Remote, data []byte) int {
	fmt.Println("OnPacketServer")
	return 0
}

func OnParseProtoHeadServer(data []byte, length int) int {
	//解析协议包头 返回长度:完整包总长度  返回0:不是完整包 返回-1:包错误
	fmt.Println("OnParseProtoHead")
	return len(data)
	//return 0
}
