package login

import (
	"github.com/75912001/xr/impl/service/login/bench"
	"github.com/75912001/xr/impl/service/login/world_mgr"
	"github.com/75912001/xr/lib/server"
)

var GServer server.Server
var GBench bench.Mgr
var GWorldMgr world_mgr.WorldMgr

func init() {
	GWorldMgr.Init()
}
