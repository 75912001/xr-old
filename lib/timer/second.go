package timer

import (
	"container/list"
	"math"
	"time"
)

// Second 秒级定时器
type Second struct {
	Millisecond
}

// AddSecond 添加秒级定时器
func (p *TimerMgr) AddSecond(cb OnTimerFun, arg interface{}, expire int64) (t *Second) {
	p.secondMutex.Lock()
	defer func() {
		p.secondMutex.Unlock()
	}()

	return p.addSecond(cb, arg, expire)
}

// DelSecond 删除秒级定时器(必须与该timerOutChan线性处理.如:在同一个goroutine select中处理数据.)
func DelSecond(t *Second) {
	t.Millisecond.inValid()
}

func (p *TimerMgr) addSecond(cb OnTimerFun, arg interface{}, expire int64) (t *Second) {
	tvecRootIdx := findTvecRootIdx(expire)

	t = &Second{
		Millisecond{
			Arg:      arg,
			Function: cb,
			expire:   expire,
			valid:    true,
		},
	}
	p.pushBackTvecRoot(t, tvecRootIdx)
	return
}

// 将秒级定时器,添加到轮转IDX的末尾.
func (p *TimerMgr) pushBackTvecRoot(timerSecond *Second, tvecRootIdx int) {
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
			t := e.Value.(*Second)
			if !t.IsValid() {
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
				t := e.Value.(*Second)
				if !t.IsValid() {
					next = e.Next()
					tr.data.Remove(e)
					continue
				}
				newIdx := findPrevTvecRootIdx(t.expire-second, idx)
				if idx != newIdx {
					next = e.Next()
					tr.data.Remove(e)
					p.pushBackTvecRoot(t, newIdx)
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
