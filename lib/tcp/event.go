package tcp

//解析协议包头 返回长度:完整包总长度  返回0:不是完整包 返回-1:包错误
type OnParseProtoHeadType func(data []byte, length int) int

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//断开链接事件
type DisConnEventClient struct {
	Client *Client
}

//处理断开链接事件
type OnDisConnClientType func(client *Client) int

//数据包事件
type PacketEventClient struct {
	Data   []byte
	Client *Client
}

//处理数据包事件
type OnPacketClientType func(client *Client, data []byte) int

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//链接成功事件
type ConnEventServer struct {
	Server *Server
	Remote *Remote
}

//处理链接成功事件
type OnConnServerType func(remote *Remote) int

//断开链接事件
type DisConnEventServer struct {
	Server *Server
	Remote *Remote
}

//处理断开链接事件
type OnDisConnServerType func(remote *Remote) int

//数据包事件
type PacketEventServer struct {
	Server *Server
	Data   []byte
	Remote *Remote
}

//处理数据包事件
type OnPacketServerType func(remote *Remote, data []byte) int
