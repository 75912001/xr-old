package client_mgr

import "github.com/75912001/xr/lib/tcp"

type CLIENT_MAP map[*tcp.Remote]*Client

type ClientMgr struct {
	ClientMap CLIENT_MAP
}

func (this *ClientMgr) Init() {
	this.ClientMap = make(CLIENT_MAP)
}

func (this *ClientMgr) AddUser(remote *tcp.Remote) (client *Client) {
	client = new(Client)

	client.Remote = remote
	this.ClientMap[remote] = client
	return
}

func (this *ClientMgr) DelUser(remote *tcp.Remote) {
	delete(this.ClientMap, remote)
}

func (this *ClientMgr) Find(remote *tcp.Remote) (client *Client) {
	client, _ = this.ClientMap[remote]
	return
}
