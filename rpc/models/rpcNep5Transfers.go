package models

type RpcNep5Transfers struct {
	Sent     []Nep5Transfer `json:"sent"`
	Received []Nep5Transfer `json:"received"`
	Address  string         `json:"address"`
}

type Nep5Transfer struct {
	Timestamp           int    `json:"timestamp"`
	AssetHash           string `json:"asset_hash"`
	TransferAddress     string `json:"transfer_address"`
	Amount              string `json:"amount"`
	BlockIndex          int    `json:"block_index"`
	TransferNotifyIndex int    `json:"transfer_notify_index"`
	TxHash              string `json:"tx_hash"`
}
