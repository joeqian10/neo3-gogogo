package rpc

import "github.com/joeqian10/neo3-gogogo/rpc/models"

type GetApplicationLogResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcApplicationLog `json:"result"`
}

type GetNep11BalancesResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcNep11Balances `json:"result"`
}

type GetNep11TransfersResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcNep11Transfers `json:"result"`
}

type GetNep11PropertiesResponse struct {
	RpcResponse
	ErrorResponse
	Result map[string]string `json:"result"`
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

// GetApplicationLog needs the ApplicationLogs plugin
func (n *RpcClient) GetApplicationLog(txId string) GetApplicationLogResponse {
	response := GetApplicationLogResponse{}
	params := []interface{}{txId}
	_ = n.makeRequest("getapplicationlog", params, &response)
	return response
}

// GetNep11Balances needs the TokensTracker plugin
func (n *RpcClient) GetNep11Balances(address string) GetNep11BalancesResponse {
	response := GetNep11BalancesResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("getnep11balances", params, &response)
	return response
}

// GetNep11Transfers needs the TokensTracker plugin
func (n *RpcClient) GetNep11Transfers(address string, startTime *int, endTime *int) GetNep11TransfersResponse {
	response := GetNep11TransfersResponse{}
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
	_ = n.makeRequest("getnep11transfers", params, &response)
	return response
}

// GetNep11Properties needs the TokensTracker plugin
func (n *RpcClient) GetNep11Properties(assetHash string, tokenId string) GetNep11PropertiesResponse {
	response := GetNep11PropertiesResponse{}
	params := []interface{}{assetHash, tokenId}
	_ = n.makeRequest("getnep11properties", params, &response)
	return response
}

// GetNep17Balances needs the TokensTracker plugin
func (n *RpcClient) GetNep17Balances(address string) GetNep17BalancesResponse {
	response := GetNep17BalancesResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("getnep17balances", params, &response)
	return response
}

// GetNep17Transfers needs the TokensTracker plugin
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
