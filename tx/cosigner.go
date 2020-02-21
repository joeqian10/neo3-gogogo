package tx

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"github.com/joeqian10/neo3-gogogo/wallet/keys"
)

const (
	MaxSubitems = 16 // This limits maximum number of AllowedContracts or AllowedGroups here.
)

type Cosigner struct {
	Account          helper.UInt160
	Scopes           WitnessScope
	AllowedContracts []helper.UInt160
	AllowedGroups    []*keys.PublicKey
}

func NewCosigner(account helper.UInt160) *Cosigner {
	return &Cosigner{
		Account:account,
	}
}

func (c *Cosigner) Size() int {
	size := 20 + 1 // Account + Scopes
	if c.Scopes&CustomContracts != 0 {
		size += helper.UInt160Slice(c.AllowedContracts).GetVarSize()
	}
	if c.Scopes&CustomGroups != 0 {
		size += keys.PublicKeySlice(c.AllowedGroups).GetVarSize()
	}
	return size
}

// Deserialize implements Serializable interface.
func (c *Cosigner) Deserialize(br *io.BinaryReader) {
	br.ReadLE(&c.Account)
	br.ReadLE(&c.Scopes)
	if c.Scopes&CustomContracts != 0 {
		length := br.ReadVarUInt(uint64(MaxSubitems))
		c.AllowedContracts = make([]helper.UInt160, length)
		for i := 0; i < int(length); i++ {
			c.AllowedContracts[i] = helper.UInt160{}
			br.ReadLE(&c.AllowedContracts[i])
		}
	} else {
		c.AllowedContracts = []helper.UInt160{}
	}
	if c.Scopes&CustomGroups != 0 {
		length := br.ReadVarUInt(uint64(MaxSubitems))
		c.AllowedGroups = make([]*keys.PublicKey, length)
		for i := 0; i < int(length); i++ {
			c.AllowedGroups[i] = &keys.PublicKey{}
			c.AllowedGroups[i].Deserialize(br)
		}
	}
}

func (c *Cosigner) Serialize(bw *io.BinaryWriter) {
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

type CosignerSlice []*Cosigner

func (cs CosignerSlice) GetVarSize() int {
	size := 0
	for _, c := range cs {
		size += c.Size()
	}
	return helper.GetVarSize(len(cs)) + size
}
