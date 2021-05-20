package handle_event

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/75912001/xr/impl/service/login"

	"github.com/75912001/xr/impl/service/common/proto_head"
	"github.com/75912001/xr/lib/tcp"
)

func OnEventConnServer(remote *tcp.Remote) int {
	//TODO 业务逻辑
	return 0
}

func OnEventDisConnServer(remote *tcp.Remote) int {
	if !remote.IsConn() {
		return 0
	}
	//TODO 业务逻辑
	return 0
}

func OnEventPacketServer(remote *tcp.Remote, data []byte) int {
	//TODO 业务逻辑

	var ph proto_head.ProtoHead
	ph.PacketLength = 24
	ph.MessageID = 1
	ph.SessionID = 2
	ph.ResultID = 3
	ph.UserID = 4

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, ph.PacketLength)
	binary.Write(buf, binary.LittleEndian, ph.MessageID)
	binary.Write(buf, binary.LittleEndian, ph.SessionID)
	binary.Write(buf, binary.LittleEndian, ph.ResultID)
	binary.Write(buf, binary.LittleEndian, ph.UserID)

	remote.Send(buf.Bytes())
	return 0
}

func OnParseProtoHeadServer(data []byte, length int) int {
	if uint32(length) < proto_head.GProtoHeadLength {
		//长度不足一个包头的长度大小
		return 0
	}
	packetLength := int(proto_head.GetPacketLength(data))
	if uint32(packetLength) < proto_head.GProtoHeadLength {
		login.GServer.Log.Error(fmt.Sprintf("packetLength:%v", packetLength))
		return -1
	}
	if login.GServer.BenchMgr.Json.Base.PacketLengthMax < uint32(packetLength) {
		login.GServer.Log.Error(fmt.Sprintf("PacketLengthMax:%v, packetLength:%v",
			login.GServer.BenchMgr.Json.Base.PacketLengthMax, packetLength))
		return -1
	}

	if length < int(packetLength) {
		return 0
	}

	return packetLength
}
