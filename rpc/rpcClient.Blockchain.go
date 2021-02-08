package rpc

import (
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"strconv"
)

type GetBestBlockHashResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetBlockResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcBlock `json:"result"`
}

type GetBlockCountResponse struct {
	RpcResponse
	ErrorResponse
	Result int `json:"result"`
}

type GetBlockHashResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetBlockHeaderResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcBlockHeader `json:"result"`
}

type GetBlockSysFeeResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetContractStateResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcContractState `json:"result"`
}

type GetRawMemPoolResponse struct {
	RpcResponse
	ErrorResponse
	Result []string `json:"result"`
}

type GetRawTransactionResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransaction `json:"result"`
}

type GetStorageResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetTransactionHeightResponse struct {
	RpcResponse
	ErrorResponse
	Result int `json:"result"`
}

type GetNextBlockValidatorsResponse struct {
	RpcResponse
	ErrorResponse
	Result []models.RpcValidator `json:"result"`
}

type GetCommitteeResponse struct {
	RpcResponse
	ErrorResponse
	Result []string `json:"result"`
}

func (n *RpcClient) GetBestBlockHash() GetBestBlockHashResponse {
	response := GetBestBlockHashResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getbestblockhash", params, &response)
	return response
}

func (n *RpcClient) GetBlock(hashOrIndex string) GetBlockResponse {
	params := []interface{}{hashOrIndex, true}
	index, err := strconv.Atoi(hashOrIndex)
	if err == nil {
		params = []interface{}{index, true}
	}
	response := GetBlockResponse{}
	_ = n.makeRequest("getblock", params, &response)
	return response
}

func (n *RpcClient) GetBlockCount() GetBlockCountResponse {
	response := GetBlockCountResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getblockcount", params, &response)
	return response
}

func (n *RpcClient) GetBlockHash(index uint32) GetBlockHashResponse {
	response := GetBlockHashResponse{}
	params := []interface{}{index}
	_ = n.makeRequest("getblockhash", params, &response)
	return response
}

func (n *RpcClient) GetBlockHeader(hashOrIndex string) GetBlockHeaderResponse {
	params := []interface{}{hashOrIndex, true}
	index, err := strconv.Atoi(hashOrIndex)
	if err == nil {
		params = []interface{}{index, true}
	}
	response := GetBlockHeaderResponse{}
	_ = n.makeRequest("getblockheader", params, &response)
	return response
}

func (n *RpcClient) GetContractState(scriptHash string) GetContractStateResponse {
	response := GetContractStateResponse{}
	params := []interface{}{scriptHash}
	_ = n.makeRequest("getcontractstate", params, &response)
	return response
}

func (n *RpcClient) GetRawMemPool() GetRawMemPoolResponse {
	response := GetRawMemPoolResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getrawmempool", params, &response)
	return response
}

func (n *RpcClient) GetRawTransaction(txid string) GetRawTransactionResponse {
	response := GetRawTransactionResponse{}
	params := []interface{}{txid, 1}
	_ = n.makeRequest("getrawtransaction", params, &response)
	return response
}

func (n *RpcClient) GetStorage(scripthash string, key string) GetStorageResponse {
	response := GetStorageResponse{}
	params := []interface{}{scripthash, key}
	_ = n.makeRequest("getstorage", params, &response)
	return response
}

func (n *RpcClient) GetTransactionHeight(txid string) GetTransactionHeightResponse {
	response := GetTransactionHeightResponse{}
	params := []interface{}{txid}
	_ = n.makeRequest("gettransactionheight", params, &response)
	return response
}

func (n *RpcClient) GetNextBlockValidators() GetNextBlockValidatorsResponse {
	response := GetNextBlockValidatorsResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getnextblockvalidators", params, &response)
	return response
}

func (n *RpcClient) GetCommittee() GetCommitteeResponse {
	response := GetCommitteeResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getcommittee", params, &response)
	return response
}
