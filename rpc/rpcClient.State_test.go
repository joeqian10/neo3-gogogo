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

func TestRpcClient_GetProof(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "2q8ACBubhEJKzaXrq9mt5PesW40qC01AEAXQMA6HZIFwAAAAwUgx5XPSf6zOCs4aYsWF30ZLdqvAMFIG5uEQkr"
		}`))),
	}, nil)

	response := rpc.GetProof("", "", "")
	r := response.Result
	assert.Equal(t, "2q8ACBubhEJKzaXrq9mt5PesW40qC01AEAXQMA6HZIFwAAAAwUgx5XPSf6zOCs4aYsWF30ZLdqvAMFIG5uEQkr", r)
}

func TestRpcClient_GetStateHeight(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"localrootindex": 1234,
				"validatedrootindex": 2345
			}
		}`))),
	}, nil)

	response := rpc.GetStateHeight()
	r := response.Result
	assert.Equal(t, uint32(1234), r.LocalRootIndex)
}

func TestRpcClient_GetStateRoot(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"version": 0,
				"index": 1234,
				"roothash": "0x3d39da5b227e3f02f5210b24690a0523162788668e490363c6a39813bb162e51",
				"witness": {
					"invocation": "DEDcGjmiHJ22R4LjUuXOF83UDtJB3FUZPy4t8Ol+dSpQovI9KAfVVOrtz/NZBmEuVGXiALkJU6vklZ9XzzDrz0PJ",
					"verification": "EQwhA6oFL7y45bM6Tu/WYlNvhoRkHwQQnx1eac3abwhIkChqEQtBMHOzuw=="
				}
			}
		}`))),
	}, nil)

	response := rpc.GetStateRoot(1234)
	r := response.Result
	assert.Equal(t, uint32(1234), r.Index)
}

func TestRpcClient_VerifyProof(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": "EQwhA6oFL7y45bM6TWYlNvhoRkHwQQnx1eac3abwhIkChqEQtBMHOzuw"
		}`))),
	}, nil)

	response := rpc.VerifyProof("", "")
	r := response.Result
	assert.Equal(t, "EQwhA6oFL7y45bM6TWYlNvhoRkHwQQnx1eac3abwhIkChqEQtBMHOzuw", r)
}
