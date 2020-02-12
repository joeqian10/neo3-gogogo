package tx

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCosigner_Deserialize(t *testing.T) {
	s := "ae716cd8bf248c38b601723ac4be2dc7979baeed" + "10" + "01" + "2b2e74b24ffc5599761499bbe67843abcb4413ad" + "00"
	br := io.NewBinaryReaderFromBuf(helper.HexTobytes(s))
	cs := Cosigner{}
	cs.Deserialize(br)
	assert.Equal(t, "edae9b97c72dbec43a7201b6388c24bfd86c71ae", cs.Account.String())
	assert.Equal(t, CustomContracts, cs.Scopes)
	assert.Equal(t, 1, len(cs.AllowedContracts))
}

func TestCosigner_Serialize(t *testing.T) {
	account, _ := helper.UInt160FromString("edae9b97c72dbec43a7201b6388c24bfd86c71ae")
	contract, _ := helper.UInt160FromString("ad1344cbab4378e6bb9914769955fc4fb2742e2b")
	cs := Cosigner{
		Account:          account,
		Scopes:           CustomContracts,
		AllowedContracts: []helper.UInt160{contract},
		AllowedGroups:    nil,
	}
	bbw := io.NewBufBinaryWriter()
	cs.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	assert.Equal(t, "ae716cd8bf248c38b601723ac4be2dc7979baeed"+"10"+"01"+"2b2e74b24ffc5599761499bbe67843abcb4413ad"+"00", helper.BytesToHex(b))
}

func TestCosigner_Size(t *testing.T) {
	account, _ := helper.UInt160FromString("edae9b97c72dbec43a7201b6388c24bfd86c71ae")
	contract, _ := helper.UInt160FromString("ad1344cbab4378e6bb9914769955fc4fb2742e2b")
	cs := Cosigner{
		Account:          account,
		Scopes:           CustomContracts,
		AllowedContracts: []helper.UInt160{contract},
		AllowedGroups:    nil,
	}
	size := cs.Size()
	assert.Equal(t, 20+1+1+20, size)
}

func TestCosignerSlice_GetVarSize(t *testing.T) {
	account, _ := helper.UInt160FromString("edae9b97c72dbec43a7201b6388c24bfd86c71ae")
	contract, _ := helper.UInt160FromString("ad1344cbab4378e6bb9914769955fc4fb2742e2b")
	cs := Cosigner{
		Account:          account,
		Scopes:           CustomContracts,
		AllowedContracts: []helper.UInt160{contract},
		AllowedGroups:    nil,
	}
	css := []*Cosigner{&cs}
	size := CosignerSlice(css).GetVarSize()
	assert.Equal(t, 1+20+1+1+20, size)
}
