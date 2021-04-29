package models

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/sc"
)

type RpcContractState struct {
	Id            int                 `json:"id"`
	UpdateCounter uint16              `json:"updatecounter"`
	Hash          string              `json:"hash"`
	Nef           RpcNefFile          `json:"nef"`
	Manifest      RpcContractManifest `json:"manifest"`
}

type RpcNefFile struct {
	Magic    uint             `json:"magic"`
	Compiler string           `json:"compiler"`
	Tokens   []RpcMethodToken `json:"tokens"`
	Script   string           `json:"script"` // base64
	CheckSum uint             `json:"checksum"`
}

type RpcMethodToken struct {
	Hash            string `json:"hash"`
	Method          string `json:"method"`
	ParametersCount uint16 `json:"paramcount"`
	HasReturnValue  bool   `json:"hasreturnvalue"`
	CallFlags       string `json:"callflags"`
}

type RpcContractManifest struct {
	Name               string                  `json:"name"`
	Groups             []RpcContractGroup      `json:"groups"`
	SupportedStandards []string                `json:"supportedstandards"`
	Abi                RpcContractAbi          `json:"abi"`
	Permissions        []RpcContractPermission `json:"permissions"`
	Trusts             []string                `json:"trusts"`
	Extra              interface{}             `json:"extra"`
}

type RpcContractGroup struct {
	PubKey    string `json:"pubkey"`
	Signature string `json:"signature"` // base64
}

type RpcContractAbi struct {
	Methods []RpcContractMethodDescriptor `json:"methods"`
	Events  []RpcContractEventDescriptor  `json:"events"`
}

type RpcContractMethodDescriptor struct {
	Name       string                           `json:"name"`
	Parameters []RpcContractParameterDefinition `json:"parameters"`
	ReturnType string                           `json:"returntype"`
	Offset     int                              `json:"offset"`
	Safe       bool                             `json:"safe"`
}

type RpcContractEventDescriptor struct {
	Name       string                           `json:"name"`
	Parameters []RpcContractParameterDefinition `json:"parameters"`
}

type RpcContractParameterDefinition struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (mp *RpcContractParameterDefinition) ToContractParameterType() (sc.ContractParameterType, error) {
	if mp == nil {
		return 0, fmt.Errorf("MethodParameter is nil")
	}
	return sc.NewContractParameterTypeFromString(mp.Type)
}

type RpcContractPermission struct {
	Contract string   `json:"contract"`
	Methods  []string `json:"methods"`
}

func (cs *RpcContractState) ToContract() (*sc.Contract, error) {
	if cs == nil {
		return nil, fmt.Errorf("ContractState is nil")
	}
	script := helper.HexToBytes(cs.Nef.Script)
	types := make([]sc.ContractParameterType, 0)
	for _, method := range cs.Manifest.Abi.Methods {
		if method.Name == "verify" {
			if method.Parameters != nil && len(method.Parameters) != 0 {
				for _, param := range method.Parameters {
					t, err := param.ToContractParameterType()
					if err != nil {
						return nil, err
					}
					types = append(types, t)
				}
			}
			break
		}
	}
	return &sc.Contract{
		Script:        script,
		ParameterList: types,
	}, nil
}
