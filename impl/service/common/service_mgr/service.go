package service_mgr

import (
	"github.com/75912001/xr/lib/addr"
	"github.com/75912001/xr/lib/tcp"
)

type Service struct {
	AddrJson addr.AddrJson
	Client   tcp.Client
}
