package tcp

import (
	"context"
	"fmt"
	"net"

	"github.com/75912001/xr/lib/log"
)

// 远端信息
type Remote struct {
	conn     *net.TCPConn     //连接
	sendChan chan interface{} //发送管道

	cancelCXT      context.Context
	cancelFuncSend context.CancelFunc
}

func (p *Remote) Start(ctx context.Context) {
	p.cancelCXT, p.cancelFuncSend = context.WithCancel(ctx)
}

//是否连接
func (p *Remote) IsConn() bool {
	return nil != p.conn
}

func (p *Remote) Stop() {
	if p.cancelFuncSend != nil {
		p.cancelFuncSend()
	}
	if p.sendChan != nil {
		close(p.sendChan)
		p.sendChan = nil
	}
	if p.IsConn() {
		p.conn.Close()
		p.conn = nil
	}
}

//发送事件
type sendEvent struct {
	data   []byte  //数据
	remote *Remote //目标端
}

//处理发送事件
//当 conn 关闭, 该函数会引发 panic ...
func (p *Remote) onSendEvent(log *log.Log) {
	log.Trace("goroutine start.")

	defer func() {
		if err := recover(); err != nil {
			log.Warn(fmt.Sprintf("onSendEvent goroutine panic:%v", err))
		}
		log.Trace("goroutine done.")
	}()
	for {
		select {
		//case <-ctxWithCancel.Done():
		//	log.Warn(fmt.Sprintf("context onSendEvent goroutine done."))
		//	return
		case v, ok := <-p.sendChan:
			if !ok {
				log.Warn(fmt.Sprintf("sendChan is close, but not nil. ok value:%v", ok))
				return
			}

			switch v.(type) {
			case *sendEvent:
				vv, ok := v.(*sendEvent)
				if ok {
					var sum int
					for {
						//超时10微妙 conn.SetWriteDeadline(time.Now().Add(time.Microsecond * 10))
						n, err := vv.remote.conn.Write(vv.data[sum:])
						if nil != err {
							log.Warn(fmt.Sprintf("send data, cnt:%v, data:%v, err:%v", n, vv.data, err))
							break
						}
						sum += n
						if len(vv.data) == sum {
							break
						}
					}
				}
			default:
				log.Crit(fmt.Sprintf("the event type could not be found. event:%v", v))
			}

		}
	}
	//for v := range p.sendChan {
	//	switch v.(type) {
	//	case *sendEvent:
	//		vv, ok := v.(*sendEvent)
	//		if ok {
	//			var sum int
	//			for {
	//				//超时10微妙 conn.SetWriteDeadline(time.Now().Add(time.Microsecond * 10))
	//				n, err := vv.remote.conn.Write(vv.data[sum:])
	//				if nil != err {
	//					log.Warn(fmt.Sprintf("send data, cnt:%v, data:%v, err:%v", n, vv.data, err))
	//					break
	//				}
	//				sum += n
	//				if len(vv.data) == sum {
	//					break
	//				}
	//			}
	//		}
	//	default:
	//		log.Crit(fmt.Sprintf("the event type could not be found. event:%v", v))
	//	}
	//}
}
