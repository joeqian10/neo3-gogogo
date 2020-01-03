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
			"result": "0x2727ab449e02150fa66943cf6d8fcdf4af349480a558a0bdbb4eea550ffeb01f"
		}`))),
	}, nil)

	response := rpc.GetBestBlockHash()
	r := response.Result
	assert.Equal(t, "0x2727ab449e02150fa66943cf6d8fcdf4af349480a558a0bdbb4eea550ffeb01f", r)
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

func TestRpcClient_ValidateAddress(t *testing.T) {
	var client = new(HttpClientMock)
	var rpc = RpcClient{Endpoint: new(url.URL), httpClient: client}
	client.On("Do", mock.Anything).Return(&http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"jsonrpc": "2.0",
			"id": 1,
			"result": {
				"address": "AQVh2pG732YvtNaxEGkQUei3YA4cvo7d2i",
				"isvalid": true
			}
		}`))),
	}, nil)

	response := rpc.ValidateAddress("")
	r := response.Result
	assert.Equal(t, "AQVh2pG732YvtNaxEGkQUei3YA4cvo7d2i", r.Address)
	assert.Equal(t, true, r.IsValid)
}
