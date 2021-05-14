package tcp

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"
)

type addEventDisConn2EventChanFunc func()
type addEventPacket2EventChanFunc func(data []byte, packetLength int)

// 远端信息
type Remote struct {
	conn                      *net.TCPConn     //连接
	sendChan                  chan interface{} //发送管道
	cancelFunc                context.CancelFunc
	waitGroupGoroutineDone    sync.WaitGroup
	addEventDisConn2EventChan addEventDisConn2EventChanFunc
	addEventPacket2EventChan  addEventPacket2EventChanFunc
}

func (p *Remote) start(conn *net.TCPConn,
	rwBuffLen uint32, sendChanCapacity uint32, recvPacketMaxLen uint32, onParseProtoHead OnParseProtoHeadFunc) {

	p.conn = conn
	p.conn.SetNoDelay(true)
	p.conn.SetReadBuffer(int(rwBuffLen))
	p.conn.SetWriteBuffer(int(rwBuffLen))

	p.sendChan = make(chan interface{}, sendChanCapacity)

	ctx := context.Background()
	ctxWithCancel, cancelFunc := context.WithCancel(ctx)
	p.cancelFunc = cancelFunc

	p.waitGroupGoroutineDone.Add(2)

	go p.onSendEvent(ctxWithCancel)

	go p.onRecvEvent(recvPacketMaxLen, onParseProtoHead)
}

//是否连接
func (p *Remote) IsConn() bool {
	return nil != p.conn
}

//发送数据(data 数据不可修改)(必须在处理EventChan事件中调用)
func (p *Remote) Send(data []byte) (err error) {
	if !p.IsConn() {
		return errors.New("[ERROR]link disconnect.")
	}
	p.sendChan <- &sendEvent{
		data:   data,
		remote: p,
	}
	return
}

func (p *Remote) stop() {
	if p.sendChan != nil {
		close(p.sendChan)
	}

	if p.IsConn() {
		p.conn.Close()
	}

	if p.cancelFunc != nil {
		p.cancelFunc()
		p.waitGroupGoroutineDone.Wait()
	}

	p.cancelFunc = nil
	p.conn = nil
	p.sendChan = nil
}

//发送事件
type sendEvent struct {
	remote *Remote //目标端
	data   []byte  //数据
}

//处理发送事件
//当 conn 关闭, 该函数会引发 panic ...
func (p *Remote) onSendEvent(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("onSendEvent goroutine panic:%v", err)
		}
		p.waitGroupGoroutineDone.Done()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case v1, ok := <-p.sendChan:
			if !ok {
				//log.Printf("sendChan is close, but not nil. ok value:%v", ok)
				return
			}
			switch v1.(type) {
			case *sendEvent:
				v2, ok2 := v1.(*sendEvent)
				if ok2 {
					var sum int
					for {
						//超时10微妙 conn.SetWriteDeadline(time.Now().Add(time.Microsecond * 10))
						n, err := v2.remote.conn.Write(v2.data[sum:])
						if nil != err {
							log.Printf("send data, cnt:%v, data:%v, err:%v", n, v2.data, err)
							return
						}
						sum += n
						if len(v2.data) == sum {
							break
						}
					}
				}
			default:
				log.Printf("the event type could not be found. event:%v", v1)
			}
		}
	}
}

//接收数据
func (p *Remote) onRecvEvent(recvPacketMaxLen uint32, onParseProtoHead OnParseProtoHeadFunc) {
	defer func() { //断开链接
		if err := recover(); err != nil {
			log.Printf("onRecvEvent goroutine panic:%v", err)
		} else { //断开链接
			p.addEventDisConn2EventChan()
		}
		p.waitGroupGoroutineDone.Done()
	}()

	//TODO [improvement] 环形缓冲
	buf := make([]byte, recvPacketMaxLen)

	var readIndex int
	for {
	LoopRead:
		readNum, err := p.conn.Read(buf[readIndex:])
		if nil != err {
			//			log.Printf("Conn.Read, read num:%v, err:%v", readNum, err)
			return
		}
		readIndex += readNum
		for {
			packetLength := onParseProtoHead(buf, readIndex)
			if 0 == packetLength {
				goto LoopRead
			}

			if -1 == packetLength {
				log.Printf("packetLength:%v, readIndex:%v, Data:%v", packetLength, readIndex, buf)
				return
			}
			p.addEventPacket2EventChan(buf, packetLength)

			copy(buf, buf[packetLength:readIndex])
			readIndex -= packetLength

			if 0 == readIndex {
				goto LoopRead
			}
		}
	}
}
