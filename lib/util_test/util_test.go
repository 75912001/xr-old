package util_test

import (
	"testing"

	"github.com/75912001/xr/lib/util"
)

func TestGetCurrentFuncName(t *testing.T) {
	t.Logf("CurrentFuncName:%v", util.GetCurrentFuncName())
}

func TestGetCurrentPath(t *testing.T) {
	currentPath, err := util.GetCurrentPath()
	if err != nil {
		t.Errorf("GetCurrentPath err:%v", err)
	}
	t.Logf("CurrentPath:%v", currentPath)
}
