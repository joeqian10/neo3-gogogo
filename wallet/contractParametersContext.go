package wallet

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/tx"
	"sort"
)

type ContextItem struct {
	Script     []byte
	Parameters []sc.ContractParameter
	Signatures map[string][]byte // Dictionary<ECPoint, byte[]>
}

func NewContextItem(contract *sc.Contract) *ContextItem {
	params := make([]sc.ContractParameter, len(contract.ParameterList))
	for i := 0; i < len(contract.ParameterList); i++ {
		params[i] = sc.ContractParameter{Type: contract.ParameterList[i]}
	}
	return &ContextItem{
		Script:     contract.Script,
		Parameters: params,
		Signatures: make(map[string][]byte, 0),
	}
}

type signatureHelper struct {
	Index     int
	Signature []byte
}

type signatureHelperSlice []signatureHelper

func (s signatureHelperSlice) Len() int {
	return len(s)
}

func (s signatureHelperSlice) Less(i int, j int) bool {
	return s[i].Index < s[j].Index
}

func (s signatureHelperSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ContractParametersContext struct {
	Verifiable   tx.IVerifiable // transaction ?
	ContextItems map[helper.UInt160]*ContextItem

	scriptHashes []helper.UInt160
}

func NewContractParametersContract(verifiable tx.IVerifiable) *ContractParametersContext {
	return &ContractParametersContext{
		Verifiable:   verifiable,
		ContextItems: make(map[helper.UInt160]*ContextItem, 0),
	}
}

func (c *ContractParametersContext) GetCompleted() bool {
	if len(c.ContextItems) < len(c.scriptHashes) {
		return false
	}
	for k, v := range c.ContextItems {
		if k.Equals(helper.UInt160Zero) {
			return false
		}
		for _, param := range v.Parameters {
			if param.Value == nil {
				return false
			}
		}
	}
	return true
}

func (c *ContractParametersContext) GetScriptHashes() []helper.UInt160 {
	if c.scriptHashes == nil {
		c.scriptHashes = c.Verifiable.GetScriptHashesForVerifying()
	}
	return c.scriptHashes
}

func (c *ContractParametersContext) AddItemWithIndex(contract *sc.Contract, index int, parameter interface{}) bool {
	item := c.createItem(contract)
	if item == nil {
		return false
	}
	item.Parameters[index].Value = parameter
	return true
}

func (c *ContractParametersContext) AddItemWithParams(contract *sc.Contract, parameters []interface{}) bool {
	item := c.createItem(contract)
	if item == nil {
		return false
	}
	if parameters != nil && len(parameters) != 0 {
		for index := 0; index < len(parameters); index++ {
			item.Parameters[index].Value = parameters[index]
		}
	}
	return true
}

func (c *ContractParametersContext) AddSignature(contract *sc.Contract, pubKey *crypto.ECPoint, signature []byte) (bool, error) {
	bs := sc.ByteSlice(contract.Script)
	if b, _, points := bs.IsMultiSigContractWithPoints(); b {
		if !pubKey.ExistsIn(points) {
			return false, nil
		}
		item := c.createItem(contract)
		if item == nil {
			return false, nil
		}

		all := true
		for _, param := range item.Parameters {
			if param.Value == nil {
				all = false
				break
			}
		}
		if all {
			return false, nil
		}

		if _, ok := item.Signatures[pubKey.String()]; ok {
			return false, nil
		}
		item.Signatures[pubKey.String()] = signature

		if len(item.Signatures) == len(contract.ParameterList) {
			dic := map[string]int{}
			for i, p := range points {
				dic[p.String()] = i
			}

			signatureHelpers := make([]signatureHelper, len(item.Signatures))
			in := 0
			for k, v := range item.Signatures {
				signatureHelpers[in] = signatureHelper{
					Index:     dic[k],
					Signature: v,
				}
				in++
			}
			// sort by descending
			sort.Sort(sort.Reverse(signatureHelperSlice(signatureHelpers)))

			for i := 0; i < len(signatureHelpers); i++ {
				if !c.AddItemWithIndex(contract, i, signatureHelpers[i]) {
					return false, fmt.Errorf("invalid operation when adding item")
				}
			}
			item.Signatures = nil
		}
		return true, nil
	} else {
		index := -1
		for i := 0; i < len(contract.ParameterList); i++ {
			if contract.ParameterList[i] == sc.Signature {
				if index >= 0 {
					return false, fmt.Errorf("not supported operation")
				} else {
					index = i
				}
			}
		}
		if index == -1 {
			// unable to find ContractParameterType.Signature in contract.ParameterList
			// return now to prevent array index out of bounds exception
			return false, nil
		}
		item := c.createItem(contract)
		if item == nil {
			return false, nil
		}
		if _, ok := item.Signatures[pubKey.String()]; ok {
			return false, nil
		}
		item.Signatures[pubKey.String()] = signature
		item.Parameters[index].Value = signature
		return true, nil
	}
}

func (c *ContractParametersContext) createItem(contract *sc.Contract) *ContextItem {
	scriptHash := *contract.GetScriptHash()
	if item, ok := c.ContextItems[scriptHash]; ok {
		return item
	}
	if !contract.GetScriptHash().ExistsIn(c.scriptHashes) {
		return nil
	}
	item := NewContextItem(contract)
	c.ContextItems[scriptHash] = item
	return item
}

func (c *ContractParametersContext) GetParameter(scriptHash *helper.UInt160, index int) *sc.ContractParameter {
	params := c.GetParameters(scriptHash)
	if params == nil {
		return nil
	}
	return &params[index]
}

func (c *ContractParametersContext) GetParameters(scriptHash *helper.UInt160) []sc.ContractParameter {
	if item, ok := c.ContextItems[*scriptHash]; ok {
		return item.Parameters
	}
	return nil
}

func (c *ContractParametersContext) GetSignatures(scriptHash *helper.UInt160) map[string][]byte {
	if item, ok := c.ContextItems[*scriptHash]; ok {
		return item.Signatures
	}
	return nil
}

func (c *ContractParametersContext) GetScript(scriptHash *helper.UInt160) []byte {
	if item, ok := c.ContextItems[*scriptHash]; ok {
		return item.Script
	}
	return nil
}

func (c *ContractParametersContext) GetWitnesses() ([]tx.Witness, error) {
	if !c.GetCompleted() {
		return nil, fmt.Errorf("invalid operation when getting witnesses")
	}
	witnesses := make([]tx.Witness, len(c.scriptHashes))
	for i := 0; i < len(c.scriptHashes); i++ {
		item := c.ContextItems[c.scriptHashes[i]]
		sb := sc.NewScriptBuilder()
		for j := len(item.Parameters) - 1; j >= 0; j-- {
			sb.EmitPushParameter(item.Parameters[j])
		}
		is, err := sb.ToArray()
		if err != nil {
			return nil, err
		}
		vs := make([]byte, 0)
		if item.Script != nil {
			vs = item.Script
		}
		witnesses[i] = tx.Witness{
			InvocationScript:   is,
			VerificationScript: vs,
		}
	}
	return witnesses, nil
}
