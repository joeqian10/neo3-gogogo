package sc

import (
	"encoding/binary"

	"github.com/joeqian10/neo3-gogogo/crypto"
)

type InteropService string

var (
	// Crypto
	ECDsaVerify        InteropService = "Neo.Crypto.ECDsaVerify"
	ECDsaCheckMultiSig InteropService = "Neo.Crypto.ECDsaCheckMultiSig"
)

// ToInteropMethodHash converts a method name to 32 bytes hash
func (p InteropService) ToInteropMethodHash() uint {
	u := binary.LittleEndian.Uint32(crypto.Sha256([]byte(p)))
	return uint(u)
}
