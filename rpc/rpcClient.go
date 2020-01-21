package rpc

import (
	"bytes"
	"encoding/json"
	"github.com/joeqian10/neo3-gogogo/helper"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// add IHttpClient for mock unit test
type IHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RpcClient struct {
	Endpoint   *url.URL
	httpClient IHttpClient
}

func NewClient(endpoint string) *RpcClient {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil
	}
	var netClient = &http.Client{
		Timeout: time.Second * 60,
	}
	return &RpcClient{Endpoint: u, httpClient: netClient}
}

func (n *RpcClient) makeRequest(method string, params []interface{}, out interface{}) error {
	request := NewRequest(method, params)
	jsonValue, _ := json.Marshal(request)
	req, err := http.NewRequest("POST", n.Endpoint.String(), bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
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

func (n *RpcClient) GetBestBlockHash() GetBestBlockHashResponse {
	response := GetBestBlockHashResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getbestblockhash", params, &response)
	return response
}

func (n *RpcClient) GetBlock(hashOrIndex string) GetBlockResponse {
	params := []interface{}{hashOrIndex, true}
	index, err := strconv.Atoi(hashOrIndex)
	if err == nil {
		params = []interface{}{index, true}
	}
	response := GetBlockResponse{}
	_ = n.makeRequest("getblock", params, &response)
	return response
}

func (n *RpcClient) GetBlockCount() GetBlockCountResponse {
	response := GetBlockCountResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getblockcount", params, &response)
	return response
}

func (n *RpcClient) GetBlockHash(index uint32) GetBlockHashResponse {
	response := GetBlockHashResponse{}
	params := []interface{}{index}
	_ = n.makeRequest("getblockhash", params, &response)
	return response
}

func (n *RpcClient) GetBlockHeader(hashOrIndex string) GetBlockHeaderResponse {
	params := []interface{}{hashOrIndex, true}
	index, err := strconv.Atoi(hashOrIndex)
	if err == nil {
		params = []interface{}{index, true}
	}
	response := GetBlockHeaderResponse{}
	_ = n.makeRequest("getblockheader", params, &response)
	return response
}

func (n *RpcClient) GetBlockSysFee(height int) GetBlockSysFeeResponse {
	response := GetBlockSysFeeResponse{}
	params := []interface{}{height}
	_ = n.makeRequest("getblocksysfee", params, &response)
	return response
}

func (n *RpcClient) GetContractState(scriptHash string) GetContractStateResponse {
	response := GetContractStateResponse{}
	params := []interface{}{scriptHash}
	_ = n.makeRequest("getcontractstate", params, &response)
	return response
}

func (n *RpcClient) GetRawMemPool() GetRawMemPoolResponse {
	response := GetRawMemPoolResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getrawmempool", params, &response)
	return response
}

func (n *RpcClient) GetRawTransaction(txid string) GetRawTransactionResponse {
	response := GetRawTransactionResponse{}
	params := []interface{}{txid, 1}
	_ = n.makeRequest("getrawtransaction", params, &response)
	return response
}

func (n *RpcClient) GetStorage(scripthash string, key string) GetStorageResponse {
	response := GetStorageResponse{}
	params := []interface{}{scripthash, key}
	_ = n.makeRequest("getstorage", params, &response)
	return response
}

func (n *RpcClient) GetTransactionHeight(txid string) GetTransactionHeightResponse {
	response := GetTransactionHeightResponse{}
	params := []interface{}{txid}
	_ = n.makeRequest("gettransactionheight", params, &response)
	return response
}

func (n *RpcClient) GetValidators() GetValidatorsResponse {
	response := GetValidatorsResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getvalidators", params, &response)
	return response
}

// Node
func (n *RpcClient) GetConnectionCount() GetConnectionCountResponse {
	response := GetConnectionCountResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getconnectioncount", params, &response)
	return response
}

func (n *RpcClient) GetPeers() GetPeersResponse {
	response := GetPeersResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getpeers", params, &response)
	return response
}

func (n *RpcClient) GetVersion() GetVersionResponse {
	response := GetVersionResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getversion", params, &response)
	return response
}

func (n *RpcClient) SendRawTransaction(rawTransactionInHex string) SendRawTransactionResponse {
	response := SendRawTransactionResponse{}
	params := []interface{}{rawTransactionInHex, 1}
	_ = n.makeRequest("sendrawtransaction", params, &response)
	return response
}

func (n *RpcClient) SubmitBlock(blockHex string) SubmitBlockResponse {
	response := SubmitBlockResponse{}
	params := []interface{}{blockHex}
	_ = n.makeRequest("submitblock", params, &response)
	return response
}

// SmartContract
func (n *RpcClient) InvokeFunction(scriptHash string, method string, args ...InvokeFunctionStackArg) InvokeResultResponse {
	response := InvokeResultResponse{}
	var params []interface{}
	if args != nil {
		params = []interface{}{scriptHash, method, args}
	} else {
		params = []interface{}{scriptHash, method}
	}
	_ = n.makeRequest("invokefunction", params, &response)
	return response
}

func (n *RpcClient) InvokeScript(scriptInHex string, scriptHashesForVerifying ...helper.UInt160) InvokeResultResponse {
	response := InvokeResultResponse{}
	params := []interface{}{scriptInHex, scriptHashesForVerifying}
	_ = n.makeRequest("invokescript", params, &response)
	return response
}

// Utilities
func (n *RpcClient) ListPlugins() ListPluginsResponse {
	response := ListPluginsResponse{}
	params := []interface{}{}
	_ = n.makeRequest("listplugins", params, &response)
	return response
}

func (n *RpcClient) ValidateAddress(address string) ValidateAddressResponse {
	response := ValidateAddressResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("validateaddress", params, &response)
	return response
}

// Plugins
// the endpoint needs to use ApplicationLogs plugin
func (n *RpcClient) GetApplicationLog(txId string) GetApplicationLogResponse {
	response := GetApplicationLogResponse{}
	params := []interface{}{txId}
	_ = n.makeRequest("getapplicationlog", params, &response)
	return response
}

// this endpoint needs RpcNep5Tracker plugin
func (n *RpcClient) GetNep5Balances(address string) GetNep5BalancesResponse {
	response := GetNep5BalancesResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("getnep5balances", params, &response)
	return response
}

// this endpoint needs RpcNep5Tracker plugin
func (n *RpcClient) GetNep5Transfers(address string, startTime *int, endTime *int) GetNep5TransfersResponse {
	response := GetNep5TransfersResponse{}
	var params []interface{}
	if startTime != nil {
		if endTime != nil {
			params = []interface{}{address, *startTime, *endTime}
		} else {
			params = []interface{}{address, *startTime}
		}
	} else {
		params = []interface{}{address}
	}
	_ = n.makeRequest("getnep5balances", params, &response)
	return response
}

// Wallet
func (n *RpcClient) ImportPrivKey(wif string) ImportPrivKeyResponse {
	response := ImportPrivKeyResponse{}
	params := []interface{}{wif}
	_ = n.makeRequest("importprivkey", params, &response)
	return response
}

func (n *RpcClient) ListAddress() ListAddressResponse {
	response := ListAddressResponse{}
	params := []interface{}{}
	_ = n.makeRequest("listaddress", params, &response)
	return response
}

func (n *RpcClient) SendFrom(assetId string, from string, to string, amount uint32, fee float32, changeAddress string) SendFromResponse {
	response := SendFromResponse{}
	params := []interface{}{assetId, from, to, amount, fee, changeAddress}
	_ = n.makeRequest("sendfrom", params, &response)
	return response
}

func (n *RpcClient) SendToAddress(assetId string, to string, amount uint32, fee float32, changeAddress string) SendToAddressResponse {
	response := SendToAddressResponse{}
	params := []interface{}{assetId, to, amount, fee, changeAddress}
	_ = n.makeRequest("sendtoaddress", params, &response)
	return response
}

func (n *RpcClient) GetCrossChainProof(blockIndex int, txID string) GetCrossChainProofResponse {
	response := GetCrossChainProofResponse{}
	params := []interface{}{blockIndex, txID}
	_ = n.makeRequest("getcrossproof", params, &response)
	return response
}
