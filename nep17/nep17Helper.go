package nep17

import (
	"math/big"

	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/sc"
)

// Nep17Helper is nep17 wrapper class, api reference: https://github.com/neo-project/proposals/tree/nep-17
type Nep17Helper struct {
	ScriptHash *helper.UInt160 // scriptHash of nep17 token
	Client     rpc.IRpcClient
}

func NewNep17Helper(scriptHash *helper.UInt160, client rpc.IRpcClient) *Nep17Helper {
	if client == nil {
		return nil
	}
	return &Nep17Helper{
		ScriptHash: scriptHash,
		Client:     client,
	}
}

func (n *Nep17Helper) Symbol() (string, error) {
	sb := sc.NewScriptBuilder()
	sb.EmitDynamicCall(n.ScriptHash, "symbol", nil)
	script, err := sb.ToArray()
	if err != nil {
		return "", err
	}
	response := n.Client.InvokeScript(helper.BytesToHex(script), nil)
	stack, err := rpc.PopInvokeStack(response)
	if err != nil {
		return "", err
	}
	p, err := stack.ToParameter()
	if err != nil {
		return "", err
	}
	return  string(p.Value.([]byte)), nil
}

func (n *Nep17Helper) Decimals() (int, error) {
	sb := sc.NewScriptBuilder()
	sb.EmitDynamicCall(n.ScriptHash, "decimals", nil)
	script, err := sb.ToArray()
	if err != nil {
		return 0, err
	}
	response := n.Client.InvokeScript(helper.BytesToHex(script), nil)
	stack, err := rpc.PopInvokeStack(response)
	if err != nil {
		return 0, err
	}
	p, err := stack.ToParameter()
	if err != nil {
		return 0, err
	}
	return int(p.Value.(*big.Int).Int64()), nil
}

func (n *Nep17Helper) TotalSupply() (*big.Int, error) {
	sb := sc.NewScriptBuilder()
	sb.EmitDynamicCall(n.ScriptHash, "totalSupply", nil)
	script, err := sb.ToArray()
	if err != nil {
		return nil, err
	}
	response := n.Client.InvokeScript(helper.BytesToHex(script), nil)
	stack, err := rpc.PopInvokeStack(response)
	if err != nil {
		return nil, err
	}
	p, err := stack.ToParameter()
	if err != nil {
		return nil, err
	}
	return p.Value.(*big.Int), nil
}

func (n *Nep17Helper) BalanceOf(account *helper.UInt160) (*big.Int, error) {
	sb := sc.NewScriptBuilder()
	param := sc.ContractParameter{
		Type:  sc.Hash160,
		Value: account,
	}
	sb.EmitDynamicCall(n.ScriptHash, "balanceOf", []interface{}{param})
	script, err := sb.ToArray()
	if err != nil {
		return nil, err
	}
	response := n.Client.InvokeScript(helper.BytesToHex(script), nil)
	stack, err := rpc.PopInvokeStack(response)
	if err != nil {
		return nil, err
	}
	p, err := stack.ToParameter()
	if err != nil {
		return nil, err
	}
	return p.Value.(*big.Int), nil
}
