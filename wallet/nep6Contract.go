package wallet

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/sc"
)

//type NEP6Contract struct {
//	sc.Contract
//	ParameterNames []string
//	Deployed bool
//}

type NEP6Contract struct {
	Script         string                    `json:"script"`
	Parameters     []NEP6ParameterDescriptor `json:"parameters"`
	Deployed       bool                      `json:"deployed"`
	parameterList  []sc.ContractParameterType
	parameterNames []string
}

type NEP6ParameterDescriptor struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewNEP6Contract(script []byte, paramList []sc.ContractParameterType, paramNames []string, isDeployed bool) (*NEP6Contract, error) {
	if len(paramList) != len(paramNames) {
		return nil, fmt.Errorf("parameter length does not match")
	}
	params := make([]NEP6ParameterDescriptor, len(paramList))
	for i := 0; i < len(paramList); i++ {
		params[i] = NEP6ParameterDescriptor{
			Name: paramNames[i],
			Type: paramList[i].String(),
		}
	}
	return &NEP6Contract{
		Script:         crypto.Base64Encode(script),
		parameterList:  paramList,
		parameterNames: paramNames,
		Parameters:     params,
		Deployed:       isDeployed,
	}, nil
}

func (c *NEP6Contract) GetScript() []byte {
	b, _ := crypto.Base64Decode(c.Script)
	return b
}

func (c *NEP6Contract) GetScriptHash() *helper.UInt160 {
	return helper.UInt160FromBytes(crypto.Hash160(c.GetScript()))
}

func (c *NEP6Contract) ToContract() (*sc.Contract, error) {
	paramList := []sc.ContractParameterType{}
	if len(c.Parameters) != 0 {
		var err error
		paramList = make([]sc.ContractParameterType, len(c.Parameters))
		for i:=0; i < len(c.Parameters); i++ {
			paramList[i], err = sc.NewContractParameterTypeFromString(c.Parameters[i].Type)
			if err != nil {
				return nil, err
			}
		}
	}
	return &sc.Contract{
		Script:        c.GetScript(),
		ParameterList: paramList,
	}, nil
}

//// Contract represents a subset of the smart contract to embed in the
//// NEP6Account so it's NEP-6 compliant.
//type Contract struct {
//	// Script hash of the contract deployed on the block chain.
//	Script string `json:"script"`
//
//	// A list of parameters used deploying this contract.
//	Parameters []interface{} `json:"parameters"`
//
//	// Indicates whether the contract has been deployed to the block chain.
//	Deployed bool `json:"deployed"`
//}
