package models

type RpcNep11Balances struct {
	Address  string            `json:"address"`
	Balances []RpcNep11Balance `json:"balance"`
}

type RpcNep11Balance struct {
	AssetHash string          `json:"assethash"`
	Name      string          `json:"name"`
	Symbol    string          `json:"symbol"`
	Decimals  string          `json:"decimals"`
	Tokens    []RpcNep11Token `json:"tokens"`
}

type RpcNep11Token struct {
	TokenId          string `json:"tokenid"`
	Amount           string `json:"amount"`
	LastUpdatedBlock uint32 `json:"lastupdatedblock"`
}
