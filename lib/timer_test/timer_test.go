package timer_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

package timer

import (
"context"
"fmt"
"math/rand"
"testing"
"time"

"github.com/75912001/xr/lib/util"
)

/*
 go test -v -count=1
*/
//扫描间隔(纳秒)
var scanSecondDuration time.Duration = 100000000//100毫秒
var scanMillisecondDuration time.Duration = 100000000//100毫秒

var testTimerCnt uint64 = 10000
var cbChan chan interface{} = make(chan interface{}, testTimerCnt)

//超时事件放置的channel
var eventChan chan interface{} = make(chan interface{}, testTimerCnt*10)

func cb(data interface{}) int {
	cbChan <- data.(*User)
	return 0
}

type User struct {
	ID           uint64
	pSecond      *Second
	pMilliSecond *Millisecond
	tm           *TimerMgr
}

func (p *User) AddSecond() bool {
	second := time.Now().Unix()
	p.pSecond = p.tm.AddSecond(cb, p, second+rand.Int63n(10))
	if rand.Int31n(100) < 20 {
		p.DelSecond()
		return false
	}
	return true
}

func (p *User) DelSecond() {
	DelSecond(p.pSecond)
}

func (p *User) AddMillisecond() bool {
	n := time.Now()
	millisecond := n.UnixNano() / 1000000

	p.pMilliSecond = p.tm.AddMillisecond(cb, p, millisecond+rand.Int63n(1000))

	if rand.Int31n(100) < 20 {
		p.DelMillisecond()
		return false
	}
	return true
}

func (p *User) DelMillisecond() {
	DelMillisecond(p.pMilliSecond)
}

func TestSecond(t *testing.T) {
	var trigger int = 0
	eventChan <- &trigger

	var waitCnt = testTimerCnt
	var tm TimerMgr
	tm.Start(context.Background(), scanSecondDuration, scanMillisecondDuration, eventChan)

	ctxWithCancel, cancelFunc := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		fmt.Println("timer second goroutine start.")
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("timer second goroutine painc:%v\n", err)
			}
			fmt.Println("timer second goroutine exit.")
		}()
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("context %v goroutine done.\n", util.GetFuncName())
				return
			case v, ok := <-eventChan:
				if ok {
					switch v.(type) {
					case *int:
						for i := uint64(0); i < testTimerCnt; i++ {
							user := &User{
								ID:           uint64(i),
								pSecond:      nil,
								pMilliSecond: nil,
								tm:           &tm,
							}
							if !user.AddSecond() {
								waitCnt--
							}
						}
						fmt.Println("AddSecond done.")
					case *Second:
						v1, ok1 := v.(*Second)
						if ok1 {
							if v1.IsValid() {
								v1.Function(v1.Arg)
							}
						}
					}
				}
			default:
				time.Sleep(time.Millisecond)
				//fmt.Println("sleep 1 second ...")
			}
		}
	}(ctxWithCancel)

	{ //仅用于测试 ...
		//等待所有timer结束
		for i := uint64(0); i < waitCnt; i++ {
			user := <-cbChan
			pUser := user.(*User)
			_ = pUser
		}
	}

	tm.Exit()
	cancelFunc()

	{ //仅用于测试 ...
		//	time.Sleep(time.Second * 1)
	}
	fmt.Printf("all timer cnt:%v\n", waitCnt)
}

func TestMillisecond(t *testing.T) {
	var trigger int = 0
	eventChan <- &trigger

	var waitCnt = testTimerCnt
	var tm TimerMgr
	tm.Start(context.Background(), scanSecondDuration, scanMillisecondDuration, eventChan)

	ctxWithCancel, cancelFunc := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		fmt.Println("timer millisecond goroutine start.")
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("timer millisecond goroutine painc:%v\n", err)
			}
			fmt.Println("timer millisecond goroutine exit.")
		}()
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("%v goroutine done.\n", util.GetFuncName())
				return
			case v, ok := <-eventChan:
				if ok {
					switch v.(type) {
					case *int:
						for i := uint64(0); i < testTimerCnt; i++ {
							user := &User{
								ID:           uint64(i),
								pSecond:      nil,
								pMilliSecond: nil,
								tm:           &tm,
							}
							if !user.AddMillisecond() {
								waitCnt--
							}
						}
						//fmt.Printf("AddMillisecond done.\n")
					case *Millisecond:
						v1, ok1 := v.(*Millisecond)
						if ok1 {
							if v1.IsValid() {
								v1.Function(v1.Arg)
							}
						}
					}
				}
			default:
				time.Sleep(time.Millisecond)
				//fmt.Printf("sleep 1 second ...\n")
			}
		}
	}(ctxWithCancel)

	{ //仅用于测试 ...
		//等待所有timer结束
		for i := uint64(0); i < waitCnt; i++ {
			user := <-cbChan
			pUser := user.(*User)
			_ = pUser
		}
	}

	tm.Exit()
	cancelFunc()

	{ //仅用于测试 ...
		//time.Sleep(time.Second * 1)
	}
	fmt.Printf("all timer cnt:%v\n", waitCnt)
}
