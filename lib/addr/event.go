package addr

//添加组播事件
type EventAddrMulticast struct {
	Addr     *Addr
	AddrJson AddrJson
}

//会收到除了自己的组播信息
type OnEventAddrMulticastFunc func(name string, id uint32, ip string, port uint16, data string) int
