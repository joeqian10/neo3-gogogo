package models

type RpcNep17Transfers struct {
	Address  string             `json:"address"`
	Sent     []RpcNep17Transfer `json:"sent"`
	Received []RpcNep17Transfer `json:"received"`
}

type RpcNep17Transfer struct {
	Timestamp           uint64 `json:"timestamp"`
	AssetHash           string `json:"assethash"`
	TransferAddress     string `json:"transferaddress"`
	Amount              string `json:"amount"`
	BlockIndex          uint32 `json:"blockindex"`
	TransferNotifyIndex uint32 `json:"transfernotifyindex"`
	TxHash              string `json:"txhash"`
}
