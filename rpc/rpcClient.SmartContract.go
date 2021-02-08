package rpc

import (
	"github.com/joeqian10/neo3-gogogo/helper"
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

func (n *RpcClient) InvokeContractVerify(scriptHash string, args []models.RpcContractParameter, signers models.RpcSigners) InvokeResultResponse {
	response := InvokeResultResponse{}
	params := []interface{}{scriptHash, args, signers}
	_ = n.makeRequest("invokecontractverify", params, &response)
	return response
}

func (n *RpcClient) InvokeFunction(scriptHash string, method string, args ...InvokeFunctionStackArg) InvokeResultResponse {
	response := InvokeResultResponse{}
	var params []interface{}
	if args != nil {
		params = []interface{}{scriptHash, method, args}
	} else {
		params = []interface{}{scriptHash, method}
	}
	_ = n.makeRequest("invokefunction", params, &response)
	return response
}

func (n *RpcClient) InvokeScript(scriptInHex string, scriptHashesForVerifying ...helper.UInt160) InvokeResultResponse {
	response := InvokeResultResponse{}
	params := []interface{}{scriptInHex}
	if scriptHashesForVerifying != nil {
		for i := 0; i < len(scriptHashesForVerifying); i++ {
			params = append(params, scriptHashesForVerifying[i].String())
		}
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
