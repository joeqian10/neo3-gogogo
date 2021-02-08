package rpc

import "github.com/joeqian10/neo3-gogogo/rpc/models"

type GetConnectionCountResponse struct {
	RpcResponse
	ErrorResponse
	Result int `json:"result"`
}

type GetPeersResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcPeers `json:"result"`
}

type GetVersionResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcVersion `json:"result"`
}

type SendRawTransactionResponse struct {
	RpcResponse
	ErrorResponse
	Result struct {
		Hash string `json:"hash"`
	} `json:"result"`
}

type SubmitBlockResponse struct {
	RpcResponse
	ErrorResponse
	Result struct {
		Hash string `json:"hash"`
	} `json:"result"`
}

func (n *RpcClient) GetConnectionCount() GetConnectionCountResponse {
	response := GetConnectionCountResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getconnectioncount", params, &response)
	return response
}

func (n *RpcClient) GetPeers() GetPeersResponse {
	response := GetPeersResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getpeers", params, &response)
	return response
}

func (n *RpcClient) GetVersion() GetVersionResponse {
	response := GetVersionResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getversion", params, &response)
	return response
}

func (n *RpcClient) SendRawTransaction(rawTransactionInHex string) SendRawTransactionResponse {
	response := SendRawTransactionResponse{}
	params := []interface{}{rawTransactionInHex, 1}
	_ = n.makeRequest("sendrawtransaction", params, &response)
	return response
}

func (n *RpcClient) SubmitBlock(blockHex string) SubmitBlockResponse {
	response := SubmitBlockResponse{}
	params := []interface{}{blockHex}
	_ = n.makeRequest("submitblock", params, &response)
	return response
}
