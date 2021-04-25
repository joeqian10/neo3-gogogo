package rpc

type RpcRequest struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

func NewRequest(method string, params []interface{}) RpcRequest {
	return RpcRequest{
		JsonRpc: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}
}
