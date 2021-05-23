package login

import (
	"github.com/75912001/xr/impl/service/common/proto_head"
	"github.com/75912001/xr/impl/service/login/world_mgr"
	"github.com/75912001/xr/impl/service/protobuf/login_proto"
	"google.golang.org/protobuf/proto"
)

func OnLoginMsg(protoHead interface{}, protoMessage *proto.Message, obj interface{}) (ret int) {
	var ok bool
	var ph *proto_head.ProtoHead
	{
		ph, ok = protoHead.(*proto_head.ProtoHead)
		if !ok {
			return -1
		}
	}
	GServer.Log.Trace(ph)

	var world *world_mgr.World
	{
		world, ok = obj.(*world_mgr.World)
		if !ok {
			return -1
		}
	}
	GServer.Log.Trace(world)

	in := (*protoMessage).(*login_proto.LoginMsg)
	GServer.Log.Trace(in.String())

	GWorldMgr.AddById(in.Id, world)
	world.IP = in.Ip
	world.Port = uint16(in.Port)

	return 0
}
