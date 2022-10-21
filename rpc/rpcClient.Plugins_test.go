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

func TestRpcClient_GetNep11Balances(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"address": "NNBjvfrvPTZRw3Tx2Qwm6bLKqbng4qQ41N",
				"balance": [
					{
						"assethash": "0x1234567890abcdef1234567890abcdef12345678",
						"name": "SampleNFT",
						"symbol": "sNFT",
						"decimals": "1",
						"amount": "9809309981",
						"tokens": [
							{
							"tokenid": "1234567890abcdef",
							"amount": "1",
							"lastupdatedblock": 12345
							}
						]
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.GetNep11Balances("NNBjvfrvPTZRw3Tx2Qwm6bLKqbng4qQ41N")
	r := response.Result
	assert.Equal(t, "NNBjvfrvPTZRw3Tx2Qwm6bLKqbng4qQ41N", r.Address)
	assert.Equal(t, "0x1234567890abcdef1234567890abcdef12345678", r.Balances[0].AssetHash)
	assert.Equal(t, "1234567890abcdef", r.Balances[0].Tokens[0].TokenId)
}

func TestRpcClient_GetNep11Transfers(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"address": "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ",
				"sent": [
					{
						"timestamp": 1578471997998,
						"assethash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
						"transferaddress": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
						"amount": "1",
						"blockindex": 72,
						"transfernotifyindex": 0,
						"txhash": "0xc28763714d06e80f28b431d0a24495f41961b7d2746fc4cdaec0607adf0d6749",
						"tokenid": "1234567890abcdef"
					}
				],
				"received": [
					{
						"timestamp": 1578471121898,
						"assethash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
						"transferaddress": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
						"amount": "1",
						"blockindex": 14,
						"transfernotifyindex": 0,
						"txhash": "0xfc4b8454601e3df8c9ed03765f7860fce4ae2aa3d52e0f4790fd89f208ed051b",
						"tokenid": "1234567890abcdef"
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.GetNep11Transfers("", nil, nil)
	r := response.Result
	assert.Equal(t, "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ", r.Address)
	assert.Equal(t, "1234567890abcdef", r.Received[0].TokenId)
}

func TestRpcClient_GetNep11Properties(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"assethash": "0x1234567890abcdef1234567890abcdef12345678",
				"name": "SampleNFT",
				"symbol": "sNFT",
				"decimals": "1",
				"tokenid": "1234567890abcdef"
			}
		}`))),
	}, nil)

	response := rpc.GetNep11Properties("", "")
	r := response.Result
	assert.Equal(t, "0x1234567890abcdef1234567890abcdef12345678", r["assethash"])
	assert.Equal(t, "1234567890abcdef", r["tokenid"])
}

func TestRpcClient_GetNep17Balances(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"address": "NNBjvfrvPTZRw3Tx2Qwm6bLKqbng4qQ41N",
				"balance": [
					{
						"assethash": "0xcd48b160c1bbc9d74997b803b9a7ad50a4bef020",
						"name": "Nep17Contract",
						"symbol": "fUSDT",
						"decimals": "6",
						"amount": "9809309981",
						"lastupdatedblock": 2362166
					},
					{
						"assethash": "0x78e1330db47634afdb5ea455302ba2d12b8d549f",
						"name": "SWTHToken",
						"symbol": "SWTH",
						"decimals": "8",
						"amount": "245054977090092",
						"lastupdatedblock": 2362166
					},
					{
						"assethash": "0xd2a4cff31913016155e38e474a2c06d08be276cf",
						"name": "GasToken",
						"symbol": "GAS",
						"decimals": "8",
						"amount": "679451069",
						"lastupdatedblock": 2362959
					},
					{
						"assethash": "0x340720c7107ef5721e44ed2ea8e314cce5c130fa",
						"name": "Nudes",
						"symbol": "NUDES",
						"decimals": "8",
						"amount": "37975094447324564869",
						"lastupdatedblock": 2351094
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.GetNep17Balances("NNBjvfrvPTZRw3Tx2Qwm6bLKqbng4qQ41N")
	r := response.Result
	assert.Equal(t, "NNBjvfrvPTZRw3Tx2Qwm6bLKqbng4qQ41N", r.Address)
	assert.Equal(t, "0xcd48b160c1bbc9d74997b803b9a7ad50a4bef020", r.Balances[0].AssetHash)
}

func TestRpcClient_GetNep17Transfers(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"address": "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ",
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
				]
			}
		}`))),
	}, nil)

	response := rpc.GetNep17Transfers("", nil, nil)
	r := response.Result
	assert.Equal(t, "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ", r.Address)
	assert.Equal(t, uint64(1578471997998), r.Sent[0].Timestamp)
	assert.Equal(t, "0xadc751e8fc4e7514cf2fcd623ad78a565985b5701b04961445b3d4794015e19a", r.Received[1].TxHash)
}
