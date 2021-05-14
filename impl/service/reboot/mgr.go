package reboot

import (
	"fmt"
	"log"
	"math/rand"
	"path"
	"runtime"
	"time"

	"github.com/75912001/xr/impl/service/reboot/bench"
	xrlog "github.com/75912001/xr/lib/log"
	"github.com/75912001/xr/lib/timer"
	"github.com/75912001/xr/lib/util"
)

var GRebootMgr RebootMgr

type RebootMgr struct {
	Log       xrlog.Log
	BenchMgr  bench.Mgr
	TimerMgr  timer.TimerMgr
	EventChan chan interface{}
}

func (p *RebootMgr) Init() (err error) {
	log.Printf("RebootMgr Init.")

	rand.Seed(time.Now().UnixNano())

	currentPath, err := util.GetCurrentPath()
	if err != nil {
		log.Fatalf("GetCurrentPath fatal:%v", err)
		return
	}
	log.Printf("current path:%v", currentPath)
	{ //加载bench.json文件
		err = p.BenchMgr.Parse(path.Join(currentPath, "bench.json"))
		if err != nil {
			log.Fatalf("parse bench.json err:%v", err)
			return
		}
		log.Printf("bench json:%+v", p.BenchMgr.Json)
	}
	{ //log
		err = p.Log.Init(p.BenchMgr.Json.Base.LogAbsPath, fmt.Sprintf("%v-%v",
			p.BenchMgr.Json.Base.ServiceName, p.BenchMgr.Json.Base.ServiceID))
		if err != nil {
			log.Fatalf("log init err:%v", err)
			return
		}
		p.Log.SetLevel(int(p.BenchMgr.Json.Base.LogLevel))
	}
	{ //runtime.GOMAXPROCS
		previousValue := runtime.GOMAXPROCS(int(p.BenchMgr.Json.Base.GoMaxProcs))
		p.Log.Info(fmt.Sprintf("go max procs new:%v, prviousValue:%v", p.BenchMgr.Json.Base.GoMaxProcs, previousValue))
	}
	//eventChan
	{
		p.EventChan = make(chan interface{}, p.BenchMgr.Json.Base.EventChanCnt)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					p.Log.Warn(fmt.Sprintf("handle_event goroutine panic:%v", err))
				}
				p.Log.Trace("handle_event goroutine done.")
			}()
			p.handleEvent()
		}()
	}
	//timer
	{
		if 0 != p.BenchMgr.Json.Timer.ScanSecondDuration || 0 != p.BenchMgr.Json.Timer.ScanMillisecondDuration {
			p.TimerMgr.Start(p.BenchMgr.Json.Timer.ScanSecondDuration, p.BenchMgr.Json.Timer.ScanMillisecondDuration, p.EventChan)
		}
	}

	return
}
