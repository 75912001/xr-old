package timer

import (
	"container/list"
	"math"
	"time"
)

//TimerSecond 秒级定时器
type TimerSecond struct {
	TimerMillisecond
}

func (p *TimerMgr) addSecond(cb OnTimerFun, arg interface{}, expire int64, oldTimerSecond *TimerSecond) (t *TimerSecond) {
	tvecRootIdx := p.findTvecRootIdx(expire)
	if nil == oldTimerSecond {
		oldTimerSecond = &TimerSecond{
			TimerMillisecond{
				expire,
				arg,
				cb,
				true,
			},
		}
	} else {
		oldTimerSecond.expire = expire
		oldTimerSecond.Arg = arg
		oldTimerSecond.Function = cb
	}

	p.secondVec[tvecRootIdx].data.PushBack(oldTimerSecond)

	if expire < p.secondVec[tvecRootIdx].minExpire {
		p.secondVec[tvecRootIdx].minExpire = expire
	}
	t = oldTimerSecond
	return
}

//扫描秒级定时器
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
					p.addSecond(t.Function, t.Arg, t.expire, t)
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
