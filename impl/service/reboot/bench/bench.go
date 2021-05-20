package bench

type Mgr struct {
	Json benchJson
}

type benchJson struct {
	Reboot struct {
		ServerIP   string `json:"serverIP"`
		ServerPort uint16 `json:"serverPort"`
		Cnt        uint32 `json:"cnt"`
	} `json:"reboot"`
}
