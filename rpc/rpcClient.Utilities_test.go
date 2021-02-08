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
				"isvalid": true,
				"label": "test",
				"watchonly": true
			}
		}`))),
	}, nil)

	response := rpc.ValidateAddress("NVVwFw6XyhtRCFQ8SpUTMdPyYt4Vd9A1XQ")
	r := response.Result
	assert.Equal(t, true, r.IsValid)
}
