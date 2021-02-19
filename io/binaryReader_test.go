package io

import (
	"bytes"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBinaryReaderFromIO(t *testing.T) {
	b := make([]byte, 4)
	r := bytes.NewReader(b)
	br := NewBinaryReaderFromIO(r)
	assert.NotNil(t, br)
}

func TestNewBinaryReaderFromBuf(t *testing.T) {
	b := make([]byte, 4)
	br := NewBinaryReaderFromBuf(b)
	assert.NotNil(t, br)
}

func TestBinaryReader_ReadBE(t *testing.T) {
	var (
		val    uint32 = 0xdeadbeef
		result uint32
		bin    = []byte{0xde, 0xad, 0xbe, 0xef}
	)
	br := NewBinaryReaderFromBuf(bin)
	br.ReadBE(&result)
	assert.Nil(t, br.Err)
	assert.Equal(t, val, result)
}

func TestBinaryReader_ReadLE(t *testing.T) {
	var (
		val    uint32 = 0xdeadbeef
		result uint32
		bin    = []byte{0xef, 0xbe, 0xad, 0xde}
	)
	br := NewBinaryReaderFromBuf(bin)
	br.ReadLE(&result)
	assert.Nil(t, br.Err)
	assert.Equal(t, val, result)
}

func TestBinaryReader_ReadVarUInt(t *testing.T) {
	var (
		val    uint16 = 0xdead
		result uint16
		bin    = []byte{0xfd, 0xad, 0xde}
	)
	br := NewBinaryReaderFromBuf(bin)
	result = uint16(br.ReadVarUIntWithMaxLimit(uint64(18446744073709551615)))
	assert.Nil(t, br.Err)
	assert.Equal(t, val, result)
}

func TestBinaryReader_ReadVarBytes(t *testing.T) {
	var (
		val    byte = 0xff
		result []byte
		bin    = []byte{0xfd, 0x01, 0x00, 0xff}
	)
	br := NewBinaryReaderFromBuf(bin)
	result = br.ReadVarBytes()
	assert.Nil(t, br.Err)
	assert.Equal(t, hex.EncodeToString([]byte{val}), hex.EncodeToString(result))
}

func TestBinaryReader_ReadVarString(t *testing.T) {
	var (
		val    string = "hello world"
		result string
		bin    = append([]byte{0x0b}, []byte(val)...)
	)
	br := NewBinaryReaderFromBuf(bin)
	result = br.ReadVarString(0x1000000)
	assert.Equal(t, val, result)
}
