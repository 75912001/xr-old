package timer

import (
	"context"
	"math/rand"

	//"fmt"
	"testing"
	"time"
)

/*
 go test -v -count=1
*/
var allCnt int64 = 10000000 //1000w
var iChan chan int64 = make(chan int64, allCnt)
var eachCnt int64 = 100000 //每次处理多少个打印一次日志 10w

func cb(data interface{}) int {
	cb_cnt := data.(int64)
	iChan <- cb_cnt
	return 0
}

func TestTimerSecond(t *testing.T) {
	var tm TimerMgr

	var c chan interface{} = make(chan interface{}, allCnt)
	tm.Start(context.Background(), 100, c)

	var outChan <-chan interface{}
	outChan = c
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("timer second goroutine painc:%v\n", err)
			}
		}()
		for v := range outChan {
			switch v.(type) {
			case *TimerSecond:
				tv, ok := v.(*TimerSecond)
				if ok {
					tv.Function(tv.Arg)
				}
			}
		}
	}()

	second := time.Now().Unix()
	for i := int64(1); i <= allCnt; i++ {
		tm.AddSecond(cb, i, second+rand.Int63n(10))
	}

	for i := int64(1); i <= allCnt; i++ {
		<-iChan
	}
	tm.Exit()
	close(c)
	time.Sleep(time.Second)
}

func addCB(data interface{}) int {
	tm := data.(*TimerMgr)
	n := time.Now()
	second := n.Unix()
	for i := int64(1); i <= allCnt; i++ {
		t := tm.AddSecond(cb, i, second+rand.Int63n(10))
		if i%2 == 0 {
			tm.DelSecond(t)
			tm.AddSecond(cb, i, second+rand.Int63n(10))
		}
	}
	return 0
}

func TestTimerSecondAddCBDelCB(t *testing.T) {
	var tm TimerMgr
	second := time.Now().Unix()
	var c chan interface{} = make(chan interface{}, allCnt)
	tm.Start(context.Background(), 100, c)

	var outChan <-chan interface{}
	outChan = c
	go func() {
		for v := range outChan {
			switch v.(type) {
			case *TimerSecond:
				tv, ok := v.(*TimerSecond)
				if ok {
					tv.Function(tv.Arg)
				}
			}
		}
	}()

	t1 := tm.AddSecond(addCB, &tm, second)
	tm.DelSecond(t1)
	tm.AddSecond(addCB, &tm, second)

	for i := int64(1); i <= allCnt; i++ {
		<-iChan
	}
	tm.Exit()
	close(c)
	time.Sleep(time.Second)
}

func cb2(data interface{}) int {
	cb_cnt := data.(int64)
	iChan <- cb_cnt
	return 0
}
func addCB2(data interface{}) int {
	tm := data.(*TimerMgr)
	n := time.Now()
	millisecond := n.UnixNano() / 1000000
	for i := int64(1); i <= allCnt; i++ {
		tm.AddMillisecond(cb2, i, millisecond)
	}

	return 0
}
func TestTimerMillisecond(t *testing.T) {
	var tm TimerMgr
	n := time.Now()
	millisecond := n.UnixNano() / 1000000

	var c chan interface{} = make(chan interface{}, allCnt)

	tm.Start(context.Background(), 100, c)

	tm.AddMillisecond(addCB2, &tm, millisecond)
	go func() {
		for v := range c {
			switch v.(type) {
			case *TimerMillisecond:
				tv, ok := v.(*TimerMillisecond)
				if ok {
					tv.Function(tv.Arg)
				}
			}
		}
	}()
	for i := int64(1); i <= allCnt; i++ {
		<-iChan
	}
	tm.Exit()
	close(c)
	time.Sleep(time.Second)
}
