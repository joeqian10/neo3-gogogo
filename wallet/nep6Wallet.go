package wallet

import (
	"encoding/json"
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/sc"
	"os"
	"strconv"

	"github.com/joeqian10/neo3-gogogo/keys"
)

const (
	// The current version of wallet implementations.
	Neo3WalletVersion = "3.0"
)

// NEP6Wallet represents a NEO (NEP-2, NEP-6) compliant wallet.
type NEP6Wallet struct {
	protocolSettings *helper.ProtocolSettings
	password *string
	path     string
	Name     *string           `json:"name"`
	Version  string            `json:"version"`
	Scrypt   *ScryptParameters `json:"scrypt"` // readonly
	Accounts []NEP6Account     `json:"accounts"`
	Extra    interface{}       `json:"extra"`

	accounts map[helper.UInt160]NEP6Account
}

// NewNEP6Wallet creates a NEO wallet.
func NewNEP6Wallet(path string, settings *helper.ProtocolSettings, name *string, scrypt *ScryptParameters) (*NEP6Wallet, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeExclusive)
	if err != nil {
		return &NEP6Wallet{
			protocolSettings: settings,
			path: path,
			Name:     name,
			Version:  Neo3WalletVersion,
			accounts: make(map[helper.UInt160]NEP6Account, 0),
			Scrypt:   scrypt,
			Extra:    nil,
		}, err
	}
	w := &NEP6Wallet{
		protocolSettings: settings,
		path: path,
		Scrypt: &ScryptParameters{},
		accounts: make(map[helper.UInt160]NEP6Account, 0),
	}
	if err := json.NewDecoder(file).Decode(w); err != nil {
		return nil, err
	}
	if w.Accounts != nil {
		w.accounts = make(map[helper.UInt160]NEP6Account, len(w.Accounts))
		for i, _ := range w.Accounts {
			w.Accounts[i].wallet = w
			w.Accounts[i].protocolSettings = settings
			w.accounts[*w.Accounts[i].GetScriptHash()] = w.Accounts[i]
		}
	}
	return w, nil
}

func (w *NEP6Wallet) GetName() string {
	if w.Name == nil {
		return ""
	}
	return *w.Name
}

func (w *NEP6Wallet) GetPath() string {
	return w.path
}

func (w *NEP6Wallet) GetVersion() string {
	return w.Version
}

func (w *NEP6Wallet) addAccount(acc *NEP6Account, isImport bool) {
	if accountOld, ok := w.accounts[*acc.GetScriptHash()]; ok {
		acc.Label = accountOld.Label
		acc.IsDefault = accountOld.IsDefault
		acc.Lock = accountOld.Lock
		if acc.Contract == nil {
			acc.Contract = accountOld.Contract
		} else {
			if accountOld.Contract != nil {
				acc.Contract.parameterNames = accountOld.Contract.parameterNames
				acc.Contract.Deployed = accountOld.Contract.Deployed
			}
		}
		acc.Extra = accountOld.Extra
	}
	w.accounts[*acc.scriptHash] = *acc
}

func (w *NEP6Wallet) Contains(scriptHash *helper.UInt160) bool {
	if _, ok := w.accounts[*scriptHash]; ok {
		return true
	}
	return false
}

// CreateAccount generates a new account for the end user and encrypts
// the private key with the password in the wallet.
func (w *NEP6Wallet) CreateAccount() (IAccount, error) {
	privateKey, err := helper.GenerateRandomBytes(32)
	if err != nil {
		return nil, err
	}
	return w.CreateAccountWithPrivateKey(privateKey)
}

func (w *NEP6Wallet) CreateAccountWithPrivateKey(privateKey []byte) (IAccount, error) {
	pair, err := keys.NewKeyPair(privateKey)
	if err != nil {
		return nil, err
	}
	script, err := sc.CreateSignatureRedeemScript(pair.PublicKey)
	if err != nil {
		return nil, err
	}
	contract, err := NewNEP6Contract(script, []sc.ContractParameterType{sc.Signature}, []string{"signature"}, false)
	if err != nil {
		return nil, err
	}
	if w.password == nil {
		return nil, fmt.Errorf("unlock wallet first")
	}
	acc, err := NewNEP6AccountFromKeyPair(w, contract.GetScriptHash(), pair, *w.password)
	if err != nil {
		return nil, err
	}
	acc.Contract = contract
	w.addAccount(acc, false)
	return acc, nil
}

func (w *NEP6Wallet) CreateAccountWithContract(contract *sc.Contract, pair *keys.KeyPair) (IAccount, error) {
	if contract == nil {
		return nil, fmt.Errorf("contract is nil")
	}
	paramNames := make([]string, len(contract.ParameterList))
	for i := 0; i < len(contract.ParameterList); i++ {
		paramNames[i] = "parameter" + strconv.Itoa(i)
	}
	nep6contract, err := NewNEP6Contract(contract.Script, contract.ParameterList, paramNames, false)
	if err != nil {
		return nil, err
	}
	var acc *NEP6Account
	if pair == nil {
		acc = NewNEP6Account(w, nep6contract.GetScriptHash(), nil)
	} else {
		if w.password == nil {
			return nil, fmt.Errorf("unlock wallet first")
		}
		acc, err = NewNEP6AccountFromKeyPair(w, nep6contract.GetScriptHash(), pair, *w.password)
		if err != nil {
			return nil, err
		}
	}
	acc.Contract = nep6contract
	w.addAccount(acc, false)
	return acc, nil
}

func (w *NEP6Wallet) CreateAccountWithScriptHash(scriptHash *helper.UInt160) (IAccount, error) {
	acc := NewNEP6Account(w, scriptHash, nil)
	w.addAccount(acc, false)
	return acc, nil
}

func (w *NEP6Wallet) DecryptKey(nep2Key string) (*keys.KeyPair, error) {
	priKey, err := GetPrivateKeyFromNEP2(nep2Key, *w.password, w.protocolSettings.AddressVersion, w.Scrypt.N, w.Scrypt.R, w.Scrypt.P)
	if err != nil {
		return nil, err
	}
	pair, err := keys.NewKeyPair(priKey)
	if err != nil {
		return nil, err
	}
	return pair, nil
}

func (w *NEP6Wallet) DeleteAccount(scriptHash *helper.UInt160) bool {
	if _, ok := w.accounts[*scriptHash]; ok {
		delete(w.accounts, *scriptHash)
		return true
	}
	return false
}

func (w *NEP6Wallet) GetAccountByPublicKey(pubKey *crypto.ECPoint) IAccount {
	contract, _ := sc.CreateSignatureContract(pubKey)
	return w.GetAccountByScriptHash(contract.GetScriptHash())
}

func (w *NEP6Wallet) GetAccountByScriptHash(scriptHash *helper.UInt160) IAccount {
	if acc, ok := w.accounts[*scriptHash]; ok {
		return &acc
	}
	return nil
}

func (w *NEP6Wallet) GetAccounts() []IAccount {
	accounts :=  []IAccount{}
	for _, v := range w.accounts {
		accounts = append(accounts, IAccount(&v))
	}
	return accounts
}

// Import account from WIF
func (w *NEP6Wallet) ImportFromWIF(wif string) (IAccount, error) {
	pair, err := keys.NewKeyPairFromWIF(wif)
	if err != nil {
		return nil, err
	}
	script, err := sc.CreateSignatureRedeemScript(pair.PublicKey)
	if err != nil {
		return nil, err
	}
	contract, err := NewNEP6Contract(script, []sc.ContractParameterType{sc.Signature}, []string{"signature"}, false)
	if err != nil {
		return nil, err
	}
	acc, err := NewNEP6AccountFromKeyPair(w, contract.GetScriptHash(), pair, *w.password)
	if err != nil {
		return nil, err
	}
	acc.Contract = contract
	w.addAccount(acc, true)
	return acc, nil
}

// Import account from Nep2Key
func (w *NEP6Wallet) ImportFromNEP2(nep2, passphrase string, N, R, P int) (IAccount, error) {
	pair, err := keys.NewKeyPairFromNEP2(nep2, passphrase, w.protocolSettings.AddressVersion, N, R, P)
	if err != nil {
		return nil, err
	}
	script, err := sc.CreateSignatureRedeemScript(pair.PublicKey)
	if err != nil {
		return nil, err
	}
	contract, err := NewNEP6Contract(script, []sc.ContractParameterType{sc.Signature}, []string{"signature"}, false)
	if err != nil {
		return nil, err
	}
	var acc *NEP6Account
	if w.Scrypt.N == 16384 && w.Scrypt.R == 8 && w.Scrypt.P == 8 {
		acc = NewNEP6Account(w, contract.GetScriptHash(), &nep2)
	} else {
		acc, err = NewNEP6AccountFromKeyPair(w, contract.GetScriptHash(), pair, passphrase)
		if err != nil {
			return nil, err
		}
	}
	acc.Contract = contract
	w.addAccount(acc, true)
	return acc, nil
}

func (w *NEP6Wallet) Lock() {
	w.password = nil
}

func (w *NEP6Wallet) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	if w.accounts != nil {
		w.Accounts = make([]NEP6Account, 0)
		for _, v := range w.accounts {
			w.Accounts = append(w.Accounts, v)
		}
	}
	err = json.NewEncoder(file).Encode(w)
	if err != nil {
		return err
	}
	return file.Close()
}

func (w *NEP6Wallet) Unlock(password string) error {
	if !w.VerifyPassword(password) {
		return fmt.Errorf("invalid password")
	}
	w.password = &password
	return nil
}

func (w *NEP6Wallet) VerifyPassword(password string) bool {
	var account *NEP6Account
	for _, acc := range w.accounts {
		if !acc.Decrypted() {
			account = &acc
			break
		}
	}
	if account == nil {
		for _, acc := range w.accounts {
			if acc.HasKey() {
				account = &acc
				break
			}
		}
	}
	if account == nil {
		return true
	}
	if account.Decrypted() {
		return account.VerifyPassword(password)
	} else {
		_, err := account.GetKeyFromPassword(password)
		if err != nil {
			return false
		}
		return true
	}
}

// JSON outputs a pretty JSON representation of the wallet.
func (w *NEP6Wallet) JSON() ([]byte, error) {
	return json.Marshal(w)
}
