package rpc

import (
	"github.com/joeqian10/neo3-gogogo/rpc/models"
)

type InvokeResultResponse struct {
	RpcResponse
	ErrorResponse
	Result models.InvokeResult `json:"result"`
}

type GetUnclaimedGasResponse struct {
	RpcResponse
	ErrorResponse
	Result models.UnclaimedGas `json:"result"`
}

func (n *RpcClient) InvokeFunction(scriptHash string, method string, args []models.RpcContractParameter, signers []models.RpcSigner) InvokeResultResponse {
	response := InvokeResultResponse{}
	if args == nil {
		args = []models.RpcContractParameter{}
	}
	if signers == nil {
		signers = []models.RpcSigner{}
	}
	params := []interface{}{scriptHash, method, args, signers}
	_ = n.makeRequest("invokefunction", params, &response)
	return response
}

// if there is no need to pass "signers", just pass nil
func (n *RpcClient) InvokeScript(scriptInBase64 string, signers []models.RpcSigner) InvokeResultResponse {
	response := InvokeResultResponse{}
	var params []interface{}
	if signers != nil {
		params = []interface{}{scriptInBase64, signers}
	} else {
		params = []interface{}{scriptInBase64}
	}
	_ = n.makeRequest("invokescript", params, &response)
	return response
}

func (n *RpcClient) GetUnclaimedGas(address string) GetUnclaimedGasResponse {
	response := GetUnclaimedGasResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("getunclaimedgas", params, &response)
	return response
}
