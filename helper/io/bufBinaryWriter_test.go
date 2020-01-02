package io

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBufBinaryWriter(t *testing.T) {
	bw := NewBufBinaryWriter()
	assert.NotNil(t, bw)
}

func TestBufBinaryWriter_Bytes(t *testing.T) {
	var (
		val     uint32 = 0xdeadbeef
		bin     = []byte{0xde, 0xad, 0xbe, 0xef}
	)
	bw := NewBufBinaryWriter()
	bw.WriteBE(val)
	assert.Nil(t, bw.Err)
	result := bw.Bytes()
	assert.Equal(t, bin, result)
}

func TestBufBinaryWriter_Reset(t *testing.T) {
	bw := NewBufBinaryWriter()
	for i := 0; i < 3; i++ {
		bw.WriteLE(uint32(i))
		assert.Nil(t, bw.Err)
		_ = bw.Bytes()
		assert.NotNil(t, bw.Err)
		bw.Reset()
		assert.Nil(t, bw.Err)
	}
}