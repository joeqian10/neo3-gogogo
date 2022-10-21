package nep17

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"math/big"
	"strconv"

	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc"
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
	response := n.Client.InvokeFunction(n.ScriptHash.String(), "symbol", nil, nil, false)
	stacks, err := rpc.PopInvokeStacks(response)
	if err != nil {
		return "", err
	}
	r := stacks[0]
	b, err := crypto.Base64Decode(r.Value.(string))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (n *Nep17Helper) Decimals() (int, error) {
	response := n.Client.InvokeFunction(n.ScriptHash.String(), "decimals", nil, nil, false)
	stacks, err := rpc.PopInvokeStacks(response)
	if err != nil {
		return -1, err
	}
	r := stacks[0]
	i, err := strconv.Atoi(r.Value.(string))
	if err != nil {
		return -1, err
	}
	return i, nil
}

func (n *Nep17Helper) TotalSupply() (*big.Int, error) {
	response := n.Client.InvokeFunction(n.ScriptHash.String(), "totalSupply", nil, nil, false)
	stacks, err := rpc.PopInvokeStacks(response)
	if err != nil {
		return nil, err
	}
	r := stacks[0]
	b, c := new(big.Int).SetString(r.Value.(string), 10)
	if !c {
		return nil, fmt.Errorf("converting value failed")
	}
	return b, nil
}

func (n *Nep17Helper) BalanceOf(account *helper.UInt160) (*big.Int, error) {
	args := []models.RpcContractParameter{
		{
			Type:  "Hash160",
			Value: account,
		},
	}

	response := n.Client.InvokeFunction(n.ScriptHash.String(), "balanceOf", args, nil, false)
	stacks, err := rpc.PopInvokeStacks(response)
	if err != nil {
		return nil, err
	}
	r := stacks[0]
	b, c := new(big.Int).SetString(r.Value.(string), 10)
	if !c {
		return nil, fmt.Errorf("converting value failed")
	}
	return b, nil
}
