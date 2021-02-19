package io

import (
	"bytes"
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBinaryWriterFromIO(t *testing.T) {
	b := new(bytes.Buffer)
	bw := NewBinaryWriterFromIO(b)
	assert.NotNil(t, bw)
}

func TestBinaryWriter_WriteBE(t *testing.T) {
	var (
		b          = new(bytes.Buffer)
		val uint32 = 0xdeadbeef
		bin        = []byte{0xde, 0xad, 0xbe, 0xef}
	)
	bw := NewBinaryWriterFromIO(b)
	bw.WriteBE(val) // write to the buffer
	assert.Nil(t, bw.Err)
	assert.Equal(t, b.Bytes(), bin)
}

func TestBinaryWriter_WriteLE(t *testing.T) {
	var (
		b          = new(bytes.Buffer)
		val uint32 = 0xdeadbeef
		bin        = []byte{0xef, 0xbe, 0xad, 0xde}
	)
	bw := NewBinaryWriterFromIO(b)
	bw.WriteLE(val)
	assert.Nil(t, bw.Err)
	assert.Equal(t, b.Bytes(), bin)
}

func TestBinaryWriter_WriteLE2(t *testing.T) {
	var (
		b        = new(bytes.Buffer)
		val uint = 0x01020304
		bin      = []byte{0x04, 0x03, 0x02, 0x01}
	)
	bw := NewBinaryWriterFromIO(b)
	bw.WriteLE(uint32(val)) // need to convert to uint32, or there's an error
	assert.Nil(t, bw.Err)
	assert.Equal(t, b.Bytes(), bin)
}

func TestBinaryWriter_WriteVarUInt(t *testing.T) {
	var (
		b          = new(bytes.Buffer)
		val uint32 = 0xdeadbeef
		bin        = []byte{0xfe, 0xef, 0xbe, 0xad, 0xde}
	)
	bw := NewBinaryWriterFromIO(b)
	bw.WriteVarUInt(uint64(val))
	assert.Nil(t, bw.Err)
	assert.Equal(t, b.Bytes(), bin)
}

func TestBinaryWriter_WriteVarBytes(t *testing.T) {
	var (
		b          = new(bytes.Buffer)
		val uint32 = 0xdeadbeef
		bin        = []byte{0x04, 0xef, 0xbe, 0xad, 0xde}
	)
	bw := NewBinaryWriterFromIO(b)
	var buff = make([]byte, 4)
	binary.LittleEndian.PutUint32(buff, val)
	bw.WriteVarBytes(buff)
	assert.Nil(t, bw.Err)
	assert.Equal(t, b.Bytes(), bin)
}

func TestBinaryWriter_WriteVarString(t *testing.T) {
	var (
		b          = new(bytes.Buffer)
		val string = "hello world"
		bin        = append([]byte{0x0b}, []byte(val)...)
	)
	bw := NewBinaryWriterFromIO(b)
	bw.WriteVarString(val)
	assert.Nil(t, bw.Err)
	assert.Equal(t, b.Bytes(), bin)
}
