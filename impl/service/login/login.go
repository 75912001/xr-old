package login

import (
	"github.com/75912001/xr/impl/service/login/bench"
	"github.com/75912001/xr/impl/service/login/world_service"
	"github.com/75912001/xr/lib/server"
)

var GServer server.Server
var GBench bench.Mgr
var GWorldServiceMgr world_service.WorldServiceMgr

func init() {
	GWorldServiceMgr.Init()
}
