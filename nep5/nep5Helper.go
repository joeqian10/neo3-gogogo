package nep5

import (
	"fmt"
	"math/big"

	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/tx"
)

// nep5 wrapper class, api reference: https://github.com/neo-project/proposals/blob/master/nep-5.mediawiki#name
type Nep5Helper struct {
	Client rpc.IRpcClient
}

func NewNep5Helper(client rpc.IRpcClient) *Nep5Helper {
	if client == nil {
		return nil
	}
	return &Nep5Helper{
		Client: client,
	}
}

func (n *Nep5Helper) TotalSupply(scriptHash helper.UInt160) (*big.Int, error) {
	sb := sc.NewScriptBuilder()
	sb.EmitAppCall(scriptHash, "totalSupply", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return nil, fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return nil, fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return nil, fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	return stack.ToParameter().Value.(*big.Int), nil
}

func (n *Nep5Helper) Name(scriptHash helper.UInt160) (string, error) {
	sb := sc.NewScriptBuilder()
	sb.EmitAppCall(scriptHash, "name", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return "", fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return "", fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return "", fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	return string(stack.ToParameter().Value.([]byte)), nil
}

func (n *Nep5Helper) Symbol(scriptHash helper.UInt160) (string, error) {
	sb := sc.NewScriptBuilder()
	sb.EmitAppCall(scriptHash, "symbol", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return "", fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return "", fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return "", fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	return string(stack.ToParameter().Value.([]byte)), nil
}

func (n *Nep5Helper) Decimals(scriptHash helper.UInt160) (int, error) {
	sb := sc.NewScriptBuilder()
	sb.EmitAppCall(scriptHash, "decimals", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return 0, fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return 0, fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return 0, fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	return int(stack.ToParameter().Value.(*big.Int).Int64()), nil
}

func (n *Nep5Helper) BalanceOf(scriptHash helper.UInt160, account helper.UInt160) (*big.Int, error) {
	return tx.NewTransactionBuilderFromClient(n.Client).GetBalance(scriptHash, account)
}

// CreateTransferTx creates nep-5 transfer transaction
func (n *Nep5Helper) CreateTransferTx(scriptHash helper.UInt160, from helper.UInt160, to helper.UInt160, amount *big.Int) (*tx.Transaction, error) {
	sb := sc.NewScriptBuilder()
	cp1 := sc.ContractParameter{
		Type:  sc.Hash160,
		Value: from.Bytes(),
	}
	cp2 := sc.ContractParameter{
		Type:  sc.Hash160,
		Value: to.Bytes(),
	}
	cp3 := sc.ContractParameter{
		Type:  sc.Integer,
		Value: amount,
	}
	sb.EmitAppCall(scriptHash, "transfer", []sc.ContractParameter{cp1, cp2, cp3})
	script := sb.ToArray()

	tx, err := tx.NewTransactionBuilderFromClient(n.Client).MakeTransaction(script, from, nil, []*tx.Cosigner{{Account: from, Scopes: tx.CalledByEntry}})
	if err != nil {
		return nil, err
	}
	return tx, nil
}
