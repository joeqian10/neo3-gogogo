package tx

import (
	"github.com/joeqian10/neo3-gogogo/wallet/keys"
	"math/big"
	"testing"

	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewTransactionBuilder(t *testing.T) {
	tb := NewTransactionBuilder("http://seed1.ngd.network:20332")
	if tb == nil {
		t.Fail()
	}
}

func TestNewTransactionBuilderFromClient(t *testing.T) {
	client := rpc.NewClient("http://seed1.ngd.network:20332")
	tb := NewTransactionBuilderFromClient(client)
	assert.NotNil(t, tb)
}

func TestTransactionBuilder_GetBlockHeight(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		Client: clientMock,
	}
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

	a, e := tb.GetBlockHeight()
	assert.Nil(t, nil, e)
	assert.Equal(t, uint32(1233), a)
}

func TestTransactionBuilder_CalculateNetworkFee(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		Client: clientMock,
	}

	sb := sc.NewScriptBuilder()
	_ = sb.EmitPushBytes(helper.HexTobytes("025e5eb8e89ab16cda6e5f646de54a8e9e9e8ce0a64e44db6b6ffeff8a6369f268"))
	_ = sb.Emit(sc.PUSHNULL)
	_ = sb.EmitSysCall(sc.ECDsaVerify.ToInteropMethodHash())
	script := sb.ToArray()
	var size int = 0
	fee := tb.CalculateNetworkFee(script, &size)

	assert.Equal(t, int64(180+180+30+1000000), fee)
	assert.Equal(t, 67+1+41, size)
}

func TestTransactionBuilder_GetBalance(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		Client: clientMock,
	}
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

	a, e := tb.GetBalance(GasToken, helper.UInt160{})
	assert.Nil(t, e)
	assert.Equal(t, big.NewInt(8913620128), a)
}

func TestTransactionBuilder_GetGasConsumed(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		Client: clientMock,
	}
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
			Script:      "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9",
			State:       "HALT",
			GasConsumed: "12600000",
			Stack: []models.InvokeStackResult{
				{
					Type:  "ByteArray",
					Value: "516c696e6b20546f6b656e",
				},
			},
		},
	})

	f, e := tb.GetGasConsumed([]byte{})
	assert.Nil(t, e)
	assert.Equal(t, int64(12600000), f) // 10 gas free limit
}

func TestTransactionBuilder_GetWitnessScript(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		Client: clientMock,
	}
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
		Result: models.ContractState{
			Hash:   "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
			Script: "QUXEkoQ=",
		},
	})

	u, _ := helper.UInt160FromString("9bde8f209c88dd0e7ca3bf0af0f476cdd8207789")
	b, e := tb.GetWitnessScript(u)
	assert.Nil(t, e)
	assert.Equal(t, []byte{0x41, 0x45, 0xc4, 0x92, 0x84}, b)
}

func TestTransactionBuilder_MakeTransaction(t *testing.T) {
	//todo
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		Client: clientMock,
	}
	// GetBlockHeight
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
	// GetGasConsumed
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
			Script:      "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9",
			State:       "HALT",
			GasConsumed: "12600000",
			Stack: []models.InvokeStackResult{
				{
					Type:  "Integer",
					Value: "1000000000",
				},
			},
		},
	})
	// GetWitnessScript
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
		Result: models.ContractState{
			Hash:   "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
			Script: "QUXEkoQ=",
		},
	})

	sender, _ := helper.AddressToScriptHash("NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ")
	operation := "name"
	sb := sc.NewScriptBuilder()
	_ = sb.EmitAppCall(NeoToken, operation, nil)
	script := sb.ToArray()
	tx, err := tb.MakeTransaction(script, sender, nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, sender.String(), tx.GetSender().String())
}

func TestTransactionBuilder_CalculateNetworkFeeWithSignStore(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		Client: clientMock,
	}
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
		Result: models.ContractState{
			Hash:   "",
			Script: "",
		},
	})

	tb.Tx = NewTransaction()
	networkFee := tb.CalculateNetworkFeeWithSignStore()

	assert.Equal(t, int64(45+ // header size
		1+ // attributes
		1+ // cosigners
		1+ // script
		1)*1000, // hashes
		networkFee)
}

func TestTransactionBuilder_AddSignature(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		Client: clientMock,
	}
	tb.Tx = NewTransaction()
	sender, _ := helper.UInt160FromString("46a6b608c2b729356b6fa1171c9e33cd0a203f0c")
	tb.Tx.SetSender(sender)
	keyPair, _ := keys.NewKeyPairFromWIF("L3Hab7wL43SbWLnkfnVCp6dT99xzfB4qLZxeR9dFgcWWPirwKyXp")
	err := tb.AddSignature(keyPair)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(tb.SignStore))
}

func TestTransactionBuilder_AddMultiSig(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		Client: clientMock,
	}
	tb.Tx = NewTransaction()
	sender, _ := helper.UInt160FromString("866a29833628627a9e410d5215852737f6f0937d")
	tb.Tx.SetSender(sender)
	cosigner1, _ := helper.UInt160FromString(keys.KeyCases[0].ScriptHash)
	cosigner2, _ := helper.UInt160FromString(keys.KeyCases[1].ScriptHash)
	cosigner3, _ := helper.UInt160FromString(keys.KeyCases[2].ScriptHash)
	cosigners := []*Cosigner{NewCosigner(cosigner1),NewCosigner(cosigner2),NewCosigner(cosigner3)}

	tb.Tx.SetCosigners(cosigners)

	keyPair, _ := keys.NewKeyPairFromWIF("L3Hab7wL43SbWLnkfnVCp6dT99xzfB4qLZxeR9dFgcWWPirwKyXp")
	keyPair1, _ := keys.NewKeyPairFromWIF(keys.KeyCases[0].Wif)
	keyPair2, _ := keys.NewKeyPairFromWIF(keys.KeyCases[1].Wif)
	keyPair3, _ := keys.NewKeyPairFromWIF(keys.KeyCases[2].Wif)
	keyPairs := []*keys.KeyPair{keyPair, keyPair1, keyPair2, keyPair3}
	pubKeys := []*keys.PublicKey{keyPair.PublicKey, keyPair1.PublicKey, keyPair2.PublicKey, keyPair3.PublicKey}
	err := tb.AddMultiSig(keyPairs, 3, pubKeys)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(tb.SignStore))
	assert.Equal(t, 4, len(tb.SignStore[0].KeyPairs))
}

func TestTransactionBuilder_AddSignItem(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		Client: clientMock,
	}
	tb.Tx = NewTransaction()
	sender, _ := helper.UInt160FromString("46a6b608c2b729356b6fa1171c9e33cd0a203f0c")
	tb.Tx.SetSender(sender)
	keyPair, _ := keys.NewKeyPairFromWIF("L3Hab7wL43SbWLnkfnVCp6dT99xzfB4qLZxeR9dFgcWWPirwKyXp")
	contract := keys.CreateSignatureContract(keyPair.PublicKey)
	err := tb.AddSignItem(contract, keyPair)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(tb.SignStore))
}

func TestTransactionBuilder_Sign(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		Client: clientMock,
	}
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
		Result: models.ContractState{
			Hash:   "",
			Script: "",
		},
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
	sender, _ := helper.UInt160FromString("46a6b608c2b729356b6fa1171c9e33cd0a203f0c")
	tb.Tx = NewTransaction()
	tb.Tx.SetSender(sender)
	keyPair, _ := keys.NewKeyPairFromWIF("L3Hab7wL43SbWLnkfnVCp6dT99xzfB4qLZxeR9dFgcWWPirwKyXp")
	_ = tb.AddSignature(keyPair)
	err := tb.Sign()
	assert.Nil(t, err)
}
