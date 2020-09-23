package timer

import (
	"container/list"
	"math"
	"time"
)

// 时间轮数量
const eTimerVecSize int = 9

// 时间轮持续时间
var gTvecRootDuration [eTimerVecSize]int64

func init() {
	for i := 0; i < len(gTvecRootDuration); i++ {
		gTvecRootDuration[i] = genDuration(i)
	}
}

type tvecRoot struct {
	data      *list.List
	minExpire int64 //最小到期时间
}

func (p *tvecRoot) init() {
	p.data = list.New()
	p.minExpire = math.MaxInt64
}

// 4,8,16,32,64,128,256,512,1024...
func genDuration(idx int) (duration int64) {
	duration = 1 << (uint)(idx+2)
	return duration
}

// 根据到期时间找到时间轮的序号
func (p *TimerMgr) findTvecRootIdx(expire int64) (idx int) {
	var duration = expire - time.Now().Unix()
	for k, v := range gTvecRootDuration {
		if duration <= v {
			idx = k
			return
		} else {
			idx++
		}
	}

	if len(gTvecRootDuration) <= idx {
		idx = len(gTvecRootDuration) - 1
	}
	return
}

// 向前查找符合时间差的时间轮序号
func (p *TimerMgr) findPrevTvecRootIdx(duration int64, srcIdx int) (idx int) {
	idx = srcIdx
	for {
		if 0 != srcIdx && duration <= gTvecRootDuration[srcIdx-1] {
			srcIdx--
			idx = srcIdx
		} else {
			break
		}
	}
	return
}
