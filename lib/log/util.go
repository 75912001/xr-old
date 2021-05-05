package log

import (
	"log"
	"os"
	"strconv"
	"time"
)

var logFilePerm os.FileMode = os.ModePerm
var logFileFlag int = os.O_CREATE | os.O_APPEND | os.O_RDWR
var logFlag int = log.Lmicroseconds //log.Ldate|log.Llongfile

// genYYYYMMDD 获取yyyymmdd
func genYYYYMMDD(sec int64) (yyyymmdd int) {
	strYYYYMMDD := time.Unix(sec, 0).Format("20060102")
	yyyymmdd, _ = strconv.Atoi(strYYYYMMDD)
	return
}

func genLogName(namePrefix, yyyymmdd, second string) (logName string) {
	return namePrefix + "-" + yyyymmdd + "-" + second + ".log"
}
