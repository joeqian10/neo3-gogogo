package rpc

import (
	"github.com/joeqian10/neo3-gogogo/helper"
)

// add IRpcClient for mock UT
type IRpcClient interface {
	// block chain
	GetBestBlockHash() GetBestBlockHashResponse
	GetBlock(hashOrIndex string) GetBlockResponse
	GetBlockCount() GetBlockCountResponse
	GetBlockHash(index uint32) GetBlockHashResponse
	GetBlockHeader(hashOrIndex string) GetBlockHeaderResponse
	GetBlockSysFee(height int) GetBlockSysFeeResponse
	GetContractState(hash string) GetContractStateResponse
	GetRawMemPool() GetRawMemPoolResponse
	GetRawTransaction(hash string) GetRawTransactionResponse
	GetStorage(scriptHash string, key string) GetStorageResponse
	GetTransactionHeight(hash string) GetTransactionHeightResponse
	GetValidators() GetValidatorsResponse
	// node
	GetConnectionCount() GetConnectionCountResponse
	GetPeers() GetPeersResponse
	GetVersion() GetVersionResponse
	SendRawTransaction(txHex string) SendRawTransactionResponse
	SubmitBlock(blockHex string) SubmitBlockResponse
	// smart contract
	InvokeFunction(scriptHash string, function string, args ...InvokeFunctionStackArg) InvokeResultResponse
	InvokeScript(script string, scriptHashesForVerifying ...helper.UInt160) InvokeResultResponse
	// utilities
	ListPlugins() ListPluginsResponse
	ValidateAddress(address string) ValidateAddressResponse
	// plugins
	GetApplicationLog(hash string) GetApplicationLogResponse
	GetNep5Balances(address string) GetNep5BalancesResponse
	GetNep5Transfers(address string, startTimestamp *int, endTimestamp *int) GetNep5TransfersResponse
	// wallet methods are not needed to mock

}
