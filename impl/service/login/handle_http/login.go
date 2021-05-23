package handle_http

import (
	"encoding/json"
	"fmt"
	"github.com/75912001/xr/impl/service/login/handle_event"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/75912001/xr/impl/service/protobuf/login_proto"
	"github.com/75912001/xr/lib/util"

	"github.com/75912001/xr/impl/service/login"
)

type loginJsonPart struct {
	Platform uint32 `json:"platform"`
	Account  string `json:"account"`
	Verify   string `json:"verify"`
}

type gatewayJson struct {
	Ip        string `json:"ip"`
	Port      uint16 `json:"port"`
	Session   string `json:"session"`
	ErrorCode uint32 `json:"errorcode"` //0:正常,其他失败(1:失败. 2:access_token过期)
}

/*
2017/07/06 04:04:36 [trace][/home/meng/work/project_1/trunk/login/login.go][33][main.loginHttpHandler]
&{POST /login HTTP/1.1 1 1 map[
Accept-Encoding:[identity]
Content-Type:[application/json]
User-Agent:[Dalvik/1.6.0 (Linux; U; Android 4.4.4; EC6108V9A_pub_hnylt Build/KTU84Q)]
Connection:[Keep-Alive]
Content-Length:[85]] 0xc820532980 85 [] false 139.196.55.173:22501 map[] map[] <nil> map[] 202.99.114.62:9417 /login <nil> <nil>}

2017/07/06 04:04:36 [trace][/home/meng/work/project_1/trunk/login/login.go][72][main.loginHttpHandler]
loginHttpHandler loginJson: {1 39100000117694 6d29911253ca8f22b6f9e06e222a738b}
*/
func LoginHttpHandler(w http.ResponseWriter, req *http.Request) {
	login.GServer.Log.Trace("req:", req)
	var gj gatewayJson

	defer func() {
		js, _ := json.Marshal(gj)
		w.Write(js)

		login.GServer.Log.Trace(gj)
	}()

	{ //解析参数
		err := req.ParseForm()
		if nil != err {
			login.GServer.Log.Error("LoginHttpHandler err.")
			gj.ErrorCode = 1
			return
		}
	}

	if "POST" != req.Method {
		login.GServer.Log.Error("LoginHttpHandler err req.Method:", req.Method)
		gj.ErrorCode = 1
		return
	}

	result, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	login.GServer.Log.Trace(req.Body)
	login.GServer.Log.Trace(string(result))
	login.GServer.Log.Trace([]byte(result))

	//json str 转struct
	var lj loginJsonPart
	err := json.Unmarshal([]byte(result), &lj)
	if nil != err {
		login.GServer.Log.Error(fmt.Sprintf("LoginHttpHandler loginJson err:%v, loginJson:%v", err, lj))
		gj.ErrorCode = 1
		return
	}

	if 0 == len(lj.Account) {
		login.GServer.Log.Error(fmt.Sprintf("LoginHttpHandler Account empty:%v", lj))
		gj.ErrorCode = 1
		return
	}

	login.GServer.Log.Trace(fmt.Sprintf("LoginHttpHandler loginJson:%v", lj))

	//////////////////////////////////////////////////////////////////////
	var newAccount string
	newAccount = lj.Account

	if login.GBench.Json.Platform != lj.Platform {
		login.GServer.Log.Error(fmt.Sprintf("LoginHttpHandler platform err:%v, %v", login.GBench.Json.Platform, lj.Platform))
		gj.ErrorCode = 1
		return
	}

	worldService := login.GWorldMgr.GetWorldService(lj.Platform, newAccount)
	if worldService == nil {
		login.GServer.Log.Error("LoginHttpHandler worldService empty")
		gj.ErrorCode = 1
		return
	}
	//////////////////////////////////////////////////////////////////////
	var session string
	{ //检查签名,生成session
		var verifyString string
		verifyString = genVerify(lj.Platform, lj.Account)
		if lj.Verify != verifyString {
			login.GServer.Log.Error("LoginHttpHandler genVerify err.")
			gj.ErrorCode = 1
			return
		}
		verifyString += fmt.Sprint(time.Now().UnixNano())
		session = util.GenMd5([]byte(verifyString))
	}

	{ //通知对应的服务器
		res := &login_proto.LoginMsgRes{
			Platform: lj.Platform,
			Account:  newAccount,
			Session:  session,
		}
		v := &handle_event.LoginMsgRes{
			Value:  res,
			Remote: worldService.Remote,
		}
		login.GServer.Push2EventChan(v)

		//返回http消息
		gj.Ip = worldService.IP
		gj.Port = worldService.Port
		gj.Session = session
	}
}
