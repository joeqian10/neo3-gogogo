package rpc

import (
	"github.com/joeqian10/neo3-gogogo/rpc/models"
)

// add IRpcClient for mock UT
type IRpcClient interface {
	GetUrl() string
	// block chain
	GetBestBlockHash() GetBestBlockHashResponse
	GetBlock(hashOrIndex string) GetBlockResponse
	GetBlockCount() GetBlockCountResponse
	GetBlockHash(index uint32) GetBlockHashResponse
	GetBlockHeader(hashOrIndex string) GetBlockHeaderResponse
	GetContractState(hash string) GetContractStateResponse
	GetRawMemPool() GetRawMemPoolResponse
	GetRawTransaction(hash string) GetRawTransactionResponse
	GetStorage(scriptHash string, key string) GetStorageResponse
	GetTransactionHeight(hash string) GetTransactionHeightResponse
	GetNextBlockValidators() GetNextBlockValidatorsResponse
	GetCommittee() GetCommitteeResponse

	// node
	GetConnectionCount() GetConnectionCountResponse
	GetPeers() GetPeersResponse
	GetVersion() GetVersionResponse
	SendRawTransaction(tx string) SendRawTransactionResponse
	SubmitBlock(block string) SubmitBlockResponse

	// plugins
	GetApplicationLog(txId string) GetApplicationLogResponse
	GetNep17Balances(address string) GetNep17BalancesResponse
	GetNep17Transfers(address string, startTimestamp *int, endTimestamp *int) GetNep17TransfersResponse

	// smart contract
	InvokeFunction(scriptHash string, function string, args []models.RpcContractParameter, signers []models.RpcSigner) InvokeResultResponse
	InvokeScript(script string, signers []models.RpcSigner) InvokeResultResponse
	GetUnclaimedGas(address string) GetUnclaimedGasResponse

	// state
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
	SendFrom(assetId string, from string, to string, amount string) SendFromResponse
	SendMany(fromAddress string, outputs []models.RpcTransferOut, signers ...models.RpcSigner) SendManyResponse
	SendToAddress(assetId string, to string, amount string) SendToAddressResponse
	InvokeContractVerify(scriptHash string, args []models.RpcContractParameter, signers []models.RpcSigner) InvokeResultResponse
}
