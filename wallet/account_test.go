package wallet

import (
	"testing"

	"github.com/joeqian10/neo3-gogogo/wallet/keys"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	for _, testCase := range keys.KeyCases {
		acc, err := NewAccountFromWIF(testCase.Wif)
		if err != nil {
			t.Fatal(err)
		}
		compareFields(t, testCase, acc)
	}
}

func TestDecryptAccount(t *testing.T) {
	for _, testCase := range keys.KeyCases {
		acc, err := NewAccountFromNEP2(testCase.Nep2key, testCase.Passphrase)
		if err != nil {
			t.Fatal(err)
		}
		compareFields(t, testCase, acc)
	}
}

func TestNewFromWif(t *testing.T) {
	for _, testCase := range keys.KeyCases {
		acc, err := NewAccountFromWIF(testCase.Wif)
		if err != nil {
			t.Fatal(err)
		}
		compareFields(t, testCase, acc)
	}
}

func TestEncryptAccount(t *testing.T) {
	acc, err := NewAccount()
	if err != nil {
		t.Fatal(err)
	}
	password := "AAA123456789"
	err = acc.Encrypt(password)
	if err != nil {
		t.Fatal(err)
	}

	nep2Key := acc.Nep2Key
	account := Account{Nep2Key: nep2Key}
	err = account.Decrypt(password)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, acc.KeyPair, account.KeyPair)
}

func compareFields(t *testing.T, tk keys.Ktype, acc *Account) {
	if want, have := tk.Address, acc.Address; want != have {
		t.Fatalf("expected %s got %s", want, have)
	}
	if want, have := tk.Wif, acc.KeyPair.ExportWIF(); want != have {
		t.Fatalf("expected %s got %s", want, have)
	}
	if want, have := tk.PublicKey, acc.KeyPair.PublicKey.String(); want != have {
		t.Fatalf("expected %s got %s", want, have)
	}
	if want, have := tk.PrivateKey, acc.KeyPair.String(); want != have {
		t.Fatalf("expected %s got %s", want, have)
	}
}
