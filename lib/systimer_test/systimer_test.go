package systimer_test

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/75912001/xr/lib/systimer"
)

/*
 go test -v -count=1
*/

//测试个数
var testTimerCnt int = 100000

//超时事件放置的channel
var eventChan chan interface{}

//完成的chan
var finishChan chan interface{}

var userMap map[*User]int

var waitCBDone sync.WaitGroup

var cbCnt int

func cb(arg interface{}) int {
	user := arg.(*User)
	_ = user

	cbCnt++
	waitCBDone.Done()
	return 0
}

type User struct {
	ID    int
	timer *time.Timer
}

func TestSysTimer(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	eventChan = make(chan interface{}, testTimerCnt*10)
	finishChan = make(chan interface{}, testTimerCnt)
	userMap = make(map[*User]int)
	cbCnt = 0

	var stm systimer.SysTimerMgr
	stm.Start(eventChan)

	for i := 1; i <= testTimerCnt; i++ {
		user := &User{}
		user.ID = i
		userMap[user] = user.ID

		timer := stm.AddSecond(cb, user, 0)
		user.timer = timer
	}

	for k, _ := range userMap {
		user := k
		if rand.Int31n(100) < 50 {
			if !systimer.Del(user.timer) {
				//log.Printf("del false ...")
			} else {
				testTimerCnt--
			}
		}
	}

	t.Logf("add timer done. testTimerCnt:%v", testTimerCnt)

	waitCBDone.Add(testTimerCnt)

	go func() {
		t.Logf("event goroutine start.")
		defer func() {
			if err := recover(); err != nil {
				t.Logf("event goroutine painc:%v", err)
			}
			t.Logf("event goroutine exit.")
		}()
		for v := range eventChan {
			switch v.(type) {
			case systimer.SysTimerParameter:
				v1, ok1 := v.(systimer.SysTimerParameter)
				if ok1 {
					v1.Fun(v1.Arg)
				}
			}
		}
	}()

	waitCBDone.Wait()

	t.Logf("call back func cnt:%v, testTimerCnt:%v", cbCnt, testTimerCnt)
}

//func timerFunc(parameters interface{}) func() {
//	return func() {
//		defer func() {
//			if err := recover(); err != nil {
//				fmt.Printf("timerFunc goroutine recover:%s, parameters:%s\n", err, parameters)
//			}
//		}()
//		func(data interface{}) int {
//			eventChan <- data.(int)
//			//fmt.Printf("%v\n", cb_cnt)
//			return 0
//		}(parameters)
//	}
//}
//
//func TestExample(t *testing.T) {
//	eventChan = make(chan interface{}, testTimerCnt*10)
//	for i := 1; i <= testTimerCnt; i++ {
//		f := timerFunc(i)
//		time.AfterFunc(0, f)
//	}
//	f := timerFunc(1)
//	pTimer := time.AfterFunc(0, f)
//	pTimer.stop()
//	for i := 1; i <= testTimerCnt; i++ {
//		<-eventChan
//	}
//}
