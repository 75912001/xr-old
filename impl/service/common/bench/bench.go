package bench

import (
	"time"

	"github.com/75912001/xr/lib/addr"

	"github.com/75912001/xr/lib/util"
)

type Mgr struct {
	Json benchJson
}

type benchJson struct {
	Base struct {
		ServiceName      string `json:"serviceName"`
		ServiceID        int    `json:"serviceID"`
		LogLevel         int    `json:"logLevel"`
		LogAbsPath       string `json:"logAbsPath"`
		GoMaxProcs       int    `json:"goMaxProcs"`
		EventChanCnt     int    `json:"eventChanCnt"`
		PacketLengthMax  int    `json:"packetLengthMax"`
		SendChanCapacity int    `json:"sendChanCapacity"`

		Comments string `json:"__comments__"`
	} `json:"base"`
	Timer struct {
		ScanSecondDuration      time.Duration `json:"scanSecondDuration"`
		ScanMillisecondDuration time.Duration `json:"scanMillisecondDuration"`
	} `json:"timer"`

	Server struct {
		IP   string `json:"ip"`
		Port string `json:"port"`
	} `json:"server"`

	AddrMulticast addr.AddrJson `json:"addrMulticast"`
}

func (p *Mgr) Parse(pathFile string) (err error) {
	err = util.UnmarshalJsonFile(pathFile, &p.Json)
	if err != nil {
		return
	}
	//log.Printf("bench json:%+v", p.Json)
	return
}
