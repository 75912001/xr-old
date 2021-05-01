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
	var err error
	defer p.waitGroup.Done()

	for v := range p.logChan {
		if p.yyyymmdd != v.yyyymmdd {
			p.file.Close()
			p.yyyymmdd = v.yyyymmdd
			logName := p.namePrefix + strconv.Itoa(p.yyyymmdd)

			p.file, err = os.OpenFile(logName, logFileFlag, logFilePerm)
			if err != nil {
				fmt.Printf("log onOutPut OpenFile err:%v\n", err)
				fmt.Printf("log:%v\n", v.data)
				continue
			}
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
