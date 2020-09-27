package tcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
	"time"

	"github.com/75912001/xr/lib/log"
)

/*
 go test -v -count=1
*/

var eventChan = make(chan interface{}, 128)
var recvSumCount int
var sendSumCount int

//go test -v -test.run TestClient_Connect
func TestClient_Connect(t *testing.T) {
	log.GLog = &log.Log{}
	log.GLog.Init("connect_")
	log.GLog.SetLevel(7)

	go handleEventChan()

	for i := 0; i < 100; i++ {
		var client Client
		client.Connect("127.0.0.1", 6666, 1024, eventChan, OnParseProtoHeadFun, OnCloseConnectFun, OnPacketFun)
		client.DisConnect()
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
	client.Connect("127.0.0.1", 6666, 1024, eventChan, OnParseProtoHeadFun, OnCloseConnectFun, OnPacketFun)

	for i := 0; i < 100000; i++ {
		client.Send(buf8.Bytes())
		client.Send(buf128.Bytes())
		client.Send(buf512.Bytes())
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		log.GLog.Debug(fmt.Sprintf("%v", i))
	}
	client.DisConnect()
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

	client.Connect("127.0.0.1", 6666, 102400, eventChan, OnParseProtoHeadFun, OnCloseConnectFun, OnPacketFun)

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
	client.DisConnect()
	log.GLog.Trace(fmt.Sprintf("recv sum count:%v", recvSumCount))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func handleEventChan() (err error) {
	//处理数据
	for v := range eventChan {
		switch v.(type) {
		case *CloseConnectEventChanClient:
			vv, ok := v.(*CloseConnectEventChanClient)
			if ok {
				log.GLog.Trace(fmt.Sprintf("CloseConnectEventChanClient."))
				vv.Client.DisConnect()
			} else {
				log.GLog.Crit("CloseConnectEventChanClient type error.")
			}
		case *RecvEventChanClient:
			vv, ok := v.(*RecvEventChanClient)
			if ok {
				log.GLog.Trace(fmt.Sprintf("RecvEventChanClient."))
				vv.Client.OnPacket(vv.Client, vv.Buf)
			} else {
				log.GLog.Crit("RecvEventChanClient type error.")
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
func OnCloseConnectFun(client *Client) int {
	log.GLog.Trace(fmt.Sprintf("OnCloseConnect. remote close connect, server ip:%v", client.server.conn.RemoteAddr().String()))
	//TODO
	return 0
}

//远端包
func OnPacketFun(client *Client, buf []byte) int {
	log.GLog.Trace(fmt.Sprintf("OnPacket. remote packet, server ip:%v, len:%v", client.server.conn.RemoteAddr().String(), len(buf)))
	//TODO
	recvSumCount += len(buf)
	return 0
}
