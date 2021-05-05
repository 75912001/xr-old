package tcp

import (
	"fmt"
	"net"
)

// 远端信息
type Remote struct {
	conn *net.TCPConn //连接
	//connLock sync.RWMutex     //连接锁 //TODO [discuss] 是否需要该锁
	sendChan chan interface{} //发送管道
}

//是否连接
func (p *Remote) IsConn() bool {
	return nil != p.conn
}

func (p *Remote) Close() {
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
	data []byte  //数据
	dst  *Remote //目标端
}

//处理发送事件
//当 conn 关闭, 该函数会引发 panic ...
func (p *Remote) onSendEvent() {
	GLog.Trace("goroutine start.")
	defer func() {
		if err := recover(); err != nil {
			GLog.Warn(fmt.Sprintf("onSendEvent goroutine panic:%v", err))
		}
		GLog.Trace("goroutine done.")
	}()
	for v := range p.sendChan {
		switch v.(type) {
		case *sendEvent:
			vv, ok := v.(*sendEvent)
			if ok {
				var sum int
				for {
					//超时10微妙 conn.SetWriteDeadline(time.Now().Add(time.Microsecond * 10))
					n, err := vv.dst.conn.Write(vv.data[sum:])
					if nil != err {
						GLog.Warn(fmt.Sprintf("send data, cnt:%v, data:%v, err:%v", n, vv.data, err))
						break
					}
					sum += n
					if len(vv.data) == sum {
						break
					}
				}
			}
		default:
			GLog.Crit(fmt.Sprintf("the event type could not be found. event:%v", v))
		}
	}
}

//接收数据
func (p *Server) onRecvEventChan(remote *Remote, recvPacketMaxLen uint32, onParseProtoHead OnParseProtoHeadType) {
	GLog.Trace("goroutine start.")

	defer func() {
		if err := recover(); err != nil {
			GLog.Warn(fmt.Sprintf("goroutine panic:%v", err))
		} else { //断开链接
			p.tcpChan <- &DisConnEventServer{
				Server: p,
				Remote: remote,
			}
		}
		GLog.Trace("goroutine done.")
	}()

	//TODO [improvement] 环形缓冲
	buf := make([]byte, recvPacketMaxLen)

	var readIndex int
	for {
	LoopRead:
		readNum, err := remote.conn.Read(buf[readIndex:])
		if nil != err {
			GLog.Error(fmt.Sprintf("Conn.Read, read num:%v, err:%v", readNum, err))
			return
		}
		readIndex += readNum
		for {
			packetLength := onParseProtoHead(buf, readIndex)
			if 0 == packetLength {
				goto LoopRead
			}

			if -1 == packetLength {
				GLog.Crit(fmt.Sprintf("packetLength:%v, readIndex:%v, Data:%v", packetLength, readIndex, buf))
				return
			}

			//接受数据
			pes := &PacketEventServer{
				Server: p,
				Data:   make([]byte, packetLength),
				Remote: remote,
			}
			copy(pes.Data, buf[:packetLength])
			p.tcpChan <- pes

			copy(buf, buf[packetLength:readIndex])
			readIndex -= packetLength

			if 0 == readIndex {
				goto LoopRead
			}
		}
	}
}
