package rpc

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/joeqian10/neo3-gogogo/sc"
)

type InvokeResultResponse struct {
	RpcResponse
	ErrorResponse
	Result models.InvokeResult `json:"result"`
}

type TraverseIteratorResponse struct {
	RpcResponse
	ErrorResponse
	Result []models.InvokeStack `json:"result"`
}

type TerminateSessionResponse struct {
	RpcResponse
	ErrorResponse
	Result bool `json:"result"`
}

type GetUnclaimedGasResponse struct {
	RpcResponse
	ErrorResponse
	Result models.UnclaimedGas `json:"result"`
}

// InvokeFunction params: scriptHash and method are necessary, set args = nil, signersOrWitnesses = nil, useDiagnostic = false if not necessary
func (n *RpcClient) InvokeFunction(scriptHash string, method string, args []models.RpcContractParameter,
	signersOrWitnesses interface{}, useDiagnostic bool) InvokeResultResponse {

	u, err := helper.UInt160FromString(scriptHash)
	if err != nil {
		return InvokeResultResponse{
			ErrorResponse: ErrorResponse{Error: RpcError{Code: -1, Message: "invalid contract script hash: " + err.Error()}},
		}
	}
	callArgs := make([]interface{}, len(args))
	for i, _ := range args {
		t, err := sc.NewContractParameterTypeFromString(args[i].Type)
		if err != nil {
			return InvokeResultResponse{
				ErrorResponse: ErrorResponse{Error: RpcError{Code: -1, Message: err.Error()}},
			}
		}
		callArgs[i] = &sc.ContractParameter{
			Type:  t,
			Value: args[i].Value,
		}
	}
	script, err := sc.MakeScript(u, method, callArgs)
	if err != nil {
		return InvokeResultResponse{
			ErrorResponse: ErrorResponse{Error: RpcError{Code: -1, Message: err.Error()}},
		}
	}

	return n.InvokeScript(crypto.Base64Encode(script), signersOrWitnesses, useDiagnostic)

	//response := InvokeResultResponse{}
	//params := []interface{}{scriptHash, method}
	//if args != nil {
	//	params = append(params, args) // params[2]
	//}
	//
	//if signers, ok := signersOrWitnesses.([]models.RpcSigner); ok {
	//	if len(params) == 2 {
	//		params = append(params, []models.RpcContractParameter{}) // params[2]
	//	}
	//	params = append(params, signers) // params[3]
	//} else if witnesses, ok := signersOrWitnesses.([]models.RpcWitness); ok {
	//	if len(params) == 2 {
	//		params = append(params, []models.RpcContractParameter{}) // params[2]
	//	}
	//	params = append(params, witnesses) // params[3]
	//}
	//
	//if useDiagnostic {
	//	if len(params) == 2 {
	//		params = append(params, []models.RpcContractParameter{}) // params[2]
	//		params = append(params, []models.RpcSigner{})            // params[3]
	//	} else if len(params) == 3 {
	//		params = append(params, []models.RpcSigner{}) // params[3]
	//	}
	//	params = append(params, useDiagnostic) // params[4]
	//}
	//
	//_ = n.makeRequest("invokefunction", params, &response)
	//return response
}

// InvokeScript params: scriptInBase64 is necessary, set signersOrWitnesses = nil, useDiagnostic = false if not necessary
func (n *RpcClient) InvokeScript(scriptInBase64 string, signersOrWitnesses interface{}, useDiagnostic bool) InvokeResultResponse {
	response := InvokeResultResponse{}
	params := []interface{}{scriptInBase64}

	if signers, ok := signersOrWitnesses.([]models.RpcSigner); ok {
		params = append(params, signers) // params[1]
	} else if witnesses, ok := signersOrWitnesses.([]models.RpcWitness); ok {
		params = append(params, witnesses) // params[1]
	}

	if useDiagnostic {
		if len(params) == 1 {
			params = append(params, []models.RpcSigner{}) // params[1]
		}
		params = append(params, useDiagnostic) // params[2]
	}

	_ = n.makeRequest("invokescript", params, &response)
	return response
}

func (n *RpcClient) TraverseIterator(sessionId string, iteratorId string, count int32) TraverseIteratorResponse {
	response := TraverseIteratorResponse{}
	params := []interface{}{sessionId, iteratorId, count}
	_ = n.makeRequest("traverseiterator", params, &response)
	return response
}

func (n *RpcClient) TerminateSession(sessionId string) TerminateSessionResponse {
	response := TerminateSessionResponse{}
	params := []interface{}{sessionId}
	_ = n.makeRequest("terminatesession", params, &response)
	return response
}

func (n *RpcClient) GetUnclaimedGas(address string) GetUnclaimedGasResponse {
	response := GetUnclaimedGasResponse{}
	params := []interface{}{address}
	_ = n.makeRequest("getunclaimedgas", params, &response)
	return response
}

func (n *RpcClient) InvokeFunctionAndIterate(scriptHash string, method string, args []models.RpcContractParameter,
	signersOrWitnesses interface{}, useDiagnostic bool, count int32) ([][]models.InvokeStack, error) {

	u, err := helper.UInt160FromString(scriptHash)
	if err != nil {
		return nil, err
	}
	callArgs := make([]interface{}, len(args))
	for i, _ := range args {
		t, err := sc.NewContractParameterTypeFromString(args[i].Type)
		if err != nil {
			return nil, err
		}
		callArgs[i] = sc.ContractParameter{Type: t, Value: args[i].Value}
	}
	script, err := sc.MakeScript(u, method, callArgs)
	if err != nil {
		return nil, err
	}
	return n.InvokeScriptAndIterate(crypto.Base64Encode(script), signersOrWitnesses, useDiagnostic, count)
}

func (n *RpcClient) InvokeScriptAndIterate(scriptInBase64 string, signersOrWitnesses interface{}, useDiagnostic bool,
	count int32) ([][]models.InvokeStack, error) {

	response := n.InvokeScript(scriptInBase64, signersOrWitnesses, useDiagnostic)

	invokeStacks, err := PopInvokeStacks(response)
	if err != nil {
		return nil, err
	}

	sessionId := response.Result.Session
	if sessionId == "" {
		return nil, nil
	}

	var iterateStacks [][]models.InvokeStack
	for _, invokeStack := range invokeStacks {
		if invokeStack.Type == "InteropInterface" &&
			invokeStack.Interface == "IIterator" &&
			invokeStack.Id != "" {
			// call TraverseIterator
			iterateResponse := n.TraverseIterator(sessionId, invokeStack.Id, count)
			if iterateResponse.HasError() {
				return nil, fmt.Errorf(iterateResponse.GetErrorInfo())
			}
			iterateStacks = append(iterateStacks, iterateResponse.Result)
		}
	}
	return iterateStacks, nil
}
