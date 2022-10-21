package rpc

import (
	"github.com/joeqian10/neo3-gogogo/rpc/models"
)

// IRpcClient is used for mock UT
type IRpcClient interface {
	GetUrl() string

	// Blockchain
	GetBestBlockHash() GetBestBlockHashResponse
	GetBlock(hashOrIndex string) GetBlockResponse
	GetBlockCount() GetBlockCountResponse
	GetBlockHash(index uint32) GetBlockHashResponse
	GetBlockHeader(hashOrIndex string) GetBlockHeaderResponse
	GetBlockHeaderCount() GetBlockCountResponse // todo
	GetContractState(hash string) GetContractStateResponse
	GetRawMemPool() GetRawMemPoolResponse
	GetRawTransaction(hash string) GetRawTransactionResponse
	GetStorage(scriptHash string, key string) GetStorageResponse
	GetTransactionHeight(hash string) GetTransactionHeightResponse
	GetNextBlockValidators() GetNextBlockValidatorsResponse
	GetCandidates() GetCandidatesResponse // todo
	GetCommittee() GetCommitteeResponse
	GetNativeContracts() GetNativeContractsResponse

	// Node
	GetConnectionCount() GetConnectionCountResponse
	GetPeers() GetPeersResponse
	GetVersion() GetVersionResponse
	SendRawTransaction(tx string) SendRawTransactionResponse
	SubmitBlock(block string) SubmitBlockResponse

	// Plugins
	GetApplicationLog(txId string) GetApplicationLogResponse
	GetNep11Balances(address string) GetNep11BalancesResponse
	GetNep11Transfers(address string, startTime *int, endTime *int) GetNep11TransfersResponse
	GetNep11Properties(assetHash string, tokenId string) GetNep11PropertiesResponse
	GetNep17Balances(address string) GetNep17BalancesResponse
	GetNep17Transfers(address string, startTimestamp *int, endTimestamp *int) GetNep17TransfersResponse

	// SmartContract
	InvokeFunction(scriptHash string, method string, args []models.RpcContractParameter, signersOrWitnesses interface{}, useDiagnostic bool) InvokeResultResponse
	InvokeScript(scriptInBase64 string, signersOrWitnesses interface{}, useDiagnostic bool) InvokeResultResponse
	TraverseIterator(sessionId string, iteratorId string, count int32) TraverseIteratorResponse
	TerminateSession(sessionId string) TerminateSessionResponse
	GetUnclaimedGas(address string) GetUnclaimedGasResponse

	// State
	GetProof(rootHash, contractScriptHash, storeKey string) GetProofResponse
	GetStateHeight() GetStateHeightResponse
	GetStateRoot(blockHeight uint32) GetStateRootResponse
	VerifyProof(rootHash string, proofInBase64 string) VerifyProofResponse

	// utilities
	ListPlugins() ListPluginsResponse
	ValidateAddress(address string) ValidateAddressResponse

	// wallet
	CloseWallet() CloseWalletResponse
	DumpPrivKey(address string) DumpPrivKeyResponse
	GetNewAddress() GetNewAddressResponse
	GetWalletBalance(assetId string) GetWalletBalanceResponse
	GetWalletUnclaimedGas() GetWalletUnclaimedGasResponse
	ImportPrivKey(wif string) ImportPrivKeyResponse
	CalculateNetworkFee(tx string) CalculateNetworkFeeResponse
	ListAddress() ListAddressResponse
	OpenWallet(path string, password string) OpenWalletResponse
	SendFrom(assetId string, from string, to string, amount string, signers []string) SendFromResponse
	SendMany(fromAddress string, toAddresses []string, signerAddresses []string) SendManyResponse
	SendToAddress(assetId string, toAddress string, amount string) SendToAddressResponse
	InvokeContractVerify(scriptHash string, args []models.RpcContractParameter, signersOrWitnesses interface{}) InvokeResultResponse
}
