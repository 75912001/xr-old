package handle_event

import (
	"fmt"

	"github.com/75912001/xr/impl/service/login"

	"github.com/75912001/xr/impl/service/common/proto_head"
	"github.com/75912001/xr/lib/tcp"
)

func OnEventConnServer(remote *tcp.Remote) int {
	_ = login.GWorldMgr.Add(remote)
	return 0
}

func OnEventDisConnServer(remote *tcp.Remote) int {
	if !remote.IsConn() {
		return 0
	}
	world := login.GWorldMgr.Find(remote)
	if nil == world {
		login.GServer.Log.Error("find world err:", remote)
		return 0
	}
	login.GWorldMgr.DelById(world.Id)
	login.GWorldMgr.Del(remote)
	return 0
}

func OnEventPacketServer(remote *tcp.Remote, data []byte) int {
	world := login.GWorldMgr.Find(remote)
	if world == nil {
		login.GServer.Log.Crit("OnEventPacketServer remote err.")
		return 0
	}

	ph := &proto_head.ProtoHead{}
	ph.PacketLength, ph.MessageID, ph.SessionID, ph.ResultID, ph.UserID = proto_head.GetProtoHead(data)
	login.GPbFunMgr.OnRecv(uint32(ph.MessageID), ph, data[proto_head.GProtoHeadLength:], world)
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
