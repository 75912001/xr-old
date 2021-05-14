package type_define

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//PacketLength 包总长度
type PACKET_LENGTH_T uint32

//SessionID 会话ID
type SESSION_ID_T uint32

//MessageID 消息ID
type MESSAGE_ID_T uint32

//ResultID 结果ID
type RESULT_ID_T uint32

//UserID 玩家ID
type USER_ID_T uint64

//GProtoHeadLength 包头长度
var GProtoHeadLength = 24
