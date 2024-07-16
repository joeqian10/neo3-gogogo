package keys

import (
	"crypto/elliptic"
	"fmt"
	"math/big"

	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/sc"
)

// PublicKeyToScriptHash returns a NEO-specific hash of the public key.
func PublicKeyToScriptHash(p *crypto.ECPoint) *helper.UInt160 {
	b, _ := sc.CreateSignatureRedeemScript(p)
	hash := crypto.Hash160(b)
	return helper.UInt160FromBytes(hash)
}

// PublicKeyToAddress returns a base58-encoded NEO-specific address based on the key hash.
func PublicKeyToAddress(p *crypto.ECPoint, version byte) string {
	return crypto.ScriptHashToAddress(PublicKeyToScriptHash(p), version)
}

// RecoverPubKeyFromSigOnSecp256r1 recover a point on Secp256r1 from a signature
func RecoverPubKeyFromSigOnSecp256r1(message, signature []byte) ([]*crypto.ECPoint, error) {
	if len(signature) != 64 {
		return nil, fmt.Errorf("invalid neo signature length")
	}
	return RecoverPubKeyFromSig(crypto.P256, message, signature)
}

// RecoverPubKeyFromSig recovers a point on a given curve from a signature
func RecoverPubKeyFromSig(curve elliptic.Curve, message, signature []byte) ([]*crypto.ECPoint, error) {
	hash := crypto.Sha256(message)
	//curve := crypto.P256
	bitLen := (curve.Params().BitSize + 7) / 8
	R := new(big.Int).SetBytes(signature[0:bitLen])
	S := new(big.Int).SetBytes(signature[bitLen:])
	sig := &Signature{
		R: R,
		S: S,
	}
	return recoverKeyFromSignature(crypto.P256, sig, hash, true)
}

// recoverKeyFromSignature recovers a public key from the signature "sig" on the
// given message hash "msg". Based on the algorithm found in section 5.1.5 of
// SEC 1 Ver 2.0, page 47-48 (53 and 54 in the pdf). This performs the details
// in the inner loop in Step 1.
func recoverKeyFromSignature(curve elliptic.Curve, sig *Signature, hash []byte, doChecks bool) ([]*crypto.ECPoint, error) {

	// calculate h = (Q + 1 + 2 * sqrt(Q)) / N
	h := new(big.Int).Add(curve.Params().P, big.NewInt(1))
	h.Add(h, new(big.Int).Mul(big.NewInt(2), new(big.Int).Sqrt(curve.Params().P)))
	h.Div(h, curve.Params().N)

	result := []*crypto.ECPoint{}

	for i := 0; uint64(i) <= h.Uint64(); i++ {

		// 1.1 x = (n * i) + r
		Rx := new(big.Int).Mul(curve.Params().N, new(big.Int).SetInt64(int64(i)))
		Rx.Add(Rx, sig.R)
		if Rx.Cmp(curve.Params().P) != -1 {
			continue
		}

		// convert 02<Rx> to point R. (step 1.2 and 1.3). If we are on an odd
		// iteration then 1.6 will be done with -R, so we calculate the other
		// term when uncompressing the point.
		//Ry, err := decodeCompressedY(curve, Rx, uint(i))
		for j := uint(0); j <= 1; j++ {
			Ry, err := decodeCompressedY(curve, Rx, j)
			if err != nil {
				continue
			}

			// 1.4 Check n*R is point at infinity
			if doChecks {
				nRx, nRy := curve.ScalarMult(Rx, Ry, curve.Params().N.Bytes())
				if nRx.Sign() != 0 || nRy.Sign() != 0 {
					continue
				}
			}

			// 1.5 calculate e from message using the same algorithm as ecdsa
			// signature calculation.
			e := hashToInt(hash, curve)

			// Step 1.6.1:
			// We calculate the two terms sR and eG separately multiplied by the
			// inverse of r (from the signature). We then add them to calculate
			// Q = r^-1(sR-eG)
			invr := new(big.Int).ModInverse(sig.R, curve.Params().N)

			// first term.
			invrS := new(big.Int).Mul(invr, sig.S)
			invrS.Mod(invrS, curve.Params().N)
			sRx, sRy := curve.ScalarMult(Rx, Ry, invrS.Bytes())

			// second term.
			e.Neg(e)
			e.Mod(e, curve.Params().N)
			e.Mul(e, invr)
			e.Mod(e, curve.Params().N)
			minuseGx, minuseGy := curve.ScalarBaseMult(e.Bytes())

			// TODO: this would be faster if we did a mult and add in one
			// step to prevent the jacobian conversion back and forth.
			Qx, Qy := curve.Add(sRx, sRy, minuseGx, minuseGy)
			Q, err := crypto.CreateECPoint(Qx, Qy, &curve)
			if err != nil {
				return nil, err
			}
			result = append(result, Q)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("invalid signature")
	}
	return result, nil
}

// hashToInt converts a hash value to an integer. There is some disagreement
// about how this is done. [NSA] suggests that this is done in the obvious
// manner, but [SECG] truncates the hash to the bit-length of the curve order
// first. We follow [SECG] because that's what OpenSSL does. Additionally,
// OpenSSL right shifts excess bits from the number if the hash is too large, and we mirror that too.
// This is borrowed from crypto/ecdsa.
func hashToInt(hash []byte, c elliptic.Curve) *big.Int {
	orderBits := c.Params().N.BitLen()
	orderBytes := (orderBits + 7) / 8
	if len(hash) > orderBytes {
		hash = hash[:orderBytes]
	}

	ret := new(big.Int).SetBytes(hash)
	excess := len(hash)*8 - orderBits
	if excess > 0 {
		ret.Rsh(ret, uint(excess))
	}
	return ret
}

// decodeCompressedY performs decompression of Y coordinate for given X and Y's least significant bit
func decodeCompressedY(curve elliptic.Curve, x *big.Int, ylsb uint) (*big.Int, error) {

	cp := curve.Params()
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
