package nep5

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/joeqian10/neo3-gogogo/wallet/keys"
	"math/big"

	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/tx"
)

// Nep5Helper is nep5 wrapper class, api reference: https://github.com/neo-project/proposals/blob/master/nep-5.mediawiki#name
type Nep5Helper struct {
	ScriptHash helper.UInt160
	Client     rpc.IRpcClient
}

func NewNep5Helper(scriptHash helper.UInt160, client rpc.IRpcClient) *Nep5Helper {
	if client == nil {
		return nil
	}
	return &Nep5Helper{
		ScriptHash: scriptHash,
		Client:     client,
	}
}

func (n *Nep5Helper) Name() (string, error) {
	sb := sc.NewScriptBuilder()
	_ = sb.EmitAppCall(n.ScriptHash, "name", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	stack, err := PopInvokeStack(response)
	if err != nil {
		return "", err
	}
	return string(stack.ToParameter().Value.([]byte)), nil
}

func (n *Nep5Helper) Symbol() (string, error) {
	sb := sc.NewScriptBuilder()
	_ = sb.EmitAppCall(n.ScriptHash, "symbol", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	stack, err := PopInvokeStack(response)
	if err != nil {
		return "", err
	}
	return string(stack.ToParameter().Value.([]byte)), nil
}

func (n *Nep5Helper) Decimals() (int, error) {
	sb := sc.NewScriptBuilder()
	_ = sb.EmitAppCall(n.ScriptHash, "decimals", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	stack, err := PopInvokeStack(response)
	if err != nil {
		return 0, err
	}
	return int(stack.ToParameter().Value.(*big.Int).Int64()), nil
}

func (n *Nep5Helper) BalanceOf(account helper.UInt160) (*big.Int, error) {
	return tx.NewTransactionBuilderFromClient(n.Client).GetBalance(n.ScriptHash, account)
}

func (n *Nep5Helper) TotalSupply() (*big.Int, error) {
	sb := sc.NewScriptBuilder()
	_ = sb.EmitAppCall(n.ScriptHash, "totalSupply", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	stack, err := PopInvokeStack(response)
	if err != nil {
		return nil, err
	}
	return stack.ToParameter().Value.(*big.Int), nil
}

func PopInvokeStack(response rpc.InvokeResultResponse) (*models.InvokeStackResult, error) {
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
	return &stack, nil
}

// CreateTransferTx creates nep-5 transfer transaction
func (n *Nep5Helper) CreateTransferTx(fromKey *keys.KeyPair, to helper.UInt160, amount *big.Int) (*tx.Transaction, error) {
	sb := sc.NewScriptBuilder()
	from := fromKey.PublicKey.ScriptHash()
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
	_ = sb.EmitAppCall(n.ScriptHash, "transfer", []sc.ContractParameter{cp1, cp2, cp3})
	script := sb.ToArray()

	tb := tx.NewTransactionBuilderFromClient(n.Client)
	t, err := tb.MakeTransaction(script, from, nil, []*tx.Cosigner{{Account: from, Scopes: tx.CalledByEntry}})
	if err != nil {
		return nil, err
	}
	err = tb.AddSignature(fromKey)
	if err != nil {
		return nil, err
	}
	err = tb.Sign()
	if err != nil {
		return nil, err
	}
	return t, nil
}
