package models

type RpcNep17Balances struct {
	Address  string            `json:"address"`
	Balances []RpcNep17Balance `json:"balance"`
}

type RpcNep17Balance struct {
	AssetHash        string `json:"assethash"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Decimals         string `json:"decimals"`
	Amount           string `json:"amount"`
	LastUpdatedBlock uint32 `json:"lastupdatedblock"`
}
