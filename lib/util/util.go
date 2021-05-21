package util

import (
	"os"
	"path/filepath"
	"runtime"
	"unsafe"
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

func IsLittleEndian() bool {
	var value int32 = 1 // 占4byte 转换成16进制 0x00 00 00 01
	// 大端(16进制)：00 00 00 01
	// 小端(16进制)：01 00 00 00
	pointer := unsafe.Pointer(&value)
	pb := (*byte)(pointer)
	if *pb != 1 {
		//bigEndian
		return false
	}
	//littleEndian
	return true
}
