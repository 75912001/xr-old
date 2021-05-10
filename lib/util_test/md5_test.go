package util

import (
	"os"
	"path"
	"testing"

	"github.com/75912001/xr/lib/util"
)

func TestGenMd5(t *testing.T) {
	var data string = "menglc"
	t.Logf("TestGenMd5:%v, md5sum:%v", data, util.GenMd5([]byte(data)))
}

func TestMD5File(t *testing.T) {
	fileName := "__test_file__"
	currentPath, err := util.GetCurrentPath()
	if err != nil {
		t.Errorf("GetCurrentPath err:%v", err)
	}
	testPathFile := path.Join(currentPath, fileName)
	data := []byte("menglc")
	err = util.OverWriteFile(testPathFile, data)
	if err != nil {
		t.Errorf("write file %v err:%v", testPathFile, err)
	}
	defer func() {
		os.Remove(testPathFile)
	}()
	md5sum, err := util.MD5File(testPathFile)
	if err != nil {
		t.Errorf("MD5File file %v err:%v", testPathFile, err)
	}
	t.Logf("TestMD5File:%v, md5sum:%v", data, md5sum)
}
