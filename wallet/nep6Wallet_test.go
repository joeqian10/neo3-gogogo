package wallet

import (
	"bytes"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/sc"
	"io/ioutil"
	"log"
	"testing"

	"github.com/joeqian10/neo3-gogogo/keys"
	"github.com/stretchr/testify/assert"
)

var hash = helper.UInt160FromBytes(crypto.Hash160([]byte{0x01}))
var testAccount1 = NewNEP6Account(testWallet, hash, nil)

func resetTestWallet()  {
	testWallet = &NEP6Wallet{
		protocolSettings: &helper.DefaultProtocolSettings,
		password: nil,
		Name:     &dummy,
		path:     "",
		Version:  "3.0",
		accounts: map[helper.UInt160]NEP6Account{},
		Scrypt:   NewScryptParameters(2, 1, 1),
		Extra:    nil,
	}
}

// test wallet from a file
func TestNewNEP6Wallet1(t *testing.T) {
	path := "test.json"
	wallet, err := NewNEP6Wallet(path, &helper.DefaultProtocolSettings, nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, 1,  len(wallet.Accounts))

	jsonBytes, err := wallet.JSON()
	assert.Nil(t, err)

	data, err := ioutil.ReadFile(path)
	assert.Nil(t, err)

	jsonString := string(jsonBytes)
	assert.Equal(t, string(data), jsonString)
}

// test wallet from params
func TestNewNEP6Wallet(t *testing.T) {
	s := "test"
	wallet, err := NewNEP6Wallet("", &helper.DefaultProtocolSettings, &s, DefaultScryptParameters)
	assert.NotNil(t, err)
	assert.Equal(t, "test", wallet.GetName())
	assert.Equal(t, "3.0", wallet.Version)
	assert.Equal(t, DefaultScryptParameters.P, wallet.Scrypt.P)
}

func TestGetPrivateKeyFromNEP2(t *testing.T) {
	pk, err := GetPrivateKeyFromNEP2("3vQB7B6MrGQZaxCuFg4oh", "TestGetPrivateKeyFromNEP2", helper.DefaultAddressVersion,2, 1, 1)
	assert.NotNil(t, err)
	pk, err = GetPrivateKeyFromNEP2(nep2, password,  helper.DefaultAddressVersion, 2, 1, 1)
	assert.Equal(t, true, bytes.Equal(privateKey, pk))
}

func TestGetPrivateKeyFromNEP22(t *testing.T) {
	priKey, err := GetPrivateKeyFromNEP2("6PYW3MNVdLsYt1SyMarsvU7pyBdokvSBsNMkhZ1tXW2Um3y9fRHdCNe7Bq", "qianzhuo", helper.DefaultAddressVersion, 16384, 8, 8)
	assert.Nil(t, err)
	log.Println(helper.BytesToHex(priKey))
}

func TestGetPrivateKeyFromWIF(t *testing.T) {
	pk, err := GetPrivateKeyFromWIF("3vQB7B6MrGQZaxCuFg4oh")
	assert.NotNil(t, err)

	pk, err = GetPrivateKeyFromWIF("L3tgppXLgdaeqSGSFw1Go3skBiy8vQAM7YMXvTHsKQtE16PBncSU")
	expected := []byte{199, 19, 77, 111, 216, 231, 61, 129, 158, 130, 117, 92, 100, 201, 55, 136, 216, 219, 9, 97, 146, 158, 2, 90, 83, 54, 60, 76, 192, 42, 105, 98}
	assert.Equal(t, helper.BytesToHex(expected), helper.BytesToHex(pk))
}

func TestNEP6Wallet_GetName(t *testing.T) {
	assert.Equal(t, "dummy", testWallet.GetName())
}

func TestNEP6Wallet_GetVersion(t *testing.T) {
	assert.Equal(t, "3.0", testWallet.GetVersion())
}

func TestNEP6Wallet_Contains(t *testing.T) {
	assert.Equal(t, false, testWallet.Contains(hash))
	testWallet.addAccount(testAccount1, false)
	assert.Equal(t, true, testWallet.Contains(hash))

	resetTestWallet()
}

func TestAddAccount(t *testing.T)  {
	_, err := testWallet.CreateAccountWithScriptHash(testScriptHash)
	assert.Nil(t, err)
	assert.Equal(t, true, testWallet.Contains(testScriptHash))
	acc := testWallet.GetAccountByScriptHash(testScriptHash)
	assert.Equal(t, true, acc.WatchOnly())
	assert.Equal(t, false, acc.HasKey())

	err = testWallet.Unlock(password)
	assert.Nil(t, err)
	_, err = testWallet.CreateAccountWithPrivateKey(privateKey)
	acc2 := testWallet.GetAccountByScriptHash(testScriptHash)
	assert.Equal(t, false, acc2.WatchOnly())
	assert.Equal(t, true, acc2.HasKey())

	_, err = testWallet.CreateAccountWithScriptHash(testScriptHash)
	assert.Nil(t, err)
	acc3 := testWallet.GetAccountByScriptHash(testScriptHash)
	assert.Equal(t, false, acc3.WatchOnly())
	assert.Equal(t, false, acc3.HasKey())

	_, err = testWallet.CreateAccountWithPrivateKey(privateKey)
	acc4 := testWallet.GetAccountByScriptHash(testScriptHash)
	assert.Equal(t, false, acc4.WatchOnly())
	assert.Equal(t, true, acc4.HasKey())

	resetTestWallet()
}

func TestNEP6Wallet_CreateAccount(t *testing.T) {
	assert.Equal(t, 0, len(testWallet.accounts))
	err := testWallet.Unlock(password)
	assert.Nil(t, err)
	_, err = testWallet.CreateAccount()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(testWallet.accounts))

	resetTestWallet()
}

func TestNEP6Wallet_CreateAccountWithPrivateKey(t *testing.T) {
	assert.Equal(t, false, testWallet.Contains(testScriptHash))
	err := testWallet.Unlock(password)
	assert.Nil(t, err)
	_, err = testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	assert.Equal(t, true, testWallet.Contains(testScriptHash))

	resetTestWallet()
}

func TestNEP6Wallet_CreateAccountWithContract(t *testing.T) {
	_, err := testWallet.CreateAccountWithContract(testContract, nil)
	assert.Nil(t, err)
	assert.Equal(t, true, testWallet.Contains(testScriptHash))
	testWallet.DeleteAccount(testScriptHash)
	assert.Equal(t, false, testWallet.Contains(testScriptHash))

	err = testWallet.Unlock(password)
	assert.Nil(t, err)
	_, err = testWallet.CreateAccountWithContract(testContract, pair)
	assert.Nil(t, err)
	assert.Equal(t, true, testWallet.Contains(testScriptHash))

	resetTestWallet()
}

func TestNEP6Wallet_CreateAccountWithScriptHash(t *testing.T) {
	assert.Equal(t, false, testWallet.Contains(testScriptHash))
	_, err := testWallet.CreateAccountWithScriptHash(testScriptHash)
	assert.Nil(t, err)
	assert.Equal(t, true, testWallet.Contains(testScriptHash))

	resetTestWallet()
}

func TestNEP6Wallet_DecryptKey(t *testing.T) {
	err := testWallet.Unlock(password)
	assert.Nil(t, err)
	k, err := testWallet.DecryptKey(nep2)
	assert.Nil(t, err)
	assert.Equal(t, 0, k.CompareTo(pair))

	resetTestWallet()
}

func TestNEP6Wallet_DeleteAccount(t *testing.T) {
	assert.Equal(t, false, testWallet.Contains(testScriptHash))
	_, err := testWallet.CreateAccountWithContract(testContract, nil)
	assert.Nil(t, err)
	assert.Equal(t, true, testWallet.Contains(testScriptHash))
	testWallet.DeleteAccount(testScriptHash)
	assert.Equal(t, false, testWallet.Contains(testScriptHash))

	resetTestWallet()
}

func TestNEP6Wallet_GetAccount(t *testing.T) {
	assert.Equal(t, false, testWallet.Contains(testScriptHash))
	err := testWallet.Unlock(password)
	assert.Nil(t, err)
	_, err = testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Equal(t, true, testWallet.Contains(testScriptHash))
	acc := testWallet.GetAccountByScriptHash(testScriptHash)
	c, err := sc.CreateSignatureContract(pair.PublicKey)
	assert.Nil(t, err)
	assert.Equal(t, crypto.ScriptHashToAddress(c.GetScriptHash(), helper.DefaultAddressVersion), acc.GetAddress())

	resetTestWallet()
}

func TestNEP6Wallet_GetAccounts(t *testing.T) {
	err := testWallet.Unlock(password)
	assert.Nil(t, err)

	privateKeys := make([][]byte, 5)
	m := make(map[string]keys.KeyPair, 5)
	for i, _ := range privateKeys {
		pk, err := helper.GenerateRandomBytes(32)
		assert.Nil(t, err)
		privateKeys[i] = pk
		p, _ := keys.NewKeyPair(pk)
		c, _ := sc.CreateSignatureContract(p.PublicKey)
		m[crypto.ScriptHashToAddress(c.GetScriptHash(), helper.DefaultAddressVersion)] = *p
		_, err = testWallet.CreateAccountWithPrivateKey(pk)
		assert.Nil(t, err)
	}

	accounts := testWallet.GetAccounts()
	for _, acc := range accounts {
		_, ok := m[acc.GetAddress()]
		assert.Equal(t, true, ok)
	}

	resetTestWallet()
}

func TestNEP6Wallet_ImportFromWIF(t *testing.T) {
	wif := pair.Export()
	assert.Equal(t, false, testWallet.Contains(testScriptHash))
	err := testWallet.Unlock(password)
	assert.Nil(t, err)
	_, err = testWallet.ImportFromWIF(wif)
	assert.Nil(t, err)
	assert.Equal(t, true, testWallet.Contains(testScriptHash))

	resetTestWallet()
}

func TestNEP6Wallet_ImportFromNEP2(t *testing.T) {
	assert.Equal(t, false, testWallet.Contains(testScriptHash))
	_, err := testWallet.ImportFromNEP2(nep2, password, 2, 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, true, testWallet.Contains(testScriptHash))

	resetTestWallet()
}

func TestNEP6Wallet_Lock(t *testing.T) {
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.NotNil(t, err)
	err = testWallet.Unlock(password)
	assert.Nil(t, err)
	_, err = testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	testWallet.DeleteAccount(testScriptHash)
	testWallet.Lock()
	_, err = testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.NotNil(t, err)

	resetTestWallet()
}

func TestNEP6Wallet_Save(t *testing.T) {
	path := "test.json"
	w, err := NewNEP6Wallet(path, &helper.DefaultProtocolSettings, nil, nil)
	assert.Nil(t, err)

	err = w.Save("testWrite.json")
	assert.Nil(t, err)
	testWrite, err := NewNEP6Wallet("testWrite.json", &helper.DefaultProtocolSettings, nil, nil)

	assert.Nil(t, err)
	assert.Equal(t, w.Name, testWrite.Name)
}

func TestNEP6Wallet_Unlock(t *testing.T) {
	_, err := testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.NotNil(t, err)
	err = testWallet.Unlock(password)
	assert.Nil(t, err)
	_, err = testWallet.CreateAccountWithPrivateKey(privateKey)
	assert.Nil(t, err)
	assert.Equal(t, true, testWallet.Contains(testScriptHash))
	testWallet.Lock()
	err = testWallet.Unlock("123456")
	assert.NotNil(t, err)

	resetTestWallet()
}

func TestNEP6Wallet_VerifyPassword(t *testing.T) {
	_, err := testWallet.ImportFromNEP2(nep2, password, 2, 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, true, testWallet.VerifyPassword(password))
	assert.Equal(t, false, testWallet.VerifyPassword("123456"))

	resetTestWallet()
}
