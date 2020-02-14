package rpc

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/stretchr/testify/mock"
)

type RpcClientMock struct {
	mock.Mock
}

func (r *RpcClientMock) GetUrl() string {
	args := r.Called()
	return args.Get(0).(string)
}

// block chain
func (r *RpcClientMock) GetBestBlockHash() GetBestBlockHashResponse {
	args := r.Called()
	return args.Get(0).(GetBestBlockHashResponse)
}
func (r *RpcClientMock) GetBlock(hashOrIndex string) GetBlockResponse {
	args := r.Called(hashOrIndex)
	return args.Get(0).(GetBlockResponse)
}

func (r *RpcClientMock) GetBlockCount() GetBlockCountResponse {
	args := r.Called()
	return args.Get(0).(GetBlockCountResponse)
}

func (r *RpcClientMock) GetBlockHash(n uint32) GetBlockHashResponse {
	args := r.Called(n)
	return args.Get(0).(GetBlockHashResponse)
}

func (r *RpcClientMock) GetBlockHeader(hashOrIndex string) GetBlockHeaderResponse {
	args := r.Called(hashOrIndex)
	return args.Get(0).(GetBlockHeaderResponse)
}

func (r *RpcClientMock) GetBlockSysFee(height int) GetBlockSysFeeResponse {
	args := r.Called(height)
	return args.Get(0).(GetBlockSysFeeResponse)
}

func (r *RpcClientMock) GetContractState(s string) GetContractStateResponse {
	args := r.Called(s)
	return args.Get(0).(GetContractStateResponse)
}

func (r *RpcClientMock) GetRawMemPool() GetRawMemPoolResponse {
	args := r.Called()
	return args.Get(0).(GetRawMemPoolResponse)
}

func (r *RpcClientMock) GetRawTransaction(s string) GetRawTransactionResponse {
	args := r.Called(s)
	return args.Get(0).(GetRawTransactionResponse)
}

func (r *RpcClientMock) GetStorage(s1 string, s2 string) GetStorageResponse {
	args := r.Called(s1, s2)
	return args.Get(0).(GetStorageResponse)
}

func (r *RpcClientMock) GetTransactionHeight(s string) GetTransactionHeightResponse {
	args := r.Called(s)
	return args.Get(0).(GetTransactionHeightResponse)
}

func (r *RpcClientMock) GetValidators() GetValidatorsResponse {
	args := r.Called()
	return args.Get(0).(GetValidatorsResponse)
}

// node
func (r *RpcClientMock) GetConnectionCount() GetConnectionCountResponse {
	args := r.Called()
	return args.Get(0).(GetConnectionCountResponse)
}

func (r *RpcClientMock) GetPeers() GetPeersResponse {
	args := r.Called()
	return args.Get(0).(GetPeersResponse)
}

func (r *RpcClientMock) GetVersion() GetVersionResponse {
	args := r.Called()
	return args.Get(0).(GetVersionResponse)
}

func (r *RpcClientMock) SendRawTransaction(s string) SendRawTransactionResponse {
	args := r.Called(s)
	return args.Get(0).(SendRawTransactionResponse)
}

func (r *RpcClientMock) SubmitBlock(s string) SubmitBlockResponse {
	args := r.Called(s)
	return args.Get(0).(SubmitBlockResponse)
}

// smart contract
func (r *RpcClientMock) InvokeFunction(s1 string, s2 string, a ...InvokeFunctionStackArg) InvokeResultResponse {
	args := r.Called(s1, s2, a)
	return args.Get(0).(InvokeResultResponse)
}

func (r *RpcClientMock) InvokeScript(s string, v ...helper.UInt160) InvokeResultResponse {
	args := r.Called(s)
	return args.Get(0).(InvokeResultResponse)
}

// utilities
func (r *RpcClientMock) ListPlugins() ListPluginsResponse {
	args := r.Called()
	return args.Get(0).(ListPluginsResponse)
}

func (r *RpcClientMock) ListAddress() ListAddressResponse {
	args := r.Called()
	return args.Get(0).(ListAddressResponse)
}

func (r *RpcClientMock) ValidateAddress(s string) ValidateAddressResponse {
	args := r.Called(s)
	return args.Get(0).(ValidateAddressResponse)
}

// plugins
func (r *RpcClientMock) GetApplicationLog(s string) GetApplicationLogResponse {
	args := r.Called(s)
	return args.Get(0).(GetApplicationLogResponse)
}

func (r *RpcClientMock) GetNep5Balances(s string) GetNep5BalancesResponse {
	args := r.Called(s)
	return args.Get(0).(GetNep5BalancesResponse)
}

func (r *RpcClientMock) GetNep5Transfers(s string, t1 *int, t2 *int) GetNep5TransfersResponse {
	args := r.Called(s, t1, t2)
	return args.Get(0).(GetNep5TransfersResponse)
}

// wallet methods are not needed to mock
