package world

import (
	"github.com/75912001/xr/impl/service/common/service_mgr"
	"github.com/75912001/xr/impl/service/world/bench"
	"github.com/75912001/xr/lib/mongodb"
	"github.com/75912001/xr/lib/pb_func_mgr"
	"github.com/75912001/xr/lib/server"
)

var GServer server.Server
var GMongodbMgr mongodb.MongodbMgr
var GLoginMgr service_mgr.ServiceMgr
var GBench bench.Mgr
var GPbLoginFunMgr pb_func_mgr.PbFunMgr

func init() {
	GLoginMgr.Init()
	GPbLoginFunMgr.Init()
	InitPbLogin()
}
