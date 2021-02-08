package rpc

import (
	"net/http"
	"testing"

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
