package wallet

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/stretchr/testify/mock"
	"math/big"
	"testing"

	"github.com/joeqian10/neo3-gogogo/tx"
	"github.com/stretchr/testify/assert"
)

var client = rpc.NewClient("http://test:20332")

func TestNewWalletHelperFromPrivateKey(t *testing.T) {
	wh, err := NewWalletHelperFromPrivateKey(client, privateKey)
	assert.Nil(t, err)
	assert.NotNil(t, wh)
	assert.NotNil(t, wh.Client)
	assert.NotNil(t, wh.wallet)
	assert.Equal(t, 1, len(wh.wallet.GetAccounts()))
}

func TestNewWalletHelperFromContract(t *testing.T) {
	wh, err := NewWalletHelperFromContract(client, testContract, pair)
	assert.Nil(t, err)
	assert.NotNil(t, wh)
	assert.NotNil(t, wh.Client)
	assert.NotNil(t, wh.wallet)
	assert.Equal(t, 1, len(wh.wallet.GetAccounts()))
}

func TestNewWalletHelperFromNEP2(t *testing.T) {
	wh, err := NewWalletHelperFromNEP2(client, nep2, password, 2, 1, 1)
	assert.Nil(t, err)
	assert.NotNil(t, wh)
	assert.NotNil(t, wh.Client)
	assert.NotNil(t, wh.wallet)
	assert.Equal(t, 1, len(wh.wallet.GetAccounts()))
}

func TestNewWalletHelperFromWIF(t *testing.T) {
	wh, err := NewWalletHelperFromWIF(client, wif)
	assert.Nil(t, err)
	assert.NotNil(t, wh)
	assert.NotNil(t, wh.Client)
	assert.NotNil(t, wh.wallet)
	assert.Equal(t, 1, len(wh.wallet.GetAccounts()))
}

func TestNewWalletHelperFromWallet(t *testing.T) {
	wh := NewWalletHelperFromWallet(client, testWallet)
	assert.NotNil(t, wh)
	assert.NotNil(t, wh.Client)
	assert.NotNil(t, wh.wallet)
}

func TestWalletHelper_CalculateNetworkFee(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)

	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	wh := NewWalletHelperFromWallet(clientMock, testWallet)

	trx := tx.NewTransaction()
	b, e := wh.CalculateNetworkFee(trx)
	assert.Nil(t, e)
	assert.Equal(t, uint64(29000), b)

	resetTestWallet()
}

func TestWalletHelper_CalculateNetworkFee2(t *testing.T) {
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
		Result: 6666666,
	})
	clientMock.On("GetContractState", mock.Anything).Return(rpc.GetContractStateResponse{
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
		Result: models.RpcContractState{
			Id:            8,
			UpdateCounter: 0,
			Hash:          "0x99042d380f2b754175717bb932a911bc0bb0ad7d",
			Nef: models.RpcNefFile{
				Magic:    0x3346454E,
				Compiler: "neon",
				Tokens: []models.RpcMethodToken{
					{
						Hash:            "0x99042d380f2b754175717bb932a911bc0bb0ad7d",
						Method:          "verify",
						ParametersCount: 0,
						HasReturnValue:  true,
						CallFlags:       "AllowCall",
					},
				},
				Script:   "DBQKo4e1Ppa3mJpjFDGgVt0fQKBC9kH4J+yMQDTkQFcBAAwFSGVsbG9Bm/ZnzkGSXegxcGhAVwQBEnAMF0ludm9rZSBTdHJvYWdlLlB1dCBmb3IgaBpQQXvjun0MByB0aW1lcy6Li9soQc/nR5YMBUhlbGxveFBBm/ZnzkHmPxiEDAJOb0Gb9mfOQZJd6DHYqnNrJiwMAk5vDAJOb0Gb9mfOQZJd6DFK2CYFEFBF2yERnlBBm/ZnzkHmPxiEIhMhDAJObxFQQZv2Z85B5j8YhAwCTm9Bm/ZnzkGSXegxcWlK2CYFEFBF2yEaUEF747p9chXDShAMBFB1dCDQShF40EoSDB0gaW50byBzdG9yYWdlIGNvbXBsZXRlbHkgZm9yINBKE2rQShQMBiB0aW1lc9DBShEyCJ1Ti1Ai+EXbKEHP50eWeBHADARXb3JkQZUBb2FpEcAMDkludm9rZVB1dENvdW50QZUBb2FAVwECNZL+//8Qs3BoJhYMEU5vIGF1dGhvcml6YXRpb24uOnh5UEExxjMdQFcBADVn/v//ELNwaCYWDBFObyBhdXRob3JpemF0aW9uLjohQcafHfBAVgEMFAqjh7U+lreYmmMUMaBW3R9AoEL2YEA=",
				CheckSum: 73195690102,
			},
			Manifest: models.RpcContractManifest{
				Name:               "testContract",
				Groups:             []models.RpcContractGroup{},
				SupportedStandards: []string{},
				Abi: models.RpcContractAbi{
					Methods: []models.RpcContractMethodDescriptor{
						{
							Name:       "verify",
							Parameters: []models.RpcContractParameterDefinition{},
							ReturnType: "Boolean",
							Offset:     28,
							Safe:       true,
						},
						{
							Name: "update",
							Parameters: []models.RpcContractParameterDefinition{
								{
									Name: "script",
									Type: "ByteArray",
								},
								{
									Name: "manifest",
									Type: "String",
								},
							},
							ReturnType: "Void",
							Offset:     363,
							Safe:       false,
						},
					},
					Events: []models.RpcContractEventDescriptor{},
				},
				Permissions: []models.RpcContractPermission{
					{
						Contract: "*",
						Methods:  []string{"*"},
					},
				},
				Trusts: []string{},
			},
		},
	})
	clientMock.On("InvokeScript", mock.Anything, mock.Anything).Return(rpc.InvokeResultResponse{
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
			Script:      "DAABECcMFPqJ+ywU3w9Z3d8E9uVlF/KzSq7rDBTitlMicpPpnE8pBtU1U6u0pnLfhhTAHwwIdHJhbnNmZXIMFIOrBnmtVcBQoTrUP1k26nP16x72QWJ9W1I=",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStack{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})
	clientMock.On("InvokeFunction", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(rpc.InvokeResultResponse{
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
			Script:      "DAABECcMFPqJ+ywU3w9Z3d8E9uVlF/KzSq7rDBTitlMicpPpnE8pBtU1U6u0pnLfhhTAHwwIdHJhbnNmZXIMFIOrBnmtVcBQoTrUP1k26nP16x72QWJ9W1I=",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStack{{
				Type:  "Boolean",
				Value: true,
			}},
		},
	})

	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	wh := NewWalletHelperFromWallet(clientMock, testWallet)

	script := []byte{}
	ab := []AccountAndBalance{
		{
			Account: helper.NewUInt160(),
			Value:   big.NewInt(1000000000000), // 10000 gas
		},
	}
	trx, err := wh.MakeTransaction(script, nil, nil, ab)
	assert.Nil(t, err)
	assert.Equal(t, int64(2059570), trx.GetNetworkFee())

	resetTestWallet()
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
	clientMock.On("InvokeScript", mock.Anything, mock.Anything).Return(rpc.InvokeResultResponse{
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
			Script:      "DAABECcMFPqJ+ywU3w9Z3d8E9uVlF/KzSq7rDBTitlMicpPpnE8pBtU1U6u0pnLfhhTAHwwIdHJhbnNmZXIMFIOrBnmtVcBQoTrUP1k26nP16x72QWJ9W1I=",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStack{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})
	clientMock.On("CalculateNetworkFee", mock.Anything).Return(rpc.CalculateNetworkFeeResponse{
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
		Result: models.RpcNetworkFee{NetworkFee: "2384840"},
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
	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	wh := NewWalletHelperFromWallet(clientMock, testWallet)
	h, e := wh.ClaimGas(helper.Neo3Magic_MainNet)
	assert.Nil(t, e)
	assert.Equal(t, "0x992f941c9751aabc8bab0200503e07e38f38f0884cb8b11f6c6c72d8d2fb2948", h)

	resetTestWallet()
}

func TestWalletHelper_GetAccountAndBalance(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	clientMock.On("InvokeScript", mock.Anything, mock.Anything).Return(rpc.InvokeResultResponse{
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
			Script:      "DAABECcMFPqJ+ywU3w9Z3d8E9uVlF/KzSq7rDBTitlMicpPpnE8pBtU1U6u0pnLfhhTAHwwIdHJhbnNmZXIMFIOrBnmtVcBQoTrUP1k26nP16x72QWJ9W1I=",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStack{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})

	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	wh := NewWalletHelperFromWallet(clientMock, testWallet)

	r, e := wh.GetAccountAndBalance(tx.GasToken)
	assert.Nil(t, e)
	assert.Equal(t, 1, len(r))
	assert.Equal(t, 0, r[0].Account.CompareTo(testScriptHash))
	assert.Equal(t, int64(8913620128), r[0].Value.Int64())

	resetTestWallet()
}

func TestWalletHelper_GetBalanceFromAccount(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	clientMock.On("InvokeScript", mock.Anything, mock.Anything).Return(rpc.InvokeResultResponse{
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
			Script:      "DAABECcMFPqJ+ywU3w9Z3d8E9uVlF/KzSq7rDBTitlMicpPpnE8pBtU1U6u0pnLfhhTAHwwIdHJhbnNmZXIMFIOrBnmtVcBQoTrUP1k26nP16x72QWJ9W1I=",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStack{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})
	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)

	wh := NewWalletHelperFromWallet(clientMock, testWallet)
	b, e := wh.GetBalanceFromAccount(tx.GasToken, testScriptHash)
	assert.Nil(t, e)
	assert.Equal(t, uint64(8913620128), b.Uint64())

	resetTestWallet()
}

func TestWalletHelper_GetBalanceFromWallet(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	clientMock.On("InvokeScript", mock.Anything, mock.Anything).Return(rpc.InvokeResultResponse{
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
			Script:      "DAABECcMFPqJ+ywU3w9Z3d8E9uVlF/KzSq7rDBTitlMicpPpnE8pBtU1U6u0pnLfhhTAHwwIdHJhbnNmZXIMFIOrBnmtVcBQoTrUP1k26nP16x72QWJ9W1I=",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStack{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})
	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	wh := NewWalletHelperFromWallet(clientMock, testWallet)

	b, e := wh.GetBalanceFromWallet(tx.GasToken, nil)
	assert.Nil(t, e)
	assert.Equal(t, uint64(8913620128), b.Uint64())

	resetTestWallet()
}

func TestWalletHelper_GetBlockHeight(t *testing.T) {
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
		Result: 6666666,
	})
	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	wh := NewWalletHelperFromWallet(clientMock, testWallet)

	r, e := wh.GetBlockHeight()
	assert.Nil(t, e)
	assert.Equal(t, uint32(6666665), r)

	resetTestWallet()
}

func TestWalletHelper_GetContractState(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	clientMock.On("GetContractState", mock.Anything).Return(rpc.GetContractStateResponse{
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
		Result: models.RpcContractState{
			Id:            8,
			UpdateCounter: 0,
			Hash:          "0x99042d380f2b754175717bb932a911bc0bb0ad7d",
			Nef: models.RpcNefFile{
				Magic:    0x3346454E,
				Compiler: "neon",
				Tokens: []models.RpcMethodToken{
					{
						Hash:            "0x99042d380f2b754175717bb932a911bc0bb0ad7d",
						Method:          "verify",
						ParametersCount: 0,
						HasReturnValue:  true,
						CallFlags:       "AllowCall",
					},
				},
				Script:   "DBQKo4e1Ppa3mJpjFDGgVt0fQKBC9kH4J+yMQDTkQFcBAAwFSGVsbG9Bm/ZnzkGSXegxcGhAVwQBEnAMF0ludm9rZSBTdHJvYWdlLlB1dCBmb3IgaBpQQXvjun0MByB0aW1lcy6Li9soQc/nR5YMBUhlbGxveFBBm/ZnzkHmPxiEDAJOb0Gb9mfOQZJd6DHYqnNrJiwMAk5vDAJOb0Gb9mfOQZJd6DFK2CYFEFBF2yERnlBBm/ZnzkHmPxiEIhMhDAJObxFQQZv2Z85B5j8YhAwCTm9Bm/ZnzkGSXegxcWlK2CYFEFBF2yEaUEF747p9chXDShAMBFB1dCDQShF40EoSDB0gaW50byBzdG9yYWdlIGNvbXBsZXRlbHkgZm9yINBKE2rQShQMBiB0aW1lc9DBShEyCJ1Ti1Ai+EXbKEHP50eWeBHADARXb3JkQZUBb2FpEcAMDkludm9rZVB1dENvdW50QZUBb2FAVwECNZL+//8Qs3BoJhYMEU5vIGF1dGhvcml6YXRpb24uOnh5UEExxjMdQFcBADVn/v//ELNwaCYWDBFObyBhdXRob3JpemF0aW9uLjohQcafHfBAVgEMFAqjh7U+lreYmmMUMaBW3R9AoEL2YEA=",
				CheckSum: 73195690102,
			},
			Manifest: models.RpcContractManifest{
				Name:               "testContract",
				Groups:             []models.RpcContractGroup{},
				SupportedStandards: []string{},
				Abi: models.RpcContractAbi{
					Methods: []models.RpcContractMethodDescriptor{
						{
							Name:       "verify",
							Parameters: []models.RpcContractParameterDefinition{},
							ReturnType: "Boolean",
							Offset:     28,
							Safe:       true,
						},
						{
							Name: "update",
							Parameters: []models.RpcContractParameterDefinition{
								{
									Name: "script",
									Type: "ByteArray",
								},
								{
									Name: "manifest",
									Type: "String",
								},
							},
							ReturnType: "Void",
							Offset:     363,
							Safe:       false,
						},
					},
					Events: []models.RpcContractEventDescriptor{},
				},
				Permissions: []models.RpcContractPermission{
					{
						Contract: "*",
						Methods:  []string{"*"},
					},
				},
				Trusts: []string{},
			},
		},
	})
	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	wh := NewWalletHelperFromWallet(clientMock, testWallet)

	c, e := wh.GetContractState(testScriptHash)
	assert.Nil(t, e)
	assert.Equal(t, "0x99042d380f2b754175717bb932a911bc0bb0ad7d", c.Hash)

	resetTestWallet()
}

func TestWalletHelper_GetGasConsumed(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	clientMock.On("InvokeScript", mock.Anything, mock.Anything).Return(rpc.InvokeResultResponse{
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
			Script:      "DAABECcMFPqJ+ywU3w9Z3d8E9uVlF/KzSq7rDBTitlMicpPpnE8pBtU1U6u0pnLfhhTAHwwIdHJhbnNmZXIMFIOrBnmtVcBQoTrUP1k26nP16x72QWJ9W1I=",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStack{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})
	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	wh := NewWalletHelperFromWallet(clientMock, testWallet)

	b, e := wh.GetGasConsumed([]byte{}, nil)
	assert.Nil(t, e)
	assert.Equal(t, int64(2007570), b)

	resetTestWallet()
}

func TestWalletHelper_GetUnClaimedGas(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	clientMock.On("GetUnclaimedGas", mock.Anything).Return(rpc.GetUnclaimedGasResponse{
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
		Result: models.UnclaimedGas{
			Unclaimed: "8913620128",
			Address:   "NNU67Fvdy3LEQTM374EJ9iMbCRxVExgM8Y",
		},
	})
	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	wh := NewWalletHelperFromWallet(clientMock, testWallet)

	v, e := wh.GetUnClaimedGas()
	assert.Nil(t, e)
	assert.Equal(t, uint64(8913620128), v)

	resetTestWallet()
}

func TestWalletHelper_MakeTransaction(t *testing.T) {
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
		Result: 6666666,
	})
	clientMock.On("InvokeScript", mock.Anything, mock.Anything).Return(rpc.InvokeResultResponse{
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
			Script:      "DAABECcMFPqJ+ywU3w9Z3d8E9uVlF/KzSq7rDBTitlMicpPpnE8pBtU1U6u0pnLfhhTAHwwIdHJhbnNmZXIMFIOrBnmtVcBQoTrUP1k26nP16x72QWJ9W1I=",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStack{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})
	script := []byte{}
	ab := []AccountAndBalance{
		{
			Account: testScriptHash,
			Value:   big.NewInt(1000000000000), // 10000 gas
		},
	}
	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	wh := NewWalletHelperFromWallet(clientMock, testWallet)

	trx, e := wh.MakeTransaction(script, nil, nil, ab)
	assert.Nil(t, e)
	assert.Equal(t, int64(1141520), trx.GetNetworkFee())

	resetTestWallet()
}

func TestWalletHelper_Sign(t *testing.T) {
	// tested in SignTransaction
}

func TestWalletHelper_SignTransaction(t *testing.T) {
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
		Result: 6666666,
	})
	clientMock.On("InvokeScript", mock.Anything, mock.Anything).Return(rpc.InvokeResultResponse{
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
			Script:      "DAABECcMFPqJ+ywU3w9Z3d8E9uVlF/KzSq7rDBTitlMicpPpnE8pBtU1U6u0pnLfhhTAHwwIdHJhbnNmZXIMFIOrBnmtVcBQoTrUP1k26nP16x72QWJ9W1I=",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStack{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})
	clientMock.On("CalculateNetworkFee", mock.Anything).Return(rpc.CalculateNetworkFeeResponse{
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
		Result: models.RpcNetworkFee{NetworkFee: "2384840"},
	})
	script := []byte{}
	ab := []AccountAndBalance{
		{
			Account: testScriptHash,
			Value:   big.NewInt(1000000000000), // 10000 gas
		},
	}
	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	wh := NewWalletHelperFromWallet(clientMock, testWallet)

	trx, e := wh.MakeTransaction(script, nil, nil, ab)
	assert.Nil(t, e)
	trx, e = wh.SignTransaction(trx, helper.Neo3Magic_MainNet)
	assert.Nil(t, e)
	assert.Equal(t, 1, len(trx.GetWitnesses()))

	resetTestWallet()
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
	clientMock.On("InvokeScript", mock.Anything, mock.Anything).Return(rpc.InvokeResultResponse{
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
			Script:      "DAABECcMFPqJ+ywU3w9Z3d8E9uVlF/KzSq7rDBTitlMicpPpnE8pBtU1U6u0pnLfhhTAHwwIdHJhbnNmZXIMFIOrBnmtVcBQoTrUP1k26nP16x72QWJ9W1I=",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStack{{
				Type:  "Integer",
				Value: "8913620128",
			}},
		},
	})
	clientMock.On("CalculateNetworkFee", mock.Anything).Return(rpc.CalculateNetworkFeeResponse{
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
		Result: models.RpcNetworkFee{NetworkFee: "2384840"},
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
	_ = testWallet.Unlock("")
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	wh := NewWalletHelperFromWallet(clientMock, testWallet)

	to := "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ"
	h, e := wh.Transfer(tx.GasToken, to, big.NewInt(10000), helper.Neo3Magic_MainNet)
	assert.Nil(t, e)
	assert.Equal(t, "0x992f941c9751aabc8bab0200503e07e38f38f0884cb8b11f6c6c72d8d2fb2948", h)

	resetTestWallet()
}
