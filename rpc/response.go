package rpc

type RpcResponse struct {
	JsonRpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
}

type ErrorResponse struct {
	Error RpcError `json:"error"`
}

func (r *ErrorResponse) HasError() bool {
	if len(r.Error.Message) == 0 {
		return false
	}
	return true
}

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//-----in separate plugins-----

//------------------------------
