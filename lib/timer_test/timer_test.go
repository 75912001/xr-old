package timer_test

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/75912001/xr/lib/timer"
)

/*
 go test -v -count=1
*/
//扫描间隔(纳秒)
var scanSecondDuration time.Duration = 100000000      //100毫秒
var scanMillisecondDuration time.Duration = 100000000 //100毫秒

//最大定时时长
const MaxSecond int = 9
const MaxMilliSecond int = 1000

//测试个数
var testTimerCnt int = 10000

//真实测试个数
var realTestTimerCnt int

//完成的chan
var finishChan chan interface{} = make(chan interface{}, testTimerCnt)

//超时事件放置的channel
var eventChan chan interface{}

func cb(data interface{}) int {
	user := data.(*User)
	finishChan <- user
	return 0
}

type User struct {
	ID           int
	pSecond      *timer.Second
	pMilliSecond *timer.Millisecond
	tm           *timer.TimerMgr
}

func (p *User) AddSecond() {
	second := time.Now().Unix()
	p.pSecond = p.tm.AddSecond(cb, p, second+int64(rand.Intn(MaxSecond+1)))
	if rand.Int31n(100) < 20 {
		p.DelSecond()
		return
	}
	realTestTimerCnt++
}

func (p *User) DelSecond() {
	timer.DelSecond(p.pSecond)
	p.pSecond = nil
}

func TestSecond(t *testing.T) {
	realTestTimerCnt = 0
	eventChan = make(chan interface{}, testTimerCnt*10)
	var waitGroupGoroutineDone sync.WaitGroup

	var tm timer.TimerMgr
	tm.Start(context.Background(), scanSecondDuration, scanMillisecondDuration, eventChan)

	for i := 0; i < testTimerCnt; i++ {
		user := &User{
			ID:           i,
			pSecond:      nil,
			pMilliSecond: nil,
			tm:           &tm,
		}
		user.AddSecond()
	}
	t.Logf("AddSecond done.")

	waitGroupGoroutineDone.Add(1)
	go func() {
		t.Logf("timer second goroutine start.")
		defer func() {
			if err := recover(); err != nil {
				t.Logf("timer second goroutine painc:%v", err)
			}
			t.Logf("timer second goroutine exit.")
			waitGroupGoroutineDone.Done()
		}()
		for v := range eventChan {
			switch v.(type) {
			case *timer.Second:
				v1, ok1 := v.(*timer.Second)
				if ok1 {
					if v1.IsValid() {
						v1.Function(v1.Arg)
					}
				}
			}
		}
	}()

	waitCnt := realTestTimerCnt

	//等待所有timer结束
	for {
		user := <-finishChan
		pUser := user.(*User)
		_ = pUser
		waitCnt--
		if waitCnt <= 0 {
			break
		}
	}
	tm.Exit()
	close(eventChan)
	//等待goroutine结束
	waitGroupGoroutineDone.Wait()

	t.Logf("realTestTimerCnt:%v", realTestTimerCnt)
}

func (p *User) AddMillisecond() {
	n := time.Now()
	millisecond := n.UnixNano() / 1000000

	p.pMilliSecond = p.tm.AddMillisecond(cb, p, millisecond+int64(rand.Intn(MaxMilliSecond+1)))

	if rand.Int31n(100) < 20 {
		p.DelMillisecond()
		return
	}
	realTestTimerCnt++
}

func (p *User) DelMillisecond() {
	timer.DelMillisecond(p.pMilliSecond)
	p.pMilliSecond = nil
}

func TestMillisecond(t *testing.T) {
	realTestTimerCnt = 0
	eventChan = make(chan interface{}, testTimerCnt*10)
	var waitGroupGoroutineDone sync.WaitGroup

	var tm timer.TimerMgr
	tm.Start(context.Background(), scanSecondDuration, scanMillisecondDuration, eventChan)

	for i := 0; i < testTimerCnt; i++ {
		user := &User{
			ID:           i,
			pSecond:      nil,
			pMilliSecond: nil,
			tm:           &tm,
		}
		user.AddMillisecond()
	}
	t.Logf("AddMillisecond done.")

	waitGroupGoroutineDone.Add(1)
	go func() {
		t.Logf("timer Millisecond goroutine start.")
		defer func() {
			if err := recover(); err != nil {
				t.Logf("timer Millisecond goroutine painc:%v", err)
			}
			t.Logf("timer Millisecond goroutine exit.")
			waitGroupGoroutineDone.Done()
		}()
		for v := range eventChan {
			switch v.(type) {
			case *timer.Millisecond:
				v1, ok1 := v.(*timer.Millisecond)
				if ok1 {
					if v1.IsValid() {
						v1.Function(v1.Arg)
					}
				}
			}
		}
	}()

	waitCnt := realTestTimerCnt

	//等待所有timer结束
	for {
		user := <-finishChan
		pUser := user.(*User)
		_ = pUser
		waitCnt--
		if waitCnt <= 0 {
			break
		}
	}
	tm.Exit()
	close(eventChan)
	//等待goroutine结束
	waitGroupGoroutineDone.Wait()

	t.Logf("realTestTimerCnt:%v", realTestTimerCnt)
}
