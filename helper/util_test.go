package helper

import (
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytesToHex(t *testing.T) {
	b := []byte{0xde, 0xad, 0xbe, 0xef}
	r := BytesToHex(b)
	assert.Equal(t, "deadbeef", r)
}

func TestHexTobytes(t *testing.T) {
	s := "deadbeef"
	r := HexTobytes(s)
	assert.Equal(t, []byte{0xde, 0xad, 0xbe, 0xef}, r)
}

func TestConcatBytes(t *testing.T) {
	b1 := []byte{0xde, 0xad}
	b2 := []byte{0xbe, 0xef}
	b := ConcatBytes(b1, b2)
	assert.Equal(t, []byte{0xde, 0xad, 0xbe, 0xef}, b)
}

func TestReverseBytes(t *testing.T) {
	var b = make([]byte, 0)
	r := ReverseBytes(b)
	assert.Equal(t, b, r)

	b = []byte{1}
	r = ReverseBytes(b)
	assert.Equal(t, b, r)

	b = []byte{1, 2}
	r = ReverseBytes(b)
	assert.Equal(t, []byte{2, 1}, r)

	b = []byte{1, 2, 3}
	r = ReverseBytes(b)
	assert.Equal(t, []byte{1, 2, 3}, b)
	assert.Equal(t, []byte{3, 2, 1}, r)
}

func TestAddressToScriptHash(t *testing.T) {
	r, err := AddressToScriptHash("NdtB8RXRmJ7Nhw1FPTm7E6HoDZGnDw37nf")
	assert.Nil(t, err)
	u, _ := UInt160FromBytes(crypto.Hash160([]byte{0x01}))
	assert.Equal(t, u.String(), r.String())
}

func TestScriptHashToAddress(t *testing.T) {
	u, _ := UInt160FromBytes(crypto.Hash160([]byte{0x01}))
	a := ScriptHashToAddress(u)
	assert.Equal(t, "NdtB8RXRmJ7Nhw1FPTm7E6HoDZGnDw37nf", a)
}

func TestReverseString(t *testing.T) {
	s := "deadbeef"
	r := ReverseString(s)
	assert.Equal(t, "feebdaed", r)
}

func TestAbs(t *testing.T) {
	abs := Abs(-2020)
	assert.Equal(t, int64(2020), abs)
}

func TestUInt32ToBytes(t *testing.T) {
	var u uint32 = 0xdeadbeef
	b := UInt32ToBytes(u)
	assert.Equal(t, []byte{0xef, 0xbe, 0xad, 0xde}, b)
}
