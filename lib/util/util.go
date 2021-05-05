package util

import "runtime"

func GetFuncName() (funcName string) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "GetFuncName err: nil"
	}
	return runtime.FuncForPC(pc).Name()
}
