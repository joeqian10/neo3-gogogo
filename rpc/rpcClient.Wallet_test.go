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

func TestRpcClient_GetWalletBalance(t *testing.T) {
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

	response := rpc.GetWalletBalance("0x8c23f196d8a1bfd103a9dcb1f9ccf0c611377d3b")
	r := response.Result
	assert.Equal(t, "3001101329992600", r.Balance)
}

func TestRpcClient_GetWalletUnclaimedGas(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "897299680935"
		}`))),
	}, nil)

	response := rpc.GetWalletUnclaimedGas()
	r := response.Result
	assert.Equal(t, "897299680935", r)
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

func TestRpcClient_CalculateNetworkFee(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"networkfee": "2384840"
			}
		}`))),
	}, nil)

	response := rpc.CalculateNetworkFee("AAzUzgl2c4kAAAAAAMhjJAAAAAAAmRQgAAKDHlc9J/rM4KzhpixYX/fRkt2q8ACBubhEJKzaXrq9mt5PesW40qC01AEAXQMA6HZIFwAAAAwUgx5XPSf6zOCs4aYsWF/30ZLdqvAMFIG5uEQkrNpeur2a3k96xbjSoLTUE8AMCHRyYW5zZmVyDBS8r0HWhMfUrW7g2Z2pcHudHwyOZkFifVtSOAJCDED0lByRy1/NfBDdKCFLA3RKAY+LLVeXAvut42izfO6PPsKX0JeaL959L0aucqcxBJfWNF3b+93mt9ItCxRoDnChKQwhAuj/F8Vn1i8nT+JHzIhKKmzTuP0Nd5qMWFYomlYKzKy0C0GVRA14QgxAMbiEtF4zjCUjGAzanxLckFiCY3DeREMGIxyerx5GCG/Ki0LGvNzbvPUAWeVGvbL5TVGlK55VfZECmy8voO1LsisRDCEC6P8XxWfWLydP4kfMiEoqbNO4/Q13moxYViiaVgrMrLQRC0ETje+v")
	r := response.Result
	assert.Equal(t, "2384840", r.NetworkFee)
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
			  "sysfee": "100000000",
			  "netfee": "1272390",
			  "validuntilblock": 2105487,
			  "attributes": [],
			  "signers": [
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
			  "sysfee": "100000000",
			  "netfee": "2483780",
			  "validuntilblock": 2105494,
			  "attributes": [],
			  "signers": [
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
			  "sysfee": "100000000",
			  "netfee": "2381780",
			  "validuntilblock": 2105500,
			  "attributes": [],
			  "signers": [
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

func TestRpcClient_InvokeContractVerify(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"script": "10c00c08646563696d616c730c143b7d3711c6f0ccf9b1dca903d1bfa1d896f1238c41627d5b52",
				"state": "HALT",
				"gasconsumed": "1007390",
				"stack": [
					{
						"type": "Integer",
						"value": "8"
					}
				]
			}
		}`))),
	}, nil)

	params := []models.RpcContractParameter{models.NewRpcContractParameter("Boolean", "true")}
	signers :=  []models.RpcSigner{models.RpcSigner{
		Account:          "0xf621168b1fce3a89c33a5f6bcf7e774b4657031c",
		Scopes:           "CalledByEntry",
		AllowedContracts: []string{},
		AllowedGroups:    []string{},
	}}

	response := rpc.InvokeContractVerify("0x8c23f196d8a1bfd103a9dcb1f9ccf0c611377d3b", params, signers)
	r := response.Result
	assert.Equal(t, "HALT", r.State)
	assert.Equal(t, "8", r.Stack[0].Value)
}
