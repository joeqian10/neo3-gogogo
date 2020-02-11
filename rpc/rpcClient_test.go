package rpc

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type HttpClientMock struct {
	mock.Mock
}

func (c *HttpClientMock) Do(req *http.Request) (*http.Response, error) {
	args := c.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestNewClient(t *testing.T) {
	rpcClient := NewClient("http://seed1.ngd.network:20332")
	assert.NotNil(t, rpcClient)
	endpoint := rpcClient.Endpoint
	assert.NotNil(t, endpoint)
	assert.Equal(t, "seed1.ngd.network:20332", endpoint.Host)
	assert.Equal(t, "http", endpoint.Scheme)
}

func TestRpcClient_GetBestBlockHash(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "0x340d56b071af90ad3fdc3f61c829287930e7f71852753165ddd4d717e417188b"
		}`))),
	}, nil)

	response := rpc.GetBestBlockHash()
	r := response.Result
	assert.Equal(t, "0x340d56b071af90ad3fdc3f61c829287930e7f71852753165ddd4d717e417188b", r)
}

func TestRpcClient_GetBlock(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"hash": "0x1329b78cbdcded8058d4f65c0f1f63fa79c2a4ed5fa266951734018f587f7835",
				"size": 492,
				"version": 0,
				"previousblockhash": "0x991cb1c359cdcf8129b5bc54b4c4fc8345ac17927d4825bcda4d6a8c46dcfb78",
				"merkleroot": "0xf2eb105cc5608fe076e563cb40a4e1593df9bd93a954c069fca2ad74e52f9ece",
				"time": 1578382911810,
				"index": 409,
				"nextconsensus": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
				"witnesses": [
					{
						"invocation": "DECziWRkZbbQVdakM0H1VRPq5V5+ZnWwoUlcuSZBvvP65/DiA3KHFu8mNMDvVfyAv9Q//4TI84gpscuzt3z4Ipc/",
						"verification": "EQwhA6oFL7y45bM6Tu/WYlNvhoRkHwQQnx1eac3abwhIkChqEQtBMHOzuw=="
					}
				],
				"consensus_data": {
					"primary": 0,
					"nonce": "75b6d0184621d863"
				},
				"tx": [
					{
						"hash": "0x3d39da5b227e3f02f5210b24690a0523162788668e490363c6a39813bb162e51",
						"size": 270,
						"version": 0,
						"nonce": 1233336052,
						"sender": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
						"sys_fee": "100000000",
						"net_fee": "1270450",
						"valid_until_block": 2102808,
						"attributes": [],
						"cosigners": [
							{
								"account": "0x2916eba24e652fa006f3e5eb8f9892d2c3b00399",
								"scopes": "CalledByEntry"
							}
						],
						"script": "AoCWmAAMFGklqlVHEkOanGE7oRTvo/rCPdvKDBSZA7DD0pKYj+vl8wagL2VOousWKRPADAh0cmFuc2ZlcgwUiXcg2M129PAKv6N8Dt2InCCP3ptBYn1bUjk=",
						"witnesses": [
							{
								"invocation": "DEDcGjmiHJ22R4LjUuXOF83UDtJB3FUZPy4t8Ol+dSpQovI9KAfVVOrtz/NZBmEuVGXiALkJU6vklZ9XzzDrz0PJ",
								"verification": "EQwhA6oFL7y45bM6Tu/WYlNvhoRkHwQQnx1eac3abwhIkChqEQtBMHOzuw=="
							}
						]
					}
				],
				"confirmations": 2,
				"nextblockhash": "0xb59b28a806d8f30a8c71195500bf3e834238df2fc79fd5f984e516737c8bb3cd"
			}
		}`))),
	}, nil)

	response := rpc.GetBlock("409")
	r := response.Result
	assert.Equal(t, "0x1329b78cbdcded8058d4f65c0f1f63fa79c2a4ed5fa266951734018f587f7835", r.Hash)
	assert.Equal(t, "DECziWRkZbbQVdakM0H1VRPq5V5+ZnWwoUlcuSZBvvP65/DiA3KHFu8mNMDvVfyAv9Q//4TI84gpscuzt3z4Ipc/", r.Witnesses[0].Invocation)
	assert.Equal(t, "DEDcGjmiHJ22R4LjUuXOF83UDtJB3FUZPy4t8Ol+dSpQovI9KAfVVOrtz/NZBmEuVGXiALkJU6vklZ9XzzDrz0PJ", r.Tx[0].Witnesses[0].Invocation)
}

func TestRpcClient_GetBlockCount(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": 2023
		}`))),
	}, nil)

	response := rpc.GetBlockCount()
	r := response.Result
	assert.Equal(t, 2023, r)
}

func TestRpcClient_GetBlockHash(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "0x903fe39c902ff244d3e55ed0aecc92eb675ac9616f1840d4a212671fa1c8b697"
		}`))),
	}, nil)

	response := rpc.GetBlockHash(1)
	r := response.Result
	assert.Equal(t, "0x903fe39c902ff244d3e55ed0aecc92eb675ac9616f1840d4a212671fa1c8b697", r)
}

func TestRpcClient_GetBlockHeader(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"hash": "0x1329b78cbdcded8058d4f65c0f1f63fa79c2a4ed5fa266951734018f587f7835",
				"size": 213,
				"version": 0,
				"previousblockhash": "0x991cb1c359cdcf8129b5bc54b4c4fc8345ac17927d4825bcda4d6a8c46dcfb78",
				"merkleroot": "0xf2eb105cc5608fe076e563cb40a4e1593df9bd93a954c069fca2ad74e52f9ece",
				"time": 1578382911810,
				"index": 409,
				"nextconsensus": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
				"witnesses": [
					{
						"invocation": "DECziWRkZbbQVdakM0H1VRPq5V5+ZnWwoUlcuSZBvvP65/DiA3KHFu8mNMDvVfyAv9Q//4TI84gpscuzt3z4Ipc/",
						"verification": "EQwhA6oFL7y45bM6Tu/WYlNvhoRkHwQQnx1eac3abwhIkChqEQtBMHOzuw=="
					}
				],
				"confirmations": 17,
				"nextblockhash": "0xb59b28a806d8f30a8c71195500bf3e834238df2fc79fd5f984e516737c8bb3cd"
			}
		}`))),
	}, nil)

	response := rpc.GetBlockHeader("409")
	r := response.Result
	assert.Equal(t, "0x991cb1c359cdcf8129b5bc54b4c4fc8345ac17927d4825bcda4d6a8c46dcfb78", r.PreviousBlockHash)
	assert.Equal(t, "EQwhA6oFL7y45bM6Tu/WYlNvhoRkHwQQnx1eac3abwhIkChqEQtBMHOzuw==", r.Witnesses[0].Verification)
}

func TestRpcClient_GetBlockSysFee(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "100000000"
		}`))),
	}, nil)

	response := rpc.GetBlockSysFee(1)
	r := response.Result
	assert.Equal(t, "100000000", r)
}

func TestRpcClient_GetContractState(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"hash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
				"script": "QUXEkoQ=",
				"manifest": {
					"groups": [],
					"features": {
						"storage": true,
						"payable": false
					},
					"abi": {
						"hash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
						"entryPoint": {
							"name": "Main",
							"parameters": [
								{
									"name": "operation",
									"type": "String"
								},
								{
									"name": "args",
									"type": "Array"
								}
							],
							"returnType": "Any"
						},
						"methods": [
							{
								"name": "unclaimedGas",
								"parameters": [
									{
										"name": "account",
										"type": "Hash160"
									},
									{
										"name": "end",
										"type": "Integer"
									}
								],
								"returnType": "Integer"
							},
							{
								"name": "registerValidator",
								"parameters": [
									{
										"name": "pubkey",
										"type": "PublicKey"
									}
								],
								"returnType": "Boolean"
							},
							{
								"name": "vote",
								"parameters": [
									{
										"name": "account",
										"type": "Hash160"
									},
									{
										"name": "pubkeys",
										"type": "Array"
									}
								],
								"returnType": "Boolean"
							},
							{
								"name": "getRegisteredValidators",
								"parameters": [],
								"returnType": "Array"
							},
							{
								"name": "getValidators",
								"parameters": [],
								"returnType": "Array"
							},
							{
								"name": "getNextBlockValidators",
								"parameters": [],
								"returnType": "Array"
							},
							{
								"name": "name",
								"parameters": [],
								"returnType": "String"
							},
							{
								"name": "symbol",
								"parameters": [],
								"returnType": "String"
							},
							{
								"name": "decimals",
								"parameters": [],
								"returnType": "Integer"
							},
							{
								"name": "totalSupply",
								"parameters": [],
								"returnType": "Integer"
							},
							{
								"name": "balanceOf",
								"parameters": [
									{
										"name": "account",
										"type": "Hash160"
									}
								],
								"returnType": "Integer"
							},
							{
								"name": "transfer",
								"parameters": [
									{
										"name": "from",
										"type": "Hash160"
									},
									{
										"name": "to",
										"type": "Hash160"
									},
									{
										"name": "amount",
										"type": "Integer"
									}
								],
								"returnType": "Boolean"
							},
							{
								"name": "onPersist",
								"parameters": [],
								"returnType": "Boolean"
							},
							{
								"name": "supportedStandards",
								"parameters": [],
								"returnType": "Array"
							}
						],
						"events": [
							{
								"name": "Transfer",
								"parameters": [
									{
										"name": "from",
										"type": "Hash160"
									},
									{
										"name": "to",
										"type": "Hash160"
									},
									{
										"name": "amount",
										"type": "Integer"
									}
								],
								"returnType": "Signature"
							}
						]
					},
					"permissions": [
						{
							"contract": "*",
							"methods": "*"
						}
					],
					"trusts": [],
					"safeMethods": [
						"unclaimedGas",
						"getRegisteredValidators",
						"getValidators",
						"getNextBlockValidators",
						"name",
						"symbol",
						"decimals",
						"totalSupply",
						"balanceOf",
						"supportedStandards"
					],
					"extra": null
				}
			}
		}`))),
	}, nil)

	response := rpc.GetContractState("0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789")
	r := response.Result
	assert.Equal(t, "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789", r.Hash)
	assert.Equal(t, "args", r.Manifest.Abi.EntryPoint.Parameters[1].Name)
	assert.Equal(t, false, r.Manifest.Features.Payable)
}

func TestRpcClient_GetRawMemPool(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": ["0x903fe39c902ff244d3e55ed0aecc92eb675ac9616f1840d4a212671fa1c8b697"]
		}`))),
	}, nil)

	response := rpc.GetRawMemPool()
	r := response.Result
	assert.Equal(t, "0x903fe39c902ff244d3e55ed0aecc92eb675ac9616f1840d4a212671fa1c8b697", r[0])
}

func TestRpcClient_GetRawTransaction(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"hash": "0x3d39da5b227e3f02f5210b24690a0523162788668e490363c6a39813bb162e51",
				"size": 270,
				"version": 0,
				"nonce": 1233336052,
				"sender": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
				"sys_fee": "100000000",
				"net_fee": "1270450",
				"valid_until_block": 2102808,
				"attributes": [],
				"cosigners": [
					{
						"account": "0x2916eba24e652fa006f3e5eb8f9892d2c3b00399",
						"scopes": "CalledByEntry"
					}
				],
				"script": "AoCWmAAMFGklqlVHEkOanGE7oRTvo/rCPdvKDBSZA7DD0pKYj+vl8wagL2VOousWKRPADAh0cmFuc2ZlcgwUiXcg2M129PAKv6N8Dt2InCCP3ptBYn1bUjk=",
				"witnesses": [
					{
						"invocation": "DEDcGjmiHJ22R4LjUuXOF83UDtJB3FUZPy4t8Ol+dSpQovI9KAfVVOrtz/NZBmEuVGXiALkJU6vklZ9XzzDrz0PJ",
						"verification": "EQwhA6oFL7y45bM6Tu/WYlNvhoRkHwQQnx1eac3abwhIkChqEQtBMHOzuw=="
					}
				],
				"blockhash": "0x1329b78cbdcded8058d4f65c0f1f63fa79c2a4ed5fa266951734018f587f7835",
				"confirmations": 56,
				"blocktime": 1578382911810,
				"vmState": "HALT"
			}
		}`))),
	}, nil)

	response := rpc.GetRawTransaction("0x3d39da5b227e3f02f5210b24690a0523162788668e490363c6a39813bb162e51")
	r := response.Result
	assert.Equal(t, "DEDcGjmiHJ22R4LjUuXOF83UDtJB3FUZPy4t8Ol+dSpQovI9KAfVVOrtz/NZBmEuVGXiALkJU6vklZ9XzzDrz0PJ", r.Witnesses[0].Invocation)
}

func TestRpcClient_GetStorage(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "410321048096980021020702280100"
		}`))),
	}, nil)

	response := rpc.GetStorage("0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789", "146925aa554712439a9c613ba114efa3fac23ddbca")
	r := response.Result
	assert.Equal(t, "410321048096980021020702280100", r)
}

func TestRpcClient_GetTransactionHeight(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": 409
		}`))),
	}, nil)

	response := rpc.GetTransactionHeight("0x3d39da5b227e3f02f5210b24690a0523162788668e490363c6a39813bb162e51")
	r := response.Result
	assert.Equal(t, 409, r)
}

func TestRpcClient_GetValidators(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": [
				{
					"publickey": "03aa052fbcb8e5b33a4eefd662536f8684641f04109f1d5e69cdda6f084890286a",
					"votes": "0",
					"active": true
				}
			]
		}`))),
	}, nil)

	response := rpc.GetValidators()
	r := response.Result
	assert.Equal(t, "03aa052fbcb8e5b33a4eefd662536f8684641f04109f1d5e69cdda6f084890286a", r[0].PublicKey)
}

// Node
func TestRpcClient_GetConnectionCount(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": 0
		}`))),
	}, nil)

	response := rpc.GetConnectionCount()
	r := response.Result
	assert.Equal(t, 0, r)
}

func TestRpcClient_GetPeers(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"unconnected": [],
				"bad": [],
				"connected": []
			}
		}`))),
	}, nil)

	response := rpc.GetPeers()
	r := response.Result
	assert.Equal(t, 0, len(r.Bad))
}

func TestRpcClient_GetVersion(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"tcpPort": 20333,
				"wsPort": 20334,
				"nonce": 1254705570,
				"useragent": "/Neo:3.0.0-preview1/"
			}
		}`))),
	}, nil)

	response := rpc.GetVersion()
	r := response.Result
	assert.Equal(t, 20333, r.TcpPort)
}

func TestRpcClient_SendRawTransaction(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"hash": "0x407c75ed84bc2cb70303fbdb791d45b56ccef7209813c53da1c2456c1241294a"
			}
		}`))),
	}, nil)

	response := rpc.SendRawTransaction("0090ab75156925aa554712439a9c613ba114efa3fac23ddbca00e1f50500000000ee4e130000000000b916200000016925aa554712439a9c613ba114efa3fac23ddbca015600640c146925aa554712439a9c613ba114efa3fac23ddbca0c146925aa554712439a9c613ba114efa3fac23ddbca13c00c087472616e736665720c14897720d8cd76f4f00abfa37c0edd889c208fde9b41627d5b523901420c40edb51c35ac1df891aa4683d1742378df15ea0564bc4e9b2f64af3fcf3dea9823f4de9ad695fdd457742cdde0fef27883e6d05a041e96867dc95d9f5ee6de7838290c2103aa052fbcb8e5b33a4eefd662536f8684641f04109f1d5e69cdda6f084890286a0b410a906ad4")
	r := response.Result
	assert.Equal(t, "0x407c75ed84bc2cb70303fbdb791d45b56ccef7209813c53da1c2456c1241294a", r.Hash)
}

func TestRpcClient_SubmitBlock(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"hash": "0x407c75ed84bc2cb70303fbdb791d45b56ccef7209813c53da1c2456c1241294a"
			}
		}`))),
	}, nil)

	response := rpc.SubmitBlock("")
	r := response.Result
	assert.Equal(t, "0x407c75ed84bc2cb70303fbdb791d45b56ccef7209813c53da1c2456c1241294a", r.Hash)
}

// SmartContract
func TestRpcClient_InvokeFunction(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"script": "10c00c08646563696d616c730c143b7d3711c6f0ccf9b1dca903d1bfa1d896f1238c41627d5b52",
				"state": "HALT",
				"gas_consumed": "1007390",
				"stack": [
					{
						"type": "Integer",
						"value": "8"
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.InvokeFunction("0x8c23f196d8a1bfd103a9dcb1f9ccf0c611377d3b", "decimals")
	r := response.Result
	assert.Equal(t, "HALT", r.State)
	assert.Equal(t, "8", r.Stack[0].Value)
}

func TestRpcClient_InvokeScript(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"script": "10c00c08646563696d616c730c143b7d3711c6f0ccf9b1dca903d1bfa1d896f1238c41627d5b52",
				"state": "HALT",
				"gas_consumed": "1007390",
				"stack": [
					{
						"type": "Integer",
						"value": "8"
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.InvokeScript("")
	r := response.Result
	assert.Equal(t, "HALT", r.State)
	assert.Equal(t, "8", r.Stack[0].Value)
}

// Utilities
func TestRpcClient_ListPlugins(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": [
				{
					"name": "LevelDBStore",
					"version": "3.0.0.0",
					"interfaces": [
						"IStoragePlugin"
					]
				},
				{
					"name": "RestServer",
					"version": "1.0.0.0",
					"interfaces": []
				},
				{
					"name": "RpcServer",
					"version": "3.0.0.0",
					"interfaces": []
				}
			]
		}`))),
	}, nil)

	response := rpc.ListPlugins()
	r := response.Result
	assert.Equal(t, 3, len(r))
	assert.Equal(t, "RpcServer", r[2].Name)
}

func TestRpcClient_ValidateAddress(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"address": "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ",
				"isvalid": true
			}
		}`))),
	}, nil)

	response := rpc.ValidateAddress("NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ")
	r := response.Result
	assert.Equal(t, true, r.IsValid)
}

// Plugins
func TestRpcClient_GetApplicationLog(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"txid": "0x797b6ad7c901c1a7a1506a90955759cf14f7f9fdf9f09e4e7d8902887f8ac709",
				"trigger": "Application",
				"vmstate": "HALT",
				"gas_consumed": "9007810",
				"stack": [],
				"notifications": [
					{
						"contract": "0x8c23f196d8a1bfd103a9dcb1f9ccf0c611377d3b",
						"state": {
							"type": "Array",
							"value": [
								{
									"type": "ByteArray",
									"value": "VHJhbnNmZXI="
								},
								{
									"type": "Any"
								},
								{
									"type": "ByteArray",
									"value": "aSWqVUcSQ5qcYTuhFO+j+sI928o="
								},
								{
									"type": "Integer",
									"value": "12070000000"
								}
							]
						}
					},
					{
						"contract": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
						"state": {
							"type": "Array",
							"value": [
								{
									"type": "ByteArray",
									"value": "VHJhbnNmZXI="
								},
								{
									"type": "ByteArray",
									"value": "aSWqVUcSQ5qcYTuhFO+j+sI928o="
								},
								{
									"type": "ByteArray",
									"value": "aSWqVUcSQ5qcYTuhFO+j+sI928o="
								},
								{
									"type": "Integer",
									"value": "100"
								}
							]
						}
					}
				]
			}
		}`))),
	}, nil)

	response := rpc.GetApplicationLog("0x797b6ad7c901c1a7a1506a90955759cf14f7f9fdf9f09e4e7d8902887f8ac709")
	r := response.Result
	assert.Equal(t, "0x797b6ad7c901c1a7a1506a90955759cf14f7f9fdf9f09e4e7d8902887f8ac709", r.TxId)
	assert.Equal(t, "Array", r.Notifications[1].State.Type)
}

func TestRpcClient_GetNep5Balances(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"balance": [
					{
						"asset_hash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
						"amount": "10000000",
						"last_updated_block": 14
					}
				],
				"address": "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ"
			}
		}`))),
	}, nil)

	response := rpc.GetNep5Balances("")
	r := response.Result
	assert.Equal(t, "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ", r.Address)
	assert.Equal(t, "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789", r.Balances[0].AssetHash)
}

func TestRpcClient_GetNep5Transfers(t *testing.T) {
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
						"asset_hash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
						"transfer_address": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
						"amount": "10000000",
						"block_index": 72,
						"transfer_notify_index": 0,
						"tx_hash": "0xc28763714d06e80f28b431d0a24495f41961b7d2746fc4cdaec0607adf0d6749"
					}
				],
				"received": [
					{
						"timestamp": 1578471121898,
						"asset_hash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
						"transfer_address": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
						"amount": "10000000",
						"block_index": 14,
						"transfer_notify_index": 0,
						"tx_hash": "0xfc4b8454601e3df8c9ed03765f7860fce4ae2aa3d52e0f4790fd89f208ed051b"
					},
					{
						"timestamp": 1578471952619,
						"asset_hash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
						"transfer_address": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
						"amount": "90000000",
						"block_index": 69,
						"transfer_notify_index": 0,
						"tx_hash": "0xadc751e8fc4e7514cf2fcd623ad78a565985b5701b04961445b3d4794015e19a"
					}
				],
				"address": "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ"
			}
		}`))),
	}, nil)

	response := rpc.GetNep5Transfers("", nil, nil)
	r := response.Result
	assert.Equal(t, "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ", r.Address)
	assert.Equal(t, 1578471997998, r.Sent[0].Timestamp)
	assert.Equal(t, "0xadc751e8fc4e7514cf2fcd623ad78a565985b5701b04961445b3d4794015e19a", r.Received[1].TxHash)
}

// Wallet
func TestRpcClient_CloseWallet(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": true
		}`))),
	}, nil)

	response := rpc.CloseWallet()
	r := response.Result
	assert.Equal(t, true, r)
}

func TestRpcClient_DumpPrivKey(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "KyoYyZpoccbR6KZ25eLzhMTUxREwCpJzDsnuodGTKXSG8fDW9t7x"
		}`))),
	}, nil)

	response := rpc.DumpPrivKey("")
	r := response.Result
	assert.Equal(t, "KyoYyZpoccbR6KZ25eLzhMTUxREwCpJzDsnuodGTKXSG8fDW9t7x", r)
}

func TestRpcClient_GetBalance(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
			  "balance": "3001101329992600"
			}
		  }`))),
	}, nil)

	response := rpc.GetBalance("0x8c23f196d8a1bfd103a9dcb1f9ccf0c611377d3b")
	r := response.Result
	assert.Equal(t, "3001101329992600", r.Balance)
}

func TestRpcClient_GetNewAddress(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "NXpCs9kcDkPvfyAobNYmFg8yfRZaDopDbf"
		  }`))),
	}, nil)

	response := rpc.GetNewAddress()
	r := response.Result
	assert.Equal(t, "NXpCs9kcDkPvfyAobNYmFg8yfRZaDopDbf", r)
}

func TestRpcClient_GetUnclaimedGas(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "735870007400"
		  }`))),
	}, nil)

	response := rpc.GetUnclaimedGas()
	r := response.Result
	assert.Equal(t, "735870007400", r)
}

func TestRpcClient_ImportPrivKey(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
			  "address": "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ",
			  "haskey": true,
			  "label": null,
			  "watchonly": false
			}
		  }`))),
	}, nil)

	response := rpc.ImportPrivKey("KyoYyZpoccbR6KZ25eLzhMTUxREwCpJzDsnuodGTKXSG8fDW9t7x")
	r := response.Result
	assert.Equal(t, "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ", r.Address)
	assert.Equal(t, "", r.Label)
}

func TestRpcClient_ListAddress(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": [
			  {
				"address": "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ",
				"haskey": true,
				"label": null,
				"watchonly": false
			  },
			  {
				"address": "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW",
				"haskey": true,
				"label": null,
				"watchonly": false
			  }
			]
		  }`))),
	}, nil)

	response := rpc.ListAddress()
	r := response.Result
	assert.Equal(t, "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ", r[0].Address)
	assert.Equal(t, false, r[1].WatchOnly)
	assert.Equal(t, 2, len(r))
}

func TestRpcClient_OpenWallet(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": true
		}`))),
	}, nil)

	response := rpc.OpenWallet("", "")
	r := response.Result
	assert.Equal(t, true, r)
}

func TestRpcClient_SendFrom(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
			  "hash": "0x035facc3be1fc57da1690e3d2f8214f449d368437d8557ffabb2d408caf9ad76",
			  "size": 272,
			  "version": 0,
			  "nonce": 1553700339,
			  "sender": "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ",
			  "sys_fee": "100000000",
			  "net_fee": "1272390",
			  "valid_until_block": 2105487,
			  "attributes": [],
			  "cosigners": [
				{
				  "account": "0xcadb3dc2faa3ef14a13b619c9a43124755aa2569",
				  "scopes": "CalledByEntry"
				}
			  ],
			  "script": "A+CSx1QCAAAADBSZA7DD0pKYj+vl8wagL2VOousWKQwUaSWqVUcSQ5qcYTuhFO+j+sI928oTwAwIdHJhbnNmZXIMFDt9NxHG8Mz5sdypA9G/odiW8SOMQWJ9W1I5",
			  "witnesses": [
				{
				  "invocation": "DEDOA/QF5jYT2TCl9T94fFwAncuBhVhciISaq4fZ3WqGarEoT/0iDo3RIwGjfRW0mm/SV3nAVGEQeZInLqKQ98HX",
				  "verification": "DCEDqgUvvLjlszpO79ZiU2+GhGQfBBCfHV5pzdpvCEiQKGoLQQqQatQ="
				}
			  ]
			}
		  }`))),
	}, nil)

	response := rpc.SendFrom("0x8c23f196d8a1bfd103a9dcb1f9ccf0c611377d3b", "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ", "NZs2zXSPuuv9ZF6TDGSWT1RBmE8rfGj7UW", "100.123")
	r := response.Result
	assert.Equal(t, "0x035facc3be1fc57da1690e3d2f8214f449d368437d8557ffabb2d408caf9ad76", r.Hash)
}

func TestRpcClient_SendMany(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
			  "hash": "0x542e64a9048bbe1ee565b840c41ccf9b5a1ef11f52e5a6858a523938a20c53ec",
			  "size": 483,
			  "version": 0,
			  "nonce": 34429660,
			  "sender": "NUMK37TV9yYKbJr1Gufh74nZiM623eBLqX",
			  "sys_fee": "100000000",
			  "net_fee": "2483780",
			  "valid_until_block": 2105494,
			  "attributes": [],
			  "cosigners": [
				{
				  "account": "0x36d6200fb4c9737c7b552d2b5530ab43605c5869",
				  "scopes": "CalledByEntry"
				},
				{
				  "account": "0x9a55ca1006e2c359bbc8b9b0de6458abdff98b5c",
				  "scopes": "CalledByEntry"
				}
			  ],
			  "script": "GgwUaSWqVUcSQ5qcYTuhFO+j+sI928oMFGlYXGBDqzBVKy1Ve3xzybQPINY2E8AMCHRyYW5zZmVyDBSJdyDYzXb08Aq/o3wO3YicII/em0FifVtSOQKQslsHDBSZA7DD0pKYj+vl8wagL2VOousWKQwUXIv536tYZN6wuci7WcPiBhDKVZoTwAwIdHJhbnNmZXIMFDt9NxHG8Mz5sdypA9G/odiW8SOMQWJ9W1I5",
			  "witnesses": [
				{
				  "invocation": "DECOdTEWg1WkuHN0GNV67kwxeuKADyC6TO59vTaU5dK6K1BGt8+EM6L3TdMga4qB2J+Meez8eYwZkSSRubkuvfr9",
				  "verification": "DCECeiS9CyBqFJwNKzonOs/yzajOraFep4IqFJVxBe6TesULQQqQatQ="
				},
				{
				  "invocation": "DEB1Laj6lvjoBJLTgE/RdvbJiXOmaKp6eNWDJt+p8kxnW6jbeKoaBRZWfUflqrKV7mZEE2JHA5MxrL5TkRIvsL5K",
				  "verification": "DCECkXL4gxd936eGEDt3KWfIuAsBsQcfyyBUcS8ggF6lZnwLQQqQatQ="
				}
			  ]
			}
		  }`))),
	}, nil)

	response := rpc.SendMany("NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ", []models.RpcTransferOut{
		{Asset: "",
			Value:   "10",
			Address: ""}})
	r := response.Result
	assert.Equal(t, "0x542e64a9048bbe1ee565b840c41ccf9b5a1ef11f52e5a6858a523938a20c53ec", r.Hash)
}

func TestRpcClient_SendToAddress(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
			  "hash": "0xee5fc3f57d9f9bc9695c88ecc504444aab622b1680b1cb0848d5b6e39e7fd118",
			  "size": 381,
			  "version": 0,
			  "nonce": 330056065,
			  "sender": "NUMK37TV9yYKbJr1Gufh74nZiM623eBLqX",
			  "sys_fee": "100000000",
			  "net_fee": "2381780",
			  "valid_until_block": 2105500,
			  "attributes": [],
			  "cosigners": [
				{
				  "account": "0xcadb3dc2faa3ef14a13b619c9a43124755aa2569",
				  "scopes": "CalledByEntry"
				}
			  ],
			  "script": "A+CSx1QCAAAADBRpJapVRxJDmpxhO6EU76P6wj3bygwUaSWqVUcSQ5qcYTuhFO+j+sI928oTwAwIdHJhbnNmZXIMFDt9NxHG8Mz5sdypA9G/odiW8SOMQWJ9W1I5",
			  "witnesses": [
				{
				  "invocation": "DECruSKmQKs0Y2cxplKROjPx8HKiyiYrrPn7zaV9zwHPumLzFc8DvgIo2JxmTnJsORyygN/su8mTmSLLb3PesBvY",
				  "verification": "DCECkXL4gxd936eGEDt3KWfIuAsBsQcfyyBUcS8ggF6lZnwLQQqQatQ="
				},
				{
				  "invocation": "DECS5npCs5PwsPUAQ01KyHyCev27dt3kDdT1Vi0K8PwnEoSlxYTOGGQCAwaiNEXSyBdBmT6unhZydmFnkezD7qzW",
				  "verification": "DCEDqgUvvLjlszpO79ZiU2+GhGQfBBCfHV5pzdpvCEiQKGoLQQqQatQ="
				}
			  ]
			}
		  }`))),
	}, nil)

	response := rpc.SendToAddress("0x8c23f196d8a1bfd103a9dcb1f9ccf0c611377d3b", "NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ", "100.123")
	r := response.Result
	assert.Equal(t, "0xee5fc3f57d9f9bc9695c88ecc504444aab622b1680b1cb0848d5b6e39e7fd118", r.Hash)
}
