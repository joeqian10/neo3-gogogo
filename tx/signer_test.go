package tx

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/joeqian10/neo3-gogogo/tx/conditions"
	"github.com/stretchr/testify/assert"
	"testing"
)

// todo

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

func TestSigner_Serialize_Deserialize_Global(t *testing.T) {
	hex := "000000000000000000000000000000000000000080"
	s := &Signer{
		Account: helper.UInt160Zero,
		Scopes:  Global,
	}
	bbw := io.NewBufBinaryWriter()
	s.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	assert.Equal(t, hex, helper.BytesToHex(b))

	s2 := NewDefaultSigner()
	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(hex))
	assert.Nil(t, br.Err)
	s2.Deserialize(br)
	assert.Nil(t, br.Err)
	assert.Equal(t, s.Account, s2.Account)
	assert.Equal(t, s.Scopes, s2.Scopes)
}

func TestSigner_Serialize_Deserialize_CalledByEntry(t *testing.T) {
	hex := "000000000000000000000000000000000000000001"
	s := &Signer{
		Account: helper.UInt160Zero,
		Scopes:  CalledByEntry,
	}
	bbw := io.NewBufBinaryWriter()
	s.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	assert.Equal(t, hex, helper.BytesToHex(b))

	s2 := NewDefaultSigner()
	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(hex))
	assert.Nil(t, br.Err)
	s2.Deserialize(br)
	assert.Nil(t, br.Err)
	assert.Equal(t, s.Account, s2.Account)
	assert.Equal(t, s.Scopes, s2.Scopes)
}

func TestSigner_Serialize_Deserialize_MaxNested_And(t *testing.T) {
	b := true
	s := &Signer{
		Account: helper.UInt160Zero,
		Scopes:  WitnessRules,
		Rules: []*WitnessRule{
			{
				Action: Allow,
				Condition: conditions.NewWitnessCondition(
					conditions.And,
					[]*conditions.WitnessCondition{
						conditions.NewWitnessCondition(
							conditions.And,
							[]*conditions.WitnessCondition{
								conditions.NewWitnessCondition(
									conditions.And,
									[]*conditions.WitnessCondition{
										conditions.NewWitnessCondition(conditions.Boolean, &b),
									}),
							}),
					}),
			},
		},
	}
	hex := "00000000000000000000000000000000000000004001010201020102010001"
	bbw := io.NewBufBinaryWriter()
	s.Serialize(bbw.BinaryWriter)
	bs := bbw.Bytes()

	assert.Equal(t, hex, helper.BytesToHex(bs))

	s2 := NewDefaultSigner()
	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(hex))
	assert.Nil(t, br.Err)
	s2.Deserialize(br)
	assert.NotNil(t, br.Err)
}

func TestSigner_Serialize_Deserialize_MaxNested_Or(t *testing.T) {
	b := true
	s := &Signer{
		Account: helper.UInt160Zero,
		Scopes:  WitnessRules,
		Rules: []*WitnessRule{
			{
				Action: Allow,
				Condition: conditions.NewWitnessCondition(
					conditions.Or,
					[]*conditions.WitnessCondition{
						conditions.NewWitnessCondition(
							conditions.Or,
							[]*conditions.WitnessCondition{
								conditions.NewWitnessCondition(
									conditions.Or,
									[]*conditions.WitnessCondition{
										conditions.NewWitnessCondition(conditions.Boolean, &b),
									}),
							}),
					}),
			},
		},
	}
	hex := "00000000000000000000000000000000000000004001010301030103010001"
	bbw := io.NewBufBinaryWriter()
	s.Serialize(bbw.BinaryWriter)
	bs := bbw.Bytes()

	assert.Equal(t, hex, helper.BytesToHex(bs))

	s2 := NewDefaultSigner()
	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(hex))
	assert.Nil(t, br.Err)
	s2.Deserialize(br)
	assert.NotNil(t, br.Err)
}

func TestSigner_Serialize_Deserialize_CustomContracts(t *testing.T) {
	hex := "000000000000000000000000000000000000000010010000000000000000000000000000000000000000"
	s := &Signer{
		Account:          helper.UInt160Zero,
		Scopes:           CustomContracts,
		AllowedContracts: []*helper.UInt160{helper.UInt160Zero},
	}
	bbw := io.NewBufBinaryWriter()
	s.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	assert.Equal(t, hex, helper.BytesToHex(b))

	s2 := NewDefaultSigner()
	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(hex))
	assert.Nil(t, br.Err)
	s2.Deserialize(br)
	assert.Nil(t, br.Err)
	assert.Equal(t, s.Account, s2.Account)
	assert.Equal(t, s.Scopes, s2.Scopes)
	assert.Equal(t, s.AllowedContracts[0], s2.AllowedContracts[0])
}

func TestSigner_Serialize_Deserialize_CustomGroups(t *testing.T) {
	hex := "0000000000000000000000000000000000000000200103b209fd4f53a7170ea4444e0cb0a6bb6a53c2bd016926989cf85f9b0fba17a70c"
	pk, err := crypto.NewECPointFromString("03b209fd4f53a7170ea4444e0cb0a6bb6a53c2bd016926989cf85f9b0fba17a70c")
	assert.Nil(t, err)
	s := &Signer{
		Account:       helper.UInt160Zero,
		Scopes:        CustomGroups,
		AllowedGroups: []*crypto.ECPoint{pk},
	}
	bbw := io.NewBufBinaryWriter()
	s.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	assert.Equal(t, hex, helper.BytesToHex(b))

	s2 := NewDefaultSigner()
	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(hex))
	assert.Nil(t, br.Err)
	s2.Deserialize(br)
	assert.Nil(t, br.Err)
	assert.Equal(t, s.Account, s2.Account)
	assert.Equal(t, s.Scopes, s2.Scopes)
	assert.Equal(t, s.AllowedGroups[0], s2.AllowedGroups[0])
}

func TestSigner_Size(t *testing.T) {
	account, _ := helper.UInt160FromString("edae9b97c72dbec43a7201b6388c24bfd86c71ae")
	contract, _ := helper.UInt160FromString("ad1344cbab4378e6bb9914769955fc4fb2742e2b")
	cs := &Signer{
		Account:          account,
		Scopes:           CustomContracts,
		AllowedContracts: []*helper.UInt160{contract},
	}
	size := cs.GetSize()
	assert.Equal(t, 20+1+1+20, size)
}

func TestSignerSlice_GetVarSize(t *testing.T) {
	account, _ := helper.UInt160FromString("edae9b97c72dbec43a7201b6388c24bfd86c71ae")
	contract, _ := helper.UInt160FromString("ad1344cbab4378e6bb9914769955fc4fb2742e2b")
	cs := &Signer{
		Account:          account,
		Scopes:           CustomContracts,
		AllowedContracts: []*helper.UInt160{contract},
	}
	css := []*Signer{cs}
	size := SignerSlice(css).GetVarSize()
	assert.Equal(t, 1+20+1+1+20, size)
}
