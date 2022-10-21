package models

type RpcNep11Transfers struct {
	Address  string             `json:"address"`
	Sent     []RpcNep11Transfer `json:"sent"`
	Received []RpcNep11Transfer `json:"received"`
}

type RpcNep11Transfer struct {
	Timestamp           uint64 `json:"timestamp"`
	AssetHash           string `json:"assethash"`
	TransferAddress     string `json:"transferaddress"`
	Amount              string `json:"amount"`
	BlockIndex          uint32 `json:"blockindex"`
	TransferNotifyIndex uint32 `json:"transfernotifyindex"`
	TxHash              string `json:"txhash"`
	TokenId             string `json:"tokenid"`
}
