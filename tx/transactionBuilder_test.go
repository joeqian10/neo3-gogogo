package tx

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
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

func TestTransactionBuilder_GetBalance(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var tb = TransactionBuilder{
		EndPoint: "",
		Client:   clientMock,
	}
	clientMock.On("GetNep5Balances", mock.Anything).Return(rpc.GetNep5BalancesResponse{
		RpcResponse: rpc.RpcResponse{
			JsonRpc: "2.0",
			ID: 1,
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
					AssetHash:   "9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
					Amount:      10000,
					LastUpdatedBlock: 123456,
				},
				{
					AssetHash:   "8c23f196d8a1bfd103a9dcb1f9ccf0c611377d3b",
					Amount:      8000,
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
		RpcResponse:   rpc.RpcResponse{
			JsonRpc: "2.0",
			ID: 1,
		},
		ErrorResponse: rpc.ErrorResponse{
			Error: rpc.RpcError{
				Code:    0,
				Message: "",
			},
		},
		Result:        models.InvokeResult{
			Script:"00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9",
			State:"HALT",
			GasConsumed:"12600000",
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

// todo
