package sc

import (
	"encoding/binary"

	"github.com/joeqian10/neo3-gogogo/crypto"
)

type InteropService string

var (
	// -----Contract-----
	Call                  InteropService = "System.Contract.Call"
	CallNative            InteropService = "System.Contract.CallNative"
	IsStandard            InteropService = "System.Contract.IsStandard"
	GetCallFlags          InteropService = "System.Contract.GetCallFlags"
	CreateStandardAccount InteropService = "System.Contract.CreateStandardAccount" // Calculate corresponding account scripthash for given public key. Warning: check first that input public key is valid, before creating the script.
	NativeOnPersist       InteropService = "System.Contract.NativeOnPersist"
	NativePostPersist     InteropService = "System.Contract.NativePostPersist"

	// -----Crypto-----
	RIPEMD160                       InteropService = "Neo.Crypto.RIPEMD160"
	SHA256                          InteropService = "Neo.Crypto.SHA256"
	VerifyWithECDsaSecp256r1        InteropService = "Neo.Crypto.VerifyWithECDsaSecp256r1"
	VerifyWithECDsaSecp256k1        InteropService = "Neo.Crypto.VerifyWithECDsaSecp256k1"
	CheckMultisigWithECDsaSecp256r1 InteropService = "Neo.Crypto.CheckMultisigWithECDsaSecp256r1"
	CheckMultisigWithECDsaSecp256k1 InteropService = "Neo.Crypto.CheckMultisigWithECDsaSecp256k1"
)

// ToInteropMethodHash converts a method name to 32 bytes hash
func (p InteropService) ToInteropMethodHash() uint {
	temp:= crypto.Sha256([]byte(p))
	u := binary.LittleEndian.Uint32(temp)
	return uint(u)
}
