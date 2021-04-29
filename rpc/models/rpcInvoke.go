package models

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/vm"
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

// Convert converts interface{} "Value" to string or []InvokeStack or map[InvokeStack]InvokeStack depending on the "Type"
func (s *InvokeStack) Convert() {
	switch s.Type {
	case vm.Array.String():
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
		break
	case vm.Boolean.String():
		if b, ok := s.Value.(bool); ok {
			s.Value = strconv.FormatBool(b)
		}
		break
	case vm.Buffer.String(), vm.ByteString.String():
		// nothing to handle
		break
	case vm.Integer.String():
		if num, ok := s.Value.(int); ok {
			s.Value = strconv.Itoa(num)
		}
		// else if number in string, nothing to handle
		break
	case vm.Map.String():
		vs := s.Value.([]interface{})
		result := make(map[InvokeStack]InvokeStack)
		for _, v := range vs {
			m := v.(map[string]interface{})
			key := m["key"].(map[string]interface{})
			value := m["value"].(map[string]interface{})
			s2 := InvokeStack{
				Type:  key["type"].(string),
				Value: key["value"],
			}
			s3 := InvokeStack{
				Type:  value["type"].(string),
				Value: value["value"],
			}
			s2.Convert()
			s3.Convert()
			result[s2] = s3
		}
		s.Value = result
		break
	case vm.Pointer.String():
		if num, ok := s.Value.(int); ok {
			s.Value = strconv.Itoa(num)
		}
		break
	}
}

func (s *InvokeStack) ToParameter() (*sc.ContractParameter, error) {
	var parameter *sc.ContractParameter = new(sc.ContractParameter)
	var err error
	s.Convert()
	switch s.Type {
	case vm.Array.String():
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
	case vm.Boolean.String():
		parameter.Type = sc.Boolean
		parameter.Value, err = strconv.ParseBool(s.Value.(string))
		break
	case vm.Buffer.String(), vm.ByteString.String():
		parameter.Type = sc.ByteArray
		parameter.Value, err = crypto.Base64Decode(s.Value.(string))
		break
	case vm.Integer.String():
		parameter.Type = sc.Integer
		var b bool
		parameter.Value, b = new(big.Int).SetString(s.Value.(string), 10)
		if !b {
			err = fmt.Errorf("converting vm.Integer to sc.Integer failed")
		}
		break
	case vm.Map.String():
		parameter.Type = sc.Map
		parameter.Value = s.Value // map[InvokeStack]InvokeStack
	case vm.Pointer.String():
		parameter.Type = sc.Integer
		var b bool
		parameter.Value, b = new(big.Int).SetString(s.Value.(string), 10)
		if !b {
			err = fmt.Errorf("converting vm.Pointer to sc.Integer failed")
		}
		break
	//case "Signature":
	//	parameter.Type = sc.Signature
	//	parameter.Value, err = crypto.Base64Decode(s.Value.(string))
	//	break
	//case "Hash160":
	//	parameter.Type = sc.Hash160
	//	parameter.Value, err = helper.UInt160FromString(s.Value.(string))
	//	break
	//case "Hash256":
	//	parameter.Type = sc.Hash256
	//	parameter.Value, err = helper.UInt256FromString(s.Value.(string))
	//	break
	//case "PublicKey":
	//	parameter.Type = sc.PublicKey
	//	parameter.Value, err = crypto.NewECPointFromString(s.Value.(string))
	//	break
	//case "String":
	//	v, err1 := crypto.Base64Decode(s.Value.(string))
	//	if err1 != nil {
	//		err = err1
	//		break
	//	}
	//	parameter.Type = sc.String
	//	parameter.Value = string(v)
	//	break
	default:
		err = fmt.Errorf("not supported stack item type")
	}
	return parameter, err
}
