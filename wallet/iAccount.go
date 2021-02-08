package wallet

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/keys"
)

type IAccount interface {
	GetScriptHash() *helper.UInt160
	GetLabel() string
	SetLabel(label *string)
	GetIsDefault() bool
	SetIsDefault(isDefault bool)
	GetLock() bool
	SetLock(lock bool)
	GetContract() *NEP6Contract
	SetContract(c *NEP6Contract)

	GetAddress() string
	HasKey() bool
	WatchOnly() bool

	GetKey() (*keys.KeyPair, error)
}
