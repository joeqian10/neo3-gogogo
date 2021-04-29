package crypto

import (
	"bytes"
	"crypto/elliptic"
	"encoding/hex"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/stretchr/testify/assert"
	"log"
	"math/big"
	"testing"
)

func GeneratePrivateKey(privateKeyLength int) []byte {
	privateKey := make([]byte, privateKeyLength)
	for i := 0; i < privateKeyLength; i++ {
		privateKey[i] = byte(i % 256)
	}
	return privateKey
}

var p256 = elliptic.P256()
var G, _ = CreateECPoint(p256.Params().Gx, p256.Params().Gy, &p256)

func TestCreateECPoint(t *testing.T) {
	x := p256.Params().Gx
	y := p256.Params().Gy
	p, err := CreateECPoint(x, y, &p256)
	assert.Nil(t, err)
	assert.Equal(t, 0, x.Cmp(p.X))
	assert.Equal(t, 0, y.Cmp(p.Y))
	assert.Equal(t, *p.Curve.Params().P, *p256.Params().P)
}

func TestNewECPoint(t *testing.T) {
	p, err := NewECPoint()
	assert.Nil(t, err)
	assert.Nil(t, p.X)
	assert.Nil(t, p.Y)
}

func TestNewECPointFromBytes(t *testing.T) {
	input1 := []byte{0x00}
	_, err := NewECPointFromBytes(input1)
	assert.NotNil(t, err)

	input2 := []byte{4, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150,
		79, 227, 66, 226, 254, 26, 127, 155, 142, 231, 235, 74, 124, 15, 158, 22, 43, 206, 51, 87, 107, 49, 94, 206, 203, 182, 64, 104, 55, 191, 81, 245}
	p2, err := NewECPointFromBytes(input2)
	assert.Nil(t, err)
	assert.Equal(t, G.X, p2.X)
	p21, err := NewECPointFromBytes(input2[1:]) // 64 bytes
	assert.Nil(t, err)
	assert.Equal(t, G.X, p21.X)

	input3 := []byte{3, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150}
	p3, err := NewECPointFromBytes(input3)
	assert.Nil(t, err)
	assert.Equal(t, G.X, p3.X)

	//input4 := GeneratePrivateKey(72)
	//p4, err := NewECPointFromBytes(input4)
	//assert.Nil(t, err)
	//expectedX4, b := new(big.Int).SetString("3634473727541135791764834762056624681715094789735830699031648273128038409767", 10)
	//assert.Equal(t, true, b)
	//expectedY4, b := new(big.Int).SetString("18165245710263168158644330920009617039772504630129940696140050972160274286151", 10)
	//assert.Equal(t, true, b)
	//assert.Equal(t, *expectedX4, *p4.X)
	//assert.Equal(t, *expectedY4, *p4.Y)
	//
	//input5 := GeneratePrivateKey(96)
	//p5, err := NewECPointFromBytes(input5)
	//assert.Nil(t, err)
	//expectedX5, b := new(big.Int).SetString("1780731860627700044960722568376592200742329637303199754547598369979440671", 10)
	//assert.Equal(t, true, b)
	//expectedY5, b := new(big.Int).SetString("14532552714582660066924456880521368950258152170031413196862950297402215317055", 10)
	//assert.Equal(t, true, b)
	//assert.Equal(t, *expectedX5, *p5.X)
	//assert.Equal(t, *expectedY5, *p5.Y)
	//
	//input6 := GeneratePrivateKey(104)
	//p6, err := NewECPointFromBytes(input6)
	//assert.Nil(t, err)
	//expectedX6, b := new(big.Int).SetString("3634473727541135791764834762056624681715094789735830699031648273128038409767", 10)
	//assert.Equal(t, true, b)
	//expectedY6, b := new(big.Int).SetString("18165245710263168158644330920009617039772504630129940696140050972160274286151", 10)
	//assert.Equal(t, true, b)
	//assert.Equal(t, *expectedX6, *p6.X)
	//assert.Equal(t, *expectedY6, *p6.Y)
}

func TestNewECPointFromString2(t *testing.T) {
	s := "028172918540b2b512eae1872a2a2e3a28d989c60d95dab8829ada7d7dd706d658"
	p, err := NewECPointFromString(s)
	assert.Nil(t, err)
	log.Println(hex.EncodeToString(p.EncodePoint(false)))
}

func TestNewECPointFromString(t *testing.T) {
	input1 := []byte{0x00}
	_, err := NewECPointFromString(hex.EncodeToString(input1))
	assert.NotNil(t, err)

	input2 := []byte{4, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150,
		79, 227, 66, 226, 254, 26, 127, 155, 142, 231, 235, 74, 124, 15, 158, 22, 43, 206, 51, 87, 107, 49, 94, 206, 203, 182, 64, 104, 55, 191, 81, 245}
	p2, err := NewECPointFromString(hex.EncodeToString(input2))
	assert.Nil(t, err)
	assert.Equal(t, G.X, p2.X)
	p21, err := NewECPointFromString(hex.EncodeToString(input2[1:])) // 64 bytes
	assert.Nil(t, err)
	assert.Equal(t, G.X, p21.X)

	input3 := []byte{3, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150}
	p3, err := NewECPointFromString(hex.EncodeToString(input3))
	assert.Nil(t, err)
	assert.Equal(t, G.X, p3.X)
}

func TestECPoint_CompareTo(t *testing.T) {
	x1 := big.NewInt(100)
	y1 := big.NewInt(200)
	p1, err := CreateECPoint(x1, y1, &p256)
	assert.Nil(t, err)

	x2 := big.NewInt(99)
	y2 := big.NewInt(200)
	p2, err := CreateECPoint(x2, y2, &p256)
	assert.Nil(t, err)

	x3 := big.NewInt(100)
	y3 := big.NewInt(201)
	p3, err := CreateECPoint(x3, y3, &p256)
	assert.Nil(t, err)

	assert.Equal(t, 1, p1.CompareTo(p2))
	assert.Equal(t, -1, p1.CompareTo(p3))
}

func TestDecodePoint(t *testing.T) {
	input1 := []byte{0x00}
	_, err := DecodePoint(input1, &p256)
	assert.NotNil(t, err)

	input2 := []byte{4, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150,
		79, 227, 66, 226, 254, 26, 127, 155, 142, 231, 235, 74, 124, 15, 158, 22, 43, 206, 51, 87, 107, 49, 94, 206, 203, 182, 64, 104, 55, 191, 81, 245}
	p2, err := DecodePoint(input2, &p256)
	assert.Nil(t, err)
	assert.Equal(t, G.X, p2.X)

	input3 := []byte{3, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150}
	p3, err := DecodePoint(input3, &p256)
	assert.Nil(t, err)
	assert.Equal(t, G.X, p3.X)
}

func TestDeserializeFrom(t *testing.T) {
	input1 := []byte{0x00}
	br1 := io.NewBinaryReaderFromBuf(input1)
	_, err := DeserializeFrom(br1, &p256)
	assert.NotNil(t, err)

	input2 := []byte{4, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150,
						79, 227, 66, 226, 254, 26, 127, 155, 142, 231, 235, 74, 124, 15, 158, 22, 43, 206, 51, 87, 107, 49, 94, 206, 203, 182, 64, 104, 55, 191, 81, 245}
	br2 := io.NewBinaryReaderFromBuf(input2)
	p2, err := DeserializeFrom(br2, &p256)
	assert.Nil(t, err)
	assert.Equal(t, G.X, p2.X)

	input3 := []byte{3, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150}
	br3 := io.NewBinaryReaderFromBuf(input3)
	p3, err := DeserializeFrom(br3, &p256)
	assert.Nil(t, err)
	assert.Equal(t, G.X, p3.X)
}

func TestExpectedECPointLength(t *testing.T) {
	l := ExpectedECPointLength(&p256)
	assert.Equal(t, 32, l)
}

func TestFromBytes(t *testing.T) {
	input1 := []byte{0x00}
	_, err := FromBytes(input1, &p256)
	assert.NotNil(t, err)

	input2 := []byte{4, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150,
		79, 227, 66, 226, 254, 26, 127, 155, 142, 231, 235, 74, 124, 15, 158, 22, 43, 206, 51, 87, 107, 49, 94, 206, 203, 182, 64, 104, 55, 191, 81, 245}
	p2, err := FromBytes(input2, &p256)
	assert.Nil(t, err)
	assert.Equal(t, G.X, p2.X)
	p21, err := FromBytes(input2[1:], &p256) // 64 bytes
	assert.Nil(t, err)
	assert.Equal(t, G.X, p21.X)

	input3 := []byte{3, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150}
	p3, err := FromBytes(input3, &p256)
	assert.Nil(t, err)
	assert.Equal(t, G.X, p3.X)
}

func TestParse(t *testing.T) {
	input1 := []byte{0x00}
	_, err := Parse(hex.EncodeToString(input1), &p256)
	assert.NotNil(t, err)

	input2 := []byte{4, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150,
		79, 227, 66, 226, 254, 26, 127, 155, 142, 231, 235, 74, 124, 15, 158, 22, 43, 206, 51, 87, 107, 49, 94, 206, 203, 182, 64, 104, 55, 191, 81, 245}
	p2, err := Parse(hex.EncodeToString(input2), &p256)
	assert.Nil(t, err)
	assert.Equal(t, G.X, p2.X)

	input3 := []byte{3, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150}
	p3, err := Parse(hex.EncodeToString(input3), &p256)
	assert.Nil(t, err)
	assert.Equal(t, G.X, p3.X)
}

func TestECPoint_Deserialize(t *testing.T) {
	input1 := []byte{0x00}
	br1 := io.NewBinaryReaderFromBuf(input1)
	p1, err := NewECPoint()
	assert.Nil(t, err)
	p1.Deserialize(br1)
	assert.NotNil(t, br1.Err)

	input2 := []byte{4, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150,
		79, 227, 66, 226, 254, 26, 127, 155, 142, 231, 235, 74, 124, 15, 158, 22, 43, 206, 51, 87, 107, 49, 94, 206, 203, 182, 64, 104, 55, 191, 81, 245}
	br2 := io.NewBinaryReaderFromBuf(input2)
	p2, err := NewECPoint()
	assert.Nil(t, err)
	p2.Deserialize(br2)
	assert.Nil(t, br2.Err)
	assert.Equal(t, G.X, p2.X)

	input3 := []byte{3, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150}
	br3 := io.NewBinaryReaderFromBuf(input3)
	p3, err := NewECPoint()
	assert.Nil(t, err)
	p3.Deserialize(br3)
	assert.Nil(t, br3.Err)
	assert.Equal(t, G.X, p3.X)
}

func TestECPoint_EncodePoint(t *testing.T) {
	expected1 := []byte{0x00}
	p1, err := NewECPoint()
	assert.Nil(t, err)
	r1 := p1.EncodePoint(true)
	assert.Equal(t, true, bytes.Equal(expected1, r1))

	expected2 := []byte{4, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150,
		79, 227, 66, 226, 254, 26, 127, 155, 142, 231, 235, 74, 124, 15, 158, 22, 43, 206, 51, 87, 107, 49, 94, 206, 203, 182, 64, 104, 55, 191, 81, 245}
	r2 := G.EncodePoint(false)
	assert.Equal(t, true, bytes.Equal(expected2, r2))

	expected3 := []byte{3, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150}
	r3 := G.EncodePoint(true)
	assert.Equal(t, true, bytes.Equal(expected3, r3))
}

func TestECPoint_Equals(t *testing.T) {
	p, err := NewECPoint()
	assert.Nil(t, err)
	assert.Equal(t, false, p.Equals(G))
	assert.Equal(t, true, G.Equals(G))
}

func TestECPoint_IsInfinity(t *testing.T) {
	p, err := NewECPoint()
	assert.Nil(t, err)
	assert.Equal(t, false, G.IsInfinity())
	assert.Equal(t, true, p.IsInfinity())
}

func TestECPoint_IsOnCurve(t *testing.T) {
	p, err := CreateECPoint(big.NewInt(100), big.NewInt(200), &p256)
	assert.Nil(t, err)
	b := p.IsOnCurve()
	assert.Equal(t, false, b)
}

func TestECPoint_Serialize(t *testing.T) {
	expected1 := []byte{0x00}
	p1, err := NewECPoint()
	assert.Nil(t, err)
	bw1 := io.NewBufBinaryWriter()
	p1.Serialize(bw1.BinaryWriter)
	r1 := bw1.Bytes()
	assert.Equal(t, true, bytes.Equal(expected1, r1))

	expected2 := []byte{3, 107, 23, 209, 242, 225, 44, 66, 71, 248, 188, 230, 229, 99, 164, 64, 242, 119, 3, 125, 129, 45, 235, 51, 160, 244, 161, 57, 69, 216, 152, 194, 150}
	bw2 := io.NewBufBinaryWriter()
	G.Serialize(bw2.BinaryWriter)
	r2 := bw2.Bytes()
	assert.Equal(t, true, bytes.Equal(expected2, r2))
}

func TestECPoint_Size(t *testing.T) {
	p, err := NewECPoint()
	assert.Nil(t, err)
	assert.Equal(t, 1, p.Size())

	assert.Equal(t, 33, G.Size())
}

func TestECPoint_String(t *testing.T) {
	s := "03b7a7f933199f28cc1c48d22a21c78ac3992cf7fceb038a9c670fe55444426619"
	p, err := NewECPointFromString(s)
	assert.Nil(t, err)
	assert.Equal(t, s, p.String())
}
