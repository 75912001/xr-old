package handle_event

import (
	"github.com/75912001/xr/impl/service/login/world_service"
)

func OnEventDefault(v interface{}) int {
	//TODO 业务逻辑

	switch v.(type) {
	//server
	case *world_service.LoginMsgRes:
		vv, ok := v.(*world_service.LoginMsgRes)
		if ok {
			if !vv.Remote.IsConn() {
				continue
			}
		}
	case *world_service.LoginKickMsgRes:
		vv, ok := v.(*world_service.LoginKickMsgRes)
		if ok {
			if !vv.Remote.IsConn() {
				continue
			}
		}
	default:

		//p.Log.Crit(fmt.Sprintf("non-existent event:%v", v))
	}

	return 0
}
