package wallet

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joeqian10/neo3-gogogo/wallet/keys"
)

const (
	// The current version of neo-go wallet implementations.
	walletVersion = "3.0"
)

// Wallet represents a NEO (NEP-2, NEP-6) compliant wallet.
type Wallet struct {
	// string type
	Name string `json:"name"`

	// Version of the wallet, used for later upgrades.
	Version string `json:"version"`

	Scrypt *ScryptParams `json:"scrypt"`

	// A list of accounts which describes the details of each account
	// in the wallet.
	Accounts []*Account `json:"accounts"`

	// Extra metadata can be used for storing arbitrary data.
	// This field can be empty.
	Extra interface{} `json:"extra"`
}

// ScryptParams is a json-serializable container for scrypt KDF parameters.
type ScryptParams struct {
	N int `json:"n"`
	R int `json:"r"`
	P int `json:"p"`
}

// NewWallet creates a NEO wallet.
func NewWallet() *Wallet {
	return &Wallet{
		Version:  walletVersion,
		Accounts: []*Account{},
		Scrypt:   &ScryptParams{keys.N, keys.R, keys.P},
	}
}

// CreateAccount generates a new account for the end user and encrypts
// the private key with the given passphrase.
func (w *Wallet) AddNewAccount() error {
	acc, err := NewAccount()
	if err != nil {
		return err
	}
	w.AddAccount(acc)
	return nil
}

// Import account from WIF
func (w *Wallet) ImportFromWIF(wif string) error {
	acc, err := NewAccountFromWIF(wif)
	if err != nil {
		return err
	}
	w.AddAccount(acc)
	return nil
}

// Import account from Nep2Key
func (w *Wallet) ImportFromNep2Key(nep2Key, passphare string) error {
	acc, err := NewAccountFromNep2(nep2Key, passphare)
	if err != nil {
		return err
	}
	w.AddAccount(acc)
	return nil
}

// AddAccount adds an existing Account to the wallet if the account is not in wallet
func (w *Wallet) AddAccount(acc *Account) {
	for _, account := range w.Accounts {
		if account.Address == acc.Address {
			account = acc
			return
		}
	}
	w.Accounts = append(w.Accounts, acc)
}

// encrypt all the accounts in wallet, save the nep2Key
func (w *Wallet) EncryptAll(password string) error {
	for _, acc := range w.Accounts {
		if acc.KeyPair != nil {
			err := acc.Encrypt(password)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// decrypt all the accounts in wallet, save the nep2Key
func (w *Wallet) DecryptAll(password string) error {
	for _, acc := range w.Accounts {
		if acc.KeyPair == nil && acc.Nep2Key != "" {
			err := acc.Decrypt(password)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Save saves the wallet data. It's the internal io.ReadWriter
// that is responsible for saving the data. This can
// be a buffer, file, etc..
func (w *Wallet) Save(path string) error {
	for _, acc := range w.Accounts {
		if acc.KeyPair != nil && acc.Nep2Key == "" {
			return fmt.Errorf("please encrypt the accounts before save wallet")
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(w)
	if err != nil {
		return err
	}
	return file.Close()
}

// NewWalletFromFile creates a Wallet from the given wallet file path
func NewWalletFromFile(path string) (*Wallet, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	wall := &Wallet{}
	if err := json.NewDecoder(file).Decode(wall); err != nil {
		return nil, err
	}
	return wall, nil
}

// JSON outputs a pretty JSON representation of the wallet.
func (w *Wallet) JSON() ([]byte, error) {
	return json.Marshal(w)
}
