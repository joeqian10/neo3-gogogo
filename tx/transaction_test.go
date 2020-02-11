package tx

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction_Deserialize(t *testing.T) {
	s := "00" + // version
		"04030201" + // nonce
		"0000000000000000000000000000000000000000" + // sender
		"00e1f50500000000" + // system fee (1 GAS)
		"0100000000000000" + // network fee (1 satoshi)
		"04030201" + // timelimit
		"00" + // no attributes
		"00" + // no cosigners
		"0111" + // push1 script
		"00" // no witness

	br := io.NewBinaryReaderFromBuf(helper.HexTobytes(s))
	tx := NewTransaction()
	tx.Deserialize(br)
	assert.Equal(t, uint8(0), tx.version)
	assert.Equal(t, uint32(0x01020304), tx.nonce)
	assert.Equal(t, helper.UInt160{}, tx.sender)
	assert.Equal(t, int64(100000000), tx.sysfee)
	assert.Equal(t, int64(0x0000000000000001), tx.netfee)
	assert.Equal(t, uint32(0x01020304), tx.validUntilBlock)
	assert.Equal(t, 0, len(tx.attributes))
	assert.Equal(t, 0, len(tx.cosigners))
	assert.Equal(t, 1, len(tx.script))
	assert.Equal(t, 0, len(tx.witnesses))
}

func TestTransaction_DeserializeUnsigned(t *testing.T) {
	s := "00" + // version
		"04030201" + // nonce
		"0000000000000000000000000000000000000000" + // sender
		"00e1f50500000000" + // system fee (1 GAS)
		"0100000000000000" + // network fee (1 satoshi)
		"04030201" + // timelimit
		"00" + // no attributes
		"00" + // no cosigners
		"0111" // push1 script

	br := io.NewBinaryReaderFromBuf(helper.HexTobytes(s))
	tx := NewTransaction()
	tx.DeserializeUnsigned(br)
	assert.Equal(t, uint8(0), tx.version)
	assert.Equal(t, uint32(0x01020304), tx.nonce)
	assert.Equal(t, helper.UInt160{}, tx.sender)
	assert.Equal(t, int64(100000000), tx.sysfee)
	assert.Equal(t, int64(0x0000000000000001), tx.netfee)
	assert.Equal(t, uint32(0x01020304), tx.validUntilBlock)
	assert.Equal(t, 0, len(tx.attributes))
	assert.Equal(t, 0, len(tx.cosigners))
	assert.Equal(t, 1, len(tx.script))
}

func TestTransaction_DeserializeWitnesses(t *testing.T) {
	s := "00" // no witness
	br := io.NewBinaryReaderFromBuf(helper.HexTobytes(s))
	tx := NewTransaction()
	tx.DeserializeWitnesses(br)
	assert.Equal(t, 0, len(tx.witnesses))
}

func TestTransaction_HeaderSize(t *testing.T) {
	tx := NewTransaction()
	headerSize := tx.HeaderSize()
	assert.Equal(t, 45, headerSize)
}

func TestTransaction_RawTransaction(t *testing.T) {
	tx := Transaction{
		version:         0x00,
		nonce:           0x01020304,
		sender:          helper.UInt160{},
		sysfee:          GasFactor,
		netfee:          0x0000000000000001,
		validUntilBlock: 0x01020304,
		attributes:      []*TransactionAttribute{},
		cosigners:       []*Cosigner{},
		script:          []byte{byte(sc.PUSH1)},
		witnesses:       []*Witness{},
	}
	r := tx.RawTransaction()
	expected := "00" + // version
		"04030201" + // nonce
		"0000000000000000000000000000000000000000" + // sender
		"00e1f50500000000" + // system fee (1 GAS)
		"0100000000000000" + // network fee (1 satoshi)
		"04030201" + // timelimit
		"00" + // no attributes
		"00" + // no cosigners
		"0111" + // push1 script
		"00" // no witness

	assert.Equal(t, expected, helper.BytesToHex(r))
}

func TestTransaction_Serialize(t *testing.T) {
	tx := Transaction{
		version:         0x00,
		nonce:           0x01020304,
		sender:          helper.UInt160{},
		sysfee:          GasFactor,
		netfee:          0x0000000000000001,
		validUntilBlock: 0x01020304,
		attributes:      []*TransactionAttribute{},
		cosigners:       []*Cosigner{},
		script:          []byte{byte(sc.PUSH1)},
		witnesses:       []*Witness{},
	}
	bbw := io.NewBufBinaryWriter()
	tx.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	expected := "00" + // version
		"04030201" + // nonce
		"0000000000000000000000000000000000000000" + // sender
		"00e1f50500000000" + // system fee (1 GAS)
		"0100000000000000" + // network fee (1 satoshi)
		"04030201" + // timelimit
		"00" + // no attributes
		"00" + // no cosigners
		"0111" + // push1 script
		"00" // no witness
	assert.Equal(t, expected, helper.BytesToHex(b))
}

func TestTransaction_SerializeUnsigned(t *testing.T) {
	tx := Transaction{
		version:         0x00,
		nonce:           0x01020304,
		sender:          helper.UInt160{},
		sysfee:          GasFactor,
		netfee:          0x0000000000000001,
		validUntilBlock: 0x01020304,
		attributes:      []*TransactionAttribute{},
		cosigners:       []*Cosigner{},
		script:          []byte{byte(sc.PUSH1)},
		witnesses:       []*Witness{},
	}
	bbw := io.NewBufBinaryWriter()
	tx.SerializeUnsigned(bbw.BinaryWriter)
	b := bbw.Bytes()
	expected := "00" + // version
		"04030201" + // nonce
		"0000000000000000000000000000000000000000" + // sender
		"00e1f50500000000" + // system fee (1 GAS)
		"0100000000000000" + // network fee (1 satoshi)
		"04030201" + // timelimit
		"00" + // no attributes
		"00" + // no cosigners
		"0111" // push1 script

	assert.Equal(t, expected, helper.BytesToHex(b))
}

func TestTransaction_SerializeWitnesses(t *testing.T) {
	tx := Transaction{
		version:         0x00,
		nonce:           0x01020304,
		sender:          helper.UInt160{},
		sysfee:          GasFactor,
		netfee:          0x0000000000000001,
		validUntilBlock: 0x01020304,
		attributes:      []*TransactionAttribute{},
		cosigners:       []*Cosigner{},
		script:          []byte{byte(sc.PUSH1)},
		witnesses:       []*Witness{},
	}
	bbw := io.NewBufBinaryWriter()
	tx.SerializeWitnesses(bbw.BinaryWriter)
	b := bbw.Bytes()
	expected := "00" // no witness

	assert.Equal(t, expected, helper.BytesToHex(b))
}
