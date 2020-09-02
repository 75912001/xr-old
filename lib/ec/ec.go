package ec

import (
	"fmt"
	"sync"
)

var rwmux sync.RWMutex

//注册（错误码，名称，描述）
func Register(errorCode int, name string, description string) (err error) {
	rwmux.Lock()
	defer func() {
		rwmux.Unlock()
	}()
	return addErrorCode(errorCode, name, description)
}

//描述信息
func Description(errorCode int) string {
	rwmux.RLock()
	defer func() {
		rwmux.RUnlock()
	}()

	ei := getErrorInfo(errorCode)
	if ei == nil {
		return fmt.Sprintf("unknown error code:%v", errorCode)
	}
	return ei.getDescription()
}

//详细信息
func Detail(errorCode int) string {
	rwmux.RLock()
	defer func() {
		rwmux.RUnlock()
	}()

	ei := getErrorInfo(errorCode)
	if ei == nil {
		return fmt.Sprintf("unknown error code:%v", errorCode)
	}
	return ei.getDetail()
}

type EC int

//错误信息
func (p EC) Error() string {
	errorCode := int(p)

	if ECSucess == errorCode {
		return ""
	}

	rwmux.RLock()
	defer func() {
		rwmux.RUnlock()
	}()
	ei := getErrorInfo(errorCode)
	if ei == nil {
		return fmt.Sprintf("unknown error code:%v", errorCode)
	}
	return ei.getDetail()
}
