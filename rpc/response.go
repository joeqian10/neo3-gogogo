package rpc

type RpcResponse struct {
	JsonRpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
}

type ErrorResponse struct {
	Error RpcError `json:"error"`
	NetError error
}

func (r *ErrorResponse) HasError() bool {
	if len(r.Error.Message) == 0 && r.NetError == nil {
		return false
	}
	return true
}

func (r *ErrorResponse) GetErrorInfo() string {
	if r.NetError != nil {
		return r.NetError.Error()
	}
	return r.Error.Message
}

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//-----in separate plugins-----

//-----------------------------
