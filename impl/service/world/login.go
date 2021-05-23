package world

import (
	"github.com/75912001/xr/impl/service/common/proto_head"
	"github.com/75912001/xr/impl/service/common/service_mgr"
	"github.com/75912001/xr/impl/service/protobuf/login_proto"
	"google.golang.org/protobuf/proto"
)

func OnLoginMsgRes(protoHead interface{}, protoMessage *proto.Message, obj interface{}) (ret int) {
	var ok bool
	var ph *proto_head.ProtoHead
	{
		ph, ok = protoHead.(*proto_head.ProtoHead)
		if !ok {
			return -1
		}
	}
	GServer.Log.Trace(ph)

	var service *service_mgr.Service
	{
		service, ok = obj.(*service_mgr.Service)
		if !ok {
			return -1
		}
	}
	GServer.Log.Trace(service)

	in := (*protoMessage).(*login_proto.LoginMsgRes)
	GServer.Log.Trace(in.String())

	//todo 登录 ...

	return 0
}
