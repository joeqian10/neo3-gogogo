package models

type RpcStateHeight struct {
	LocalRootIndex uint32 `json:"localrootindex"`
	ValidateRootIndex uint32 `json:"validatedrootindex"`
}
