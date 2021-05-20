package service_mgr

import (
	"github.com/75912001/xr/lib/addr"
	"github.com/75912001/xr/lib/tcp"
)

type Service struct {
	addrJson addr.AddrJson
	Client   tcp.Client
}
