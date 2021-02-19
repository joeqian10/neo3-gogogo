package io

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestSerializable struct {
	Flag  bool
	Value []byte
}

func (t *TestSerializable) Serialize(writer *BinaryWriter) {
	writer.WriteLE(t.Flag)
	writer.WriteVarBytes(t.Value)
}

func (t *TestSerializable) Deserialize(reader *BinaryReader) {
	reader.ReadLE(&t.Flag)
	t.Value = reader.ReadVarBytes()
}

func TestToArray(t *testing.T) {
	v, _ := hex.DecodeString("abcd")
	ts := &TestSerializable{
		Flag:  true,
		Value: v,
	}
	data, _ := ToArray(ts)
	assert.Equal(t, data, []byte{0x01, 0x02, 0xab, 0xcd})
}

func TestAsSerializable(t *testing.T) {
	data := []byte{0x01, 0x02, 0xab, 0xcd}
	ts := &TestSerializable{}
	err := AsSerializable(ts, data)
	assert.Nil(t, err)
	assert.Equal(t, true, ts.Flag)
	assert.Equal(t, []byte{0xab, 0xcd}, ts.Value)
}
