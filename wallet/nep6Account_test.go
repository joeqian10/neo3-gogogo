package wallet

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/sc"
	"math/big"
	"testing"

	"github.com/joeqian10/neo3-gogogo/keys"
	"github.com/stretchr/testify/assert"
)

var password = "Satoshi"
var privateKey = []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
	0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}
var pair, _ = keys.NewKeyPair(privateKey)
var wif = pair.Export()
var nep2, _ = pair.ExportWithPassword(password, helper.DefaultAddressVersion, 2, 1, 1)
var testContract, _ = sc.CreateSignatureContract(pair.PublicKey)
var testScriptHash = testContract.GetScriptHash()

var testWallet = &NEP6Wallet{
	protocolSettings: &helper.DefaultProtocolSettings,
	password:         nil,
	Name:             &dummy,
	path:             "",
	Version:          "3.0",
	accounts:         map[helper.UInt160]NEP6Account{},
	Scrypt:           NewScryptParameters(2, 1, 1),
	Extra:            nil,
}
var testAccount = NewNEP6Account(testWallet, testScriptHash, nil)

func TestNewNEP6Account(t *testing.T) {
	assert.Equal(t, 0, testAccount.scriptHash.CompareTo(testScriptHash))
	assert.Equal(t, true, testAccount.Decrypted())
	assert.Equal(t, false, testAccount.HasKey())
}

func TestNewNEP6AccountFromKeyPair(t *testing.T) {
	password := "hello world"
	account, err := NewNEP6AccountFromKeyPair(testWallet, testScriptHash, pair, password)
	assert.Nil(t, err)
	assert.Equal(t, 0, account.scriptHash.CompareTo(testScriptHash))
	assert.Equal(t, true, account.Decrypted())
	assert.Equal(t, true, account.HasKey())
}

func TestNEP6Account_GetKey(t *testing.T) {
	k, err := testAccount.GetKey()
	assert.Nil(t, err)
	assert.Nil(t, k)
	err = testWallet.Unlock(password)
	assert.Nil(t, err)
	account := NewNEP6Account(testWallet, testScriptHash, &nep2)
	k, err = account.GetKey()
	assert.Nil(t, err)
	assert.Equal(t, 0, k.CompareTo(pair))
}

func TestNEP6Account_GetKeyFromPassword(t *testing.T) {
	account := NewNEP6Account(testWallet, testScriptHash, &nep2)
	k, err := account.GetKeyFromPassword(password)
	assert.Nil(t, err)
	assert.Equal(t, 0, k.CompareTo(pair))
}

func TestAccountAndBalanceSlice_RemoveAt(t *testing.T) {
	a1, _ := helper.UInt160FromString("0000000000000000000000000000000000000001")
	a2, _ := helper.UInt160FromString("0000000000000000000000000000000000000002")
	a3, _ := helper.UInt160FromString("0000000000000000000000000000000000000003")
	us := []*AccountAndBalance{
		{
			Account: a1,
			Value:   big.NewInt(1),
		},
		{
			Account: a2,
			Value:   big.NewInt(2),
		},
		{
			Account: a3,
			Value:   big.NewInt(3),
		},
	}
	bs := AccountAndBalanceSlice(us)
	x := bs.RemoveAt(0)
	assert.Equal(t, 2, len(x))
	y := bs.RemoveAt(2)
	assert.Equal(t, 2, len(y))
}

func TestNEP6Account_VerifyPassword(t *testing.T) {
	account := NewNEP6Account(testWallet, testScriptHash, &nep2)
	assert.Equal(t, true, account.VerifyPassword(password))
	assert.Equal(t, false, account.VerifyPassword("b"))
}
