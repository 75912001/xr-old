package tcp

import (
	"fmt"
	"net"
	"sync"

	"github.com/75912001/xr/lib/log"
)

// 远端信息
type remote struct {
	conn     *net.TCPConn     //连接
	connLock sync.RWMutex     //连接锁
	sendChan chan interface{} //发送管道
}

//是否连接
func (p *remote) isConn() bool {
	return nil != p.conn
}

//发送事件
type sendEvent struct {
	data []byte  //数据
	dst  *remote //目标端
}

//处理发送事件
func (p *remote) onSendEvent() {
	log.GLog.Trace("goroutine start.")
	for v := range p.sendChan {
		switch v.(type) {
		case *sendEvent:
			vv, ok := v.(*sendEvent)
			if ok {
				vv.dst.connLock.RLock()
				if !vv.dst.isConn() {
					vv.dst.connLock.RUnlock()
					log.GLog.Warn(fmt.Sprintf("remote is disconnect, data:%v", vv.data))
					continue
				}

				var sum int
				for {
					n, err := vv.dst.conn.Write(vv.data[sum:])
					if nil != err {
						log.GLog.Warn(fmt.Sprintf("send data, cnt:%v, data:%v, err:%v", n, vv.data, err))
						break
					}
					sum += n
					if len(vv.data) == sum {
						break
					}
				}
				vv.dst.connLock.RUnlock()
			}
		default:
			log.GLog.Crit(fmt.Sprintf("the event type could not be found. event:%v", v))
		}
	}
	log.GLog.Trace("goroutine done.")
}
