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
	second int64
	data   string
}

// 写日志
func (p *Log) onOutPut() {
	var err error
	defer p.waitGroup.Done()

	for v := range p.logChan {
		yyyymmdd := genYYYYMMDD(v.second)
		if p.yyyymmdd != yyyymmdd {
			p.file.Close()
			p.yyyymmdd = yyyymmdd
			logName := p.namePrefix + strconv.Itoa(p.yyyymmdd)
			logName = genLogName(p.namePrefix, fmt.Sprintf("%v", p.yyyymmdd), fmt.Sprintf("%v", v.second))
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
	second := time.Now().Unix()
	p.logChan <- &logData{
		second: second,
		data:   "[" + *prefix + "][" + funName + "][" + strLine + "]" + *str,
	}
}
