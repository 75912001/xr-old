package systimer

import (
	"time"
)

type OnTimerFun func(arg interface{}) int

type SysTimerMgr struct {
	timerOutChan chan<- interface{} //超时的都会放入其中
}

func (p *SysTimerMgr) Start(timerOutChan chan<- interface{}) {
	p.timerOutChan = timerOutChan
}

func (p *SysTimerMgr) Exit() {
	//TODO
}

func (p *SysTimerMgr) AddSecond(cb OnTimerFun, arg interface{}, second time.Duration) (timer *time.Timer) {
	var stp SysTimerParameter
	stp.Fun = cb
	stp.Arg = arg

	f := p.timerFunc(stp)
	timer = time.AfterFunc(second*1000000000, f)
	return
}

//true:success,false:fail
func Del(timer *time.Timer) bool {
	return timer.Stop()
}

func (p *SysTimerMgr) timerFunc(parameter SysTimerParameter) func() {
	return func() {
		func(parameter interface{}) {
			p.timerOutChan <- parameter.(SysTimerParameter)
		}(parameter)
	}
}

type SysTimerParameter struct {
	Fun OnTimerFun
	Arg interface{}
}
