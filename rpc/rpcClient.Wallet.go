package rpc

import "github.com/joeqian10/neo3-gogogo/rpc/models"

type CloseWalletResponse struct {
	RpcResponse
	ErrorResponse
	Result bool `json:"result"`
}

type DumpPrivKeyResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetNewAddressResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type GetWalletBalanceResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcWalletBalance `json:"result"`
}

type GetWalletUnclaimedGasResponse struct {
	RpcResponse
	ErrorResponse
	Result string `json:"result"`
}

type ImportPrivKeyResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcAddress `json:"result"`
}

type CalculateNetworkFeeResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcNetworkFee `json:"result"`
}

type ListAddressResponse struct {
	RpcResponse
	ErrorResponse
	Result []models.RpcAddress `json:"result"`
}

type OpenWalletResponse struct {
	RpcResponse
	ErrorResponse
	Result bool `json:"result"`
}

type SendFromResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransaction `json:"result"`
}

type SendManyResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransaction `json:"result"`
}

type SendToAddressResponse struct {
	RpcResponse
	ErrorResponse
	Result models.RpcTransaction `json:"result"`
}

func (n *RpcClient) CloseWallet() CloseWalletResponse {
	response := CloseWalletResponse{}
	params := []interface{}{}
	_ = n.makeRequest("closewallet", params, &response)
	return response
}

func (n *RpcClient) DumpPrivKey(address string) DumpPrivKeyResponse {
	response := DumpPrivKeyResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("dumpprivkey", params, &response)
	return response
}

func (n *RpcClient) GetNewAddress() GetNewAddressResponse {
	response := GetNewAddressResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getnewaddress", params, &response)
	return response
}

func (n *RpcClient) GetWalletBalance(assetId string) GetWalletBalanceResponse {
	response := GetWalletBalanceResponse{}
	params := []interface{}{assetId}
	_ = n.makeRequest("getwalletbalance", params, &response)
	return response
}

func (n *RpcClient) GetWalletUnclaimedGas() GetWalletUnclaimedGasResponse {
	response := GetWalletUnclaimedGasResponse{}
	params := []interface{}{}
	_ = n.makeRequest("getwalletunclaimedgas", params, &response)
	return response
}

func (n *RpcClient) ImportPrivKey(wif string) ImportPrivKeyResponse {
	response := ImportPrivKeyResponse{}
	params := []interface{}{wif}
	_ = n.makeRequest("importprivkey", params, &response)
	return response
}

func (n *RpcClient) CalculateNetworkFee(tx string) CalculateNetworkFeeResponse {
	response := CalculateNetworkFeeResponse{}
	params := []interface{}{tx}
	_ = n.makeRequest("calculatenetworkfee", params, &response)
	return response
}

func (n *RpcClient) ListAddress() ListAddressResponse {
	response := ListAddressResponse{}
	params := []interface{}{}
	_ = n.makeRequest("listaddress", params, &response)
	return response
}

func (n *RpcClient) OpenWallet(path string, password string) OpenWalletResponse {
	response := OpenWalletResponse{}
	params := []interface{}{path, password}
	_ = n.makeRequest("openwallet", params, &response)
	return response
}

func (n *RpcClient) SendFrom(assetId string, fromAddress string, toAddress string, amount string, signerAddresses []string) SendFromResponse {
	response := SendFromResponse{}
	params := []interface{}{assetId, fromAddress, toAddress, amount}
	if len(signerAddresses) > 0 {
		params = append(params, signerAddresses)
	}
	_ = n.makeRequest("sendfrom", params, &response)
	return response
}

func (n *RpcClient) SendMany(fromAddress string, toAddresses []string, signerAddresses []string) SendManyResponse {
	response := SendManyResponse{}
	var params []interface{}
	if fromAddress != "" {
		params = append(params, fromAddress)
	}
	params = append(params, toAddresses)
	if len(signerAddresses) > 0 {
		params = append(params, signerAddresses)
	}
	_ = n.makeRequest("sendmany", params, &response)
	return response
}

func (n *RpcClient) SendToAddress(assetId string, toAddress string, amount string) SendToAddressResponse {
	response := SendToAddressResponse{}
	params := []interface{}{assetId, toAddress, amount}
	_ = n.makeRequest("sendtoaddress", params, &response)
	return response
}

func (n *RpcClient) InvokeContractVerify(scriptHash string, args []models.RpcContractParameter,
	signersOrWitnesses interface{}) InvokeResultResponse {

	response := InvokeResultResponse{}
	params := []interface{}{scriptHash}
	if args != nil {
		params = append(params, args)
	}
	if signers, ok := signersOrWitnesses.([]models.RpcSigner); ok {
		if len(params) == 1 {
			params = append(params, []models.RpcContractParameter{})
		}
		params = append(params, signers)
	} else if witnesses, ok := signersOrWitnesses.([]models.RpcWitness); ok {
		if len(params) == 1 {
			params = append(params, []models.RpcContractParameter{})
		}
		params = append(params, witnesses)
	}
	_ = n.makeRequest("invokecontractverify", params, &response)
	return response
}
