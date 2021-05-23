package world

import "github.com/75912001/xr/impl/service/protobuf/login_proto"

func InitPbLogin() {
	GPbLoginFunMgr.Register(uint32(login_proto.CMD_LOGIN_MSG), OnLoginMsgRes, new(login_proto.LoginMsgRes))
}
