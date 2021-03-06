package timer

import (
	"context"
	"fmt"
	"math/rand"

	//"fmt"
	"testing"
	"time"
)

/*
 go test -v -count=1
*/
//扫描间隔(毫秒)
var scanIntervalMillisecond time.Duration = 100

//超时事件放置的channel
var timerOutChan chan interface{} = make(chan interface{}, 10000)
var timerCnt int64 = 10000
var cbChan chan int64 = make(chan int64, timerCnt)

func cb(data interface{}) int {
	cbChan <- data.(int64)
	//fmt.Printf("%v\n", cb_cnt)
	return 0
}

func TestExampleTimerSecond(t *testing.T) {
	var tm TimerMgr

	tm.Start(context.Background(), scanIntervalMillisecond, timerOutChan)

	var outChan <-chan interface{}
	outChan = timerOutChan
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("second timerout goroutine recover:%s\n", err)
			}
		}()
		for v := range outChan {
			switch v.(type) {
			case *Second:
				tv, ok := v.(*Second)
				if ok {
					tv.Function(tv.Arg)
				}
			}
		}
		fmt.Printf("TestExampleTimerSecond goroutine done.\n")
	}()

	second := time.Now().Unix()
	for i := int64(1); i <= timerCnt; i++ {
		//tm.AddSecond(cb, i, second+rand.Int63n(10))
		tm.AddSecond(cb, i, second)
	}

	pSecond := tm.AddSecond(cb, 1, second)
	DelSecond(pSecond)

	//等待所有timer结束
	for i := int64(1); i <= timerCnt; i++ {
		<-cbChan
	}
	tm.Exit()
	close(timerOutChan)
	close(cbChan)
}

func TestExampleTimerMillisecond(t *testing.T) {
	var tm TimerMgr

	tm.Start(context.Background(), scanIntervalMillisecond, timerOutChan)

	var outChan <-chan interface{}
	outChan = timerOutChan

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("millisecond timerout goroutine recover:%s\n", err)
			}
		}()
		for v := range outChan {
			switch v.(type) {
			case *Millisecond:
				tv, ok := v.(*Millisecond)
				if ok {
					tv.Function(tv.Arg)
				}
			}
		}
		fmt.Printf("TestExampleTimerMillisecond goroutine done.\n")
	}()

	n := time.Now()
	millisecond := n.UnixNano() / 1000000

	for i := int64(1); i <= timerCnt; i++ {
		tm.AddMillisecond(cb, i, millisecond)
	}

	pMillisecond := tm.AddMillisecond(cb, 1, millisecond)
	DelMillisecond(pMillisecond)

	//等待所有timer结束
	for i := int64(1); i <= timerCnt; i++ {
		<-cbChan
	}
	tm.Exit()
	close(timerOutChan)
	close(cbChan)
}

func cbTimerSecondAddDel(data interface{}) int {
	tm := data.(*TimerMgr)
	n := time.Now()
	second := n.Unix()
	for i := int64(1); i <= timerCnt; i++ {
		t := tm.AddSecond(cb, i, second+rand.Int63n(10))
		if i%2 == 0 {
			DelSecond(t)
			tm.AddSecond(cb, i, second+rand.Int63n(10))
		}
	}
	return 0
}

func TestTimerSecondAddDel(t *testing.T) {
	var tm TimerMgr
	second := time.Now().Unix()

	tm.Start(context.Background(), scanIntervalMillisecond, timerOutChan)

	var outChan <-chan interface{}
	outChan = timerOutChan
	go func() {
		for v := range outChan {
			switch v.(type) {
			case *Second:
				tv, ok := v.(*Second)
				if ok {
					tv.Function(tv.Arg)
				}
			}
		}
	}()

	tm.AddSecond(cbTimerSecondAddDel, &tm, second)

	for i := int64(1); i <= timerCnt; i++ {
		<-cbChan
	}
	tm.Exit()
	close(timerOutChan)
	close(cbChan)
}

///////////////////////////////////////////////////////////////////////
func cbTimerMillisecondAddDel(data interface{}) int {
	tm := data.(*TimerMgr)
	n := time.Now()
	millisecond := n.UnixNano() / 1000000
	for i := int64(1); i <= timerCnt; i++ {
		t := tm.AddMillisecond(cb, i, millisecond)
		if i%2 == 0 {
			DelMillisecond(t)
			tm.AddMillisecond(cb, i, millisecond)
		}
	}

	return 0
}
func TestTimerMillisecondAddDel(t *testing.T) {
	var tm TimerMgr
	n := time.Now()
	millisecond := n.UnixNano() / 1000000

	tm.Start(context.Background(), scanIntervalMillisecond, timerOutChan)

	tm.AddMillisecond(cbTimerMillisecondAddDel, &tm, millisecond)
	go func() {
		for v := range timerOutChan {
			switch v.(type) {
			case *Millisecond:
				tv, ok := v.(*Millisecond)
				if ok {
					tv.Function(tv.Arg)
				}
			}
		}
	}()
	for i := int64(1); i <= timerCnt; i++ {
		<-cbChan
	}
	tm.Exit()
	close(timerOutChan)
	close(cbChan)
}
