package models

type RpcContractParameter struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func NewRpcContractParameter(t string, v interface{}) RpcContractParameter {
	return RpcContractParameter{Type: t, Value: v}
}
