package crypto

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
)

const BASE58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const PREFIX rune = '1'

// Encode ... ref: neo/Cryptography/Base58.cs
func Encode(input []byte) string {
	b := []byte{0}
	//tmp := helper.ConcatBytes(b, input) // input is big-endian
	tmp := append(b, input[:]...)
	x := new(big.Int).SetBytes(tmp)
	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)
	encoded := ""

	for x.Cmp(zero) > 0 {
		x.QuoRem(x, m, r)
		encoded = string(BASE58[r.Int64()]) + encoded
	}
	// leading zeros
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			encoded = string(BASE58[0]) + encoded
		} else {
			break
		}
	}

	return encoded
}

// Decode...
func Decode(input string) ([]byte, error) {
	var (
		startIndex = 0
		zero       = 0
	)
	for i, c := range input {
		if c == PREFIX {
			zero++
		} else {
			startIndex = i
			break
		}
	}
	bi := big.NewInt(0)
	base := big.NewInt(58)
	for _, c := range input[startIndex:] {
		index := strings.IndexRune(BASE58, c)
		if index == -1 {
			return nil, fmt.Errorf(
				"invalid character '%c' when decoding this base58 string: '%s'", c, input,
			)
		}
		bi.Mul(bi, base)
		bi.Add(bi, big.NewInt(int64(index)))
	}
	ba := bi.Bytes() // ba is big-endian
	// add leading zeros
	i := 0
	for i < len(input) && input[i] == '1' {
		i++
	}
	// strip Sign Byte
	stripSignByte := 0
	if len(ba) > 0 && ba[0] == 0 && ba[1] >= 0x80 {
		stripSignByte = 1
	}

	r := make([]byte, len(ba)-stripSignByte+i)
	copy(r[i:], ba[stripSignByte:])
	return r, nil
}

// Base58CheckEncode ...
func Base58CheckEncode(input []byte) string {
	hash := Hash256(input)
	value := append(input, hash[:4]...)
	return Encode(value)
}

// Base58CheckDecode ...
func Base58CheckDecode(input string) ([]byte, error) {
	ba, err := Decode(input)
	if err != nil {
		return nil, err
	}
	if len(ba) < 4 {
		return nil, fmt.Errorf("invalid base58 check string: missing checksum")
	}

	checkSum := Hash256(ba[:len(ba)-4])
	if bytes.Compare(checkSum[0:4], ba[len(ba)-4:]) != 0 {
		return nil, fmt.Errorf("invalid base58 check string: invalid checksum")
	}

	return ba[:len(ba)-4], nil
}
