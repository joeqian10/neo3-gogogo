package tx

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewTransactionBuilder(t *testing.T) {
	tb := NewTransactionBuilder("http://seed1.ngd.network:20332")
	if tb == nil {
		t.Fail()
	}
	assert.Equal(t, "http://seed1.ngd.network:20332", tb.EndPoint)
}

func TestTransactionBuilder_GetBlockHeight(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
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

func TestTransactionBuilder_CalculateNetWorkFee(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
	}

	sb := sc.NewScriptBuilder()
	_ = sb.EmitPushBytes(helper.HexTobytes("025e5eb8e89ab16cda6e5f646de54a8e9e9e8ce0a64e44db6b6ffeff8a6369f268"))
	_ = sb.Emit(sc.PUSHNULL)
	_ = sb.EmitSysCall(sc.ECDsaVerify.ToInteropMethodHash())
	script := sb.ToArray()
	var size int = 0
	fee := tb.CalculateNetWorkFee(script, &size)

	assert.Equal(t, int64(180+180+30+1000000), fee)
	assert.Equal(t, 67+1+41, size)
}

func TestTransactionBuilder_GetBalance(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
	}
	clientMock.On("GetNep5Balances", mock.Anything).Return(rpc.GetNep5BalancesResponse{
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
		Result: models.RpcNep5Balances{
			Balances: []models.Nep5Balance{
				{
					AssetHash:        "9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
					Amount:           10000,
					LastUpdatedBlock: 123456,
				},
				{
					AssetHash:        "8c23f196d8a1bfd103a9dcb1f9ccf0c611377d3b",
					Amount:           8000,
					LastUpdatedBlock: 135790,
				},
			},
			Address: "AGofsxAUDwt52KjaB664GYsqVAkULYvKNt",
		},
	})

	a, e := tb.GetBalance(helper.UInt160{}, GasToken)
	assert.Nil(t, e)
	assert.Equal(t, int64(8000), a)
}

func TestTransactionBuilder_GetGasConsumed(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
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
		EndPoint: "",
		Client:   clientMock,
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
			Hash: "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
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
}