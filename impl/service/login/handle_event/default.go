package handle_event

import (
	"fmt"
	"github.com/75912001/xr/impl/service/common/proto"
	"github.com/75912001/xr/impl/service/common/proto_head"
	"github.com/75912001/xr/impl/service/login"
	"github.com/75912001/xr/impl/service/protobuf/login_proto"
)

func OnEventDefault(v interface{}) int {
	switch v.(type) {
	//server
	case *LoginMsgRes:
		vv, ok := v.(*LoginMsgRes)
		if ok {
			if !vv.Remote.IsConn() {
				break
			}
			err := proto.Send(vv.Remote, vv.Value.(*login_proto.LoginMsgRes),
				proto_head.MessageID(login_proto.CMD_LOGIN_MSG), 0, 0, 0)
			if err != nil {
				login.GServer.Log.Error("send err:", err)
				return 0
			}
		}
	default:
		login.GServer.Log.Crit(fmt.Sprintf("non-existent event:%v", v))
	}

	return 0
}
