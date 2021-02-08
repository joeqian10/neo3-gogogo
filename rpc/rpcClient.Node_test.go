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
				"useragent": "/Neo:3.0.0-preview5/"
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
