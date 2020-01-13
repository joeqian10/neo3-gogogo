package wallet

import (
	"github.com/joeqian10/neo3-gogogo/wallet/keys"
)

// Account represents a NEO account. It holds the private and public key
// along with some metadata.
type Account struct {
	// NEO  KeyPair.
	KeyPair *keys.KeyPair `json:"-"`

	// NEO public address.
	Address string `json:"address"`

	// Label is a label the user had made for this account.
	Label string `json:"label"`

	// Indicates whether the account is the default change account.
	Default bool `json:"isDefault"`

	// Indicates whether the account is locked by the user.
	// the client shouldn't spend the funds in a locked account.
	Locked bool `json:"lock"`

	// Encrypted Key of the account also known as the key.
	Nep2Key string `json:"key"`

	// contract is a Contract object which describes the details of the contract.
	// This field can be null (for watch-only address).
	Contract *Contract `json:"contract"`

	// This field can be empty.
	Extra interface{} `json:"extra"`
}

// Contract represents a subset of the smart contract to embed in the
// Account so it's NEP-6 compliant.
type Contract struct {
	// Script hash of the contract deployed on the block chain.
	Script string `json:"script"`

	// A list of parameters used deploying this contract.
	Parameters []interface{} `json:"parameters"`

	// Indicates whether the contract has been deployed to the block chain.
	Deployed bool `json:"deployed"`
}

// NewAccountFromKeyPair created a wallet from the given PrivateKey.
func NewAccountFromKeyPair(p *keys.KeyPair) *Account {
	pubAddr := p.PublicKey.Address()
	a := &Account{
		KeyPair: p,
		Address: pubAddr,
	}
	return a
}

// NewAccount creates a new Account with a random generated PrivateKey.
func NewAccount() (*Account, error) {
	privateKey, err := keys.GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	return NewAccountFromKeyPair(privateKey), nil
}

// NewAccountFromWIF creates a new Account from the given WIF.
func NewAccountFromWIF(wif string) (*Account, error) {
	privateKey, err := keys.NewKeyPairFromWIF(wif)
	if err != nil {
		return nil, err
	}
	return NewAccountFromKeyPair(privateKey), nil
}

// NewAccountFromNep2 decrypts the nep2Key with the given passphrase and
// returns the decrypted Account.
func NewAccountFromNep2(nep2Key, passphrase string) (*Account, error) {
	wif, err := keys.NEP2Decrypt(nep2Key, passphrase)
	if err != nil {
		return nil, err
	}
	return NewAccountFromKeyPair(wif), nil
}

// Encrypt encrypts the wallet's PrivateKey with the given passphrase under the NEP-2 standard.
func (a *Account) Encrypt(passphrase string) (err error) {
	if a.Nep2Key, err = keys.NEP2Encrypt(a.KeyPair, passphrase); err != nil {
		return err
	}

	if a.Address == "" {
		a.Address = a.KeyPair.PublicKey.Address()
	}
	return nil
}

// Decrypt encrypts the wallet's PrivateKey with the given passphrase under the NEP-2 standard.
func (a *Account) Decrypt(passphrase string) (err error) {
	if a.KeyPair == nil {
		a.KeyPair, err = keys.NEP2Decrypt(a.Nep2Key, passphrase)
	}

	if a.Address == "" {
		a.Address = a.KeyPair.PublicKey.Address()
	}
	return err
}
