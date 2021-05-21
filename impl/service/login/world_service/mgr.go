package world_service

import (
	"github.com/75912001/xr/impl/service/common/client_mgr"
)

type WorldServiceMgr struct {
	clientMgr client_mgr.ClientMgr
}

func (p *WorldServiceMgr) GetRandWorldService() (worldService *client_mgr.Client) {
	for k, _ := range p.clientMgr.ClientMap {
		worldService = p.clientMgr.Find(k)
		break
	}
	return
}

func (p *WorldServiceMgr) Init() {
	p.clientMgr.Init()
}

//数据包事件
type LoginKickMsgRes struct {
	Data []byte
}

type LoginMsgRes struct {
	Data []byte
}
