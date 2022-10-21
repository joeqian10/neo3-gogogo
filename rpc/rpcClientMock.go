package rpc

import (
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/stretchr/testify/mock"
)

type RpcClientMock struct {
	mock.Mock
}

func (r *RpcClientMock) GetUrl() string {
	args := r.Called()
	return args.Get(0).(string)
}

// ---------------- start section: Blockchain ----------------

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

func (r *RpcClientMock) GetBlockHeaderCount() GetBlockCountResponse {
	args := r.Called()
	return args.Get(0).(GetBlockCountResponse)
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

func (r *RpcClientMock) GetNextBlockValidators() GetNextBlockValidatorsResponse {
	args := r.Called()
	return args.Get(0).(GetNextBlockValidatorsResponse)
}

func (r *RpcClientMock) GetCandidates() GetCandidatesResponse {
	args := r.Called()
	return args.Get(0).(GetCandidatesResponse)
}

func (r *RpcClientMock) GetCommittee() GetCommitteeResponse {
	args := r.Called()
	return args.Get(0).(GetCommitteeResponse)
}

func (r *RpcClientMock) GetNativeContracts() GetNativeContractsResponse {
	args := r.Called()
	return args.Get(0).(GetNativeContractsResponse)
}

// ---------------- start section: Node ----------------

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

// ---------------- start section: Plugins ----------------

func (r *RpcClientMock) GetApplicationLog(s string) GetApplicationLogResponse {
	args := r.Called(s)
	return args.Get(0).(GetApplicationLogResponse)
}

func (r *RpcClientMock) GetNep11Balances(s string) GetNep11BalancesResponse {
	args := r.Called(s)
	return args.Get(0).(GetNep11BalancesResponse)
}

func (r *RpcClientMock) GetNep11Transfers(s string, t1 *int, t2 *int) GetNep11TransfersResponse {
	args := r.Called(s, t1, t2)
	return args.Get(0).(GetNep11TransfersResponse)
}

func (r *RpcClientMock) GetNep11Properties(s1 string, s2 string) GetNep11PropertiesResponse {
	args := r.Called(s1, s2)
	return args.Get(0).(GetNep11PropertiesResponse)
}

func (r *RpcClientMock) GetNep17Balances(s string) GetNep17BalancesResponse {
	args := r.Called(s)
	return args.Get(0).(GetNep17BalancesResponse)
}

func (r *RpcClientMock) GetNep17Transfers(s string, t1 *int, t2 *int) GetNep17TransfersResponse {
	args := r.Called(s, t1, t2)
	return args.Get(0).(GetNep17TransfersResponse)
}

// ---------------- start section: SmartContract ----------------

func (r *RpcClientMock) InvokeFunction(s1 string, s2 string, a []models.RpcContractParameter,
	signersOrWitnesses interface{}, b bool) InvokeResultResponse {
	args := r.Called(s1, s2, a, signersOrWitnesses, b)
	return args.Get(0).(InvokeResultResponse)
}

func (r *RpcClientMock) InvokeScript(s string, signersOrWitnesses interface{}, b bool) InvokeResultResponse {
	args := r.Called(s, signersOrWitnesses, b)
	return args.Get(0).(InvokeResultResponse)
}

func (r *RpcClientMock) TraverseIterator(s1 string, s2 string, c int32) TraverseIteratorResponse {
	args := r.Called(s1, s2, c)
	return args.Get(0).(TraverseIteratorResponse)
}

func (r *RpcClientMock) TerminateSession(s string) TerminateSessionResponse {
	args := r.Called(s)
	return args.Get(0).(TerminateSessionResponse)
}

func (r *RpcClientMock) GetUnclaimedGas(s string) GetUnclaimedGasResponse {
	args := r.Called(s)
	return args.Get(0).(GetUnclaimedGasResponse)
}

// ---------------- start section: State ----------------

func (r *RpcClientMock) GetProof(s1, s2, s3 string) GetProofResponse {
	args := r.Called(s1, s2, s3)
	return args.Get(0).(GetProofResponse)
}

func (r *RpcClientMock) GetStateHeight() GetStateHeightResponse {
	args := r.Called()
	return args.Get(0).(GetStateHeightResponse)
}

func (r *RpcClientMock) GetStateRoot(u uint32) GetStateRootResponse {
	args := r.Called(u)
	return args.Get(0).(GetStateRootResponse)
}

func (r *RpcClientMock) VerifyProof(s string, p string) VerifyProofResponse {
	args := r.Called(s, p)
	return args.Get(0).(VerifyProofResponse)
}

// ---------------- start section: Utilities ----------------

func (r *RpcClientMock) ListPlugins() ListPluginsResponse {
	args := r.Called()
	return args.Get(0).(ListPluginsResponse)
}

func (r *RpcClientMock) ValidateAddress(s string) ValidateAddressResponse {
	args := r.Called(s)
	return args.Get(0).(ValidateAddressResponse)
}

// ---------------- start section: Wallet ----------------

func (r *RpcClientMock) CloseWallet() CloseWalletResponse {
	args := r.Called()
	return args.Get(0).(CloseWalletResponse)
}

func (r *RpcClientMock) DumpPrivKey(s string) DumpPrivKeyResponse {
	args := r.Called(s)
	return args.Get(0).(DumpPrivKeyResponse)
}

func (r *RpcClientMock) GetNewAddress() GetNewAddressResponse {
	args := r.Called()
	return args.Get(0).(GetNewAddressResponse)
}

func (r *RpcClientMock) GetWalletBalance(s string) GetWalletBalanceResponse {
	args := r.Called(s)
	return args.Get(0).(GetWalletBalanceResponse)
}

func (r *RpcClientMock) GetWalletUnclaimedGas() GetWalletUnclaimedGasResponse {
	args := r.Called()
	return args.Get(0).(GetWalletUnclaimedGasResponse)
}

func (r *RpcClientMock) ImportPrivKey(s string) ImportPrivKeyResponse {
	args := r.Called(s)
	return args.Get(0).(ImportPrivKeyResponse)
}

func (r *RpcClientMock) CalculateNetworkFee(s string) CalculateNetworkFeeResponse {
	args := r.Called(s)
	return args.Get(0).(CalculateNetworkFeeResponse)
}

func (r *RpcClientMock) ListAddress() ListAddressResponse {
	args := r.Called()
	return args.Get(0).(ListAddressResponse)
}

func (r *RpcClientMock) OpenWallet(s1 string, s2 string) OpenWalletResponse {
	args := r.Called(s1, s2)
	return args.Get(0).(OpenWalletResponse)
}

func (r *RpcClientMock) SendFrom(s1, s2, s3, s4 string, ss []string) SendFromResponse {
	args := r.Called(s1, s2, s3, s4, ss)
	return args.Get(0).(SendFromResponse)
}

func (r *RpcClientMock) SendMany(s string, ss []string, sn []string) SendManyResponse {
	args := r.Called(s, s, sn)
	return args.Get(0).(SendManyResponse)
}

func (r *RpcClientMock) SendToAddress(s1, s2, s3 string) SendToAddressResponse {
	args := r.Called(s1, s2, s3)
	return args.Get(0).(SendToAddressResponse)
}

func (r *RpcClientMock) InvokeContractVerify(s string, a []models.RpcContractParameter, sw interface{}) InvokeResultResponse {
	args := r.Called(s, a, sw)
	return args.Get(0).(InvokeResultResponse)
}
