package rpc

import "github.com/joeqian10/neo3-gogogo/rpc/models"

type ListPluginsResponse struct {
	RpcResponse
	ErrorResponse
	Result []models.RpcListPlugin `json:"result"`
}

type ValidateAddressResponse struct {
	RpcResponse
	ErrorResponse
	Result models.ValidateAddress `json:"result"`
}

func (n *RpcClient) ListPlugins() ListPluginsResponse {
	response := ListPluginsResponse{}
	params := []interface{}{}
	_ = n.makeRequest("listplugins", params, &response)
	return response
}

func (n *RpcClient) ValidateAddress(address string) ValidateAddressResponse {
	response := ValidateAddressResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("validateaddress", params, &response)
	return response
}
