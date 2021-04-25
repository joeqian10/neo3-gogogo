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

	response := rpc.InvokeFunction("0x8c23f196d8a1bfd103a9dcb1f9ccf0c611377d3b", "decimals", nil, nil)
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

	response := rpc.InvokeScript("", nil)
	r := response.Result
	assert.Equal(t, "HALT", r.State)
	assert.Equal(t, "8", r.Stack[0].Value)
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
