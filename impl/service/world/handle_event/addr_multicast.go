package handle_event

import (
	"fmt"

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
		return 0
	case common.MONGODB_NAME:
		return 0
	}
	return 0
}
