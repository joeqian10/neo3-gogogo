package nep17

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math/big"
	"testing"
)

func TestNewNep17Helper(t *testing.T) {
	nep17helper := NewNep17Helper(helper.NewUInt160(), rpc.NewClient("http://seed1.ngd.network:20332"))
	assert.NotNil(t, nep17helper)
}

func TestNep17Helper_Symbol(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep17Helper{
		ScriptHash: helper.NewUInt160(),
		Client:     clientMock,
	}
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
			Script:      "10c00c0673796d626f6c0c14897720d8cd76f4f00abfa37c0edd889c208fde9b41627d5b52 ",
			State:       "HALT",
			GasConsumed: "1007390",
			Stack: []models.InvokeStack{
				{
					Type:  "ByteString",
					Value: "bmVv",
				},
			},
		},
	})

	symbol, err := nh.Symbol()
	assert.Nil(t, err)
	assert.Equal(t, "neo", symbol)
}

func TestNep17Helper_Decimals(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep17Helper{
		ScriptHash: helper.NewUInt160(),
		Client:     clientMock,
	}
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
			Script:      "10c00c08646563696d616c730c143b7d3711c6f0ccf9b1dca903d1bfa1d896f1238c41627d5b52 ",
			State:       "HALT",
			GasConsumed: "1007390",
			Stack: []models.InvokeStack{
				{
					Type:  "Integer",
					Value: "8",
				},
			},
		},
	})

	d, err := nh.Decimals()
	assert.Nil(t, err)
	assert.Equal(t, 8, d)
}


func TestNep17Helper_TotalSupply(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep17Helper{
		ScriptHash: helper.NewUInt160(),
		Client:     clientMock,
	}
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
			Script:      "10c00c0b746f74616c537570706c790c14897720d8cd76f4f00abfa37c0edd889c208fde9b41627d5b52",
			State:       "HALT",
			GasConsumed: "2007390",
			Stack: []models.InvokeStack{
				{
					Type:  "Integer",
					Value: "100000000",
				},
			},
		},
	})

	s, e := nh.TotalSupply()
	assert.Nil(t, e)
	assert.Equal(t, big.NewInt(100000000), s)
}

func TestNep17Helper_BalanceOf(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep17Helper{
		ScriptHash: helper.NewUInt160(),
		Client:     clientMock,
	}
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
			Script:      "0c146925aa554712439a9c613ba114efa3fac23ddbca11c00c0962616c616e63654f660c143b7d3711c6f0ccf9b1dca903d1bfa1d896f1238c41627d5b52",
			State:       "HALT",
			GasConsumed: "2007570",
			Stack: []models.InvokeStack{
				{
					Type:  "Integer",
					Value: "8913620128",
				},
			},
		},
	})

	b, err := nh.BalanceOf(helper.NewUInt160())
	assert.Nil(t, err)
	assert.Equal(t, big.NewInt(8913620128), b)
}
