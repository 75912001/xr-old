package client_mgr

import "github.com/75912001/xr/lib/tcp"

type Client struct {
	Remote *tcp.Remote
	IP     string //TODO 赋值
	Port   uint16 //TODO 赋值
}
