package tcp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/75912001/xr/lib/addrmulticase"

	"github.com/75912001/xr/lib/log"
	"github.com/75912001/xr/lib/tcp"
)

var eventChan = make(chan interface{}, 10000)
var GLog *log.Log

func init() {
	GLog = new(log.Log)
	GLog.Init("test_log")
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//server
func OnConnServer(remote *tcp.Remote) int {
	//	fmt.Println("OnConnServer")
	return 0
}

func OnDisConnServer(remote *tcp.Remote) int {
	fmt.Println("OnDisconnServer")
	if !remote.IsConn() {
		GLog.Warn("duplicate shutdowns")
		return 0
	}
	return 0
}

func OnPacketServer(remote *tcp.Remote, data []byte) int {
	fmt.Println("OnPacketServer")
	return 0
}

func OnParseProtoHeadServer(data []byte, length int) int {
	//解析协议包头 返回长度:完整包总长度  返回0:不是完整包 返回-1:包错误
	fmt.Println("OnParseProtoHead")
	return len(data)
	//return 0
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//client
func OnDisConnClient(client *tcp.Client) int {
	fmt.Println("OnDisconnClient")
	if !client.Remote.IsConn() {
		GLog.Warn("duplicate shutdowns")
		return 0
	}
	return 0
}

func OnPacketClient(client *tcp.Client, data []byte) int {
	fmt.Println("OnPacketClient")
	return 0
}

func OnParseProtoHeadClient(data []byte, length int) int {
	//解析协议包头 返回长度:完整包总长度  返回0:不是完整包 返回-1:包错误
	fmt.Println("OnParseProtoHeadClient")
	return len(data)
}

func handleEvent() {
	for v := range eventChan {
		switch v.(type) {
		//server
		case *tcp.ConnEventServer:
			vv, ok := v.(*tcp.ConnEventServer)
			if ok {
				GLog.Debug(fmt.Sprintf("ConnEventServer, remote:%v", vv.Remote))
				vv.Server.OnConn(vv.Remote)

			}
		case *tcp.DisConnEventServer:
			vv, ok := v.(*tcp.DisConnEventServer)
			if ok {
				GLog.Debug(fmt.Sprintf("DisConnEventServer, remote:%v", vv.Remote))
				vv.Server.OnDisConn(vv.Remote)
				if vv.Remote.IsConn() {
					vv.Remote.Close()
				}
			}
		case *tcp.PacketEventServer:
			vv, ok := v.(*tcp.PacketEventServer)
			if ok {
				if !vv.Remote.IsConn() {
					continue
				}
				vv.Server.OnPacket(vv.Remote, vv.Data)
			}
			//client
		case *tcp.DisConnEventClient:
			vv, ok := v.(*tcp.DisConnEventClient)
			if ok {
				GLog.Debug(fmt.Sprintf("DisconnEventClient, remote:%v", vv.Client.Remote))
				vv.Client.OnDisConn(vv.Client)
				if vv.Client.Remote.IsConn() {
					vv.Client.Remote.Close()
				}
			}
		case *tcp.PacketEventClient:
			vv, ok := v.(*tcp.PacketEventClient)
			if ok {
				if !vv.Client.Remote.IsConn() {
					continue
				}
				vv.Client.OnPacket(vv.Client, vv.Data)
			}
		//addrMulticase
		case *addrmulticase.AddrMulticast:

		default:
			GLog.Crit(fmt.Sprintf("non-existent event:%v", v))
		}
	}
}

func TestServer(t *testing.T) {
	defer func() {
		GLog.Exit()
	}()

	var s tcp.Server
	go func() {
		defer func() {
			if err := recover(); err != nil {
				GLog.Warn(fmt.Sprintf("handleEvent goroutine panic:%v", err))
			}
			GLog.Trace("handleEvent goroutine done.")
		}()
		handleEvent()
	}()

	address := "127.0.0.1:8787"
	rwBuffLen := 1000
	var recvPacketMaxLen uint32 = 1000
	var sendChanCapacity uint32 = 1000
	err := s.Strat(address, GLog, rwBuffLen, recvPacketMaxLen, eventChan,
		OnConnServer, OnDisConnServer, OnPacketServer, OnParseProtoHeadServer, sendChanCapacity)
	if err != nil {
		t.Fatalf("server start err:%v", err)
		return
	}

	for {
		time.Sleep(time.Second)
		sendChanCnt, recvChanCnt := s.Info()
		fmt.Println(fmt.Sprintf("server sendChanCnt:%v, recvChanCnt:%v", sendChanCnt, recvChanCnt))
	}
}

func TestClient(t *testing.T) {
	defer func() {
		GLog.Exit()
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				GLog.Warn(fmt.Sprintf("handleEvent goroutine panic:%v", err))
			}
			GLog.Trace("handleEvent goroutine done.")
		}()
		handleEvent()
	}()

	address := "127.0.0.1:8787"
	rwBuffLen := 1000
	var recvPacketMaxLen uint32 = 1000
	var sendChanCapacity uint32 = 1000
	for i := 0; i < 10000; i++ {
		var c tcp.Client
		err := c.Connect(address, GLog, rwBuffLen, recvPacketMaxLen,
			eventChan, OnDisConnClient, OnPacketClient, OnParseProtoHeadClient, sendChanCapacity)
		if err != nil {
			t.Fatalf("server start err:%v", err)
			return
		}
	}
	for {
		time.Sleep(time.Second)
	}
}

//func CheckPort(port string) error {
//	var err error
//
//	tcpAddress, err := net.ResolveTCPAddr("tcp4", ":"+port)
//	if err != nil {
//		return err
//	}
//
//	for i := 0; i < 3; i++ {
//		listener, err := net.ListenTCP("tcp", tcpAddress)
//		if err != nil {
//			time.Sleep(time.Duration(100) * time.Millisecond)
//			if i == 3 {
//				return err
//			}
//			continue
//		} else {
//			listener.Close()
//			break
//		}
//	}
//
//	return nil
//}

///////////////////////////////////
//client

/*
 go test -v -count=1
*/
/*
var eventChan = make(chan interface{}, 128)
var recvSumCount int
var sendSumCount int

//go test -v -test.run TestClient_Connect
func TestClient_Connect(t *testing.T) {
	log.GLog = &log.Log{}
	log.GLog.Init("connect_")

	go handleEventChan()

	for i := 0; i < 1000; i++ {
		var client Client
		client.Connect("127.0.0.1:6666", 1024, eventChan, OnParseProtoHeadFun, OnDisconnectFun, OnPacketFun, 1024)
		time.Sleep(time.Nanosecond * 5) //立即disconnect, 可能会在关闭connect后,才调用协程中的conn.Read, 这时conn为nil.
		time.Sleep(time.Nanosecond * 5)
		client.DisConn()
	}

	for i := 0; i < 2; i++ {
		time.Sleep(time.Second)
	}
}

//go test -v -test.run TestClient_Send
func TestClient_Send(t *testing.T) {
	log.GLog = &log.Log{}
	log.GLog.Init("send_")
	//log.GLog.SetLevel(7)

	buf8 := new(bytes.Buffer)
	binary.Write(buf8, binary.LittleEndian, uint32(8))
	binary.Write(buf8, binary.LittleEndian, uint32(100))

	buf128 := new(bytes.Buffer)
	binary.Write(buf128, binary.LittleEndian, uint32(128))
	for i := 0; i < 31; i++ {
		binary.Write(buf128, binary.LittleEndian, uint32(i))
	}

	buf512 := new(bytes.Buffer)
	binary.Write(buf512, binary.LittleEndian, uint32(512))
	for i := 0; i < 127; i++ {
		binary.Write(buf512, binary.LittleEndian, uint32(i))
	}

	go handleEventChan()

	var client Client
	client.Connect("127.0.0.1:6666", 1024, eventChan, OnParseProtoHeadFun, OnDisconnectFun, OnPacketFun, 1024)

	for i := 0; i < 100000; i++ {
		client.Send(buf8.Bytes())
		client.Send(buf128.Bytes())
		client.Send(buf512.Bytes())
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		log.GLog.Debug(fmt.Sprintf("%v", i))
	}
	client.DisConn()
}

//go test -v -test.run TestClient_Recv
func TestClient_Recv(t *testing.T) {
	log.GLog = &log.Log{}
	log.GLog.Init("recv_")

	buf8 := new(bytes.Buffer)
	binary.Write(buf8, binary.LittleEndian, uint32(8))
	binary.Write(buf8, binary.LittleEndian, uint32(100))

	buf128 := new(bytes.Buffer)
	binary.Write(buf128, binary.LittleEndian, uint32(128))
	for i := 0; i < 31; i++ {
		binary.Write(buf128, binary.LittleEndian, uint32(i))
	}

	buf512 := new(bytes.Buffer)
	binary.Write(buf512, binary.LittleEndian, uint32(512))
	for i := 0; i < 127; i++ {
		binary.Write(buf512, binary.LittleEndian, uint32(i))
	}

	go handleEventChan()

	var client Client

	client.Connect("127.0.0.1:6666", 1024, eventChan, OnParseProtoHeadFun, OnDisconnectFun, OnPacketFun, 1024)

	for i := 0; i < 100000; i++ {
		client.Send(buf8.Bytes())
		client.Send(buf128.Bytes())
		client.Send(buf512.Bytes())
		sendSumCount += 8 + 128 + 512
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		log.GLog.Debug(fmt.Sprintf("%v", i))
	}
	client.DisConn()
	log.GLog.Trace(fmt.Sprintf("recv sum count:%v", recvSumCount))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func handleEventChan() (err error) {
	//处理数据
	for v := range eventChan {
		switch v.(type) {
		case *DisconnEventClient:
			vv, ok := v.(*DisconnEventClient)
			if ok {
				log.GLog.Trace(fmt.Sprintf("CloseConnectEventChan."))
				vv.Client.OnDisConn()
			} else {
				log.GLog.Crit("CloseConnectEventChan type error.")
			}
		case *PacketEventClient:
			vv, ok := v.(*PacketEventClient)
			if ok {
				if !vv.Client.server.isConn() {
					continue
				}
				log.GLog.Trace(fmt.Sprintf("RecvEventChan."))
				vv.Client.OnPacket(vv.Client, vv.Data)
			} else {
				log.GLog.Crit("RecvEventChan type error.")
			}
		default:
			log.GLog.Crit(fmt.Sprintf("not find event, event:%v", v))
		}
	}
	return err
}

func OnParseProtoHeadFun(buf []byte, length int) int {
	if length < 4 { //长度不足一个包头的长度大小
		return 0
	}

	packetLength := int(parseProtoHeadPacketLength(buf))

	if int(packetLength) < 4 {
		log.GLog.Crit(fmt.Sprintf("PacketLength:%v", packetLength))
		return -1
	}

	if length < int(packetLength) {
		return 0
	}

	return packetLength
}

//解析协议包头长度
func parseProtoHeadPacketLength(buf []byte) (packetLength uint32) {
	buf1 := bytes.NewBuffer(buf[0:4])
	binary.Read(buf1, binary.LittleEndian, &packetLength)
	return packetLength
}

//远端链接关闭
func OnDisconnectFun(client *Client) int {
	log.GLog.Trace(fmt.Sprintf("disconnect, server ip:%v", client.server.conn.RemoteAddr().String()))
	//TODO
	return 0
}

//远端包
func OnPacketFun(client *Client, buf []byte) int {
	log.GLog.Trace(fmt.Sprintf("packet, server ip:%v, len:%v", client.server.conn.RemoteAddr().String(), len(buf)))
	//TODO
	recvSumCount += len(buf)
	return 0
}
*/
