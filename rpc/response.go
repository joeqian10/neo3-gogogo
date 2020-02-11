package rpc

import "github.com/joeqian10/neo3-gogogo/rpc/models"

type RpcResponse struct {
	JsonRpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
}

type ErrorResponse struct {
	Error RpcError `json:"error"`
}

func (r *ErrorResponse) HasError() bool {
	if len(r.Error.Message) == 0 {
		return false
	}
	return true
}

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// block chain
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
	Result models.ContractState `json:"result"`
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

type GetValidatorsResponse struct {
	RpcResponse
	ErrorResponse
	Result []models.RpcValidator `json:"result"`
}

// node
type GetConnectionCountResponse struct {
	RpcResponse
	ErrorResponse
	Result int `json:"result"`
}

type GetPeersResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcPeers `json:"result"`
}

type GetVersionResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcVersion `json:"result"`
}

type SendRawTransactionResponse struct {
	RpcResponse
	ErrorResponse
	Result struct {
		Hash string `json:"hash"`
	} `json:"result"`
}

type SubmitBlockResponse struct {
	RpcResponse
	ErrorResponse
	Result struct {
		Hash string `json:"hash"`
	} `json:"result"`
}

// smart contract
type InvokeResultResponse struct {
	RpcResponse
	ErrorResponse
	Result models.InvokeResult `json:"result"`
}

// utilities
type ListPluginsResponse struct {
	RpcResponse
	ErrorResponse
	Result []models.RpcListPlugin `json:"result"`
}

type ValidateAddressResponse struct {
	RpcResponse
	ErrorResponse
	Result models.ValidateAddress `json:"result"`
}

// plugins
type GetApplicationLogResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcApplicationLog `json:"result"`
}

type GetNep5BalancesResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcNep5Balances `json:"result"`
}

type GetNep5TransfersResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcNep5Transfers `json:"result"`
}

// wallet
type CloseWalletResponse struct {
	RpcResponse
	ErrorResponse
	Result bool `json:"result"`
}

type DumpPrivKeyResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetBalanceResponse struct {
	RpcResponse
	ErrorResponse
	Result struct {
		Balance string `json:"balance"`
	} `json:"result"`
}

type GetNewAddressResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetUnclaimedGasResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type ImportPrivKeyResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcAddress `json:"result"`
}

type ListAddressResponse struct {
	RpcResponse
	ErrorResponse
	Result []models.RpcAddress `json:"result"`
}

type OpenWalletResponse struct {
	RpcResponse
	ErrorResponse
	Result bool `json:"result"`
}

type SendFromResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransaction `json:"result"`
}

type SendManyResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransaction `json:"result"`
}

type SendToAddressResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransaction `json:"result"`
}

type GetCrossChainProofResponse struct {
	RpcResponse
	ErrorResponse
	CrosschainProof string `json:"result"`
}
