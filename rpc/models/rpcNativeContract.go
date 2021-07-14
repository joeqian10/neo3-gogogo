package models

type RpcNativeContract struct {
	Id            int                 `json:"id"`
	Hash          string              `json:"hash"`
	Nef           RpcNefFile          `json:"nef"`
	Manifest      RpcContractManifest `json:"manifest"`
	UpdateHistory []uint			  `json:"updatehistory"`
}
