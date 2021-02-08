package rpc

import "github.com/joeqian10/neo3-gogogo/rpc/models"

type GetApplicationLogResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcApplicationLog `json:"result"`
}

type GetNep17BalancesResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcNep17Balances `json:"result"`
}

type GetNep17TransfersResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcNep17Transfers `json:"result"`
}

// the endpoint needs to use ApplicationLogs plugin
func (n *RpcClient) GetApplicationLog(txId string) GetApplicationLogResponse {
	response := GetApplicationLogResponse{}
	params := []interface{}{txId}
	_ = n.makeRequest("getapplicationlog", params, &response)
	return response
}

// this endpoint needs RpcNep17Tracker plugin
func (n *RpcClient) GetNep17Balances(address string) GetNep17BalancesResponse {
	response := GetNep17BalancesResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("getnep17balances", params, &response)
	return response
}

// this endpoint needs RpcNep17Tracker plugin
func (n *RpcClient) GetNep17Transfers(address string, startTime *int, endTime *int) GetNep17TransfersResponse {
	response := GetNep17TransfersResponse{}
	var params []interface{}
	if startTime != nil {
		if endTime != nil {
			params = []interface{}{address, *startTime, *endTime}
		} else {
			params = []interface{}{address, *startTime}
		}
	} else {
		params = []interface{}{address}
	}
	_ = n.makeRequest("getnep17balances", params, &response)
	return response
}