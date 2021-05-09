package timer

//
//优先级:加入顺序,到期
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
	timerOutChan     chan<- interface{}       //超时的*Second/*Millisecond都会放入其中
	secondMutex      sync.Mutex
	milliSecondMutex sync.Mutex
	cancelFunc       context.CancelFunc
}

//OnTimerFun 回调定时器函数(使用协程回调)
type OnTimerFun func(data interface{}) int

//Start scanSecondDuration:扫描秒级定时器, 纳秒间隔(如100000000,则每100毫秒扫描一次秒定时器)
//Start scanMillisecondDuration:扫描毫秒级定时器, 纳秒间隔(如100000000,则每100毫秒扫描一次毫秒定时器)
//timerOutChan 是超时事件放置的channel,由外部传入(处理定时器相关数据,必须与该timerOutChan线性处理.如:在同一个goroutine select中处理数据.)
func (p *TimerMgr) Start(ctx context.Context, scanSecondDuration time.Duration, scanMillisecondDuration time.Duration,
	timerOutChan chan<- interface{}) {
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
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("timer second goroutine panic:%v\n", err)
			}
		}()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context timer second goroutine done.")
				return
			case <-time.After(scanSecondDuration * time.Nanosecond):
				p.secondMutex.Lock()
				p.scanSecond()
				p.secondMutex.Unlock()
			}
		}
	}(ctxWithCancel)

	//每millisecond个毫秒更新
	go func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("timer millisecond goroutine painc:%v\n", err)
			}
		}()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context timer millisecond goroutine done.")
				return
			case <-time.After(scanMillisecondDuration * time.Nanosecond):
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

// AddSecond 添加秒级定时器
func (p *TimerMgr) AddSecond(cb OnTimerFun, arg interface{}, expire int64) (t *Second) {
	p.secondMutex.Lock()
	defer func() {
		p.secondMutex.Unlock()
	}()

	return p.addSecond(cb, arg, expire)
}

// DelSecond 删除秒级定时器
func DelSecond(t *Second) {
	t.Millisecond.inValid()
}

//AddMillisecond 添加毫秒级定时器
func (p *TimerMgr) AddMillisecond(cb OnTimerFun, arg interface{}, expireMillisecond int64) (t *Millisecond) {
	t = &Millisecond{
		Arg:      arg,
		Function: cb,
		expire:   expireMillisecond,
		valid:    true,
	}

	p.milliSecondMutex.Lock()
	defer func() {
		p.milliSecondMutex.Unlock()
	}()

	p.millisecondList.PushBack(t)
	return
}

//DelMillisecond 删除毫秒级定时器
func DelMillisecond(t *Millisecond) {
	t.inValid()
}
