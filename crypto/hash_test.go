package crypto

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSha256(t *testing.T) {
	value := []byte("hello world")
	result := Sha256(value)
	assert.Equal(t, "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9", hex.EncodeToString(result))
}

func TestHash160(t *testing.T) {
	value := []byte("hello world")
	result := Hash160(value)
	assert.Equal(t, "d7d5ee7824ff93f94c3055af9382c86c68b5ca92", hex.EncodeToString(result))
}

func TestHash256(t *testing.T) {
	value := []byte("hello world")
	result := Hash256(value)
	assert.Equal(t, "bc62d4b80d9e36da29c16c5d4d9f11731f36052c72401a76c23c0fb5a9b74423", hex.EncodeToString(result))
}
