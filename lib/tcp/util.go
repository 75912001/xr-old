package tcp

import (
	"errors"

	"github.com/75912001/xr/lib/log"
)

var GLog *log.Log

//发送数据(必须在处理EventChan事件中调用)
func Send(remote *Remote, data []byte) (err error) {
	if !remote.IsConn() {
		GLog.Warn("link disconnect.")
		return errors.New("[ERROR]link disconnect.")
	}
	remote.sendChan <- &sendEvent{
		data: data,
		dst:  remote,
	}
	return
}
