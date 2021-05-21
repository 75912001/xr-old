package handle_event

import "github.com/75912001/xr/lib/tcp"

//数据包事件

type LoginMsgRes struct {
	Value  interface{}
	Remote *tcp.Remote
}
