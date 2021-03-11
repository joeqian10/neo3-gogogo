package rpc

import (
	"github.com/joeqian10/neo3-gogogo/mpt"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
)

type GetProofResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetStateHeightResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcStateHeight `json:"result"`
}

type GetStateRootResponse struct {
	RpcResponse
	ErrorResponse
	Result mpt.StateRoot `json:"result"`
}

type VerifyProofResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"` // base64
}

func (n *RpcClient) GetProof(rootHash, contractScriptHash, storeKey string) GetProofResponse {
	response := GetProofResponse{}
	params := []interface{}{rootHash, contractScriptHash, storeKey}
	_ = n.makeRequest("getproof", params, &response)
	return response
}

func (n *RpcClient) GetStateHeight() GetStateHeightResponse {
	response := GetStateHeightResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getstateheight", params, &response)
	return response
}

func (n *RpcClient) GetStateRoot(blockHeight uint32) GetStateRootResponse {
	response := GetStateRootResponse{}
	params := []interface{}{blockHeight}
	_ = n.makeRequest("getstateroot", params, &response)
	return response
}

func (n *RpcClient) VerifyProof(rootHash string, proofInBase64 string) VerifyProofResponse {
	response := VerifyProofResponse{}
	params := []interface{}{rootHash, proofInBase64}
	_ = n.makeRequest("verifyproof", params, &response)
	return response
}
