package handle_http

import (
	"fmt"

	"github.com/75912001/xr/lib/util"
)

func genVerify(platform uint32, account string) (verifyString string) {
	var s string
	s += "platform=" + fmt.Sprint(platform)
	s += "account=" + account
	s += "miyin"

	verifyString = util.GenMd5([]byte(s))
	return
}
