package addr

import (
	"encoding/json"
	"fmt"

	"github.com/75912001/xr/lib/log"
)

//TODO [improvement] 移除 服务信息

//会收到除了自己的组播信息
type OnAddrType func(name string, id uint32, ip string, port uint16, data string) int

//添加组播事件
type AddrEvent struct {
	Addr     *Addr
	AddrJson addrJson
}

type Addr struct {
	OnAddr          OnAddrType
	addrChan        chan<- interface{} //服务处理的事件
	log             *log.Log
	serverMap       serverNameMap //服务器地址信息
	addrFirstBuffer string        //同步的服务器地址信息(发送数据)标记第一次发送数据
	addrBuffer      string        //同步的服务器地址信息(发送数据)
	selfAddr        addrJson      //自己服务器地址信息
	multicast       multicast
}

//multicastIP:239.0.0.8
//multicastPort:8890
//netName:eth0
func (p *Addr) Start(log *log.Log, eventChan chan<- interface{}, onAddr OnAddrType,
	multicastIP string, multicastPort uint16, netName string,
	addrName string, addrID uint32, addrIP string, addrPort uint16, addrData string) (err error) {
	p.serverMap = make(serverNameMap)

	p.log = log
	p.addrChan = eventChan
	p.OnAddr = onAddr

	p.selfAddr.Cmd = 0
	p.selfAddr.Name = addrName
	p.selfAddr.ID = addrID
	p.selfAddr.IP = addrIP
	p.selfAddr.Port = addrPort
	p.selfAddr.Data = addrData

	aj := p.selfAddr
	{
		data, err := json.Marshal(aj)
		if err != nil {
			p.log.Crit("json Marshal err:", err)
			return err
		}
		p.addrFirstBuffer = string(data)
	}
	{
		aj.Cmd = 1
		data, err := json.Marshal(aj)
		if err != nil {
			p.log.Crit("json Marshal err:", err)
			return err
		}
		p.addrBuffer = string(data)
	}

	err = p.multicast.start(multicastIP, multicastPort, netName, p.log, p)
	if err != nil {
		p.log.Crit("multicast start err:", err)
		return err
	}
	return
}

/*
//json
{
	"cmd":123,
	"name":"loginService",
	"id":1,
	"ip":"127.0.0.1",
	"port":7878,
	"data":"this is data."
}
*/

type addrJson struct {
	//cmd:[0,第一次发送]
	//[1,平时发送]
	Cmd  uint32 `json:"cmd"`
	Name string `json:"name"`
	ID   uint32 `json:"id"`
	IP   string `json:"ip"`
	Port uint16 `json:"port"`
	Data string `json:"data"`
}

func (p *Addr) parse(data []byte) (err error) {
	var aj addrJson
	err = json.Unmarshal(data, &aj)
	if err != nil {
		p.log.Crit(fmt.Sprintf("json Marshal err:%v, data:%v", err, data))
		return
	}
	if p.selfAddr.Name != aj.Name && p.selfAddr.ID != aj.ID {
		if 0 == aj.Cmd {
			p.multicast.doAddrSYN([]byte(p.addrBuffer))
			p.add(aj.Name, aj.ID, &aj)
		} else {
			if nil == p.find(aj.Name, aj.ID) {
				p.multicast.doAddrSYN([]byte(p.addrBuffer))
				p.add(aj.Name, aj.ID, &aj)
			}
		}

		p.addrChan <- &AddrEvent{
			Addr:     p,
			AddrJson: aj,
		}
	}
	return
}

type serverIDMap map[uint32]addrJson
type serverNameMap map[string]serverIDMap

//添加到内存中
func (p *Addr) add(name string, id uint32, aj *addrJson) {
	_, valid := p.serverMap[name]
	if valid {
		p.serverMap[name][id] = *aj
	} else {
		serverIDMap := make(serverIDMap)
		serverIDMap[id] = *aj
		p.serverMap[name] = serverIDMap
	}
}

func (p *Addr) find(name string, id uint32) (aj *addrJson) {
	value, valid := p.serverMap[name]
	if valid {
		value2, valid2 := value[id]
		if valid2 {
			return &value2
		}
	}
	return
}
