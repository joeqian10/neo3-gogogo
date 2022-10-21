package models

type RpcContractParameterContext struct {
	Type    string                    `json:"type"`
	Hash    string                    `json:"hash"`
	Data    string                    `json:"data"`
	Items   map[string]RpcContextItem `json:"items"`
	Network uint32                    `json:"network"`
}

type RpcContextItem struct {
	Script     string                 `json:"script"`
	Parameters []RpcContractParameter `json:"parameters"`
	Signatures map[string]string      `json:"signatures"`
}
