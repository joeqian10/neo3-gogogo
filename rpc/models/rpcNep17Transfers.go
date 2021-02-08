package models

type RpcNep17Transfers struct {
	Sent     []RpcNep17Transfer `json:"sent"`
	Received []RpcNep17Transfer `json:"received"`
	Address  string             `json:"address"`
}

type RpcNep17Transfer struct {
	Timestamp           int    `json:"timestamp"`
	AssetHash           string `json:"assethash"`
	TransferAddress     string `json:"transferaddress"`
	Amount              string `json:"amount"`
	BlockIndex          int    `json:"blockindex"`
	TransferNotifyIndex int    `json:"transfernotifyindex"`
	TxHash              string `json:"txhash"`
}
