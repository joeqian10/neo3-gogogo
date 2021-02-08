package wallet

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/keys"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/tx"
)

type IWallet interface {
	GetName() string
	GetPath() string
	GetVersion() string

	ChangePassword(oldPassword string, newPassword string) bool
	Contains(scriptHash *helper.UInt160) bool

	CreateAccount() *IAccount
	CreateAccountWithPrivateKey(privateKey []byte) IAccount
	CreateAccountWithContract(contract *sc.Contract, key ...keys.KeyPair) IAccount
	CreateAccountWithScriptHash(scriptHash *helper.UInt160) IAccount
	DeleteAccount(scriptHash *helper.UInt160) bool
	GetAccount(scriptHash *helper.UInt160) IAccount
	GetAccounts() []IAccount

	//MakeTransaction(script []byte, sender *helper.UInt160, cosigners []tx.Signer, attributes []tx.ITransactionAttribute) (*tx.Transaction, error)
	//Sign(ctx ContractParametersContext) bool
}

func getSigners(sender *helper.UInt160, cosigners []tx.Signer) []tx.Signer {
	for i := 0; i < len(cosigners); i++ {
		if cosigners[i].Account.Equals(sender) {
			if i == 0 {
				return cosigners
			}
			result := make([]tx.Signer, len(cosigners))
			result[0] = cosigners[i]
			if i == len(cosigners)-1 {
				copy(result[1:], cosigners[0:i])
				return result
			} else {
				copy(result[1:i+1], cosigners[0:i])
				copy(result[i+1:], cosigners[i+1:])
				return result
			}
		}
	}
	signer := tx.NewSigner(sender, tx.None)
	return append([]tx.Signer{*signer}, cosigners...)
}

func Sign(verifiable tx.IVerifiable, pair *keys.KeyPair) ([]byte, error) {
	return SignWithMagic(verifiable, pair, tx.Neo3Magic)
}

func SignWithMagic(verifiable tx.IVerifiable, pair *keys.KeyPair, magic uint32) ([]byte, error) {
	return pair.Sign(tx.GetHashDataWithMagic(verifiable, magic))
}
