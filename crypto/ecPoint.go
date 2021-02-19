package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"math/big"
)

type ECPoint struct {
	X, Y *big.Int
	Curve elliptic.Curve

	compressedPoint   []byte
	uncompressedPoint []byte
}

var P256 = elliptic.P256()

func CreateECPoint(x *big.Int, y *big.Int, curve *elliptic.Curve) (*ECPoint, error) {
	if (x == nil) != (y == nil) || curve == nil {
		return nil, fmt.Errorf("exactly one of the parameters is nil")
	}
	if (x != nil ) && (y != nil ) { // both are nil means infinite point
		P := (*curve).Params().P
		if x.Cmp(P) >= 0 || y.Cmp(P) >= 0 {
			return nil, fmt.Errorf("invalid parameter: X or Y is bigger than P")
		}
	}
	p := &ECPoint{
			X:     x,
			Y:     y,
			Curve: *curve,
	}
	return p, nil
}

func NewECPoint() (*ECPoint, error) {
	return CreateECPoint(nil, nil, &P256)
}

//NewPublicKey return a public key created from the given []byte.
func NewECPointFromBytes(data []byte) (*ECPoint, error) {
	p256 :=  elliptic.P256() // Secp256r1
	return FromBytes(data, &p256)
}

// NewPublicKeyFromString return a public key created from the given hex string.
func NewECPointFromString(s string) (*ECPoint, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return NewECPointFromBytes(b)
}


// IsInfinity checks if point P is infinity on EllipticCurve ec.
func (p *ECPoint) IsInfinity() bool {
	return p.X == nil && p.Y == nil
}

// IsOnCurve checks if this point is on the specified curve
func (p *ECPoint) IsOnCurve() bool {
	return p.Curve.IsOnCurve(p.X, p.Y)
}

func (p *ECPoint) Size() int {
	if p.IsInfinity() {
		return 1
	}
	return 33
}

// Compare two points
func (p *ECPoint) CompareTo(other *ECPoint) int {
	if p.Curve != other.Curve {
		panic("invalid comparison for points with different curves")
	}
	if p == other {
		return 0
	}
	xLess := p.X.Cmp(other.X)
	if xLess != 0 {
		return xLess
	}
	return p.Y.Cmp(other.Y)
}

func DecodePoint(encoded []byte, curve *elliptic.Curve) (*ECPoint, error) {
	var p *ECPoint
	var err error
	expectedPointLength := ((*curve).Params().BitSize + 7) / 8
	switch encoded[0] {
	case 0x02, 0x03: // compressed
		if len(encoded) != expectedPointLength+1 {
			return nil, fmt.Errorf("incorrect length for compressed encoding")
		}
		yTilde := int(encoded[0] & 1)
		x := new(big.Int).SetBytes(encoded[1:])
		p, err = decompressPoint(yTilde, x, curve)
		if err != nil {
			return nil, err
		}
		p.compressedPoint = encoded
		break
	case 0x04: // uncompressed
		if len(encoded) != 2*expectedPointLength+1 {
			return nil, fmt.Errorf("incorrect length for uncompressed/hybrid encoding")
		}
		x := new(big.Int).SetBytes(encoded[1 : expectedPointLength+1])
		y := new(big.Int).SetBytes(encoded[expectedPointLength+1:])
		p, err = CreateECPoint(x, y, curve)
		if err != nil {
			return nil, err
		}
		p.uncompressedPoint = encoded
		break
	default:
		return nil, fmt.Errorf("invalid point encoding")
	}
	return p, nil
}

// y**2 = x**3 + a*x + b  % p, a = -3
// x³ - 3x + b
func decompressPoint(yTilde int, x *big.Int, curve *elliptic.Curve) (*ECPoint, error) {
	A := big.NewInt(3)
	B := (*curve).Params().B
	P := (*curve).Params().P

	xCubed := new(big.Int).Exp(x, A, P)          // x^2
	threeX := new(big.Int).Mul(x, A)             // x^3
	threeX.Mod(threeX, P)                        // 3*x
	ySquared := new(big.Int).Sub(xCubed, threeX) // x^3 - 3*x
	ySquared.Add(ySquared, B)                    // x^3 - 3*x + B
	ySquared.Mod(ySquared, P)                    // x^3 - 3*x + B % P
	y := new(big.Int).ModSqrt(ySquared, P)
	if y == nil {
		return nil, fmt.Errorf("error computing Y for compressed point")
	}
	if y.Bit(0) != uint(yTilde) {
		y.Neg(y)
		y.Mod(y, P)
	}
	return CreateECPoint(x, y, curve)
}

// decodeCompressedY performs decompression of Y coordinate for given X and Y's least significant bit
func decodeCompressedY(x *big.Int, ylsb uint) (*big.Int, error) {
	c := elliptic.P256()
	cp := c.Params()
	three := big.NewInt(3)

	xCubed := new(big.Int).Exp(x, three, cp.P)
	threeX := new(big.Int).Mul(x, three)
	threeX.Mod(threeX, cp.P)
	ySquared := new(big.Int).Sub(xCubed, threeX)
	ySquared.Add(ySquared, cp.B)
	ySquared.Mod(ySquared, cp.P)
	y := new(big.Int).ModSqrt(ySquared, cp.P)
	if y == nil {
		return nil, fmt.Errorf("error computing Y for compressed point")
	}
	if y.Bit(0) != ylsb {
		y.Neg(y)
		y.Mod(y, cp.P)
	}
	return y, nil
}

// Deserialize a PublicKey from the given io.Reader.
func (p *ECPoint) Deserialize(br *io.BinaryReader) {
	q, err := DeserializeFrom(br, &p.Curve)
	if err != nil {
		br.Err = err
		return
	}
	p.X = q.X
	p.Y = q.Y
}

func DeserializeFrom(br *io.BinaryReader, curve *elliptic.Curve) (*ECPoint, error) {
	buffer := make([]byte, 1+ExpectedECPointLength(curve)*2)
	buffer[0] = br.ReadByte()
	switch buffer[0] {
	case 0x02, 0x03:
		br.ReadLE(buffer[1 : 1+ExpectedECPointLength(curve)])
		return DecodePoint(buffer[:1+ExpectedECPointLength(curve)], curve)
	case 0x04:
		br.ReadLE(buffer[1 : 1+ExpectedECPointLength(curve)*2])
		return DecodePoint(buffer, curve)
	default:
		return nil, fmt.Errorf("invalid point encoding")
	}
}

// EncodePoint encodes the point to a byte array
func (p *ECPoint) EncodePoint(compressed bool) []byte {
	if p.IsInfinity() {
		return []byte{0x00}
	}
	var data []byte
	if compressed {
		if p.compressedPoint != nil {
			return p.compressedPoint
		}
		data = make([]byte, 33)
	} else {
		if p.uncompressedPoint != nil {
			return p.uncompressedPoint
		}
		data = make([]byte, 65)
		yBytes := p.Y.Bytes()
		paddedY := append(bytes.Repeat([]byte{0x00}, 32-len(yBytes)), yBytes...)
		copy(data[65-len(paddedY):], paddedY)
	}
	xBytes := p.X.Bytes()
	paddedX := append(bytes.Repeat([]byte{0x00}, 32-len(xBytes)), xBytes...)
	copy(data[33-len(paddedX):], paddedX)
	if compressed {
		if p.Y.Bit(0) == 0 {
			data[0] = 0x02
		} else {
			data[0]	= 0x03
		}
	} else {
		data[0]	= 0x04
	}
	return data
}

func (p *ECPoint) Equals(other *ECPoint) bool {
	if p == other {return true}
	if other == nil {return false}
	if p.IsInfinity() && other.IsInfinity() {return true}
	if p.IsInfinity() || other.IsInfinity() {return false}
	return p.X.Cmp(other.X) == 0 && p.Y.Cmp(other.Y) == 0
}

func (p *ECPoint) ExistsIn(points []ECPoint) bool {
	for _, point := range points {
		if p.Equals(&point) {
			return true
		}
	}
	return false
}

// expected byte array length
func ExpectedECPointLength(curve *elliptic.Curve) int {
	return ((*curve).Params().BitSize + 7) / 8
}

func FromBytes(data []byte, curve *elliptic.Curve) (*ECPoint, error) {
	switch len(data) {
	case 33, 65:
		return DecodePoint(data, curve)
	case 64, 72:
		l := len(data)
		return DecodePoint(append([]byte{0x04}, data[l-64:]...), curve)
	case 96, 104:
		l := len(data)
		return DecodePoint(append([]byte{0x04}, data[l-96:l-32]...), curve)
	default:
		return nil, fmt.Errorf("invalid parameter format")
	}
}

func Parse(value string, curve *elliptic.Curve) (*ECPoint,error) {
	return DecodePoint(helper.HexToBytes(value), curve)
}

// Serialize encodes a PublicKey to the given io.Writer.
func (p *ECPoint) Serialize(bw *io.BinaryWriter) {
	bw.WriteLE(p.EncodePoint(true))
}

// String implements the Stringer interface.
func (p *ECPoint) String() string {
	return helper.BytesToHex(p.EncodePoint(true))
}

// ToECDsa returns a ecdsa.PublicKey
func (p *ECPoint) ToECDsa() *ecdsa.PublicKey {
	return &ecdsa.PublicKey{
		Curve: p.Curve,
		X:     p.X,
		Y:     p.Y,
	}
}




// PublicKeys is a list of public keys.
type PublicKeySlice []ECPoint

func (keys PublicKeySlice) Len() int           { return len(keys) }
func (keys PublicKeySlice) Swap(i, j int)      { keys[i], keys[j] = keys[j], keys[i] }
func (keys PublicKeySlice) Less(i, j int) bool { return (&keys[i]).CompareTo(&keys[j]) == -1 }

func (keys PublicKeySlice) GetVarSize() int {
	var size int = 0
	for _, k := range keys {
		size += k.Size()
	}
	return helper.GetVarSize(len(keys)) + size
}
