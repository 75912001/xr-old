package handle_event

import (
	"fmt"
	"github.com/75912001/xr/impl/service/common/proto"
	"github.com/75912001/xr/impl/service/common/proto_head"
	"github.com/75912001/xr/impl/service/protobuf/login_proto"

	"github.com/75912001/xr/impl/service/common"
	"github.com/75912001/xr/impl/service/common/service_mgr"
	"github.com/75912001/xr/impl/service/world"
)

func OnEventAddrMulticast(name string, id uint32, ip string, port uint16, data string) int {
	switch name {
	case common.LOGIN_NAME:
		loginService := world.GLoginMgr.FindById(id)
		if loginService != nil {
			return 0
		}
		address := fmt.Sprintf("%v:%v", ip, port)
		j := &world.GServer.BenchMgr.Json
		var service service_mgr.Service
		err := service.Client.Connect(address, j.Base.PacketLengthMax, world.GServer.GetEventChan(),
			OnEventDisConnClient, OnEventPacketClient, OnParseProtoHeadClient, j.Base.SendChanCapacity)
		if err != nil {
			world.GServer.Log.Error(fmt.Printf("connect login service err:%v, name:%v, id:%v, ip:%v, port:%v, data:%v",
				err, name, id, ip, port, data))
			return 0
		}
		world.GLoginMgr.AddService(&service)
		{
			req := &login_proto.LoginMsg{
				Ip:   world.GServer.BenchMgr.Json.Server.IP,
				Port: uint32(world.GServer.BenchMgr.Json.Server.Port),
				Id:   world.GServer.BenchMgr.Json.Base.ServiceID,
			}
			proto.Send(&service.Client.Remote, req, proto_head.MessageID(login_proto.CMD_LOGIN_MSG), 0, 0, 0)
		}
		return 0
	case common.MONGODB_NAME:
		return 0
	}
	return 0
}
