package rpc

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestRpcClient_GetApplicationLog(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"txid": "0xd6ea48f1c33defc1815562b3ace4ead99bf33a8ae67b2642cf73c2f192a717e5",
				"executions": [
					{
						"trigger": "Application",
						"vmstate": "HALT",
						"gasconsumed": "9007990",
						"stack": [],
						"notifications": [
							{
								"contract": "0x668e0c1f9d7b70a99dd9e06eadd4c784d641afbc",
								"eventname": "Transfer",
								"state": {
									"type": "Array",
									"value": [
										{
											"type": "Any"
										},
										{
											"type": "ByteString",
											"value": "9S37k0BBDIaRxjEhW0Sk+9lDN4s="
										},
										{
											"type": "Integer",
											"value": "400000000"
										}
									]
								}
							},
							{
								"contract": "0xde5f57d430d3dece511cf975a8d37848cb9e0525",
								"eventname": "Transfer",
								"state": {
									"type": "Array",
									"value": [
										{
											"type": "ByteString",
											"value": "9S37k0BBDIaRxjEhW0Sk+9lDN4s="
										},
										{
											"type": "ByteString",
											"value": "1rSxahaE1EDW2TzNNlNk0rjQEpI="
										},
										{
											"type": "Integer",
											"value": "1"
										}
									]
								}
							}
						]
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.GetApplicationLog("0xd6ea48f1c33defc1815562b3ace4ead99bf33a8ae67b2642cf73c2f192a717e5")
	r := response.Result
	assert.Equal(t, "0xd6ea48f1c33defc1815562b3ace4ead99bf33a8ae67b2642cf73c2f192a717e5", r.TxId)
	assert.Equal(t, "Array", r.Executions[0].Notifications[0].State.Type)
}

func TestRpcClient_GetNep17Balances(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"balance": [
					{
						"assethash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
						"amount": "10000000",
						"lastupdatedblock": 14
					},
					{
						"assethash": "0x9c33bbf2f5afbbc8fe271dd37508acd93573cffc",
						"amount": "9995000000000000",
						"lastupdatedblock": 17145
					}
				],
				"address": "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ"
			}
		}`))),
	}, nil)

	response := rpc.GetNep17Balances("NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ")
	r := response.Result
	assert.Equal(t, "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ", r.Address)
	assert.Equal(t, "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789", r.Balances[0].AssetHash)
}

func TestRpcClient_GetNep17Transfers(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"sent": [
					{
						"timestamp": 1578471997998,
						"assethash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
						"transferaddress": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
						"amount": "10000000",
						"blockindex": 72,
						"transfernotifyindex": 0,
						"txhash": "0xc28763714d06e80f28b431d0a24495f41961b7d2746fc4cdaec0607adf0d6749"
					}
				],
				"received": [
					{
						"timestamp": 1578471121898,
						"assethash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
						"transferaddress": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
						"amount": "10000000",
						"blockindex": 14,
						"transfernotifyindex": 0,
						"txhash": "0xfc4b8454601e3df8c9ed03765f7860fce4ae2aa3d52e0f4790fd89f208ed051b"
					},
					{
						"timestamp": 1578471952619,
						"assethash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
						"transferaddress": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
						"amount": "90000000",
						"blockindex": 69,
						"transfernotifyindex": 0,
						"txhash": "0xadc751e8fc4e7514cf2fcd623ad78a565985b5701b04961445b3d4794015e19a"
					}
				],
				"address": "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ"
			}
		}`))),
	}, nil)

	response := rpc.GetNep17Transfers("", nil, nil)
	r := response.Result
	assert.Equal(t, "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ", r.Address)
	assert.Equal(t, 1578471997998, r.Sent[0].Timestamp)
	assert.Equal(t, "0xadc751e8fc4e7514cf2fcd623ad78a565985b5701b04961445b3d4794015e19a", r.Received[1].TxHash)
}
