package addr

import (
	"context"
	"encoding/json"
	"log"
)

//TODO [improvement] 移除 服务信息

//会收到除了自己的组播信息
type OnAddrFunc func(name string, id uint32, ip string, port uint16, data string) int

//添加组播事件
type AddrEvent struct {
	Addr     *Addr
	AddrJson AddrJson
}

type Addr struct {
	OnAddr          OnAddrFunc
	addrChan        chan<- interface{} //服务处理的事件
	serverMap       serverNameMap      //服务器地址信息
	addrFirstBuffer string             //同步的服务器地址信息(发送数据)标记第一次发送数据
	addrBuffer      string             //同步的服务器地址信息(发送数据)
	selfAddr        AddrJson           //自己服务器地址信息
	multicast       multicast
}

//multicastIP:239.0.0.8
//multicastPort:8890
//netName:eth0
func (p *Addr) Start(eventChan chan<- interface{}, onAddrFunc OnAddrFunc,
	multicastIP string, multicastPort uint16, netName string,
	addrName string, addrID uint32, addrIP string, addrPort uint16, addrData string) (err error) {
	p.serverMap = make(serverNameMap)

	p.addrChan = eventChan
	p.OnAddr = onAddrFunc

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
			log.Printf("json Marshal err:%v", err)
			return err
		}
		p.addrFirstBuffer = string(data)
	}
	{
		aj.Cmd = 1
		data, err := json.Marshal(aj)
		if err != nil {
			log.Printf("json Marshal err:%v", err)
			return err
		}
		p.addrBuffer = string(data)
	}

	err = p.multicast.start(context.Background(), multicastIP, multicastPort, netName, p)
	if err != nil {
		log.Printf("multicast start err:%v", err)
		return err
	}
	return
}

func (p *Addr) Exit() {
	p.multicast.exit()
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

type AddrJson struct {
	//cmd:[0,第一次发送]
	//[1,平时发送]
	Cmd  uint32 `json:"cmd"`
	Name string `json:"name"`
	ID   uint32 `json:"id"`
	IP   string `json:"ip"`
	Port uint16 `json:"port"`
	Data string `json:"data"`
}

func (p *Addr) handleAddrMulticast(data []byte) (err error) {
	var aj AddrJson
	err = json.Unmarshal(data, &aj)
	if err != nil {
		log.Printf("json Marshal err:%v, data:%v", err, data)
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

type serverIDMap map[uint32]AddrJson
type serverNameMap map[string]serverIDMap

//添加到内存中
func (p *Addr) add(name string, id uint32, aj *AddrJson) {
	_, valid := p.serverMap[name]
	if valid {
		p.serverMap[name][id] = *aj
	} else {
		serverIDMap := make(serverIDMap)
		serverIDMap[id] = *aj
		p.serverMap[name] = serverIDMap
	}
}

func (p *Addr) find(name string, id uint32) (aj *AddrJson) {
	value, valid := p.serverMap[name]
	if valid {
		value2, valid2 := value[id]
		if valid2 {
			return &value2
		}
	}
	return
}
