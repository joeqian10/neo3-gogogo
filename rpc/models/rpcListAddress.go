package models

type RpcListAddress struct {
	AddressList []RpcAddress
}

type RpcAddress struct {
	Address   string `json:"address"`
	HasKey    bool   `json:"haskey"`
	Label     string `json:"label"`
	WatchOnly bool   `json:"watchonly"`
}

type ValidateAddress struct {
	Address string `json:"address"`
	IsValid bool   `json:"isvalid"`
}
