package wallet

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/keys"
	"github.com/joeqian10/neo3-gogogo/sc"
	"math/big"
)

// NEP6Account represents a NEO account. It holds the private and public key
// along with some metadata.
type NEP6Account struct {
	protocolSettings *helper.ProtocolSettings
	scriptHash *helper.UInt160
	Address    string        `json:"address"`
	Label      *string       `json:"label"`
	IsDefault  bool          `json:"isdefault"`
	Lock       bool          `json:"lock"`
	Nep2Key    *string       `json:"key"`
	nep2KeyNew *string
	Contract   *NEP6Contract `json:"contract"`
	wallet     *NEP6Wallet
	key        *keys.KeyPair
	Extra      interface{} `json:"extra"`
}

func NewNEP6Account(wallet *NEP6Wallet, scriptHash *helper.UInt160, nep2Key *string) *NEP6Account {
	return &NEP6Account{
		protocolSettings: wallet.protocolSettings,
		scriptHash: scriptHash,
		wallet:     wallet,
		Nep2Key:    nep2Key,
		Address:    crypto.ScriptHashToAddress(scriptHash, wallet.protocolSettings.AddressVersion), //
	}
}

func NewNEP6AccountFromKeyPair(wallet *NEP6Wallet, scriptHash *helper.UInt160, pair *keys.KeyPair, password string) (*NEP6Account, error) {
	nep2Key, err := pair.ExportWithPassword(password, wallet.protocolSettings.AddressVersion, wallet.Scrypt.N, wallet.Scrypt.R, wallet.Scrypt.P)
	if err != nil {
		return nil, err
	}
	a := NewNEP6Account(wallet, scriptHash, &nep2Key)
	a.key = pair
	return a, nil
}

func (a *NEP6Account) Decrypted() bool {
	return a.Nep2Key == nil || a.key != nil
}

func (a *NEP6Account) GetScriptHash() *helper.UInt160 {
	if a.scriptHash == nil {
		a.scriptHash = a.Contract.GetScriptHash()
	}
	return a.scriptHash
}

func (a *NEP6Account) GetLabel() string {
	if a.Label == nil {
		return ""
	}
	return *a.Label
}

func (a *NEP6Account) SetLabel(label *string) {
	a.Label = label
}

func (a *NEP6Account) GetIsDefault() bool {
	return a.IsDefault
}

func (a *NEP6Account) SetIsDefault(isDefault bool) {
	a.IsDefault = isDefault
}

func (a *NEP6Account) GetLock() bool {
	return a.Lock
}

func (a *NEP6Account) SetLock(lock bool) {
	a.Lock = lock
}

func (a *NEP6Account) GetContract() *NEP6Contract {
	return a.Contract
}

func (a *NEP6Account) SetContract(contract *NEP6Contract) {
	a.Contract = contract
}

func (a *NEP6Account) GetAddress() string {
	return a.Address
}

func (a *NEP6Account) HasKey() bool {
	return a.Nep2Key != nil
}

func (a *NEP6Account) WatchOnly() bool {
	return a.Contract == nil
}

func (a *NEP6Account) GetKey() (*keys.KeyPair, error) {
	if a.Nep2Key == nil {
		return nil, nil
	}
	if a.key == nil {
		var err error
		a.key, err = a.wallet.DecryptKey(*a.Nep2Key)
		if err != nil {
			return nil, err
		}
	}
	return a.key, nil
}

func (a *NEP6Account) GetKeyFromPassword(password string) (*keys.KeyPair, error) {
	if a.Nep2Key == nil {
		return nil, nil
	}
	if a.key == nil {
		priKey, err := GetPrivateKeyFromNEP2(*a.Nep2Key, password, a.protocolSettings.AddressVersion, a.wallet.Scrypt.N, a.wallet.Scrypt.R, a.wallet.Scrypt.P)
		if err != nil {
			return nil, err
		}
		a.key, err = keys.NewKeyPair(priKey)
		if err != nil {
			return nil, err
		}
	}
	return a.key, nil
}

// GenerateNEP6Account creates a new NEP6Account with a random generated PrivateKey.
func GenerateNEP6Account(password string) (*NEP6Account, error) {
	pair, err := keys.GenerateKeyPair()
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
	acc, err := NewNEP6AccountFromKeyPair(nil, contract.GetScriptHash(), pair, password)
	if err != nil {
		return nil, err
	}
	acc.Contract = contract
	return acc, nil
}

func (a *NEP6Account) VerifyPassword(password string) bool {
	_, err := GetPrivateKeyFromNEP2(*a.Nep2Key, password, a.protocolSettings.AddressVersion, a.wallet.Scrypt.N, a.wallet.Scrypt.R, a.wallet.Scrypt.P)
	if err != nil {
		return false
	}
	return true
}

type AccountAndBalance struct {
	Account *helper.UInt160
	Value   *big.Int
}

type AccountAndBalanceSlice []AccountAndBalance

func (us AccountAndBalanceSlice) Len() int {
	return len(us)
}

func (us AccountAndBalanceSlice) Less(i int, j int) bool {
	return us[i].Value.Cmp(us[j].Value) < 0
}

func (us AccountAndBalanceSlice) Swap(i, j int) {
	us[i], us[j] = us[j], us[i]
}

func (us AccountAndBalanceSlice) RemoveAt(index int) []AccountAndBalance {
	length := len(us)
	if index < 0 || index >= length {
		return us
	}
	tmp := append(us[:index], us[index+1:]...)
	return tmp
}

func findPayingAccounts(orderedAccounts []AccountAndBalance, amount *big.Int) []AccountAndBalance {
	result := make([]AccountAndBalance, 0)
	sum := big.NewInt(0)
	for _, ab := range orderedAccounts {
		sum = sum.Add(sum, ab.Value)
	}
	if sum.Cmp(amount) == 0 {
		return orderedAccounts
	} else {
		for i, _ := range orderedAccounts {
			if orderedAccounts[i].Value.Cmp(amount) < 0 {
				continue
			}
			if orderedAccounts[i].Value.Cmp(amount) == 0 {
				result = append(result, orderedAccounts[i])
				orderedAccounts = AccountAndBalanceSlice(orderedAccounts).RemoveAt(i)
			} else {
				result = append(result, AccountAndBalance{
					Account: orderedAccounts[i].Account,
					Value:   amount,
				})
				orderedAccounts[i] = AccountAndBalance{
					Account: orderedAccounts[i].Account,
					Value:   new(big.Int).Sub(orderedAccounts[i].Value, amount),
				}
			}
			break
		}
		if len(result) == 0 {
			i := len(orderedAccounts) - 1
			for orderedAccounts[i].Value.Cmp(amount) <= 0 {
				result = append(result, orderedAccounts[i])
				amount = amount.Sub(amount, orderedAccounts[i].Value)
				orderedAccounts = AccountAndBalanceSlice(orderedAccounts).RemoveAt(i)
				i--
			}
			for i = 0; i < len(orderedAccounts); i++ {
				if orderedAccounts[i].Value.Cmp(amount) < 0 {
					continue
				}
				if orderedAccounts[i].Value.Cmp(amount) < 0 {
					result = append(result, orderedAccounts[i])
					orderedAccounts = AccountAndBalanceSlice(orderedAccounts).RemoveAt(i)
				} else {
					result = append(result, AccountAndBalance{
						Account: orderedAccounts[i].Account,
						Value:   amount,
					})
					orderedAccounts[i] = AccountAndBalance{
						Account: orderedAccounts[i].Account,
						Value:   new(big.Int).Sub(orderedAccounts[i].Value, amount),
					}
				}
				break
			}
		}
	}
	return result
}
