package wallet

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/keys"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/tx"
)

type IWallet interface {
	GetName() string
	GetPath() string
	GetVersion() string

	//ChangePassword(oldPassword string, newPassword string) bool
	Contains(scriptHash *helper.UInt160) bool

	CreateAccount() (IAccount, error)
	CreateAccountWithPrivateKey(privateKey []byte) (IAccount, error)
	CreateAccountWithContract(contract *sc.Contract, key *keys.KeyPair) (IAccount, error)
	CreateAccountWithScriptHash(scriptHash *helper.UInt160) (IAccount, error)

	DeleteAccount(scriptHash *helper.UInt160) bool

	GetAccountByPublicKey(pubKey *crypto.ECPoint) IAccount
	GetAccountByScriptHash(scriptHash *helper.UInt160) IAccount

	GetAccounts() []IAccount
}

func GetPrivateKeyFromNEP2(nep2 string, passphrase string, version byte, N, R, P int) ([]byte, error) {
	pair, err := keys.NewKeyPairFromNEP2(nep2, passphrase, version, N, R, P)
	if err != nil {
		return nil, err
	}
	return pair.PrivateKey, nil
}

func GetPrivateKeyFromWIF(wif string) ([]byte, error) {
	pair, err := keys.NewKeyPairFromWIF(wif)
	if err != nil {
		return nil, err
	}
	return pair.PrivateKey, nil
}

func getSigners(sender *helper.UInt160, cosigners []*tx.Signer) []*tx.Signer {
	for i := 0; i < len(cosigners); i++ {
		if cosigners[i].Account.Equals(sender) {
			if i == 0 {
				return cosigners
			}
			result := make([]*tx.Signer, len(cosigners))
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
	return append([]*tx.Signer{signer}, cosigners...)
}

func Sign(verifiable tx.IVerifiable, pair *keys.KeyPair, magic uint32) ([]byte, error) {
	return pair.Sign(tx.GetSignData(verifiable, magic))
}
