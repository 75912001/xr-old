package bench

type Mgr struct {
	Json benchJson
}

type benchJson struct {
	DB struct {
		IP   string `json:"ip"`
		Port uint16 `json:"port"`
	} `json:"db"`
}
