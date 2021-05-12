package timer_test

import (
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
const MaxSecond int = 0
const MaxMilliSecond int = 0

//测试个数
var testTimerCnt int = 100000

//真实测试个数
var realTestTimerCnt int

//超时事件放置的channel
var eventChan chan interface{}

var waitCBDone sync.WaitGroup

func cb(arg interface{}) int {
	user := arg.(*User)
	_ = user
	waitCBDone.Done()
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
	if rand.Int31n(100) < 50 {
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
	rand.Seed(time.Now().UnixNano())
	realTestTimerCnt = 0
	eventChan = make(chan interface{}, testTimerCnt*10)
	var waitGroupGoroutineDone sync.WaitGroup

	var tm timer.TimerMgr
	tm.Start(scanSecondDuration, scanMillisecondDuration, eventChan)

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
	waitCBDone.Add(realTestTimerCnt)

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

	waitCBDone.Wait()

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

	if rand.Int31n(100) < 50 {
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
	rand.Seed(time.Now().UnixNano())
	realTestTimerCnt = 0
	eventChan = make(chan interface{}, testTimerCnt*10)
	var waitGroupGoroutineDone sync.WaitGroup

	var tm timer.TimerMgr
	tm.Start(scanSecondDuration, scanMillisecondDuration, eventChan)

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
	waitCBDone.Add(realTestTimerCnt)

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

	waitCBDone.Wait()

	tm.Exit()
	close(eventChan)
	//等待goroutine结束
	waitGroupGoroutineDone.Wait()

	t.Logf("realTestTimerCnt:%v", realTestTimerCnt)
}
