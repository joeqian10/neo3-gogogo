package rpc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

// IHttpClient for mock unit test
type IHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RpcClient struct {
	Endpoint   *url.URL
	_url       string
	httpClient IHttpClient
	userName   string
	password   string
}

func NewClient(endpoint string) *RpcClient {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil
	}
	var netClient = &http.Client{
		Timeout: time.Second * 60,
	}
	return &RpcClient{Endpoint: u, httpClient: netClient, _url: endpoint}
}

func (n *RpcClient) SetBasicAuth(user string, pass string) {
	n.userName = user
	n.password = pass
}

func (n *RpcClient) GetUrl() string {
	return n._url
}

func (n *RpcClient) makeRequest(method string, params []interface{}, out interface{}) error {
	request := NewRequest(method, params)
	jsonValue, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", n.Endpoint.String(), bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	if n.userName != "" && n.password != "" {
		req.SetBasicAuth(n.userName, n.password)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Set("Connection", "close")
	req.Close = true
	res, err := n.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		return err
	}
	return nil
}

func getRpcName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.Function
	//fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}

// Plugins


//-------------------------------------------------------------------

type GetCrossChainProofResponse struct {
	RpcResponse
	ErrorResponse
	CrossChainProof string `json:"result"`
}

func (n *RpcClient) GetCrossChainProof(blockIndex int, txID string) GetCrossChainProofResponse {
	response := GetCrossChainProofResponse{}
	params := []interface{}{blockIndex, txID}
	_ = n.makeRequest("getcrossproof", params, &response)
	return response
}
