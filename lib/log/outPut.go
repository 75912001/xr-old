package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

type logData struct {
	yyyymmdd int
	data     string
}

// 写日志
func (p *Log) onOutPut() {
	defer func() {
		p.waitGroup.Done()
		if err := recover(); err != nil {
			fmt.Printf("log onOutPut goroutine recover:%s\n", err)
		}
	}()
	for v := range p.logChan {
		if p.yyyymmdd != v.yyyymmdd {
			p.file.Close()
			p.yyyymmdd = v.yyyymmdd
			logName := p.namePrefix + strconv.Itoa(p.yyyymmdd)
			p.file, _ = os.OpenFile(logName, logFileFlag, logFilePerm)
			p.logger = log.New(p.file, "", logFlag)
		}
		p.logger.Print(v.data)
	}
}

// 路径,文件名,行数,函数名称
func (p *Log) outPut(calldepth int, prefix *string, str *string) {
	pc, _, line, ok := runtime.Caller(calldepth)
	if !ok {
		return
	}
	funName := runtime.FuncForPC(pc).Name()
	var strLine = strconv.Itoa(line)
	p.logChan <- &logData{
		yyyymmdd: genYYYYMMDD(time.Now().Unix()),
		data:     "[" + *prefix + "][" + funName + "][" + strLine + "]" + *str,
	}
}
