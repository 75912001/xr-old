package timer

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

//TimerMgr 定时器管理器
type TimerMgr struct {
	secondVec        [eTimerVecSize]*tvecRoot //秒,数据
	millisecondList  *list.List               //毫秒,数据
	timerOutChan     chan<- interface{}       //超时的*TimerSecond/*TimerMillisecond都会放入其中
	exit             bool                     //退出(true:退出,该管理器退出)
	secondMutex      sync.Mutex
	milliSecondMutex sync.Mutex
}

//OnTimerFun 回调定时器函数(使用协程回调)
type OnTimerFun func(data interface{}) int

//Start millisecond:毫秒间隔(如50,则每50毫秒扫描一次毫秒定时器)
//timerOutChan 是超时事件放置的channel,由外部传入
func (p *TimerMgr) Start(millisecond int64, timerOutChan chan<- interface{}) {
	for idx := range p.secondVec {
		p.secondVec[idx] = &tvecRoot{}
		p.secondVec[idx].init()
	}
	p.millisecondList = list.New()
	p.timerOutChan = timerOutChan

	//每秒更新
	go func() {
		for !p.exit{
			time.Sleep(time.Second)

			p.secondMutex.Lock()

			p.scanSecond()

			p.secondMutex.Unlock()
		}
		fmt.Println("second timer go fun exit")
	}()
	//每millisecond个毫秒更新
	go func() {
		for !p.exit{
			time.Sleep(time.Duration(millisecond) * time.Millisecond)

			p.milliSecondMutex.Lock()

			p.scanMillisecond()

			p.milliSecondMutex.Unlock()
		}
	}()
}

func (p *TimerMgr) Exit() {
	p.exit = true
}

//AddSecond 添加秒级定时器
func (p *TimerMgr) AddSecond(cb OnTimerFun, arg interface{}, expire int64) (t *TimerSecond) {
	p.secondMutex.Lock()
	defer func() {
		p.secondMutex.Unlock()
	}()

	return p.addSecond(cb, arg, expire, nil)
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
