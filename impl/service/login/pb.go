package login

import (
	"github.com/75912001/xr/impl/service/login/pb"
	"github.com/75912001/xr/impl/service/protobuf/login_proto"
	"github.com/75912001/xr/lib/pb_func_mgr"
)

func Init() {
	GPbFunMgr.Register(pb_func_mgr.MessageID(login_proto.CMD_LOGIN_MSG), pb.OnLoginMsg, new(login_proto.LoginMsg))
}
