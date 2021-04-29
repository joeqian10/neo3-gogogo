package rpc

import (
	"bytes"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestPopInvokeStack(t *testing.T) {
	r := InvokeResultResponse{
		RpcResponse: RpcResponse{
			JsonRpc: "2.0",
			ID:      1,
		},
		ErrorResponse: ErrorResponse{
			Error: RpcError{
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
	}
	p, e := PopInvokeStack(r)
	assert.Nil(t, e)
	assert.Equal(t, "Integer", p.Type)
}

func TestInvokeStack_Convert(t *testing.T) {
	var client = new(HttpClientMock)
	var rpcClient = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"script": "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9",
				"state": "HALT",
				"gasconsumed": "0.126",
				"exception": null,
				"stack": [
					{
						"type": "Map",
						"value": [
							{
								"key": 
								{
									"type": "Integer",
									"value": "1"
								},
								"value": 
								{
									"type": "Pointer",
									"value": 0
								}
							}
						]
					}
				]
			}
		}`))),
	}, nil)

	response := rpcClient.InvokeScript("", nil)
	assert.False(t, response.HasError())
}

