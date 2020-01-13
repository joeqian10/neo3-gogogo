package keys

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"io"
	"math/big"
	"sort"

	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/sc"
)

// PublicKeys is a list of public keys.
type PublicKeys []*PublicKey

func (keys PublicKeys) Len() int           { return len(keys) }
func (keys PublicKeys) Swap(i, j int)      { keys[i], keys[j] = keys[j], keys[i] }
func (keys PublicKeys) Less(i, j int) bool { return keys[i].Compare(keys[j]) == -1 }

// PublicKey represents a public key and provides a high level
// API around the X/Y point.
type PublicKey struct {
	X *big.Int
	Y *big.Int
}

// NewPublicKey return a public key created from the given []byte.
func NewPublicKey(data []byte) (*PublicKey, error) {
	pubKey := new(PublicKey)
	if err := pubKey.Deserialize(bytes.NewReader(data)); err != nil {
		return nil, err
	}
	return pubKey, nil
}

// NewPublicKeyFromString return a public key created from the
// given hex string.
func NewPublicKeyFromString(s string) (*PublicKey, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return NewPublicKey(b)
}

// Bytes returns the byte array representation of the public key.
func (p *PublicKey) ecdsa() *ecdsa.PublicKey {
	pubKey := ecdsa.PublicKey{X: p.X, Y: p.Y}
	pubKey.Curve = elliptic.P256()
	return &pubKey
}

// Bytes returns the byte array representation of the public key.
func (p *PublicKey) EncodeCompression() []byte {
	if p.isInfinity() {
		return []byte{0x00}
	}

	var (
		x       = p.X.Bytes()
		paddedX = append(bytes.Repeat([]byte{0x00}, 32-len(x)), x...)
		prefix  = byte(0x03)
	)

	if p.Y.Bit(0) == 0 {
		prefix = byte(0x02)
	}

	return append([]byte{prefix}, paddedX...)
}

// decodeCompressedY performs decompression of Y coordinate for given X and Y's least significant bit
func decodeCompressedY(x *big.Int, ylsb uint) (*big.Int, error) {
	c := elliptic.P256()
	cp := c.Params()
	three := big.NewInt(3)
	/* y**2 = x**3 + a*x + b  % p */
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
func (p *PublicKey) Deserialize(r io.Reader) error {
	var prefix uint8
	var x, y *big.Int
	var err error

	if err = binary.Read(r, binary.LittleEndian, &prefix); err != nil {
		return err
	}

	// Infinity
	switch prefix {
	case 0x00:
		// noop, initialized to nil
		return nil
	case 0x02, 0x03:
		// Compressed public keys
		xbytes := make([]byte, 32)
		if _, err := io.ReadFull(r, xbytes); err != nil {
			return err
		}
		x = new(big.Int).SetBytes(xbytes)
		ylsb := uint(prefix & 0x1)
		y, err = decodeCompressedY(x, ylsb)
		if err != nil {
			return err
		}
	case 0x04:
		xbytes := make([]byte, 32)
		ybytes := make([]byte, 32)
		if _, err = io.ReadFull(r, xbytes); err != nil {
			return err
		}
		if _, err = io.ReadFull(r, ybytes); err != nil {
			return err
		}
		x = new(big.Int).SetBytes(xbytes)
		y = new(big.Int).SetBytes(ybytes)
	default:
		return fmt.Errorf("invalid prefix %d", prefix)
	}
	c := elliptic.P256()
	cp := c.Params()
	if !c.IsOnCurve(x, y) {
		return fmt.Errorf("encoded point is not on the P256 curve")
	}
	if x.Cmp(cp.P) >= 0 || y.Cmp(cp.P) >= 0 {
		return fmt.Errorf("encoded point is not correct (X or Y is bigger than P")
	}
	p.X, p.Y = x, y

	return nil
}

// Serialize encodes a PublicKey to the given io.Writer.
func (p *PublicKey) Serialize(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, p.EncodeCompression())
}

// Signature returns a NEO-specific hash of the key.
func (p *PublicKey) ScriptHash() helper.UInt160 {
	b := CreateSignatureRedeemScript(p)
	hash := crypto.Hash160(b)
	hash160, _ := helper.UInt160FromBytes(hash)
	return hash160
}

// Address returns a base58-encoded NEO-specific address based on the key hash.
func (p *PublicKey) Address() string {
	return helper.ScriptHashToAddress(p.ScriptHash())
}

// isInfinity checks if point P is infinity on EllipticCurve ec.
func (p *PublicKey) isInfinity() bool {
	return p.X == nil && p.Y == nil
}

// String implements the Stringer interface.
func (p *PublicKey) String() string {
	return helper.BytesToHex(p.EncodeCompression())
}

// Compare q to q
func (p *PublicKey) Compare(q *PublicKey) int {
	xLess := p.X.Cmp(q.X)
	if xLess != 0 {
		return xLess
	}
	return p.Y.Cmp(p.Y)
}

// create signature check script
func CreateSignatureRedeemScript(p *PublicKey) []byte {
	builder := sc.NewScriptBuilder()
	_ = builder.EmitPushBytes(p.EncodeCompression())
	_ = builder.Emit(sc.PUSHNULL)
	_ = builder.EmitSysCall(sc.ECDsaVerify.ToInteropMethodHash())
	return builder.ToArray()
}

// create multi-signature check script
func CreateMultiSigRedeemScript(m int, ps ...*PublicKey) ([]byte, error) {
	if !(m >= 1 && m < len(ps) && len(ps) <= 1024) {
		return nil, fmt.Errorf("argument exception: %v,%v", m, len(ps))
	}

	builder := sc.NewScriptBuilder()
	err := builder.EmitPushInt(m)
	if err != nil {
		return nil, err
	}
	pubKeys := PublicKeys(ps)
	sort.Sort(pubKeys)
	for _, p := range pubKeys {
		err = builder.EmitPushBytes(p.EncodeCompression())
		if err != nil {
			return nil, err
		}
	}
	err = builder.EmitPushInt(pubKeys.Len())
	if err != nil {
		return nil, err
	}

	err = builder.Emit(sc.PUSHNULL)
	err = builder.EmitSysCall(sc.ECDsaCheckMultiSig.ToInteropMethodHash())
	if err != nil {
		return nil, err
	}
	return builder.ToArray(), nil
}
