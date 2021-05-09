package log

import (
	"log"
	"os"
	"strconv"
	"strings"
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

//if 15:04:05  return 150405
func genHHMMSS(sec int64) (hhmmss string) {
	hhmmss = time.Unix(sec, 0).Format("15:04:05")
	hhmmss = strings.Replace(hhmmss, ":", "", -1)
	return
}

func genLogName(namePrefix, yyyymmdd, second string) (logName string) {
	return namePrefix + "-" + yyyymmdd + "-" + second + ".log"
}
