package log

import (
	"testing"
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
*/
func TestExample(t *testing.T) {
	var log *Log = new(Log)
	log.Init("test_log")

	log.Trace("trace")
	log.Debug("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warn("warn")
	log.Error("error")
	log.Crit("crit")
	log.Emerg("emerg")

	log.Exit()
}
