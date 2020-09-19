package timer

import (
	"container/list"
	"context"
	"fmt"
	"sync"
	"time"
)

//TimerMgr 定时器管理器
type TimerMgr struct {
	secondVec        [eTimerVecSize]*tvecRoot //秒,数据
	millisecondList  *list.List               //毫秒,数据
	timerOutChan     chan<- interface{}       //超时的*TimerSecond/*TimerMillisecond都会放入其中
	secondMutex      sync.Mutex
	milliSecondMutex sync.Mutex
	cancelFunc       context.CancelFunc
}

//OnTimerFun 回调定时器函数(使用协程回调)
type OnTimerFun func(data interface{}) int

//Start millisecond:毫秒间隔(如50,则每50毫秒扫描一次毫秒定时器)
//timerOutChan 是超时事件放置的channel,由外部传入
func (p *TimerMgr) Start(ctx context.Context, millisecond int64, timerOutChan chan<- interface{}) {
	for idx := range p.secondVec {
		p.secondVec[idx] = &tvecRoot{}
		p.secondVec[idx].init()
	}
	p.millisecondList = list.New()
	p.timerOutChan = timerOutChan

	ctxWithCancel, cancelFunc := context.WithCancel(ctx)
	p.cancelFunc = cancelFunc
	//每秒更新
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("timer second goroutine done.")
				return
			case <-time.After(time.Second):
				time.Sleep(time.Second)
				p.secondMutex.Lock()
				p.scanSecond()
				p.secondMutex.Unlock()
			}
		}
	}(ctxWithCancel)

	//每millisecond个毫秒更新
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("timer millisecond goroutine done.")
				return
			case <-time.After(time.Duration(millisecond) * time.Millisecond):
				p.milliSecondMutex.Lock()
				p.scanMillisecond()
				p.milliSecondMutex.Unlock()
			}
		}
	}(ctxWithCancel)
}

func (p *TimerMgr) Exit() {
	p.cancelFunc()
}

//AddSecond 添加秒级定时器
func (p *TimerMgr) AddSecond(cb OnTimerFun, arg interface{}, expire int64) (t *TimerSecond) {
	p.secondMutex.Lock()
	defer func() {
		p.secondMutex.Unlock()
	}()

	return p.addSecond(cb, arg, expire)
}

//DelSecond 删除秒级定时器
func (p *TimerMgr) DelSecond(t *TimerSecond) {
	t.valid = false
}

//AddMillisecond 添加毫秒级定时器
func (p *TimerMgr) AddMillisecond(cb OnTimerFun, arg interface{}, expireMillisecond int64) (t *TimerMillisecond) {
	t = &TimerMillisecond{
		expireMillisecond,
		arg,
		cb,
		true,
	}

	p.milliSecondMutex.Lock()
	defer func() {
		p.milliSecondMutex.Unlock()
	}()

	p.millisecondList.PushBack(t)
	return
}

//DelMillisecond 删除毫秒级定时器
func (p *TimerMgr) DelMillisecond(t *TimerMillisecond) {
	t.valid = false
}