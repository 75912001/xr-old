package pb_func_mgr

import (
	"github.com/75912001/xr/lib/ec"
	"google.golang.org/protobuf/proto"
	"log"
)

//ProtoBufFun 协议function
type ProtoBufFun func(protoHead interface{}, protoMessage *proto.Message, obj interface{}) (ret int)

//PbFunMgr 管理器
type PbFunMgr struct {
	pbFunMap ProtoBufFunMap
}

//Init 初始化管理器
func (p *PbFunMgr) Init() {
	p.pbFunMap = make(ProtoBufFunMap)
}

//Register 注册消息
func (p *PbFunMgr) Register(messageID uint32, pbFun ProtoBufFun,
	protoMessage proto.Message) (ret int) {
	{
		pbFunHandle := p.find(messageID)
		if nil != pbFunHandle {
			log.Fatalf("pb func register messageID existent:%v", messageID)
			return ec.ECPBMessageIdExistent
		}
	}
	{
		pb := &pbFunHandle{
			pbFun:        pbFun,
			protoMessage: &protoMessage,
		}
		p.pbFunMap[messageID] = pb
	}

	return 0
}

//OnRecv 收到消息
func (p *PbFunMgr) OnRecv(messageID uint32, protoHead interface{}, bodyBuf []byte, obj interface{}) (ret int) {
	pbFunHandle, ok := p.pbFunMap[messageID]
	if !ok {
		return ec.ECPBMessageIdNonExistent
	}

	err := proto.Unmarshal(bodyBuf, *pbFunHandle.protoMessage)
	if nil != err {
		return ec.ECPBUnmarshal
	}
	return pbFunHandle.pbFun(protoHead, pbFunHandle.protoMessage, obj)
}

type pbFunHandle struct {
	pbFun        ProtoBufFun
	protoMessage *proto.Message
}

//ProtoBufFunMap 协议function map
type ProtoBufFunMap map[uint32]*pbFunHandle

func (p *PbFunMgr) find(messageID uint32) (pbFunHandle *pbFunHandle) {
	pbFunHandle, _ = p.pbFunMap[messageID]
	return pbFunHandle
}
