package util

import (
	"encoding/json"
	"io/ioutil"
)

func ParseJson(pathFile string, v interface{}) (err error) {
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

func SaveJson(pathFile string, v interface{}) (err error) {
	saveData, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		return
	}
	err = OverWriteFile(pathFile, string(saveData))
	if err != nil {
		return
	}
	return
}
