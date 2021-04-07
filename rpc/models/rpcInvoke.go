package models

import (
	"encoding/base64"
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/sc"
	"math/big"
	"strconv"
)

type InvokeResult struct {
	Script      string        `json:"script"`
	State       string        `json:"state"`
	GasConsumed string        `json:"gasconsumed"`
	Exception   string        `json:"exception"`
	Stack       []InvokeStack `json:"stack"`
	Tx          string        `json:"tx"`
}

type InvokeStack struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// Convert converts interface{} "Value" to string or []InvokeStack depending on the "Type"
func (s *InvokeStack) Convert() {
	if s.Type != "Array" {
		switch s.Type {
		case "Boolean":
			if b, ok := s.Value.(bool); ok {
				s.Value = strconv.FormatBool(b)
			}
			break
		case "Integer":
			if num, ok := s.Value.(int); ok {
				s.Value = strconv.Itoa(num)
			}
			break
		}
	} else {
		vs := s.Value.([]interface{})
		result := make([]InvokeStack, len(vs))
		for i, v := range vs {
			m := v.(map[string]interface{})
			s2 := InvokeStack{
				Type:  m["type"].(string),
				Value: m["value"],
			}
			s2.Convert()
			result[i] = s2
		}
		s.Value = result
	}
}

func (s *InvokeStack) ToParameter() (*sc.ContractParameter, error) {
	var parameter *sc.ContractParameter = new(sc.ContractParameter)
	var err error
	s.Convert()
	switch s.Type {
	case "Signature":
		parameter.Type = sc.Signature
		parameter.Value, err = crypto.Base64Decode(s.Value.(string))
		break
	case "ByteArray":
		parameter.Type = sc.ByteArray
		parameter.Value, err = crypto.Base64Decode(s.Value.(string))
		break
	case "Boolean":
		parameter.Type = sc.Boolean
		parameter.Value, err = strconv.ParseBool(s.Value.(string))
		break
	case "Integer":
		parameter.Type = sc.Integer
		var b bool
		parameter.Value, b = new(big.Int).SetString(s.Value.(string), 10)
		if !b {
			err = fmt.Errorf("convert to Integer failed")
		}
		break
	case "Hash160":
		parameter.Type = sc.Hash160
		parameter.Value, err = helper.UInt160FromString(s.Value.(string))
		break
	case "Hash256":
		parameter.Type = sc.Hash256
		parameter.Value, err = helper.UInt256FromString(s.Value.(string))
		break
	case "PublicKey":
		parameter.Type = sc.PublicKey
		parameter.Value, err = crypto.NewECPointFromString(s.Value.(string))
		break
	case "String":
		v, err1 := crypto.Base64Decode(s.Value.(string))
		if err1 != nil {
			err = err1
			break
		}
		parameter.Type = sc.String
		parameter.Value = string(v)
		break
	case "Array":
		parameter.Type = sc.Array
		a := s.Value.([]InvokeStack)
		r := make([]sc.ContractParameter, len(a))
		for i, _ := range a {
			t, err1 := a[i].ToParameter()
			if err1 != nil {
				err = err1
				break
			}
			r[i] = *t
		}
		break
	default:
		err = fmt.Errorf("invalid type")
	}
	return parameter, err
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
		parameter.Value, _ = crypto.NewECPointFromString(r.Value)
	case "String":
		parameter.Type = sc.String
		parameter.Value = r.Value
	}
	return parameter
}
