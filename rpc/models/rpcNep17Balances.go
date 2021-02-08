package models

type RpcNep17Balances struct {
	Balances []RpcNep17Balance `json:"balance"`
	Address  string            `json:"address"`
}

type RpcNep17Balance struct {
	AssetHash        string `json:"assethash"`
	Amount           uint64 `json:"amount"`
	LastUpdatedBlock int    `json:"lastupdatedblock"`
}
