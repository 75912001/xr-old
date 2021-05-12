package log

//使用系统log,自带锁
//使用协程操作io输出日志
//每天自动创建新的日志文件

import (
	"log"
	"os"
	"sync"

	"github.com/75912001/xr/lib/util"
)

// Log 日志
type Log struct {
	level           int      //日志等级
	file            *os.File //日志文件
	logger          *log.Logger
	logChan         chan *logData
	yyyymmdd        int    //日志年月日
	namePrefix      string //日志文件名称前缀
	waitGroupOutPut sync.WaitGroup
	absPath         string //绝对路径
}

// Init 初始化
// name:日志前缀名称
func (p *Log) Init(absPath, namePrefix string) (err error) {
	err = util.MkdirAll(absPath)
	if err != nil {
		return
	}

	p.absPath = absPath
	p.level = LevelOn
	p.namePrefix = namePrefix

	p.logChan = make(chan *logData, 100000)

	p.waitGroupOutPut.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("log onOutPut goroutine painc:%v", err)
			}
			p.waitGroupOutPut.Done()
		}()
		p.onOutPut()
	}()
	return err
}

// SetLevel 设置日志等级
func (p *Log) SetLevel(level int) {
	if level < LevelOff || LevelOn < level {
		log.Printf("log SetLevel level err:%v", level)
		return
	}
	p.level = level
}

// 退出
func (p *Log) Exit() {
	if p.logChan != nil {
		//close chan, for range 读完chan会退出.
		close(p.logChan)
		//等待logChan 的for range 退出.
		p.waitGroupOutPut.Wait()
		p.file.Close()
		//goroutine 退出,再设置chan为nil, (如果没有退出就设置为nil, 读chan == nil  会 block)
		p.logChan = nil
	}
}

//Trace 踪迹日志
func (p *Log) Trace(v ...interface{}) {
	p.levelFunc(LevelTrace, v)
}

// Debug 调试日志
func (p *Log) Debug(v ...interface{}) {
	p.levelFunc(LevelDebug, v)
}

// Info 报告日志
func (p *Log) Info(v ...interface{}) {
	p.levelFunc(LevelInfo, v)
}

// Notice 公告日志
func (p *Log) Notice(v ...interface{}) {
	p.levelFunc(LevelNotice, v)
}

// Warn 警告日志
func (p *Log) Warn(v ...interface{}) {
	p.levelFunc(LevelWarn, v)
}

// Error 错误日志
func (p *Log) Error(v ...interface{}) {
	p.levelFunc(LevelError, v)
}

// Crit 临界日志
func (p *Log) Crit(v ...interface{}) {
	p.levelFunc(LevelCrit, v)
}

// Emerg 不可用日志
func (p *Log) Emerg(v ...interface{}) {
	p.levelFunc(LevelEmerg, v)
}
