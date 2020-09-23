package timer

import (
	"container/list"
	"time"
)

// TimerMillisecond 毫秒级定时器
type TimerMillisecond struct {
	expire   int64 //过期时间戳
	Arg      interface{}
	Function OnTimerFun //超时调用的函数
	valid    bool       //有效(false:不执行,扫描时自动删除)
}

// 扫描毫秒级定时器
func (p *TimerMgr) scanMillisecond() {
	t := time.Now()
	millisecond := t.UnixNano() / 1000000

	var next *list.Element
	for e := p.millisecondList.Front(); e != nil; e = next {
		timerMillisecond := e.Value.(*TimerMillisecond)
		if !timerMillisecond.valid {
			next = e.Next()
			p.millisecondList.Remove(e)
			continue
		}
		if timerMillisecond.expire <= millisecond {
			p.timerOutChan <- timerMillisecond
			next = e.Next()
			p.millisecondList.Remove(e)
		} else {
			next = e.Next()
		}
	}
}
