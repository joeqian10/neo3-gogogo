package nep5

import (
	"github.com/joeqian10/neo3-gogogo/tx"
	"github.com/joeqian10/neo3-gogogo/wallet/keys"
	"math/big"
	"testing"

	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewNep5Helper(t *testing.T) {
	nep5helper := NewNep5Helper(helper.UInt160{}, rpc.NewClient("http://seed1.ngd.network:20332"))
	assert.NotNil(t, nep5helper)
}

func TestNep5Helper_TotalSupply(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
		ScriptHash:helper.UInt160{},
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
			Script:      "10c00c0b746f74616c537570706c790c14897720d8cd76f4f00abfa37c0edd889c208fde9b41627d5b52",
			State:       "HALT",
			GasConsumed: "2007390",
			Stack: []models.InvokeStackResult{
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

func TestNep5Helper_Decimals(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
		ScriptHash:helper.UInt160{},
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
			Script:      "10c00c08646563696d616c730c143b7d3711c6f0ccf9b1dca903d1bfa1d896f1238c41627d5b52 ",
			State:       "HALT",
			GasConsumed: "1007390",
			Stack: []models.InvokeStackResult{
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

func TestNep5Helper_Name(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
		ScriptHash:helper.UInt160{},
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
			Script:      "10c00c046e616d650c14897720d8cd76f4f00abfa37c0edd889c208fde9b41627d5b52 ",
			State:       "HALT",
			GasConsumed: "1007390",
			Stack: []models.InvokeStackResult{
				{
					Type:  "ByteArray",
					Value: "TkVP",
				},
			},
		},
	})

	name, err := nh.Name()
	assert.Nil(t, err)
	assert.Equal(t, "NEO", name)
}

func TestNep5Helper_Symbol(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
		ScriptHash:helper.UInt160{},
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
			Script:      "10c00c0673796d626f6c0c14897720d8cd76f4f00abfa37c0edd889c208fde9b41627d5b52 ",
			State:       "HALT",
			GasConsumed: "1007390",
			Stack: []models.InvokeStackResult{
				{
					Type:  "ByteArray",
					Value: "bmVv",
				},
			},
		},
	})

	symbol, err := nh.Symbol()
	assert.Nil(t, err)
	assert.Equal(t, "neo", symbol)
}

func TestNep5Helper_BalanceOf(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
		ScriptHash:helper.UInt160{},
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

	b, err := nh.BalanceOf(helper.UInt160{})
	assert.Nil(t, err)
	assert.Equal(t, big.NewInt(8913620128), b)
}

func TestNep5Helper_CreateTransferTx(t *testing.T) {
	var clientMock = new(rpc.RpcClientMock)
	var nh = Nep5Helper{
		ScriptHash:helper.UInt160{},
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
	keyPair, _ := keys.NewKeyPairFromWIF("L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z")
	tr, e := nh.CreateTransferTx(keyPair, helper.UInt160{}, big.NewInt(10))
	assert.Nil(t, e)
	assert.Equal(t, 1233+tx.MaxValidUntilBlockIncrement, tr.GetValidUntilBlock())
}

func TestPopInvokeStack(t *testing.T) {
	r := rpc.InvokeResultResponse{
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
	}
	p, e := PopInvokeStack(r)
	assert.Nil(t, e)
	assert.Equal(t, "Integer", p.Type)
}
