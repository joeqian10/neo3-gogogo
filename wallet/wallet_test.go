package wallet

import (
	"io/ioutil"
	"testing"

	"github.com/joeqian10/neo3-gogogo/wallet/keys"
	"github.com/stretchr/testify/assert"
)

func TestNewWalletFromFile(t *testing.T) {
	path := "test.json"
	testWallet, err := NewWalletFromFile(path)
	assert.Nil(t, err)
	assert.Equal(t, len(testWallet.Accounts), 5)

	jsonBytes, err := testWallet.JSON()
	assert.Nil(t, err)

	data, err := ioutil.ReadFile(path)
	assert.Nil(t, err)

	jsonString := string(jsonBytes)
	assert.Equal(t, string(data), jsonString)

}

func TestWallet_Save(t *testing.T) {
	path := "test.json"
	testWallet, err := NewWalletFromFile(path)
	assert.Nil(t, err)

	err = testWallet.Save("testWrite.json")
	assert.Nil(t, err)
	testWrite, err := NewWalletFromFile("testWrite.json")

	assert.Nil(t, err)
	assert.Equal(t, testWallet, testWrite)
}

func TestNewWallet(t *testing.T) {
	testWallet := NewWallet()
	assert.Equal(t, 0, len(testWallet.Accounts))

	err := testWallet.AddNewAccount()
	assert.Nil(t, err)
	assert.Equal(t, len(testWallet.Accounts), 1)
	assert.True(t, testWallet.Accounts[0].Address != "")

	err = testWallet.ImportFromWIF(keys.KeyCases[0].Wif)
	assert.Nil(t, err)
	assert.Equal(t, len(testWallet.Accounts), 2)
	assert.Equal(t, testWallet.Accounts[1].Address, keys.KeyCases[0].Address)

	err = testWallet.ImportFromNEP2Key(keys.KeyCases[1].Nep2key, keys.KeyCases[1].Passphrase)
	assert.Nil(t, err)
	assert.Equal(t, len(testWallet.Accounts), 3)
	assert.Equal(t, testWallet.Accounts[2].Address, keys.KeyCases[1].Address)

	// add duplicate account, the length does not increase
	err = testWallet.ImportFromNEP2Key(keys.KeyCases[1].Nep2key, keys.KeyCases[1].Passphrase)
	assert.Nil(t, err)
	assert.Equal(t, len(testWallet.Accounts), 3)
}

func TestEncryptWallet(t *testing.T) {
	testWallet := NewWallet()
	assert.Equal(t, len(testWallet.Accounts), 0)

	err := testWallet.ImportFromWIF(keys.KeyCases[0].Wif)
	assert.Nil(t, err)
	assert.Equal(t, len(testWallet.Accounts), 1)
	assert.Equal(t, testWallet.Accounts[0].Address, keys.KeyCases[0].Address)

	path := "testWrite.json"
	err = testWallet.Save(path)
	assert.NotNil(t, err)
	assert.Equal(t, "please encrypt the accounts before save wallet", err.Error())

	password := "@#$%"
	err = testWallet.EncryptAll(password)
	err = testWallet.Save(path)
	assert.Nil(t, err)

	wallet2, err := NewWalletFromFile(path)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(wallet2.Accounts))
	assert.Equal(t, keys.KeyCases[0].Address, wallet2.Accounts[0].Address)

	_ = wallet2.DecryptAll(password)
	assert.Equal(t, keys.KeyCases[0].PrivateKey, wallet2.Accounts[0].KeyPair.String())
}
