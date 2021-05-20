package bench

type Mgr struct {
	Json benchJson
}

type benchJson struct {
	LoginHttp struct {
		Pattern string `json:"pattern"`
		IP      string `json:"ip"`
		Port    uint16 `json:"port"`
	} `json:"loginHttp"`
	Platform uint32 `json:"platform"`
}
