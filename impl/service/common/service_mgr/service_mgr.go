package service_mgr

import (
	"github.com/75912001/xr/lib/tcp"
)

type SERVICE_MAP map[*tcp.Client]*Service
type SERVICE_ID_MAP map[uint32]*Service //key:service id, val:service

type ServiceMgr struct {
	serviceMap   SERVICE_MAP
	serviceIDMap SERVICE_ID_MAP
}

func (p *ServiceMgr) Init() {
	p.serviceMap = make(SERVICE_MAP)
	p.serviceIDMap = make(SERVICE_ID_MAP)
}

func (p *ServiceMgr) AddService(service *Service) {
	p.serviceMap[&service.Client] = service
	p.serviceIDMap[service.addrJson.ID] = service
	return
}

func (p *ServiceMgr) DelService(service *Service) {
	delete(p.serviceMap, &service.Client)
	delete(p.serviceIDMap, service.addrJson.ID)
}

func (p *ServiceMgr) Find(client *tcp.Client) (service *Service) {
	service, _ = p.serviceMap[client]
	return
}

func (p *ServiceMgr) FindById(id uint32) (service *Service) {
	service, _ = p.serviceIDMap[id]
	return
}
