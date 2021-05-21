package world_mgr

import (
	"github.com/75912001/xr/lib/tcp"
)

type World struct {
	Remote *tcp.Remote
	IP     string
	Port   uint16
	Id     uint32
}
