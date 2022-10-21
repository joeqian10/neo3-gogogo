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

	response := rpc.InvokeFunction("0x8c23f196d8a1bfd103a9dcb1f9ccf0c611377d3b", "decimals", nil, nil, false)
	stacks, err := PopInvokeStacks(response)
	assert.Nil(t, err)
	assert.Equal(t, "8", stacks[0].Value.(string))
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

	response := rpc.InvokeScript("", nil, false)
	stacks, err := PopInvokeStacks(response)
	assert.Nil(t, err)
	assert.Equal(t, "8", stacks[0].Value.(string))
}

func TestRpcClient_TraverseIterator(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": [
				{
					"type": "ByteString",
					"value": "TWV0YVBhbmFjZWEgIzE3LTAy"
				},
				{
					"type": "ByteString",
					"value": "TWV0YVBhbmFjZWEgIzctMDI="
				},
				{
					"type": "ByteString",
					"value": "TWV0YVBhbmFjZWEgIzgtMTA="
				}
			]
		}`))),
	}, nil)

	response := rpc.TraverseIterator("9f267840-216b-47d0-aeee-665df3657a6e", "6f3a3284-0975-4774-9dfa-aab23dc9b42e", 100)
	r := response.Result
	assert.Equal(t, 3, len(r))
}

func TestRpcClient_TerminateSession(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": false
		}`))),
	}, nil)

	response := rpc.TerminateSession("9f267840-216b-47d0-aeee-665df3657a6e")
	r := response.Result
	assert.Equal(t, false, r)
}

func TestRpcClient_GetUnclaimedGas(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"unclaimed": "9693738",
				"address": "NNU67Fvdy3LEQTM374EJ9iMbCRxVExgM8Y"
			}
		}`))),
	}, nil)

	response := rpc.GetUnclaimedGas("NNU67Fvdy3LEQTM374EJ9iMbCRxVExgM8Y")
	r := response.Result
	assert.Equal(t, "9693738", r.Unclaimed)
}
