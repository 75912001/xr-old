package util

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestGetFuncName(t *testing.T) {
	fmt.Printf("FuncName:%v\n", GetFuncName())
}

func TestGenYYYYMMDD(t *testing.T) {
	yyyymmdd := GenYYYYMMDD(time.Now().Unix())
	fmt.Printf("YYYYMMDD:%v\n", yyyymmdd)
}

func TestGenMd5(t *testing.T) {
	var v string = "kevin meng"
	fmt.Printf("kevin meng md5sum:%v\n", GenMd5(&v))
}

func TestHASH(t *testing.T) {
	var v string = "kevin meng"
	fmt.Printf("kevin meng hase32:%v, hash64:%v\n", HASH32(&v), HASH64(&v))
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
	/*
			var strJson string = `
		{
		    "appConfigMaps": [
		        {
		            "component": "sdsPilot",
		            "version": "sdspilot-r.2.0.1.0",
		            "namespace": ["urn:mavenir:npn-config-mgmt:1.0"],

		            "configApplyPriority": 0,
		            "installPriority": 0,
		            "InstallEnabled": 0
		        },
		         {
		            "component": "k8sCluster",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ns:yang:k8sCluster:1.2.4.0"],
		            "bridgeTemplate": [
		                "bridge_E1.tmpl",
		                "bridge_N2.tmpl",
		                "bridge_N4C.tmpl",
		                "bridge_N4U.tmpl",
		                "bridge_N26.tmpl",
		                "bridge_RX.tmpl",
		                "bridge_S5C.tmpl",
		                "bridge_UPF.tmpl"
		            ],
		            "bridgeConfigScript": "bridgeConfig.sh",
		            "preInstallScript": "vdu_preInstall.sh",
		            "installScript": "installK8SCluster.sh",

		            "configApplyPriority": 1,
		            "installPriority": 1,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "eck-operator",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ns:yang:eck-operator:1.2.4.0"],

		            "installScript": "mwpChartInstall.sh",
		            "postInstallScript": "",

		            "configApplyPriority": 2,
		            "installPriority": 2,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "elastic-stack",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ns:yang:elastic-stack:1.2.4.0"],

		            "installScript": "mwpChartInstall.sh",

		            "configApplyPriority": 3,
		            "installPriority": 3,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "fluent-bit",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ns:yang:fluent-bit:1.2.4.0"],

		            "installScript": "mwpChartInstall.sh",

		            "configApplyPriority": 4,
		            "installPriority": 4,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "istio-operator",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ns:yang:istio-operator:1.2.4.0"],

		            "installScript": "mwpChartInstall.sh",

		            "configApplyPriority": 5,
		            "installPriority": 5,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "istio-stack",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ns:yang:istio-stack:1.2.4.0"],

		            "installScript": "mwpChartInstall.sh",

		            "configApplyPriority": 6,
		            "installPriority": 6,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "kube-prometheus-stack",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ns:yang:kube-prometheus-stack:1.2.4.0"],

		            "installScript": "mwpChartInstall.sh",

		            "configApplyPriority": 7,
		            "installPriority": 7,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "logstash",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ns:yang:logstash:1.2.4.0"],

		            "installScript": "mwpChartInstall.sh",

		            "configApplyPriority": 8,
		            "installPriority": 8,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "strimzi-kafka-operator",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ns:yang:strimzi-kafka-operator:1.2.4.0"],

		            "installScript": "mwpChartInstall.sh",

		            "configApplyPriority": 9,
		            "installPriority": 9,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "strimzi-kafka-stack",
		            "version": "1.2.4.2",
		            "installScript": "mwpChartInstall.sh",
		            "configApplyPriority": 10,
		            "installPriority": 10,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "traefik",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ns:yang:traefik:1.2.4.0"],

		            "installScript": "mwpChartInstall.sh",

		            "configApplyPriority": 11,
		            "installPriority": 11,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "mtcil",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ran:mavenir_data_container"],

		            "installScript": "mwpChartInstall.sh",

		            "configApplyScript": "mtcil_configApply.sh",

		            "configApplyPriority": 20,
		            "installPriority": 20,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "SRIOV",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ran:sriov"],

		            "postBridgeConfigScript": "config-sriov.sh",
		            "preInstallScript": "sriov_preInstall.sh",
		            "installScript": "sriov_install.sh",

		            "configApplyPriority": 23,
		            "installPriority": 23,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "5GC-CRDL",
		            "version": "1.2.4.2",
		            "namespace": ["urn:mavenir:ns:yang:5gc-crdl:1.2.4.0"],

		            "installScript": "5gc_crdl_install.sh",
		            "postInstallScript": "5gc_crdl_postInstall.sh",

		            "configApplyPriority": 24,
		            "installPriority": 24,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "udm",
		            "version": "1.2.4.2",

		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 26,
		            "installPriority": 26,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "udr",
		            "version": "1.2.4.2",

		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 27,
		            "installPriority": 27,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "amf",
		            "version": "1.2.4.2",
					"namespace": [
						"urn:com:mavenir:_5gc:yang:mavenir-amf-profile:",
						"urn:com:mavenir:_5gc:yang:mavenir-amf:"
					],

		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 28,
		            "installPriority": 28,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "smf",
		            "version": "1.2.4.2",
		            "namespace": [
						"urn:com:mavenir:_5gc:yang:mavenir-smf-profile:",
						"urn:com:mavenir:_5gc:yang:mavenir-smf:"
					],
		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 29,
		            "installPriority": 29,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "upf",
		            "version": "1.2.4.2",
					"namespace": [
						"urn:com:mavenir:_5gc:yang:mavenir-upf-profile:",
						"urn:com:mavenir:_5gc:yang:mavenir-upf:"
					],
		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 30,
		            "installPriority": 30,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "ksync",
		            "version": "1.2.4.2",

		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 31,
		            "installPriority": 31,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "udsf",
		            "version": "1.2.4.2",
					"namespace": [
						"urn:com:mavenir:_5gc:yang:mavenir-udsf:"
					],
		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 33,
		            "installPriority": 33,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "ausf",
		            "version": "1.2.4.2",
					"namespace": [
						"urn:com:mavenir:_5gc:yang:mavenir-ausf:"
					],
		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 34,
		            "installPriority": 34,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "nrf",
		            "version": "1.2.4.2",
					"namespace": [
						"urn:com:mavenir:_5gc:yang:mavenir-nrf:"
					],
		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 35,
		            "installPriority": 35,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "nssf",
		            "version": "1.2.4.2",
					"namespace": [
						"urn:com:mavenir:_5gc:yang:mavenir-nssf:"
					],
		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 36,
		            "installPriority": 36,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "sp",
		            "version": "1.2.4.2",
					"namespace": [
						"urn:com:mavenir:_5gc:yang:mavenir-sp:"
					],
		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 37,
		            "installPriority": 37,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "cucp",
		            "version": "1.2.4.2",
					"namespace": [
						"urn:mavenir:ns:yang:cucp:gnb:"
					],
		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 38,
		            "installPriority": 38,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "cuup",
		            "version": "1.2.4.2",
					"namespace": [
						"urn:mavenir:ns:yang:cuup:gnb"
					],
		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 40,
		            "installPriority": 40,
		            "InstallEnabled": 1
		        },
		        {
		            "component": "du",
		            "version": "1.2.4.2",
					"namespace": [
						"urn:mavenir:ns:yang:du:gnb:"
					],
		            "installScript": "chartInstall.sh",

		            "configApplyPriority": 41,
		            "installPriority": 41,
		            "InstallEnabled": 1
		        }

		    ]
		}
		`
	*/
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
