package util

import (
	"os"
	"path"
	"testing"

	"github.com/75912001/xr/lib/util"
)

func TestUnmarshalJsonFile(t *testing.T) {
	fileName := "__test_file__.json"
	currentPath, err := util.GetCurrentPath()
	if err != nil {
		t.Errorf("GetCurrentPath err:%v", err)
	}
	testPathFile := path.Join(currentPath, fileName)
	data := []byte(__test_json_file_content__)
	err = util.OverWriteFile(testPathFile, data)
	if err != nil {
		t.Errorf("write file %v err:%v", testPathFile, err)
	}
	defer func() {
		os.Remove(testPathFile)
	}()
	var tj testJson
	err = util.UnmarshalJsonFile(testPathFile, &tj)
	if err != nil {
		t.Errorf("UnmarshalJsonFile err:%v", err)
	}
}

func TestSaveJson(t *testing.T) {
	var tj testJson
	tj.ServiceName = "name"
	tj.ServiceID = 123
	tj.Server.IP = "127.0.0.1"
	tj.Server.Port = "8899"

	fileName := "__test_file__.json"
	currentPath, err := util.GetCurrentPath()
	if err != nil {
		t.Errorf("GetCurrentPath err:%v", err)
	}
	testPathFile := path.Join(currentPath, fileName)

	err = util.SaveJson(&tj, testPathFile)
	if err != nil {
		t.Errorf("SaveJson file %v err:%v", testPathFile, err)
	}
	defer func() {
		os.Remove(testPathFile)
	}()
}

type testJson struct {
	ServiceName string `json:"serviceName"`
	ServiceID   int    `json:"serviceID"`
	Server      struct {
		IP   string `json:"ip"`
		Port string `json:"port"`
	} `json:"server"`
}

const __test_json_file_content__ string = `
{
  "serviceName": "world",
  "serviceID": 1,
  "__comments__":"无,则不启用该功能",
  "server": {
    "ip": "127.0.0.1",
    "port": "3001"
  }
}
`

/*
/////////////////////////////////////////////////////////////////////////////
//TODO 新功能
//JSON2map JSON转换成为Map
func JSON2map(strJSON *string) (s map[string]interface{}, err error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(*strJSON), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func TestJson2map(t *testing.T) {
	//	var strJson string = `
	//{
	//"tradeNo":"5c84ad403373ec0803dbddddc77246b1",
	//"productId":"tjlhxkgddj0o1",
	//"k1":1,
	//"k2":"v2",
	//"k3":"v3",
	//"k4":"v4"
	//}
	//`
	var strJson string = `
	{
		"k1":1,
		"k2":"2",
		"k3":[3],
		"k4":["4"],
		"k5":[5,55],
		"k6":["6","66"],
		"k7":["7",77],
		"k8-1":[
			{
				"k8-2-1":821,
				"k8-2-2":822
			},
			{
				"k8-2-10":821,
				"k8-2-20":822
			}
		]
	}
`

	//var jsonMap map[string]interface{}
	//jsonMap = make(map[string]interface{}, 0)

	jsonMap, err := JSON2map(&strJson)
	if nil == err {
		//成功
		fmt.Println("parse json success:", jsonMap)
	} else {
		//失败
		fmt.Println("parse json err:", err)
	}
	{
		v, ok := jsonMap["k1"]
		if ok {
			vv := v.(float64)
			fmt.Println("value:", vv)
		} else {
			//non-existent
			fmt.Println("non-existent")
		}
	}
	{
		v, ok := jsonMap["k2"]
		if ok {
			vv := v.(string)
			fmt.Println("value:", vv)
		} else {
			//non-existent
			fmt.Println("non-existent")
		}
	}
	{
		for k, v := range jsonMap {
			fmt.Println(k, v)
		}
	}
}
*/
