package util

import (
	"testing"
	"time"

	"github.com/75912001/xr/lib/util"
)

func TestGenYYYYMMDD(t *testing.T) {
	yyyymmdd := util.GenYYYYMMDD(time.Now().Unix())
	t.Logf("YYYYMMDD:%v", yyyymmdd)
}
