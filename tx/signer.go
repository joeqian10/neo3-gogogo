package tx

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
)

const (
	MaxSubitems = 16 // This limits maximum number of AllowedContracts or AllowedGroups here.
)

type Signer struct {
	Account          *helper.UInt160
	Scopes           WitnessScope
	AllowedContracts []helper.UInt160
	AllowedGroups    []crypto.ECPoint
}

func NewSigner(account *helper.UInt160, scopes WitnessScope) *Signer {
	return &Signer{
		Account: account,
		Scopes: scopes,
	}
}

func NewDefaultSigner() *Signer {
	return NewSigner(helper.NewUInt160(), None)
}

func (c *Signer) Size() int {
	size := 20 + 1 // Account + Scopes
	if c.Scopes&CustomContracts != 0 {
		size += helper.UInt160Slice(c.AllowedContracts).GetVarSize()
	}
	if c.Scopes&CustomGroups != 0 {
		size += crypto.PublicKeySlice(c.AllowedGroups).GetVarSize()
	}
	return size
}

func (c *Signer) CompareTo(d *Signer) int {
	r := c.Account.CompareTo(d.Account)
	if r != 0 {return r}
	r = c.Scopes.CompareTo(d.Scopes)
	return r
}

// Deserialize implements Serializable interface.
func (c *Signer) Deserialize(br *io.BinaryReader) {
	br.ReadLE(c.Account)
	br.ReadLE(&c.Scopes)
	if c.Scopes&CustomContracts != 0 {
		length := br.ReadVarUIntWithMaxLimit(uint64(MaxSubitems))
		c.AllowedContracts = make([]helper.UInt160, length)
		for i := 0; i < int(length); i++ {
			c.AllowedContracts[i] = helper.UInt160{}
			br.ReadLE(&c.AllowedContracts[i])
		}
	} else {
		c.AllowedContracts = []helper.UInt160{}
	}
	if c.Scopes&CustomGroups != 0 {
		length := br.ReadVarUIntWithMaxLimit(uint64(MaxSubitems))
		c.AllowedGroups = make([]crypto.ECPoint, length)
		for i := 0; i < int(length); i++ {
			c.AllowedGroups[i] = crypto.ECPoint{}
			c.AllowedGroups[i].Deserialize(br)
		}
	}
}

func (c *Signer) Serialize(bw *io.BinaryWriter) {
	bw.WriteLE(c.Account)
	bw.WriteLE(c.Scopes)
	if c.Scopes&CustomContracts != 0 {
		bw.WriteVarUInt(uint64(len(c.AllowedContracts)))
		for _, ac := range c.AllowedContracts {
			bw.WriteLE(ac)
		}
	}
	if c.Scopes&CustomContracts != 0 {
		bw.WriteVarUInt(uint64(len(c.AllowedGroups)))
		for _, ag := range c.AllowedGroups {
			ag.Serialize(bw)
		}
	}
}

type SignerSlice []Signer

func (cs SignerSlice) GetVarSize() int {
	size := 0
	for _, c := range cs {
		size += c.Size()
	}
	return helper.GetVarSize(len(cs)) + size
}
