package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

// Sha256 gets the SHA-256 hash value of b
func Sha256(b []byte) []byte {
	sha := sha256.New()
	sha.Write(b)
	return sha.Sum(nil)
}

// Hash256 gets the twice SHA-256 hash value of ba
func Hash256(ba []byte) []byte {
	sha := sha256.New()
	sha.Write(ba)
	ba = sha.Sum(nil)
	sha.Reset()
	sha.Write(ba)
	return sha.Sum(nil)
}

// Hash160 first calculate SHA-256 hash result of ba, then RIPEMD-160 hash of the result
func Hash160(ba []byte) []byte {
	sha := sha256.New()
	sha.Write(ba)
	ba = sha.Sum(nil)
	ripemd := ripemd160.New()
	ripemd.Write(ba)
	return ripemd.Sum(nil)
}
