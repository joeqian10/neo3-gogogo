package tx

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSigner(t *testing.T) {
	signer := NewDefaultSigner()
	assert.NotNil(t, signer)
	assert.Equal(t, 0, signer.Account.CompareTo(helper.UInt160Zero))
	assert.Equal(t, None, signer.Scopes)
}

func TestSigner_CompareTo(t *testing.T) {
	signer1 := NewSigner(helper.UInt160Zero, Global)
	signer2 := NewSigner(helper.UInt160FromBytes([]byte{0x01}), Global)
	assert.Equal(t, -1, signer1.CompareTo(signer2))
}

func TestSigner_Deserialize(t *testing.T) {
	s := "ae716cd8bf248c38b601723ac4be2dc7979baeed" + "10" + "01" + "2b2e74b24ffc5599761499bbe67843abcb4413ad" + "00"
	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(s))
	assert.Nil(t, br.Err)
	cs := NewDefaultSigner()
	cs.Deserialize(br)
	assert.Equal(t, "edae9b97c72dbec43a7201b6388c24bfd86c71ae", cs.Account.String())
	assert.Equal(t, CustomContracts, cs.Scopes)
	assert.Equal(t, 1, len(cs.AllowedContracts))
}

func TestSigner_Serialize(t *testing.T) {
	account, _ := helper.UInt160FromString("edae9b97c72dbec43a7201b6388c24bfd86c71ae")
	contract, _ := helper.UInt160FromString("ad1344cbab4378e6bb9914769955fc4fb2742e2b")
	cs := Signer{
		Account:          account,
		Scopes:           CustomContracts,
		AllowedContracts: []helper.UInt160{*contract},
		AllowedGroups:    nil,
	}
	bbw := io.NewBufBinaryWriter()
	cs.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	assert.Equal(t, "ae716cd8bf248c38b601723ac4be2dc7979baeed"+"10"+"01"+"2b2e74b24ffc5599761499bbe67843abcb4413ad"+"00", helper.BytesToHex(b))
}

func TestSigner_Size(t *testing.T) {
	account, _ := helper.UInt160FromString("edae9b97c72dbec43a7201b6388c24bfd86c71ae")
	contract, _ := helper.UInt160FromString("ad1344cbab4378e6bb9914769955fc4fb2742e2b")
	cs := Signer{
		Account:          account,
		Scopes:           CustomContracts,
		AllowedContracts: []helper.UInt160{*contract},
		AllowedGroups:    nil,
	}
	size := cs.Size()
	assert.Equal(t, 20+1+1+20, size)
}

func TestSignerSlice_GetVarSize(t *testing.T) {
	account, _ := helper.UInt160FromString("edae9b97c72dbec43a7201b6388c24bfd86c71ae")
	contract, _ := helper.UInt160FromString("ad1344cbab4378e6bb9914769955fc4fb2742e2b")
	cs := Signer{
		Account:          account,
		Scopes:           CustomContracts,
		AllowedContracts: []helper.UInt160{*contract},
		AllowedGroups:    nil,
	}
	css := []Signer{cs}
	size := SignerSlice(css).GetVarSize()
	assert.Equal(t, 1+20+1+1+20, size)
}

func TestSigner_Deserialize2(t *testing.T) {
	signer := NewSigner(helper.NewUInt160(), Global)
	bw := io.NewBufBinaryWriter()
	signer.Serialize(bw.BinaryWriter)
	expected := "000000000000000000000000000000000000000080"
	assert.Equal(t, expected, helper.BytesToHex(bw.Bytes()))

	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(expected))
	signer1 := NewDefaultSigner()
	signer1.Deserialize(br)
	assert.Equal(t, 0, signer.CompareTo(signer1))
}
