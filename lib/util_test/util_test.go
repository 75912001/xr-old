package util_test

import (
	"fmt"
	"testing"

	"github.com/75912001/xr/lib/util"
)

func TestGetCurrentFuncName(t *testing.T) {
	fmt.Printf("CurrentFuncName:%v\n", util.GetCurrentFuncName())
}

/*
func TestGetCurrentPath(t *testing.T) {
	fmt.Printf("CurrentPath:%v\n", util.GetCurrentPath())
}

func TestGenYYYYMMDD(t *testing.T) {
	yyyymmdd := util.GenYYYYMMDD(time.Now().Unix())
	fmt.Printf("YYYYMMDD:%v\n", yyyymmdd)
}

func TestGenMd5(t *testing.T) {
	var v string = "kevin meng"
	fmt.Printf("kevin meng md5sum:%v\n", util.GenMd5(&v))
}

func TestHASH(t *testing.T) {
	var v string = "kevin meng"
	fmt.Printf("kevin meng hase32:%v, hash64:%v\n", util.HASH32(&v), util.HASH64(&v))
}

/////////////////////////////////////////////////////////////////////////////
//TODO
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
