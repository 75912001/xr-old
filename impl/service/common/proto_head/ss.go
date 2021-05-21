package proto_head

import (
	"encoding/binary"
)

//PacketLength 包总长度
type PacketLength uint32

//MessageID 消息ID
type MessageID uint32

//SessionID 会话ID
type SessionID uint32

//ResultID 结果ID
type ResultID uint32

//UserID 玩家ID
type UserID uint64

//GProtoHeadLength 包头长度
const GProtoHeadLength uint32 = 24

//ProtoHead 协议包头
type ProtoHead struct {
	PacketLength PacketLength //总包长度,包含包头＋包体长度
	MessageID    MessageID    //消息号
	SessionID    SessionID    //会话id
	ResultID     ResultID     //结果id
	UserID       UserID       //用户id
}

////////////////////////////////////////////////////////////////////////////////
//获取协议包头长度
func GetPacketLength(buf []byte) (packetLength PacketLength) {
	packetLength = PacketLength(binary.LittleEndian.Uint32(buf[0:4]))
	return
	//buf1 := bytes.NewBuffer(buf[0:4])
	//binary.Read(buf1, binary.LittleEndian, &packetLength)
}

//获取协议包头
func GetProtoHead(buf []byte) (packetLength PacketLength, messageID MessageID, sessionID SessionID, resultID ResultID, userID UserID) {
	packetLength = PacketLength(binary.LittleEndian.Uint32(buf[0:4]))
	messageID = MessageID(binary.LittleEndian.Uint32(buf[4:8]))
	sessionID = SessionID(binary.LittleEndian.Uint32(buf[8:12]))
	resultID = ResultID(binary.LittleEndian.Uint32(buf[12:16]))
	userID = UserID(binary.LittleEndian.Uint64(buf[16:24]))

	return
}
