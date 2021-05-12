package log

import (
	"fmt"
	"log"
	"os"
	"path"
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
	for v := range p.logChan {
		yyyymmdd := genYYYYMMDD(v.second)
		if p.yyyymmdd != yyyymmdd {
			if p.file != nil {
				p.file.Close()
				p.file = nil
			}

			p.yyyymmdd = yyyymmdd
			logName := genLogName(p.namePrefix, fmt.Sprintf("%v", p.yyyymmdd), genHHMMSS(v.second))
			p.file, err = os.OpenFile(path.Join(p.absPath, logName), logFileFlag, logFilePerm)
			if err != nil {
				log.Printf("log onOutPut OpenFile err:%v", err)
				log.Printf("log:%v", v.data)
				continue
			}
			p.logger = log.New(p.file, "", logFlag)
		}
		p.logger.Print(v.data)
	}
}

// 路径,文件名,行数,函数名称
func (p *Log) levelFunc(level int, v ...interface{}) {
	if p.level < level {
		return
	}

	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		log.Printf("runtime.Caller err.")
		return
	}
	funName := runtime.FuncForPC(pc).Name()
	var strLine = strconv.Itoa(line)
	second := time.Now().Unix()
	p.logChan <- &logData{
		second: second,
		data:   "[" + levelTag[level] + "][" + file + "][" + funName + "][" + strLine + "]" + fmt.Sprintln(v...),
	}
}
