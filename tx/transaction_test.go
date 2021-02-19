package tx

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction_GetScript(t *testing.T) {
	x := Transaction{}
	assert.Nil(t, x.GetScript())
}

func TestTransaction_SetScript(t *testing.T) {
	val := make([]byte, 32)
	val[0] = 0x42
	x := Transaction{}
	x.SetScript(val)
	assert.Equal(t, 32, len(x.GetScript()))
}

func TestTransaction_GetSystemFee(t *testing.T) {
	x := Transaction{}
	assert.Equal(t, int64(0), x.GetSystemFee())
}

func TestTransaction_SetSystemFee(t *testing.T) {
	var val int64 = 4200000000
	x := Transaction{}
	x.SetSystemFee(val)
	assert.Equal(t, val, x.GetSystemFee())
}

func TestTransaction_GetSize(t *testing.T) {
	x := Transaction{}
	val := make([]byte, 32)
	val[0] = 0x42
	x.SetScript(val)
	x.SetSigners([]Signer{})
	x.SetAttributes([]ITransactionAttribute{})
	x.SetWitnesses([]Witness{{
		InvocationScript:   []byte{},
		VerificationScript: []byte{},
	}})

	assert.Equal(t, uint8(0), x.GetVersion())
	assert.Equal(t, 32, len(x.GetScript()))
	assert.Equal(t, 33, sc.ByteSlice(x.GetScript()).GetVarSize())
	assert.Equal(t, 63, x.GetSize())
}

func TestTransaction_GetScriptHashesForVerifying(t *testing.T) {
	x := Transaction{}
	x.SetSigners([]Signer{{Account: helper.UInt160Zero, Scopes: Global}})
	hashes := x.GetScriptHashesForVerifying()
	assert.Equal(t, 1, len(hashes))
}

func TestTransaction_Deserialize(t *testing.T) {
	s := "00" + // version
		"04030201" + // nonce
		"00e1f50500000000" + // system fee (1 GAS)
		"0100000000000000" + // network fee (1 satoshi)
		"04030201" + // timelimit
		"01000000000000000000000000000000000000000000" + // empty signer
		"00" + // no attributes
		"0111" + // push1 script
		"00" // no witness

	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(s))
	tx := NewTransaction()
	tx.Deserialize(br)
	assert.Equal(t, uint8(0), tx.version)
	assert.Equal(t, uint32(0x01020304), tx.nonce)
	assert.Equal(t, int64(100000000), tx.sysfee)
	assert.Equal(t, int64(0x0000000000000001), tx.netfee)
	assert.Equal(t, uint32(0x01020304), tx.validUntilBlock)
	assert.Equal(t, 1, len(tx.signers))
	assert.Equal(t, 0, len(tx.attributes))
	assert.Equal(t, 1, len(tx.script))
	assert.Equal(t, 0, len(tx.witnesses))
}

func TestTransaction_DeserializeUnsigned(t *testing.T) {
	s := "00" + // version
		"04030201" + // nonce
		"00e1f50500000000" + // system fee (1 GAS)
		"0100000000000000" + // network fee (1 satoshi)
		"04030201" + // timelimit
		"01000000000000000000000000000000000000000000" + // empty signer
		"00" + // no attributes
		"0111" + // push1 script
		"00" // no witness

	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(s))
	tx := NewTransaction()
	tx.DeserializeUnsigned(br)
	assert.Equal(t, uint8(0), tx.version)
	assert.Equal(t, uint32(0x01020304), tx.nonce)
	assert.Equal(t, int64(100000000), tx.sysfee)
	assert.Equal(t, int64(0x0000000000000001), tx.netfee)
	assert.Equal(t, uint32(0x01020304), tx.validUntilBlock)
	assert.Equal(t, 1, len(tx.signers))
	assert.Equal(t, 0, len(tx.attributes))
	assert.Equal(t, 1, len(tx.script))
}

func TestTransaction_DeserializeWitnesses(t *testing.T) {
	s := "00" // no witness
	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(s))
	tx := NewTransaction()
	tx.DeserializeWitnesses(br)
	assert.Equal(t, 0, len(tx.witnesses))
}

func TestTransaction_HeaderSize(t *testing.T) {
	tx := NewTransaction()
	headerSize := tx.HeaderSize()
	assert.Equal(t, 25, headerSize)
}

func TestTransaction_ToByteArray(t *testing.T) {
	tx := Transaction{
		version:         0x00,
		nonce:           0x01020304,
		sysfee:          GasFactor,
		netfee:          0x0000000000000001,
		validUntilBlock: 0x01020304,
		signers:         []Signer{},
		attributes:      []ITransactionAttribute{},
		script:          []byte{byte(sc.PUSH1)},
		witnesses:       []Witness{},
	}
	r := tx.ToByteArray()
	expected := "00" + // version
		"04030201" + // nonce
		"00e1f50500000000" + // system fee (1 GAS)
		"0100000000000000" + // network fee (1 satoshi)
		"04030201" + // timelimit
		"00" + // no signers
		"00" + // no attributes
		"0111" + // push1 script
		"00" // no witness

	assert.Equal(t, expected, helper.BytesToHex(r))
}

func TestTransaction_Serialize(t *testing.T) {
	tx := Transaction{
		version:         0x00,
		nonce:           0x01020304,
		sysfee:          GasFactor,
		netfee:          0x0000000000000001,
		validUntilBlock: 0x01020304,
		signers:         []Signer{},
		attributes:      []ITransactionAttribute{},
		script:          []byte{byte(sc.PUSH1)},
		witnesses:       []Witness{},
	}
	bbw := io.NewBufBinaryWriter()
	tx.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	expected := "00" + // version
		"04030201" + // nonce
		"00e1f50500000000" + // system fee (1 GAS)
		"0100000000000000" + // network fee (1 satoshi)
		"04030201" + // timelimit
		"00" + // no signers
		"00" + // no attributes
		"0111" + // push1 script
		"00" // no witness
	assert.Equal(t, expected, helper.BytesToHex(b))
}

func TestTransaction_SerializeUnsigned(t *testing.T) {
	tx := Transaction{
		version:         0x00,
		nonce:           0x01020304,
		sysfee:          GasFactor,
		netfee:          0x0000000000000001,
		validUntilBlock: 0x01020304,
		signers:         []Signer{},
		attributes:      []ITransactionAttribute{},
		script:          []byte{byte(sc.PUSH1)},
		witnesses:       []Witness{},
	}
	bbw := io.NewBufBinaryWriter()
	tx.SerializeUnsigned(bbw.BinaryWriter)
	b := bbw.Bytes()
	expected := "00" + // version
		"04030201" + // nonce
		"00e1f50500000000" + // system fee (1 GAS)
		"0100000000000000" + // network fee (1 satoshi)
		"04030201" + // timelimit
		"00" + // no signers
		"00" + // no attributes
		"0111" // push1 script

	assert.Equal(t, expected, helper.BytesToHex(b))
}

func TestTransaction_SerializeWitnesses(t *testing.T) {
	tx := Transaction{
		version:         0x00,
		nonce:           0x01020304,
		sysfee:          GasFactor,
		netfee:          0x0000000000000001,
		validUntilBlock: 0x01020304,
		signers:         []Signer{},
		attributes:      []ITransactionAttribute{},
		script:          []byte{byte(sc.PUSH1)},
		witnesses:       []Witness{},
	}
	bbw := io.NewBufBinaryWriter()
	tx.SerializeWitnesses(bbw.BinaryWriter)
	b := bbw.Bytes()
	expected := "00" // no witness

	assert.Equal(t, expected, helper.BytesToHex(b))
}
