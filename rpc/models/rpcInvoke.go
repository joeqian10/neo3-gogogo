package models

import (
	"encoding/base64"
	"math/big"
	"strconv"

	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/wallet/keys"
)

type InvokeResult struct {
	Script      string              `json:"script"`
	State       string              `json:"state"`
	GasConsumed string              `json:"gas_consumed"`
	Stack       []InvokeStackResult `json:"stack"`
}

type InvokeStackResult struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (r *InvokeStackResult) ToParameter() sc.ContractParameter {
	var parameter sc.ContractParameter

	switch r.Type {
	case "Signature":
		parameter.Type = sc.Signature
		parameter.Value, _ = base64.StdEncoding.DecodeString(r.Value)
	case "ByteArray":
		parameter.Type = sc.ByteArray
		parameter.Value, _ = base64.StdEncoding.DecodeString(r.Value)
	case "Boolean":
		parameter.Type = sc.Boolean
		parameter.Value, _ = strconv.ParseBool(r.Value)
	case "Integer":
		parameter.Type = sc.Integer
		parameter.Value, _ = new(big.Int).SetString(r.Value, 10)
	case "Hash160":
		parameter.Type = sc.Hash160
		parameter.Value, _ = helper.UInt160FromString(r.Value)
	case "Hash256":
		parameter.Type = sc.Hash256
		parameter.Value, _ = helper.UInt256FromString(r.Value)
	case "PublicKey":
		parameter.Type = sc.PublicKey
		parameter.Value, _ = keys.NewPublicKeyFromString(r.Value)
	case "String":
		parameter.Type = sc.String
		parameter.Value = r.Value
	}

	return parameter
}

type InvokeFunctionStackArg struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewInvokeFunctionStackArg(t string, v string) InvokeFunctionStackArg {
	return InvokeFunctionStackArg{Type: t, Value: v}
}

func NewInvokeFunctionStackByteArray(value string) InvokeFunctionStackArg {
	return InvokeFunctionStackArg{Type: "ByteArray", Value: value}
}
