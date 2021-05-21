package proto

import (
	"github.com/75912001/xr/impl/service/common/proto_head"
	"github.com/75912001/xr/lib/tcp"
	"google.golang.org/protobuf/proto"
)

//发送消息
func Send(remote *tcp.Remote, pb proto.Message, messageID proto_head.MessageID,
	sessionID proto_head.SessionID, resultID proto_head.ResultID, userID proto_head.UserID) (err error) {
	var b [proto_head.GProtoHeadLength]byte
	//type MessageID uint32
	b[4] = byte(messageID)
	b[5] = byte(messageID >> 8)
	b[6] = byte(messageID >> 16)
	b[7] = byte(messageID >> 24)
	//type SessionID uint32
	b[8] = byte(sessionID)
	b[9] = byte(sessionID >> 8)
	b[10] = byte(sessionID >> 16)
	b[11] = byte(sessionID >> 24)
	//type ResultID uint32
	b[12] = byte(resultID)
	b[13] = byte(resultID >> 8)
	b[14] = byte(resultID >> 16)
	b[15] = byte(resultID >> 24)
	//type UserID uint64
	b[16] = byte(userID)
	b[17] = byte(userID >> 8)
	b[18] = byte(userID >> 16)
	b[19] = byte(userID >> 24)
	b[20] = byte(userID >> 32)
	b[21] = byte(userID >> 40)
	b[22] = byte(userID >> 48)
	b[23] = byte(userID >> 56)

	msgBuf, err := proto.Marshal(pb)
	if nil != err {
		return
	}

	sendBufAllLength := uint32(len(msgBuf)) + proto_head.GProtoHeadLength

	//type PacketLength uint32
	b[0] = byte(sendBufAllLength)
	b[1] = byte(sendBufAllLength >> 8)
	b[2] = byte(sendBufAllLength >> 16)
	b[3] = byte(sendBufAllLength >> 24)

	msg := b[:]
	msg = append(msg, msgBuf...)
	remote.Send(msg)

	//headBuf := new(bytes.Buffer)
	//binary.Write(headBuf, binary.LittleEndian, sendBufAllLength)
	//binary.Write(headBuf, binary.LittleEndian, messageID)
	//binary.Write(headBuf, binary.LittleEndian, sessionID)
	//binary.Write(headBuf, binary.LittleEndian, resultID)
	//binary.Write(headBuf, binary.LittleEndian, userID)
	return
}
