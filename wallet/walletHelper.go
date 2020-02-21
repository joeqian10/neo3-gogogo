package wallet

import (
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/nep5"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/tx"
	"math/big"
)

type WalletHelper struct {
	Client  rpc.IRpcClient
	Account *Account
}

func NewWalletHelper(rpc rpc.IRpcClient, account *Account) *WalletHelper {
	return &WalletHelper{
		Client:  rpc,
		Account: account,
	}
}

// Transfer is used to transfer neo or gas or other nep5 asset, from Account
func (w *WalletHelper) Transfer(assetId helper.UInt160, to string, amount *big.Int) (string, error) {
	t, err := helper.AddressToScriptHash(to)
	if err != nil {
		return "", err
	}

	tr, err := nep5.NewNep5Helper(assetId, w.Client).CreateTransferTx(w.Account.KeyPair, t, amount)
	if err != nil {
		return "", err
	}

	// use RPC to send the tx
	response := w.Client.SendRawTransaction(hex.EncodeToString(tr.ToByteArray()))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return "", fmt.Errorf(msg)
	}
	return response.Result.Hash, nil
}

// ClaimGas for Account
func (w *WalletHelper) ClaimGas() (string, error) {
	f, err := helper.AddressToScriptHash(w.Account.Address)
	if err != nil {
		return "", err
	}
	balance, err := w.GetBalance(tx.NeoToken)
	if err != nil {
		return "", err
	}
	t, err := nep5.NewNep5Helper(tx.NeoToken, w.Client).CreateTransferTx(w.Account.KeyPair, f, balance)
	if err != nil {
		return "", err
	}
	// use RPC to send the tx
	response := w.Client.SendRawTransaction(hex.EncodeToString(t.ToByteArray()))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return "", fmt.Errorf(msg)
	}
	return response.Result.Hash, nil
}

// GetBalance is used to get balance of nep5 asset, include neo and gas
func (w *WalletHelper) GetBalance(assetId helper.UInt160) (balance *big.Int, err error) {
	account, err := helper.AddressToScriptHash(w.Account.Address)
	if err != nil {
		return nil, err
	}
	return nep5.NewNep5Helper(assetId, w.Client).BalanceOf(account)
}

// GetUnClaimedGas for Account
func (w *WalletHelper) GetUnClaimedGas() (float64, error) {
	hash, err := helper.AddressToScriptHash(w.Account.Address)
	if err != nil {
		return 0, err
	}
	height := int64(w.Client.GetBlockCount().Result - 1)
	sb := sc.NewScriptBuilder()
	_ = sb.EmitAppCall(tx.NeoToken, "unclaimedGas", []sc.ContractParameter{
		{Type: sc.Hash160, Value: hash.Bytes()},
		{Type: sc.Integer, Value: big.NewInt(height)}})
	script := hex.EncodeToString(sb.ToArray())

	response := w.Client.InvokeScript(script)
	stack, err := nep5.PopInvokeStack(response)
	if err != nil {
		return 0, err
	}
	value := float64(stack.ToParameter().Value.(*big.Int).Int64()) / tx.GasFactor
	return value, nil
}
