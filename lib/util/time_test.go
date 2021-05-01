package util

import (
	"fmt"
	"testing"
	"time"
)

func TestGenYYYYMMDD(t *testing.T) {
	yyyymmdd := GenYYYYMMDD(time.Now().Unix())
	fmt.Printf("YYYYMMDD:%v\n", yyyymmdd)
}

