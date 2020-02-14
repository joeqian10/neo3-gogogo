package wallet

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/nep5"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/tx"
)

type WalletHelper struct {
	RpcClient rpc.IRpcClient
	Account   *Account
}

func NewWalletHelper(rpc rpc.IRpcClient, account *Account) *WalletHelper {
	return &WalletHelper{
		RpcClient: rpc,
		Account:   account,
	}
}

// GetBalance is used to get balance of nep5 asset, include neo and gas
func (w *WalletHelper) GetBalance(assetId string, address string) (balance *big.Int, err error) {
	account, err := helper.AddressToScriptHash(address)
	if err != nil {
		return nil, err
	}
	assetHash, err := helper.UInt160FromString(assetId)
	if err != nil {
		return nil, err
	}
	return nep5.NewNep5Helper(w.RpcClient).BalanceOf(assetHash, account)
}

// Transfer is used to transfer neo or gas or other utxo asset, single signature
func (w *WalletHelper) Transfer(assetId helper.UInt160, from string, to string, amount *big.Int) (bool, error) {
	f, err := helper.AddressToScriptHash(from)
	if err != nil {
		return false, err
	}
	t, err := helper.AddressToScriptHash(to)
	if err != nil {
		return false, err
	}

	tx, err := nep5.NewNep5Helper(w.RpcClient).CreateTransferTx(assetId, f, t, amount)
	if err != nil {
		return false, err
	}
	// sign
	err = tx.AddSignature(w.Account.KeyPair)
	if err != nil {
		return false, err
	}
	// use RPC to send the tx
	response := w.RpcClient.SendRawTransaction(hex.EncodeToString(tx.GetHashData()))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return false, fmt.Errorf(msg)
	}
	return true, nil
}

// ClaimGas for Account
func (w *WalletHelper) ClaimGas() (bool, error) {
	f, err := helper.AddressToScriptHash(w.Account.Address)
	if err != nil {
		return false, err
	}
	balance, err := w.GetBalance(tx.NeoTokenId, w.Account.Address)
	if err != nil {
		return false, err
	}
	tx, err := nep5.NewNep5Helper(w.RpcClient).CreateTransferTx(tx.NeoToken, f, f, balance)
	if err != nil {
		return false, err
	}
	// sign
	err = tx.AddSignature(w.Account.KeyPair)
	if err != nil {
		return false, err
	}
	// use RPC to send the tx
	response := w.RpcClient.SendRawTransaction(hex.EncodeToString(tx.GetHashData()))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return false, fmt.Errorf(msg)
	}
	return true, nil
}
