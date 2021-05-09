package util

import (
	"os"
	"path"
	"testing"

	"github.com/75912001/xr/lib/util"
)

func TestOverWriteFile(t *testing.T) {
	fileName := "__test_file__"
	currentPath, err := util.GetCurrentPath()
	if err != nil {
		t.Errorf("GetCurrentPath err:%v", err)
	}
	testPathFile := path.Join(currentPath, fileName)
	err = util.OverWriteFile(testPathFile, __test_file_content__)
	if err != nil {
		t.Errorf("write file %v err:%v", testPathFile, err)
	}
	defer func() {
		os.Remove(testPathFile)
	}()
}

func TestPathFileExists(t *testing.T) {
	fileName := "__test_file__"
	currentPath, err := util.GetCurrentPath()
	if err != nil {
		t.Errorf("GetCurrentPath err:%v", err)
	}
	testPathFile := path.Join(currentPath, fileName)

	//检测不存在的文件
	os.Remove(testPathFile)
	exists := util.PathFileExists(testPathFile)
	if exists {
		t.Errorf("file exists %v", testPathFile)
	}

	//检测存在的文件
	err = util.OverWriteFile(testPathFile, __test_file_content__)
	if err != nil {
		t.Errorf("write file %v err:%v", testPathFile, err)
	}
	defer func() {
		os.Remove(testPathFile)
	}()
	exists = util.PathFileExists(testPathFile)
	if !exists {
		t.Errorf("file non-exists %v", testPathFile)
	}
}

func TestMkdirAll(t *testing.T) {
	dirName := "__test_dir1__"
	currentPath, err := util.GetCurrentPath()
	if err != nil {
		t.Errorf("GetCurrentPath err:%v", err)
	}
	testPath1 := path.Join(currentPath, dirName)
	os.RemoveAll(testPath1)

	testPath2 := path.Join(testPath1, "__test_dir2__")
	err = util.MkdirAll(testPath2)
	if err != nil {
		t.Errorf("MkdirAll err:%v", err)
	}
	defer func() {
		os.RemoveAll(testPath1)
	}()
	if !util.PathFileExists(testPath2) {
		t.Errorf("%v non-exists", testPath2)
	}
}

const __test_file_content__ string = `
this is test file content.
`
