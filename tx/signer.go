package tx

import (
	"fmt"
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
	AllowedContracts []*helper.UInt160
	AllowedGroups    []*crypto.ECPoint
	Rules            []*WitnessRule
}

func NewSigner(account *helper.UInt160, scopes WitnessScope) *Signer {
	return &Signer{
		Account: account,
		Scopes:  scopes,
	}
}

func NewDefaultSigner() *Signer {
	return NewSigner(helper.NewUInt160(), None)
}

func (c *Signer) GetSize() int {
	size := 20 + 1 // Account + Scopes
	if c.Scopes&CustomContracts != 0 {
		size += helper.UInt160Slice(c.AllowedContracts).GetVarSize()
	}
	if c.Scopes&CustomGroups != 0 {
		size += crypto.PublicKeySlice(c.AllowedGroups).GetVarSize()
	}
	if c.Scopes&WitnessRules != 0 {

	}
	return size
}

func (c *Signer) CompareTo(d *Signer) int {
	r := c.Account.CompareTo(d.Account)
	if r != 0 {
		return r
	}
	r = c.Scopes.CompareTo(d.Scopes)
	return r
}

// Deserialize implements Serializable interface.
func (c *Signer) Deserialize(br *io.BinaryReader) {
	br.ReadLE(c.Account)
	br.ReadLE(&c.Scopes)
	if (c.Scopes & ^(CalledByEntry | CustomContracts | CustomGroups | WitnessRules | Global)) != 0 {
		br.Err = fmt.Errorf("invalid witness scopes: %s", c.Scopes.String())
		return
	}
	if c.Scopes&Global != 0 && c.Scopes != Global {
		br.Err = fmt.Errorf("invalid witness scopes: %s", c.Scopes.String())
		return
	}
	if c.Scopes&CustomContracts != 0 {
		length := br.ReadVarUIntWithMaxLimit(uint64(MaxSubitems))
		c.AllowedContracts = make([]*helper.UInt160, length)
		for i := 0; i < int(length); i++ {
			c.AllowedContracts[i] = helper.NewUInt160()
			br.ReadLE(c.AllowedContracts[i])
		}
	} else {
		c.AllowedContracts = []*helper.UInt160{}
	}
	if c.Scopes&CustomGroups != 0 {
		length := br.ReadVarUIntWithMaxLimit(uint64(MaxSubitems))
		c.AllowedGroups = make([]*crypto.ECPoint, length)
		for i := 0; i < int(length); i++ {
			c.AllowedGroups[i], _ = crypto.NewECPoint()
			c.AllowedGroups[i].Deserialize(br)
		}
	} else {
		c.AllowedGroups = []*crypto.ECPoint{}
	}
	if c.Scopes&WitnessRules != 0 {
		length := br.ReadVarUIntWithMaxLimit(uint64(MaxSubitems))
		c.Rules = make([]*WitnessRule, length)
		for i := 0; i < int(length); i++ {
			c.Rules[i] = &WitnessRule{}
			c.Rules[i].Deserialize(br)
		}
	} else {
		c.Rules = []*WitnessRule{}
	}
}

func (c *Signer) Serialize(bw *io.BinaryWriter) {
	bw.WriteLE(c.Account)
	bw.WriteLE(c.Scopes)
	if c.Scopes&CustomContracts != 0 {
		bw.WriteVarUInt(uint64(len(c.AllowedContracts)))
		for _, ac := range c.AllowedContracts {
			ac.Serialize(bw)
		}
	}
	if c.Scopes&CustomGroups != 0 {
		bw.WriteVarUInt(uint64(len(c.AllowedGroups)))
		for _, ag := range c.AllowedGroups {
			ag.Serialize(bw)
		}
	}
	if c.Scopes&WitnessRules != 0 {
		bw.WriteVarUInt(uint64(len(c.Rules)))
		for _, r := range c.Rules {
			r.Serialize(bw)
		}
	}
}

type SignerSlice []*Signer

func (cs SignerSlice) GetVarSize() int {
	size := 0
	for _, c := range cs {
		size += c.GetSize()
	}
	return helper.GetVarSize(len(cs)) + size
}
