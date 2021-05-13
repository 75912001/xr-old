package tcp

//解析协议包头 返回长度:完整包总长度  返回0:不是完整包 返回-1:包错误
type OnParseProtoHeadFunc func(data []byte, length int) int

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//client
//数据包事件
type EventPacketClient struct {
	Client *Client
	Data   []byte
}

//处理数据包事件
type OnEventPacketClientFunc func(client *Client, data []byte) int

//断开链接事件
type EventDisConnClient struct {
	Client *Client
}

//处理断开链接事件
type OnEventDisConnClientFunc func(client *Client) int

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//server
//链接成功事件
type EventConnServer struct {
	Server *Server
	Remote *Remote
}

//处理链接成功事件
type OnEventConnServerFunc func(remote *Remote) int

//数据包事件
type EventPacketServer struct {
	Server *Server
	Remote *Remote
	Data   []byte
}

//处理数据包事件
type OnEventPacketServerFunc func(remote *Remote, data []byte) int

//断开链接事件
type EventDisConnServer struct {
	Server *Server
	Remote *Remote
}

//处理断开链接事件
type OnEventDisConnServerFunc func(remote *Remote) int
