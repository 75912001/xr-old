package log

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/75912001/xr/lib/util"
)

type logData struct {
	YYYYMMDD int
	data     string
}

// 写日志
func (p *Log) onOutPut() {
	defer p.waitGroup.Done()
	for v := range p.logChan {
		if p.yyyymmdd != v.YYYYMMDD {
			p.file.Close()
			p.yyyymmdd = v.YYYYMMDD
			logName := p.namePrefix + strconv.Itoa(p.yyyymmdd)
			p.file, _ = os.OpenFile(logName, p.fileFlag, p.perm)
			p.logger = log.New(p.file, "", p.logFlag)
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
	p.logChan <- logData{
		YYYYMMDD: util.GenYYYYMMDD(time.Now().Unix()),
		data:     "[" + *prefix + "][" + funName + "][" + strLine + "]" + *str,
	}
}
