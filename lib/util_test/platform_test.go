package util

import (
	"testing"

	"github.com/75912001/xr/lib/util"
)

func TestIsWindows(t *testing.T) {
	t.Logf("IsWindows:%v", util.IsWindows())
}

func TestIsLinux(t *testing.T) {
	t.Logf("IsLinux:%v", util.IsLinux())
}

func TestIsDarwin(t *testing.T) {
	t.Logf("IsDarwin:%v", util.IsDarwin())
}
