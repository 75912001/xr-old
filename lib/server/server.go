package server

import (
	"fmt"
	"log"
	"math/rand"
	"path"
	"runtime"
	"time"

	"github.com/75912001/xr/lib/addr"

	"github.com/75912001/xr/lib/timer"

	"github.com/75912001/xr/lib/tcp"

	"github.com/75912001/xr/lib/bench"
	xrlog "github.com/75912001/xr/lib/log"
	"github.com/75912001/xr/lib/util"
)

type Server struct {
	GLog       xrlog.Log
	BenchMgr   bench.Mgr
	TimerMgr   timer.TimerMgr
	TcpService tcp.Server
	Addr       addr.Addr
	eventChan  chan interface{}
}

func (p *Server) Init(onEventConnServerFunc tcp.OnEventConnServer_F,
	onEventDisConnServerFunc tcp.OnEventDisConnServerFunc,
	onEventPacketServerFunc tcp.OnEventPacketServerFunc,
	onParseProtoHeadFunc tcp.OnParseProtoHead_F,
	onEventAddrMulticastFunc addr.OnEventAddrMulticastFunc) (err error) {
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
		p.GLog.SetLevel(int(p.BenchMgr.Json.Base.LogLevel))
	}
	{ //runtime.GOMAXPROCS
		previousValue := runtime.GOMAXPROCS(int(p.BenchMgr.Json.Base.GoMaxProcs))
		p.GLog.Info(fmt.Sprintf("go max procs new:%v, prviousValue:%v", p.BenchMgr.Json.Base.GoMaxProcs, previousValue))
	}
	//eventChan
	{
		p.eventChan = make(chan interface{}, p.BenchMgr.Json.Base.EventChanCnt)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					p.GLog.Warn(fmt.Sprintf("handle_event goroutine panic:%v", err))
				}
				p.GLog.Trace("handle_event goroutine done.")
			}()
			p.handleEvent()
		}()
	}
	//timer
	{
		if 0 != p.BenchMgr.Json.Timer.ScanSecondDuration || 0 != p.BenchMgr.Json.Timer.ScanMillisecondDuration {
			p.TimerMgr.Start(p.BenchMgr.Json.Timer.ScanSecondDuration, p.BenchMgr.Json.Timer.ScanMillisecondDuration, p.eventChan)
		}
	}
	//tcp service
	{
		if 0 != len(p.BenchMgr.Json.Server.IP) || 0 != len(p.BenchMgr.Json.Server.Port) {
			address := p.BenchMgr.Json.Server.IP + ":" + p.BenchMgr.Json.Server.Port

			err = p.TcpService.Strat(address, p.BenchMgr.Json.Base.PacketLengthMax, p.eventChan,
				onEventConnServerFunc, onEventDisConnServerFunc, onEventPacketServerFunc, onParseProtoHeadFunc, p.BenchMgr.Json.Base.SendChanCapacity)
			if err != nil {
				p.GLog.Crit("StartTcpService err:", err)
				return
			}
		}
	}
	//add multicast
	{
		am := &p.BenchMgr.Json.AddrMulticast
		m := &p.BenchMgr.Json.Multicast
		if 0 != len(am.Name) && 0 != am.ID && 0 != len(am.IP) && 0 != am.Port &&
			0 != len(m.IP) && 0 != m.Port && 0 != len(m.NetworkInterfacenName) {
			err = p.Addr.Start(p.eventChan, onEventAddrMulticastFunc, m.IP, m.Port, m.NetworkInterfacenName,
				am.Name, am.ID, am.IP, am.Port, am.Data)
			if err != nil {
				p.GLog.Crit(fmt.Printf("addr multicase err:%v", err))
				return
			}
		}
	}

	return
}

func (p *Server) Stop() (err error) {
	p.TimerMgr.Stop()
	p.TcpService.Stop()
	p.Addr.Stop()
	p.GLog.Stop()

	return
}
