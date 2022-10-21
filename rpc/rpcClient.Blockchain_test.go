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
						"sysfee": "100000000",
						"netfee": "1270450",
						"validuntilblock": 2102808,
						"attributes": [],
						"signers": [
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
			"result": 2077
		}`))),
	}, nil)

	response := rpc.GetBlockCount()
	r := response.Result
	assert.Equal(t, 2077, r)
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

func TestRpcClient_GetBlockHeaderCount(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": 2077
		}`))),
	}, nil)

	response := rpc.GetBlockHeaderCount()
	r := response.Result
	assert.Equal(t, 2077, r)
}

func TestRpcClient_GetContractState(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"id": 8,
				"updatecounter": 0,
				"hash": "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789",
				"nef": {
					"magic": 6321442516,
					"compiler": "neon",
					"tokens": [
						{
							"hash": "0x99042d380f2b754175717bb932a911bc0bb0ad7d",
							"method": "verify",
							"paramcount": 0,
							"hasreturnvalue": true,
							"callflags": "AllowCall"
						}
					],
					"script": "DBQKo4e1Ppa3mJpjFDGgVt0fQKBC9kH4J+yMQDTkQFcBAAwFSGVsbG9Bm/ZnzkGSXegxcGhAVwQBEnAMF0ludm9rZSBTdHJvYWdlLlB1dCBmb3IgaBpQQXvjun0MByB0aW1lcy6Li9soQc/nR5YMBUhlbGxveFBBm/ZnzkHmPxiEDAJOb0Gb9mfOQZJd6DHYqnNrJiwMAk5vDAJOb0Gb9mfOQZJd6DFK2CYFEFBF2yERnlBBm/ZnzkHmPxiEIhMhDAJObxFQQZv2Z85B5j8YhAwCTm9Bm/ZnzkGSXegxcWlK2CYFEFBF2yEaUEF747p9chXDShAMBFB1dCDQShF40EoSDB0gaW50byBzdG9yYWdlIGNvbXBsZXRlbHkgZm9yINBKE2rQShQMBiB0aW1lc9DBShEyCJ1Ti1Ai+EXbKEHP50eWeBHADARXb3JkQZUBb2FpEcAMDkludm9rZVB1dENvdW50QZUBb2FAVwECNZL+//8Qs3BoJhYMEU5vIGF1dGhvcml6YXRpb24uOnh5UEExxjMdQFcBADVn/v//ELNwaCYWDBFObyBhdXRob3JpemF0aW9uLjohQcafHfBAVgEMFAqjh7U+lreYmmMUMaBW3R9AoEL2YEA=",
					"checksum": 73195690102
				},
				"manifest": {
					"name": "testContract",
					"groups": [],
					"supportedstandards": [],
					"abi": {
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
								"returnType": "Integer",
								"offset": 28,
								"safe": true
							}
						],
						"events": []
					},
					"permissions": [
						{
							"contract": "*",
							"methods": [
								"*"
							]
						}
					],
					"trusts": []
				}
			}
		}`))),
	}, nil)

	response := rpc.GetContractState("0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789")
	r := response.Result
	assert.Equal(t, "0x9bde8f209c88dd0e7ca3bf0af0f476cdd8207789", r.Hash)
	assert.Equal(t, "unclaimedGas", r.Manifest.Abi.Methods[0].Name)
}

func TestRpcClient_GetRawMemPool(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": [
				"0x9786cce0dddb524c40ddbdd5e31a41ed1f6b5c8a683c122f627ca4a007a7cf4e",
				"0xb488ad25eb474f89d5ca3f985cc047ca96bc7373a6d3da8c0f192722896c1cd7",
				"0xf86f6f2c08fbf766ebe59dc84bc3b8829f1053f0a01deb26bf7960d99fa86cd6"]
		}`))),
	}, nil)

	response := rpc.GetRawMemPool()
	r := response.Result
	assert.Equal(t, "0x9786cce0dddb524c40ddbdd5e31a41ed1f6b5c8a683c122f627ca4a007a7cf4e", r[0])
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
				"sysfee": "100000000",
				"netfee": "1270450",
				"validuntilblock": 2102808,
				"attributes": [],
				"signers": [
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

func TestRpcClient_GetNextBlockValidators(t *testing.T) {
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

	response := rpc.GetNextBlockValidators()
	r := response.Result
	assert.Equal(t, "03aa052fbcb8e5b33a4eefd662536f8684641f04109f1d5e69cdda6f084890286a", r[0].PublicKey)
}

func TestRpcClient_GetCandidates(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": [
				{
					"publickey": "02237309a0633ff930d51856db01d17c829a5b2e5cc2638e9c03b4cfa8e9c9f971",
					"votes": "634270",
					"active": false
				},
				{
					"publickey": "022f1beae94cf0d266d7d26691b431958c8d13768103ab20aed817b57568da293f",
					"votes": "1803",
					"active": false
				},
				{
					"publickey": "0239a37436652f41b3b802ca44cbcb7d65d3aa0b88c9a0380243bdbe1aaa5cb35b",
					"votes": "2026211",
					"active": true
				},
				{
					"publickey": "02486fd15702c4490a26703112a5cc1d0923fd697a33406bd5a1c00e0013b09a70",
					"votes": "2079259",
					"active": true
				},
				{
					"publickey": "0248a37e04c7a5fb9fdc9f0323b2955a94cbb2296d2ad1feacea24ec774a87c4a4",
					"votes": "714040",
					"active": false
				},
				{
					"publickey": "024c7b7fb6c310fccf1ba33b082519d82964ea93868d676662d4a59ad548df0e7d",
					"votes": "150",
					"active": false
				},
				{
					"publickey": "024c8b6b34990a164fc7860dc497babd71911d049bdf8e186f5f51258b98509434",
					"votes": "462",
					"active": false
				},
				{
					"publickey": "02784c0e81f6e3eadbd1516f9311d1fe0662f9c360ea331d9d6a4934800c96ed47",
					"votes": "1006166",
					"active": false
				},
				{
					"publickey": "02946248f71bdf14933e6735da9867e81cc9eea0b5895329aa7f71e7745cf40659",
					"votes": "620884",
					"active": false
				},
				{
					"publickey": "029b46bf20b19fb3da1a4591bac09acd86c460ae29cb74c7634d658702205ec2f7",
					"votes": "1",
					"active": false
				},
				{
					"publickey": "02aaec38470f6aad0042c6e877cfd8087d2676b0f516fddd362801b9bd3936399e",
					"votes": "2031956",
					"active": true
				},
				{
					"publickey": "02ada21c39ccc5ff4a0547270bee2eb940c394c3afe0d8f5470e2fe5b56759cee2",
					"votes": "0",
					"active": false
				},
				{
					"publickey": "02beaf473e48740f8ac1b70ff2b6cdb850a7c247b9d036508d6f0bdaa1e750eb3f",
					"votes": "380",
					"active": false
				},
				{
					"publickey": "02ca0e27697b9c248f6f16e085fd0061e26f44da85b58ee835c110caa5ec3ba554",
					"votes": "2035673",
					"active": true
				},
				{
					"publickey": "02cc10d0e929ca752cfd3408bedda06463e2d93fd435e4c2b86a895b3792dee4c8",
					"votes": "660112",
					"active": false
				},
				{
					"publickey": "02df48f60e8f3e01c48ff40b9b7f1310d7a8b2a193188befe1c2e3df740e895093",
					"votes": "1541",
					"active": false
				},
				{
					"publickey": "02e4e0db9314a42bebcb7e348a7e1ab8d4a87f518371485b0d66dbc0368f8cf58d",
					"votes": "643413",
					"active": false
				},
				{
					"publickey": "02ec143f00b88524caf36a0121c2de09eef0519ddbe1c710a00f0e2663201ee4c0",
					"votes": "3802422",
					"active": true
				},
				{
					"publickey": "034f7ea4ca4506ef288c5d5ed61686b9f39a0bc5f7670858305e32ea504ab543f3",
					"votes": "1838887",
					"active": false
				},
				{
					"publickey": "0353f7aff015d3d204eecd508f4aa67f447df4d15ec4ba649f4fa91bdb7a78a354",
					"votes": "1501",
					"active": false
				},
				{
					"publickey": "035d574cc6a904e82dfd82d7f6fc9c2ca042d4410a4910ecc8c07a07db49dc6513",
					"votes": "665975",
					"active": false
				},
				{
					"publickey": "03734d4b637dbac04d0eb45198bfe9c5a42aca907e8fd1e741eb52def583347257",
					"votes": "622773",
					"active": false
				},
				{
					"publickey": "0389ba00856f58bdc2b7c3fd6c2e73ed4829252f38083c1e295574ca599e93fe82",
					"votes": "1001849",
					"active": false
				},
				{
					"publickey": "0392fbd1d809a3c62f7dcde8f25454a1570830a21e4b014b3f362a79baf413e115",
					"votes": "622057",
					"active": false
				},
				{
					"publickey": "03a3aba8edacd820f0b6c55f0bad25b733af00694c8194c04b71ce628c197fbe98",
					"votes": "625891",
					"active": false
				},
				{
					"publickey": "03b209fd4f53a7170ea4444e0cb0a6bb6a53c2bd016926989cf85f9b0fba17a70c",
					"votes": "3021704",
					"active": true
				},
				{
					"publickey": "03b8d9d5771d8f513aa0869b9cc8d50986403b78c6da36890638c3d46a5adce04a",
					"votes": "2017769",
					"active": false
				},
				{
					"publickey": "03d9e8b16bd9b22d3345d6d4cde31be1c3e1d161532e3d0ccecb95ece2eb58336e",
					"votes": "12052968",
					"active": true
				},
				{
					"publickey": "03fd04de983f4e04c9629ab3cfc83f41be7431b96bf852a91873c38ca8f737ee2c",
					"votes": "624224",
					"active": false
				}
			]
		}`))),
	}, nil)

	response := rpc.GetCandidates()
	r := response.Result
	assert.Equal(t, 29, len(r))
}

func TestRpcClient_GetCommittee(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
    		"id": 1,
    		"result": [
				"020f2887f41474cfeb11fd262e982051c1541418137c02a0f4961af911045de639",
				"03204223f8c86b8cd5c89ef12e4f0dbb314172e9241e30c9ef2293790793537cf0",
				"0222038884bbd1d8ff109ed3bdef3542e768eef76c1247aea8bc8171f532928c30",
				"0226933336f1b75baa42d42b71d9091508b638046d19abd67f4e119bf64a7cfb4d",
				"023a36c72844610b4d34d1968662424011bf783ca9d984efa19a20babf5582f3fe",
				"03409f31f0d66bdc2f70a9730b66fe186658f84a8018204db01c106edc36553cd0",
				"02486fd15702c4490a26703112a5cc1d0923fd697a33406bd5a1c00e0013b09a70",
				"024c7b7fb6c310fccf1ba33b082519d82964ea93868d676662d4a59ad548df0e7d",
				"02504acbc1f4b3bdad1d86d6e1a08603771db135a73e61c9d565ae06a1938cd2ad",
				"03708b860c1de5d87f5b151a12c2a99feebd2e8b315ee8e7cf8aa19692a9e18379",
				"0288342b141c30dc8ffcde0204929bb46aed5756b41ef4a56778d15ada8f0c6654",
				"02a62c915cf19c7f19a50ec217e79fac2439bbaad658493de0c7d8ffa92ab0aa62",
				"02aaec38470f6aad0042c6e877cfd8087d2676b0f516fddd362801b9bd3936399e",
				"03b209fd4f53a7170ea4444e0cb0a6bb6a53c2bd016926989cf85f9b0fba17a70c",
				"03b8d9d5771d8f513aa0869b9cc8d50986403b78c6da36890638c3d46a5adce04a",
				"03c6aa6e12638b36e88adc1ccdceac4db9929575c3e03576c617c49cce7114a050",
				"02ca0e27697b9c248f6f16e085fd0061e26f44da85b58ee835c110caa5ec3ba554",
				"02cd5a5547119e24feaa7c2a0f37b8c9366216bab7054de0065c9be42084003c8a",
				"03cdcea66032b82f5c30450e381e5295cae85c5e6943af716cc6b646352a6067dc",
				"03d281b42002647f0113f36c7b8efb30db66078dfaaa9ab3ff76d043a98d512fde",
				"02df48f60e8f3e01c48ff40b9b7f1310d7a8b2a193188befe1c2e3df740e895093"
    		]
		}`))),
	}, nil)

	response := rpc.GetCommittee()
	r := response.Result
	assert.Equal(t, 21, len(r))
	assert.Equal(t, "020f2887f41474cfeb11fd262e982051c1541418137c02a0f4961af911045de639", r[0])
}

func TestRpcClient_GetNativeContracts(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
		  "jsonrpc": "2.0",
		  "id": 1,
		  "result": [
			{
			  "id": -1,
			  "hash": "0xfffdc93764dbaddd97c48f252a53ea4643faa3fd",
			  "nef": {
				"magic": 860243278,
				"compiler": "neo-core-v3.0",
				"tokens": [],
				"script": "EEEa93tnQBBBGvd7Z0AQQRr3e2dAEEEa93tnQBBBGvd7Z0AQQRr3e2dAEEEa93tnQBBBGvd7Z0A=",
				"checksum": 1110259869
			  },
			  "manifest": {
				"name": "ContractManagement",
				"groups": [],
				"supportedstandards": [],
				"abi": {
				  "methods": [
					{
					  "name": "deploy",
					  "parameters": [
						{
						  "name": "nefFile",
						  "type": "ByteArray"
						},
						{
						  "name": "manifest",
						  "type": "ByteArray"
						}
					  ],
					  "returntype": "Array",
					  "offset": 0,
					  "safe": false
					},
					{
					  "name": "deploy",
					  "parameters": [
						{
						  "name": "nefFile",
						  "type": "ByteArray"
						},
						{
						  "name": "manifest",
						  "type": "ByteArray"
						},
						{
						  "name": "data",
						  "type": "Any"
						}
					  ],
					  "returntype": "Array",
					  "offset": 7,
					  "safe": false
					},
					{
					  "name": "destroy",
					  "parameters": [],
					  "returntype": "Void",
					  "offset": 14,
					  "safe": false
					},
					{
					  "name": "getContract",
					  "parameters": [
						{
						  "name": "hash",
						  "type": "Hash160"
						}
					  ],
					  "returntype": "Array",
					  "offset": 21,
					  "safe": true
					},
					{
					  "name": "getMinimumDeploymentFee",
					  "parameters": [],
					  "returntype": "Integer",
					  "offset": 28,
					  "safe": true
					},
					{
					  "name": "setMinimumDeploymentFee",
					  "parameters": [
						{
						  "name": "value",
						  "type": "Integer"
						}
					  ],
					  "returntype": "Void",
					  "offset": 35,
					  "safe": false
					},
					{
					  "name": "update",
					  "parameters": [
						{
						  "name": "nefFile",
						  "type": "ByteArray"
						},
						{
						  "name": "manifest",
						  "type": "ByteArray"
						}
					  ],
					  "returntype": "Void",
					  "offset": 42,
					  "safe": false
					},
					{
					  "name": "update",
					  "parameters": [
						{
						  "name": "nefFile",
						  "type": "ByteArray"
						},
						{
						  "name": "manifest",
						  "type": "ByteArray"
						},
						{
						  "name": "data",
						  "type": "Any"
						}
					  ],
					  "returntype": "Void",
					  "offset": 49,
					  "safe": false
					}
				  ],
				  "events": [
					{
					  "name": "Deploy",
					  "parameters": [
						{
						  "name": "Hash",
						  "type": "Hash160"
						}
					  ]
					},
					{
					  "name": "Update",
					  "parameters": [
						{
						  "name": "Hash",
						  "type": "Hash160"
						}
					  ]
					},
					{
					  "name": "Destroy",
					  "parameters": [
						{
						  "name": "Hash",
						  "type": "Hash160"
						}
					  ]
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
				"extra": null
			  },
			  "updatehistory": [
				0
			  ]
			},
			{
			  "id": -2,
			  "hash": "0xacce6fd80d44e1796aa0c2c625e9e4e0ce39efc0",
			  "nef": {
				"magic": 860243278,
				"compiler": "neo-core-v3.0",
				"tokens": [],
				"script": "EEEa93tnQBBBGvd7Z0AQQRr3e2dAEEEa93tnQBBBGvd7Z0AQQRr3e2dAEEEa93tnQBBBGvd7Z0AQQRr3e2dAEEEa93tnQA==",
				"checksum": 2135988409
			  },
			  "manifest": {
				"name": "StdLib",
				"groups": [],
				"supportedstandards": [],
				"abi": {
				  "methods": [
					{
					  "name": "atoi",
					  "parameters": [
						{
						  "name": "value",
						  "type": "String"
						},
						{
						  "name": "base",
						  "type": "Integer"
						}
					  ],
					  "returntype": "Integer",
					  "offset": 0,
					  "safe": true
					},
					{
					  "name": "base58Decode",
					  "parameters": [
						{
						  "name": "s",
						  "type": "String"
						}
					  ],
					  "returntype": "ByteArray",
					  "offset": 7,
					  "safe": true
					},
					{
					  "name": "base58Encode",
					  "parameters": [
						{
						  "name": "data",
						  "type": "ByteArray"
						}
					  ],
					  "returntype": "String",
					  "offset": 14,
					  "safe": true
					},
					{
					  "name": "base64Decode",
					  "parameters": [
						{
						  "name": "s",
						  "type": "String"
						}
					  ],
					  "returntype": "ByteArray",
					  "offset": 21,
					  "safe": true
					},
					{
					  "name": "base64Encode",
					  "parameters": [
						{
						  "name": "data",
						  "type": "ByteArray"
						}
					  ],
					  "returntype": "String",
					  "offset": 28,
					  "safe": true
					},
					{
					  "name": "deserialize",
					  "parameters": [
						{
						  "name": "data",
						  "type": "ByteArray"
						}
					  ],
					  "returntype": "Any",
					  "offset": 35,
					  "safe": true
					},
					{
					  "name": "itoa",
					  "parameters": [
						{
						  "name": "value",
						  "type": "Integer"
						},
						{
						  "name": "base",
						  "type": "Integer"
						}
					  ],
					  "returntype": "String",
					  "offset": 42,
					  "safe": true
					},
					{
					  "name": "jsonDeserialize",
					  "parameters": [
						{
						  "name": "json",
						  "type": "ByteArray"
						}
					  ],
					  "returntype": "Any",
					  "offset": 49,
					  "safe": true
					},
					{
					  "name": "jsonSerialize",
					  "parameters": [
						{
						  "name": "item",
						  "type": "Any"
						}
					  ],
					  "returntype": "ByteArray",
					  "offset": 56,
					  "safe": true
					},
					{
					  "name": "serialize",
					  "parameters": [
						{
						  "name": "item",
						  "type": "Any"
						}
					  ],
					  "returntype": "ByteArray",
					  "offset": 63,
					  "safe": true
					}
				  ],
				  "events": []
				},
				"permissions": [
				  {
					"contract": "*",
					"methods": "*"
				  }
				],
				"trusts": [],
				"extra": null
			  },
			  "updatehistory": [
				0
			  ]
			}
		  ]
		}`))),
	}, nil)

	response := rpc.GetNativeContracts()
	r := response.Result
	assert.Equal(t, 2, len(r))
}
