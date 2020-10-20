package timer

import (
	"container/list"
	"math"
	"time"
)

// TimerSecond 秒级定时器
type TimerSecond struct {
	TimerMillisecond
}

func (p *TimerMgr) addSecond(cb OnTimerFun, arg interface{}, expire int64) (t *TimerSecond) {
	tvecRootIdx := p.findTvecRootIdx(expire)

	t = &TimerSecond{
		TimerMillisecond{
			Arg:      arg,
			Function: cb,
			expire:   expire,
			valid:    true,
		},
	}

	p.secondVec[tvecRootIdx].data.PushBack(t)

	if expire < p.secondVec[tvecRootIdx].minExpire {
		p.secondVec[tvecRootIdx].minExpire = expire
	}

	return
}

// 换挡,秒级定时器. 将定时器,添加到轮转IDX中.
func (p *TimerMgr) shiftTvecRoot(timerSecond *TimerSecond, tvecRootIdx int) {
	p.secondVec[tvecRootIdx].data.PushBack(timerSecond)

	if timerSecond.expire < p.secondVec[tvecRootIdx].minExpire {
		p.secondVec[tvecRootIdx].minExpire = timerSecond.expire
	}
}

// 扫描秒级定时器
func (p *TimerMgr) scanSecond() {
	second := time.Now().Unix()

	var next *list.Element

	tr0 := p.secondVec[0]
	if tr0.minExpire <= second {
		//更新最小过期时间戳
		tr0.minExpire = math.MaxInt64
		for e := tr0.data.Front(); nil != e; e = next {
			t := e.Value.(*TimerSecond)
			if !t.valid {
				next = e.Next()
				tr0.data.Remove(e)
				continue
			}
			if t.expire <= second {
				p.timerOutChan <- t
				next = e.Next()
				tr0.data.Remove(e)
			} else {
				if t.expire < tr0.minExpire {
					tr0.minExpire = t.expire
				}
				next = e.Next()
			}
		}
	}

	//更新时间轮,从序号为1的数组开始
	for idx := 1; idx < eTimerVecSize; idx++ {
		tr := p.secondVec[idx]
		if (tr.minExpire - second) <= gTvecRootDuration[idx-1] {
			tr.minExpire = math.MaxInt64
			for e := tr.data.Front(); e != nil; e = next {
				t := e.Value.(*TimerSecond)
				if !t.valid {
					next = e.Next()
					tr.data.Remove(e)
					continue
				}
				newIdx := p.findPrevTvecRootIdx(t.expire-second, idx)
				if idx != newIdx {
					next = e.Next()
					tr.data.Remove(e)
					p.shiftTvecRoot(t, newIdx)
				} else {
					if t.expire < tr.minExpire {
						tr.minExpire = t.expire
					}
					next = e.Next()
				}
			}
		}
	}
}
