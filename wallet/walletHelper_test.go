package wallet

import (
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/stretchr/testify/mock"
	"math/big"
	"testing"

	"github.com/joeqian10/neo3-gogogo/tx"
	"github.com/stretchr/testify/assert"
)

func TestNewWalletHelper(t *testing.T) {
	txBuilder := tx.NewTransactionBuilder("http://seed1.ngd.network:20332")
	account, err := NewAccountFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
	assert.NotNil(t, txBuilder)
	assert.NotNil(t, account)
	assert.Nil(t, err)
	walletHelper := NewWalletHelper(txBuilder.Client, account)
	assert.NotNil(t, walletHelper)
	assert.Equal(t, "http://seed1.ngd.network:20332", walletHelper.Client.GetUrl())
	assert.Equal(t, "03b7a7f933199f28cc1c48d22a21c78ac3992cf7fceb038a9c670fe55444426619", walletHelper.Account.KeyPair.PublicKey.String())
}

func TestWalletHelper_ClaimGas(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	clientMock.On("GetBlockCount", mock.Anything).Return(rpc.GetBlockCountResponse{
		RpcResponse: rpc.RpcResponse{
			JsonRpc: "2.0",
			ID:      1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result: 1234,
	})
	clientMock.On("InvokeScript", mock.Anything).Return(rpc.InvokeResultResponse{
		RpcResponse: rpc.RpcResponse{
			JsonRpc: "2.0",
			ID:      1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result: models.InvokeResult{
			Script:      "0c146925aa554712439a9c613ba114efa3fac23ddbca11c00c0962616c616e63654f660c143b7d3711c6f0ccf9b1dca903d1bfa1d896f1238c41627d5b52",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStackResult{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})
	clientMock.On("SendRawTransaction", mock.Anything).Return(rpc.SendRawTransactionResponse{
		RpcResponse: rpc.RpcResponse{
			JsonRpc: "2.0",
			ID:      1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result: struct {
			Hash string `json:"hash"`
		}{Hash: "0x992f941c9751aabc8bab0200503e07e38f38f0884cb8b11f6c6c72d8d2fb2948"},
	})
	account, _ := NewAccountFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
	var w = NewWalletHelper(clientMock, account)
	h, e := w.ClaimGas()
	assert.Nil(t, e)
	assert.Equal(t, "0x992f941c9751aabc8bab0200503e07e38f38f0884cb8b11f6c6c72d8d2fb2948", h)
}

func TestWalletHelper_GetBalance(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	clientMock.On("InvokeScript", mock.Anything).Return(rpc.InvokeResultResponse{
		RpcResponse: rpc.RpcResponse{
			JsonRpc: "2.0",
			ID:      1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result: models.InvokeResult{
			Script:      "0c146925aa554712439a9c613ba114efa3fac23ddbca11c00c0962616c616e63654f660c143b7d3711c6f0ccf9b1dca903d1bfa1d896f1238c41627d5b52",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStackResult{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})
	account, _ := NewAccountFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
	var w = NewWalletHelper(clientMock, account)
	b, e := w.GetBalance(tx.GasToken)
	assert.Nil(t, e)
	assert.Equal(t, uint64(8913620128), b.Uint64())
}

func TestWalletHelper_GetUnClaimedGas(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	clientMock.On("GetBlockCount", mock.Anything).Return(rpc.GetBlockCountResponse{
		RpcResponse: rpc.RpcResponse{
			JsonRpc: "2.0",
			ID:      1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result: 1234,
	})
	clientMock.On("InvokeScript", mock.Anything).Return(rpc.InvokeResultResponse{
		RpcResponse: rpc.RpcResponse{
			JsonRpc: "2.0",
			ID:      1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result: models.InvokeResult{
			Script:      "0c146925aa554712439a9c613ba114efa3fac23ddbca11c00c0962616c616e63654f660c143b7d3711c6f0ccf9b1dca903d1bfa1d896f1238c41627d5b52",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStackResult{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})
	account, _ := NewAccountFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
	var w = NewWalletHelper(clientMock, account)
	v, err := w.GetUnClaimedGas()
	assert.Nil(t, err)
	assert.Equal(t, 89.13620128, v)
}

func TestWalletHelper_Transfer(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	clientMock.On("GetBlockCount", mock.Anything).Return(rpc.GetBlockCountResponse{
		RpcResponse: rpc.RpcResponse{
			JsonRpc: "2.0",
			ID:      1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result: 1234,
	})
	clientMock.On("InvokeScript", mock.Anything).Return(rpc.InvokeResultResponse{
		RpcResponse: rpc.RpcResponse{
			JsonRpc: "2.0",
			ID:      1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result: models.InvokeResult{
			Script:      "0c146925aa554712439a9c613ba114efa3fac23ddbca11c00c0962616c616e63654f660c143b7d3711c6f0ccf9b1dca903d1bfa1d896f1238c41627d5b52",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStackResult{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})
	clientMock.On("SendRawTransaction", mock.Anything).Return(rpc.SendRawTransactionResponse{
		RpcResponse: rpc.RpcResponse{
			JsonRpc: "2.0",
			ID:      1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result: struct {
			Hash string `json:"hash"`
		}{Hash: "0x992f941c9751aabc8bab0200503e07e38f38f0884cb8b11f6c6c72d8d2fb2948"},
	})
	account, _ := NewAccountFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
	var w = NewWalletHelper(clientMock, account)
	to := "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ"
	h, e := w.Transfer(tx.GasToken, to, big.NewInt(10000))
	assert.Nil(t, e)
	assert.Equal(t, "0x992f941c9751aabc8bab0200503e07e38f38f0884cb8b11f6c6c72d8d2fb2948", h)
}

//func TestWallet_Transfer(t *testing.T) {
//	client := rpc.NewClient("http://127.0.0.1:10332")
//	client.SetBasicAuth("krain", "123456")
//	account, _ := NewAccountFromWIF("KyoYyZpoccbR6KZ25eLzhMTUxREwCpJzDsnuodGTKXSG8fDW9t7x")
//	address := "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ"
//	api := NewWalletHelper(client, account)
//	gasBalace, _ := api.GetBalance(tx.GasToken, address)
//
//	assert.Equal(t, true, gasBalace.Int64() > 0)
//
//	result, err := api.Transfer(tx.NeoToken, "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW", big.NewInt(100))
//	assert.Nil(t, err)
//	assert.True(t, len(result) > 0)
//	fmt.Println(result)
//
//	claimable, _ := api.GetUnClaimedGas(address)
//	assert.True(t, claimable > 0)
//	fmt.Println(claimable)
//
//	result, _ = api.ClaimGas()
//	assert.True(t, len(result) > 0)
//	fmt.Println(result)
//}

// func TestWallet_NEP5Helper(t *testing.T) {
// 	client := rpc.NewClient("http://127.0.0.1:10332")
// 	nep5Api := nep5.NewNep5Helper(client)

// 	total, err := nep5Api.TotalSupply(tx.NeoToken)
// 	assert.Nil(t, err)
// 	assert.Equal(t, big.NewInt(100000000), total)

// 	name, err := nep5Api.Name(tx.NeoToken)
// 	assert.Nil(t, err)
// 	assert.Equal(t, "NEO", name)

// 	tokensymbol, err := nep5Api.Symbol(tx.NeoToken)
// 	assert.Nil(t, err)
// 	assert.Equal(t, "neo", tokensymbol)

// 	decimals, err := nep5Api.Decimals(tx.GasToken)
// 	assert.Nil(t, err)
// 	assert.Equal(t, 8, decimals)

// 	acc, _ := helper.AddressToScriptHash("NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ")
// 	balance, err := nep5Api.BalanceOf(tx.GasToken, acc)
// 	assert.Nil(t, err)
// 	assert.Equal(t, true, balance.Int64() > 0)

// }
