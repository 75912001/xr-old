package ec

import (
	"errors"
	"fmt"
)

//错误信息
var errorInfoMap map[int]*errorInfo = make(map[int]*errorInfo)

func addErrorCode(errorCode int, name string, description string) (err error) {
	if existsErrorCode(errorCode) {
		err = errors.New(fmt.Sprintf("error code value exists. error code:%v, name:%v, description:%v", errorCode, name, description))
		return
	}
	errorInfoMap[errorCode] = &errorInfo{errorCode, name, description}
	return
}

func existsErrorCode(errorCode int) (exists bool) {
	_, exists = errorInfoMap[errorCode]
	return
}

func getErrorInfo(errorCode int) (errorInfo *errorInfo) {
	errorInfo, exists := errorInfoMap[errorCode]
	if !exists {
		return
	}
	return
}
