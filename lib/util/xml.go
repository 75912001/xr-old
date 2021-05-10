package util

import (
	"encoding/xml"
	"io/ioutil"
)

func UnmarshalXmlFile(pathFile string, v interface{}) (err error) {
	data, err := ioutil.ReadFile(pathFile)
	if err != nil {
		return
	}

	err = xml.Unmarshal(data, v)
	if err != nil {
		return
	}
	return
}
