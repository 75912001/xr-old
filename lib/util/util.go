package util

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetCurrentFuncName() (funcName string) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "GetCurrentFuncName err: nil"
	}
	return runtime.FuncForPC(pc).Name()
}

//TODO 注意:不支持 link/快捷方式
func GetCurrentPath() (currentPath string, err error) {
	currentPath, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}
	return
}
