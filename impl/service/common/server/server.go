package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"path"
	"runtime"
	"time"

	"github.com/75912001/xr/lib/timer"

	"github.com/75912001/xr/lib/tcp"

	"github.com/75912001/xr/impl/service/common/bench"
	xrlog "github.com/75912001/xr/lib/log"
	"github.com/75912001/xr/lib/util"
)

type Server struct {
	GLog       xrlog.Log
	BenchMgr   bench.Mgr
	TimerMgr   timer.TimerMgr
	TcpService tcp.Server
	eventChan  chan interface{}
}

func (p *Server) Init(onConn tcp.OnConnServerType,
	onDisconn tcp.OnDisConnServerType,
	onPacket tcp.OnPacketServerType,
	onParseProtoHead tcp.OnParseProtoHeadType) (err error) {
	log.Printf("service Init.")

	rand.Seed(time.Now().UnixNano())

	currentPath, err := util.GetCurrentPath()
	if err != nil {
		log.Fatalf("GetCurrentPath fatal:%v", err)
		return
	}
	log.Printf("service current path:%v", currentPath)

	{ //加载bench.json文件
		err = p.BenchMgr.Parse(path.Join(currentPath, "bench.json"))
		if err != nil {
			log.Fatalf("parse bench.json err:%v", err)
			return
		}
		log.Printf("bench json:%+v", p.BenchMgr.Json)
	}

	{ //log
		err = p.GLog.Init(p.BenchMgr.Json.Base.LogAbsPath, fmt.Sprintf("%v-%v",
			p.BenchMgr.Json.Base.ServiceName, p.BenchMgr.Json.Base.ServiceID))
		if err != nil {
			log.Fatalf("log init err:%v", err)
			return
		}
		p.GLog.SetLevel(xrlog.LevelOn)
	}

	previousValue := runtime.GOMAXPROCS(p.BenchMgr.Json.Base.GoMaxProcs)
	p.GLog.Info(fmt.Sprintf("go max procs new:%v, prviousValue:%v", p.BenchMgr.Json.Base.GoMaxProcs, previousValue))

	//eventChan
	{
		p.eventChan = make(chan interface{}, p.BenchMgr.Json.Base.EventChanCnt)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					p.GLog.Warn(fmt.Sprintf("handleEvent goroutine panic:%v", err))
				}
				p.GLog.Trace("handleEvent goroutine done.")
			}()
			p.handleEvent()
		}()
	}
	//timer
	if 0 != p.BenchMgr.Json.Timer.ScanSecondDuration || 0 != p.BenchMgr.Json.Timer.ScanMillisecondDuration {
		p.TimerMgr.Start(context.Background(), p.BenchMgr.Json.Timer.ScanSecondDuration, p.BenchMgr.Json.Timer.ScanMillisecondDuration, p.eventChan)
	}

	//tcp service
	if 0 != len(p.BenchMgr.Json.Server.IP) || 0 != len(p.BenchMgr.Json.Server.Port) {
		address := p.BenchMgr.Json.Server.IP + ":" + p.BenchMgr.Json.Server.Port

		err = p.TcpService.Strat(address, &p.GLog, p.BenchMgr.Json.Base.PacketLengthMax, p.eventChan,
			onConn, onDisconn, onPacket, onParseProtoHead, p.BenchMgr.Json.Base.SendChanCapacity)
		if err != nil {
			p.GLog.Crit("StartTcpService err:", err)
			return
		}
	}

	return
}

func (p *Server) Exit() (err error) {
	p.TimerMgr.Exit()
	p.TcpService.Exit()
	p.GLog.Exit()

	return
}
