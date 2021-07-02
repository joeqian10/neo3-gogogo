package helper

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"math/big"
)

// bytes to hex string
func BytesToHex(b []byte) string {
	return hex.EncodeToString(b)
}

// Simple hex string to bytes
func HexToBytes(s string) (b []byte) {
	b, _ = hex.DecodeString(s)
	return b
}

// ConcatBytes ...
func ConcatBytes(b1 []byte, b2 []byte) []byte {
	var buffer bytes.Buffer //Buffer: length changeable, writable, readable
	buffer.Write(b1)
	buffer.Write(b2)
	return buffer.Bytes()
}

// ReverseBytes without change original slice
func ReverseBytes(data []byte) []byte {
	b := make([]byte, len(data))
	copy(b, data)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return b
}

// ReverseString, "abcd" to "dcba"
func ReverseString(input string) string {
	return string(ReverseBytes([]byte(input)))
}

// UInt16ToBytes converts uint16 to byte array
func UInt16ToBytes(n uint16) []byte {
	var buff = make([]byte, 2)
	binary.LittleEndian.PutUint16(buff, n)
	return buff
}

// UInt32ToBytes converts uint32 to byte array
func UInt32ToBytes(n uint32) []byte {
	var buff = make([]byte, 4)
	binary.LittleEndian.PutUint32(buff, n)
	return buff
}

// UInt64ToBytes converts uint32 to byte array
func UInt64ToBytes(n uint64) []byte {
	var buff = make([]byte, 8)
	binary.LittleEndian.PutUint64(buff, n)
	return buff
}

// IntToBytes ...
func IntToBytes(n int) []byte {
	var buff = make([]byte, 4)
	binary.LittleEndian.PutUint32(buff, uint32(n))
	return buff
}

// Int16ToBytes ...
func Int16ToBytes(n int16) []byte {
	var buff = make([]byte, 2)
	binary.LittleEndian.PutUint16(buff, uint16(n))
	return buff
}

// Int64ToBytes ...
func Int64ToBytes(n int64) []byte {
	var buff = make([]byte, 8)
	binary.LittleEndian.PutUint64(buff, uint64(n))
	return buff
}

func BytesToUInt64(bs []byte) uint64 {
	bs = PadRight(bs, 8, false)
	return binary.LittleEndian.Uint64(bs)
}

func BytesToUInt32(bs []byte) uint32 {
	bs = PadRight(bs, 4, false)
	return binary.LittleEndian.Uint32(bs)
}

func PadRight(data []byte, length int, negative bool) []byte {
	if len(data) >= length {
		return data[:length] // return the most left bytes of length
	}
	newData := make([]byte, length)
	for i := 0; i < len(data); i++ {
		newData[i] = data[i]
	}
	if negative {
		for i := len(data); i < length; i++ {
			newData[i] = 0xff
		}
	}
	return newData
}

func Abs(x int64) int64 {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetVarSize(value int) int {
	if value < 0xFD {
		return 1 // sizeof(byte)
	} else if value <= 0xFFFF {
		return 1 + 2 // sizeof(byte) + sizeof(ushort)
	} else {
		return 1 + 4 // sizeof(byte) + sizeof(uint)
	}
}

func BigIntToNeoBytes(data *big.Int) []byte {
	bs := data.Bytes()
	if len(bs) == 0 {
		return []byte{}
	}
	// golang big.Int use big-endian
	bs = ReverseBytes(bs)
	// bs now is little-endian
	if data.Sign() < 0 {
		for i, b := range bs {
			bs[i] = ^b
		}
		for i := 0; i < len(bs); i++ {
			if bs[i] == 255 {
				bs[i] = 0
			} else {
				bs[i] += 1
				break
			}
		}
		if bs[len(bs)-1] < 128 {
			bs = append(bs, 255)
		}
	} else {
		if bs[len(bs)-1] >= 128 {
			bs = append(bs, 0)
		}
	}
	return bs
}

var bigOne = big.NewInt(1)

func BigIntFromNeoBytes(ba []byte) *big.Int {
	res := big.NewInt(0)
	l := len(ba)
	if l == 0 {
		return res
	}

	bs := make([]byte, 0, l)
	bs = append(bs, ba...)
	bs = ReverseBytes(bs)

	if bs[0]>>7 == 1 {
		for i, b := range bs {
			bs[i] = ^b
		}

		temp := big.NewInt(0)
		temp.SetBytes(bs)
		temp.Add(temp, bigOne)
		bs = temp.Bytes()
		res.SetBytes(bs)
		return res.Neg(res)
	}

	res.SetBytes(bs)
	return res
}

// a>b, returns 1
// a==b, returns 0
// a<b, returns -1
func CompareTo(a, b uint64) int {
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}

func XOR(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("cannot XOR unequal length arrays")
	}
	dst := make([]byte, len(a))
	for i := 0; i < len(dst); i++ {
		dst[i] = a[i] ^ b[i]
	}
	return dst
}
