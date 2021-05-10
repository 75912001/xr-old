package util

import (
	"encoding/json"
	"io/ioutil"
)

func UnmarshalJsonFile(pathFile string, v interface{}) (err error) {
	data, err := ioutil.ReadFile(pathFile)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
	return
}

func SaveJson(v interface{}, dstPathFile string) (err error) {
	data, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		return
	}

	err = OverWriteFile(dstPathFile, data)
	if err != nil {
		return
	}
	return
}
