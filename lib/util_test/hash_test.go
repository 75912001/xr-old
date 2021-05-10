package util

import (
	"testing"

	"github.com/75912001/xr/lib/util"
)

func TestHASH32(t *testing.T) {
	var content string = "menglc"
	data := []byte(content)
	value := util.HASH32(data)
	t.Logf("%v hase32:%v", content, value)
}

func TestHASH64(t *testing.T) {
	var content string = "menglc"
	data := []byte(content)
	value := util.HASH64(data)
	t.Logf("%v hash64:%v", content, value)
}
