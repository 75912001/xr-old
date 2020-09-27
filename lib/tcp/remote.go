package tcp

import (
	"fmt"
	"net"
	"sync"

	"github.com/75912001/xr/lib/log"
)

// remote 远端连接信息
type remote struct {
	conn             *net.TCPConn //连接
	lock             sync.RWMutex
	sendChan         chan interface{} //需要发送
	sendEventChanCnt uint32           //发送chan大小
}

//连接是否有效
func (p *remote) isConnect() bool {
	return nil != p.conn
}

//发送数据事件channel
type sendEventChan struct {
	buf []byte
	dst *remote
}

//处理待发的数据
func onEventSend(sendChan chan interface{}) {
	log.GLog.Trace("onEventSend goroutine start.")
	for v := range sendChan {
		switch v.(type) {
		case *sendEventChan:
			vv, ok := v.(*sendEventChan)
			if ok {
				vv.dst.lock.RLock()
				if !vv.dst.isConnect() {
					vv.dst.lock.RUnlock()
					log.GLog.Error(fmt.Sprintf("remote is valid, Buf:%v", vv.buf))
					continue
				}

				var sum int
				for {
					n, err := vv.dst.conn.Write(vv.buf[sum:])
					if nil != err {
						log.GLog.Error(fmt.Sprintf("send chan, cnt:%v, Buf:%v, err:%v", n, vv.buf, err))
						break
					}
					sum += n
					if len(vv.buf) == sum {
						break
					}
				}
				vv.dst.lock.RUnlock()
			}
		default:
			log.GLog.Crit(fmt.Sprintf("no find send event:%v", v))
		}
	}
	log.GLog.Trace("onEventSend goroutine done.")
}
