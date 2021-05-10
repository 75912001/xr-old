package timer

//
//优先级:加入顺序,到期
import (
	"container/list"
	"context"
	"log"
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
	waitGroup        sync.WaitGroup
}

//OnTimerFun 回调定时器函数(使用协程回调)
type OnTimerFun func(arg interface{}) int

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

	p.waitGroup.Add(2)

	ctxWithCancel, cancelFunc := context.WithCancel(ctx)
	p.cancelFunc = cancelFunc
	//每秒更新
	go func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("timer second goroutine panic:%v", err)
			}
			p.waitGroup.Done()
		}()
		for {
			select {
			case <-ctx.Done():
				log.Printf("context timer second goroutine done.")
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
				log.Printf("timer millisecond goroutine painc:%v", err)
			}
			p.waitGroup.Done()
		}()
		for {
			select {
			case <-ctx.Done():
				log.Printf("context timer millisecond goroutine done.")
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
	//等待 second, milliSecond goroutine退出.
	p.waitGroup.Wait()
}
