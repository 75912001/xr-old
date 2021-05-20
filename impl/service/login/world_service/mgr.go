package world_service

import (
	"github.com/75912001/xr/impl/service/common/client_mgr"
)

type WorldServiceMgr struct {
	client_mgr.ClientMgr

}

func (p *WorldServiceMgr) GetRandWorldService() (worldService *client_mgr.Client) {
	for k, _ := range p.ClientMap {
		worldService = p.Find(k)
		break
	}
	return
}
