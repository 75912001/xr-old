package tcp

import (
	"errors"
)

//发送数据(必须在处理EventChan事件中调用)
func Send(remote *Remote, data []byte) (err error) {
	if !remote.IsConn() {
		return errors.New("[ERROR]link disconnect.")
	}
	remote.sendChan <- &sendEvent{
		data: data,
		dst:  remote,
	}
	return
}
