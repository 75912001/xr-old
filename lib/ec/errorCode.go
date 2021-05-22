package ec

import "fmt"

// error code
const (
	ECSucess = iota
	ECLink
	ECSYS
	ECParam
	ECPacket
	ECTimeOut
	ECChanFull
	ECChanEmpty
	ECOutOfRange
	ECInvValue
	ECConflict
	ECType
	ECInvPointer
	ECUnknown
	ECNonExistent
	ECPBMessageIdExistent
	ECPBMessageIdNonExistent
	ECPBUnmarshal
	ECMAX
	//0xff
)

// defines error information
type errorInfo struct {
	code        int
	name        string
	description string
}

//return:desc
func (p *errorInfo) getDescription() string {
	return p.description
}

//return:[name:code]desc
func (p *errorInfo) getDetail() string {
	return fmt.Sprintf("[%v:%v]%v", p.name, p.code, p.description)
}

var errorInformation = []errorInfo{
	ECSucess:                 {ECSucess, "ECSucess", "sucess"},
	ECLink:                   {ECLink, "ECLink", "link error"},
	ECSYS:                    {ECSYS, "ECSYS", "system error"},
	ECParam:                  {ECParam, "ECParam", "parameter error"},
	ECPacket:                 {ECPacket, "ECPacket", "data pack error"},
	ECTimeOut:                {ECTimeOut, "ECTimeOut", "time out"},
	ECChanFull:               {ECChanFull, "ECChanFull", "chan full"},
	ECChanEmpty:              {ECChanEmpty, "ECChanEmpty", "chan empty"},
	ECOutOfRange:             {ECOutOfRange, "ECOutOfRange", "value out of range"},
	ECInvValue:               {ECInvValue, "ECInvValue", "invalid value"},
	ECConflict:               {ECConflict, "ECConflict", "conflict"},
	ECType:                   {ECType, "ECType", "type mismatch"},
	ECInvPointer:             {ECInvPointer, "ECInvPointer", "invalid pointer"},
	ECUnknown:                {ECUnknown, "ECUnknown", "Unknown"},
	ECNonExistent:            {ECNonExistent, "ECNonExistent", "non-existent"},
	ECPBMessageIdExistent:    {ECPBMessageIdExistent, "ECPBMessageIdExistent", "message id existent"},
	ECPBMessageIdNonExistent: {ECPBMessageIdNonExistent, "ECPBMessageIdNonExistent", "message id non-existent"},
	ECPBUnmarshal:            {ECPBUnmarshal, "ECPBUnmarshal", "message id unmarshal"},
}

func init() {
	for errorCode := ECSucess; errorCode < ECMAX; errorCode++ {
		err := Register(errorCode, errorInformation[errorCode].name, errorInformation[errorCode].description)
		if err != nil {
			panic(err)
		}
	}
}
