package util

import (
	"os"
	"path"
	"testing"

	"github.com/75912001/xr/lib/util"
)

func TestUnmarshalXmlFile(t *testing.T) {
	fileName := "__test_file__.xml"
	currentPath, err := util.GetCurrentPath()
	if err != nil {
		t.Errorf("GetCurrentPath err:%v", err)
	}
	testPathFile := path.Join(currentPath, fileName)
	data := []byte(__test_xml_file_content__)
	err = util.OverWriteFile(testPathFile, data)
	if err != nil {
		t.Errorf("write file %v err:%v", testPathFile, err)
	}
	defer func() {
		os.Remove(testPathFile)
	}()
	var x testXML
	err = util.UnmarshalXmlFile(testPathFile, &x)
	if err != nil {
		t.Errorf("UnmarshalXmlFile err:%v", err)
	}
}

type testXML struct {
	ServiceName string `xml:"serviceName"`
	ServiceID   int    `xml:"serviceID"`
	Server      struct {
		IP   string `xml:"ip"`
		Port string `xml:"port"`
	} `xml:"server"`
}

const __test_xml_file_content__ string = `
<xml>
	<serviceName>name</serviceName>
	<serviceID>123</serviceID>
	<server>
		<ip>127.0.0.1</ip>
		<port>8899</port>
	</server>
</xml>
`
