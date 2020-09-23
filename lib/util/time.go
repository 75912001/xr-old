package util

import (
	"strconv"
	"time"
)

// GenYYYYMMDD 获取yyyymmdd
func GenYYYYMMDD(sec int64) (yyyymmdd int) {
	strYYYYMMDD := time.Unix(sec, 0).Format("20060102")
	yyyymmdd, _ = strconv.Atoi(strYYYYMMDD)
	return
}
