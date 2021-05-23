package login

import (
	"github.com/75912001/xr/impl/service/protobuf/login_proto"
)

func Init() {
	GPbFunMgr.Register(uint32(login_proto.CMD_LOGIN_MSG), OnLoginMsg, new(login_proto.LoginMsg))
}
