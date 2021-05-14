package user

import (
	"github.com/75912001/xr/impl/service/common/type_define"
	"github.com/75912001/xr/lib/tcp"
)

type User struct {
	remote *tcp.Remote
	uid    type_define.USER_ID
}
