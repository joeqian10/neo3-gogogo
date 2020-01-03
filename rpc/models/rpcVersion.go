package models

type RpcVersion struct {
	TcpPort   int    `json:"tcpPort"`
	WsPort    int    `json:"wsPort"`
	Nonce     string `json:"nonce"`
	UserAgent string `json:"useragent"`
}
