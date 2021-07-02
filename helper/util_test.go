package helper

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestBytesToHex(t *testing.T) {
	b := []byte{0xde, 0xad, 0xbe, 0xef}
	r := BytesToHex(b)
	assert.Equal(t, "deadbeef", r)
}

func TestHexToBytes(t *testing.T) {
	s := "deadbeef"
	r := HexToBytes(s)
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

func TestReverseString(t *testing.T) {
	s := "deadbeef"
	r := ReverseString(s)
	assert.Equal(t, "feebdaed", r)
}

func TestAbs(t *testing.T) {
	abs := Abs(-2020)
	assert.Equal(t, int64(2020), abs)
}

func TestUInt16ToBytes(t *testing.T) {
	var u uint16 = 0xbeef
	b := UInt16ToBytes(u)
	assert.Equal(t, []byte{0xef, 0xbe}, b)
}

func TestUInt32ToBytes(t *testing.T) {
	var u uint32 = 0xdeadbeef
	b := UInt32ToBytes(u)
	assert.Equal(t, []byte{0xef, 0xbe, 0xad, 0xde}, b)
}

func TestUInt64ToBytes(t *testing.T) {
	var u uint64 = 0xfeedabeedeadbeef
	b := UInt64ToBytes(u)
	assert.Equal(t, []byte{0xef, 0xbe, 0xad, 0xde, 0xee, 0xab, 0xed, 0xfe}, b)
}

func TestIntToBytes(t *testing.T) {
	var u int = 0xdeadbeef
	b := IntToBytes(u)
	assert.Equal(t, []byte{0xef, 0xbe, 0xad, 0xde}, b)
}

func TestInt64ToBytes(t *testing.T) {
	var u int64 = 0x0000feedabeebeef
	b := Int64ToBytes(u)
	assert.Equal(t, []byte{0xef, 0xbe, 0xee, 0xab, 0xed, 0xfe, 0x00, 0x00}, b)
}

func TestGenerateRandomBytes(t *testing.T) {
	l := 8
	b, err := GenerateRandomBytes(l)
	assert.Nil(t, err)
	assert.Equal(t, l, len(b))
}

func TestGetVarSize(t *testing.T) {
	var r int
	r = GetVarSize(0xF9)
	assert.Equal(t, 1, r)
	r = GetVarSize(0xFFFF)
	assert.Equal(t, 3, r)
	r = GetVarSize(0xFFFFFF)
	assert.Equal(t, 5, r)
}

func TestBigIntToNeoBytes(t *testing.T) {
	b := big.NewInt(-200)
	bs := BigIntToNeoBytes(b)
	assert.Equal(t, []byte{0x38, 0xff}, bs)
}

func TestBigIntFromNeoBytes(t *testing.T) {
	bs := []byte{0x38, 0xff}
	b := BigIntFromNeoBytes(bs)
	assert.Equal(t, 0, big.NewInt(-200).Cmp(b))
}
