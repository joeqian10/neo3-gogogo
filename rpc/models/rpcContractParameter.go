package models

type RpcContractParameter struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewRpcContractParameter(t string, v string) RpcContractParameter {
	return RpcContractParameter{Type: t, Value: v}
}
