package models

type RpcTransferOut struct {
	Asset   string `json:"asset"`
	Value   string `json:"value"`
	Address string `json:"address"`
}
