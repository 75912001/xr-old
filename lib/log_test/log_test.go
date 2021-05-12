package log_test

import (
	"fmt"
	"testing"

	"github.com/75912001/xr/lib/util"

	xrlog "github.com/75912001/xr/lib/log"
)

/*
 go test -v -count=1
*/

/*
型号名称：	MacBook Pro
型号标识符：	MacBookPro11,4
处理器名称：	Intel Core i7
处理器速度：	2.2 GHz
处理器数目：	1
核总数：	4
L2 缓存（每个核）：	256 KB
L3 缓存：	6 MB
内存：	16 GB

每行65字节
共65,000,000 byte=>62M
////////////////////////////////////////////////////////////////////////////////
100W 7.721s=>129516次/s=>130次/ms

new:
120704k  7.01s
80W
*/

func TestExample(t *testing.T) {
	absPath, err := util.GetCurrentPath()
	if err != nil {
		t.Errorf("GetCurrentPath err:%v", err)
	}
	t.Logf("absPath:%v", absPath)
	var log *xrlog.Log = new(xrlog.Log)
	err = log.Init(absPath, "test_log")
	if err != nil {
		t.Errorf("log Init err:%v", err)
	}

	log.SetLevel(xrlog.LevelOn)
	for i := 0; i < 10000; i++ {
		log.Trace(fmt.Sprintf("LevelOn trace:%v", 1))
		log.Debug(fmt.Sprintf("LevelOn debug:%v,%v", "2", 3))
		log.Info("LevelOn info")
		log.Notice("LevelOn notice")
		log.Warn("LevelOn warn")
		log.Error("LevelOn error")
		log.Crit("LevelOn crit")
		log.Emerg("LevelOn emerg")
	}

	log.SetLevel(xrlog.LevelOff)
	log.Trace(fmt.Sprintf("LevelOff trace:%v", 1))
	log.Debug(fmt.Sprintf("LevelOff debug:%v,%v", "2", 3))
	log.Info("LevelOff info")
	log.Notice("LevelOff notice")
	log.Warn("LevelOff warn")
	log.Error("LevelOff error")
	log.Crit("LevelOff crit")
	log.Emerg("LevelOff emerg")

	//for test coverage
	log.SetLevel(100)
	log.Exit()
}
