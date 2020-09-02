package timer

import (
	//"fmt"
	"testing"
	"time"
)

var allCnt int64 = 10000000 //1000w
var iChan chan int64 = make(chan int64, allCnt)
var eachCnt int64 = 100000 //每次处理多少个打印一次日志 10w

func cb(data interface{}) int {
	cb_cnt := data.(int64)
//	if 0 == cb_cnt%eachCnt {
//		fmt.Println(cb_cnt)
//	}
	iChan <- cb_cnt
	return 0
}

func TestTimerSecond2(t *testing.T) {
	var tm TimerMgr
	second := time.Now().Unix()
	var c chan interface{} = make(chan interface{}, allCnt)
	tm.Start(100, c)

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

	for i := int64(1); i <= allCnt; i++ {
		tm.AddSecond(cb, i, second)
	}

	for i := int64(1); i <= allCnt; i++ {
		<-iChan
	}
}

func addCB(data interface{}) int {
	tm := data.(*TimerMgr)
	n := time.Now()
	second := n.Unix()
	for i := int64(1); i <= allCnt; i++ {
		t := tm.AddSecond(cb, i, second)
		if i%2 == 0 {
			tm.DelSecond(t)
			tm.AddSecond(cb, i, second + 10)
		}
	}
	return 0
}

func TestTimerSecond3(t *testing.T) {
	var tm TimerMgr
	second := time.Now().Unix()
	var c chan interface{} = make(chan interface{}, allCnt)
	tm.Start(100, c)

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

	t1 := tm.AddSecond(addCB, &tm, second+10)
	tm.DelSecond(t1)
	tm.AddSecond(addCB, &tm, second)

	for i := int64(1); i <= allCnt; i++ {
		<-iChan
	}
	tm.Exit()
}

func cb2(data interface{}) int {
	cb_cnt := data.(int64)

	//if 0 == cb_cnt%eachCnt {
	//	fmt.Println(cb_cnt)
	//}
	iChan <- cb_cnt
	return 0
}
func addCB2(data interface{}) int {
	tm := data.(*TimerMgr)
	n := time.Now()
//	second := n.Unix()
	millisecond := n.UnixNano() / 1000000
//	fmt.Println("begin:", second)
	for i := int64(1); i <= allCnt; i++ {
		tm.AddMillisecond(cb2, i, millisecond)
	}
	//n = time.Now()
//	second = n.Unix()
//	fmt.Println("end:", second)
	return 0
}
func TestTimerMillisecond(t *testing.T) {
	var tm TimerMgr
	n := time.Now()
	millisecond := n.UnixNano() / 1000000

	var c chan interface{} = make(chan interface{}, allCnt)

	tm.Start(100, c)

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
}

