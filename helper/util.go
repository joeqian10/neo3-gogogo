package helper

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
)

// bytes to hex string
func BytesToHex(b []byte) string {
	return hex.EncodeToString(b)
}

// Simple hex string to bytes
func HexTobytes(hexstring string) (b []byte) {
	b, _ = hex.DecodeString(hexstring)
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

func ScriptHashToAddress(scriptHash UInt160) string {
	var addressVersion byte = 0x35
	data := append([]byte{addressVersion}, scriptHash.Bytes()...)
	return crypto.Base58CheckEncode(data)
}

func AddressToScriptHash(address string) (UInt160, error) {
	data, err := crypto.Base58CheckDecode(address)
	var u UInt160
	if err != nil {
		return u, err
	}
	if data == nil || len(data) != 21 || data[0] != 0x35 {
		return u, fmt.Errorf("invalid address string")
	}
	return UInt160FromBytes(data[1:])
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

// Int64ToBytes ...
func Int64ToBytes(n int64) []byte {
	var buff = make([]byte, 8)
	binary.LittleEndian.PutUint64(buff, uint64(n))
	return buff
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
