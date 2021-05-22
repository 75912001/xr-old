package login

import (
	"github.com/75912001/xr/impl/service/login/bench"
	"github.com/75912001/xr/impl/service/login/world_mgr"
	"github.com/75912001/xr/lib/pb_func_mgr"
	"github.com/75912001/xr/lib/server"
)

var GServer server.Server
var GBench bench.Mgr
var GWorldMgr world_mgr.WorldMgr
var GPbFunMgr pb_func_mgr.PbFunMgr

func init() {
	GWorldMgr.Init()
	GPbFunMgr.Init()
	Init()
}
