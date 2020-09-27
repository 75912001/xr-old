package log

//使用系统log,自带锁
//使用协程操作io输出日志
//每天自动创建新的日志文件

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/75912001/xr/lib/util"
)

var GLog *Log

// Log 日志
type Log struct {
	level      int      //日志等级
	file       *os.File //日志文件
	logger     *log.Logger
	logChan    chan logData
	yyyymmdd   int    //日志年月日
	namePrefix string //日志文件名称前缀
	perm       os.FileMode
	fileFlag   int
	logFlag    int
	waitGroup  sync.WaitGroup
}

// Init 初始化
// name:日志前缀名称
func (p *Log) Init(namePrefix string) (err error) {
	p.perm = os.ModePerm
	p.fileFlag = os.O_CREATE | os.O_APPEND | os.O_RDWR
	p.logFlag = log.Ltime //log.Ldate|log.Llongfile
	p.level = levelOn
	p.namePrefix = namePrefix
	p.yyyymmdd = util.GenYYYYMMDD(time.Now().Unix())

	logName := p.namePrefix + strconv.Itoa(p.yyyymmdd)
	p.file, err = os.OpenFile(logName, p.fileFlag, p.perm)
	if nil != err {
		return err
	}
	p.logger = log.New(p.file, "", p.logFlag)

	p.logChan = make(chan logData, 10000)

	p.waitGroup.Add(1)
	go p.onOutPut()
	return err
}

// SetLevel 设置日志等级
func (p *Log) SetLevel(level int) {
	p.level = level
}

// 退出
func (p *Log) Exit() {
	close(p.logChan)
	p.waitGroup.Wait()
	p.file.Close()
}

// Trace 踪迹日志
func (p *Log) Trace(v ...interface{}) {
	if p.level < levelTrace {
		return
	}
	body := fmt.Sprintln(v...)
	p.outPut(2, &strTrace, &body)
}

// Debug 调试日志
func (p *Log) Debug(v ...interface{}) {
	if p.level < levelDebug {
		return
	}
	body := fmt.Sprintln(v...)
	p.outPut(2, &strDebug, &body)
}

// Info 报告日志
func (p *Log) Info(v ...interface{}) {
	if p.level < levelInfo {
		return
	}
	body := fmt.Sprintln(v...)
	p.outPut(2, &strInfo, &body)
}

// Notice 公告日志
func (p *Log) Notice(v ...interface{}) {
	if p.level < levelNotice {
		return
	}
	body := fmt.Sprintln(v...)
	p.outPut(2, &strNotice, &body)
}

// Warn 警告日志
func (p *Log) Warn(v ...interface{}) {
	if p.level < levelWarn {
		return
	}
	body := fmt.Sprintln(v...)
	p.outPut(2, &strWarn, &body)
}

// Error 错误日志
func (p *Log) Error(v ...interface{}) {
	if p.level < levelError {
		return
	}
	body := fmt.Sprintln(v...)
	p.outPut(2, &strError, &body)
}

// Crit 临界日志
func (p *Log) Crit(v ...interface{}) {
	if p.level < levelCrit {
		return
	}
	body := fmt.Sprintln(v...)
	p.outPut(2, &strCrit, &body)
}

// Emerg 不可用日志
func (p *Log) Emerg(v ...interface{}) {
	if p.level < levelEmerg {
		return
	}
	body := fmt.Sprintln(v...)
	p.outPut(2, &strEmerg, &body)
}
