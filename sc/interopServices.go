package sc

import (
	"encoding/binary"

	"github.com/joeqian10/neo3-gogogo/crypto"
)

type InteropService string

var (
	// -----Contract-----
	System_Contract_Call         InteropService = "System.Contract.Call"
	System_Contract_CallNative   InteropService = "System.Contract.CallNative"
	System_Contract_IsStandard   InteropService = "System.Contract.IsStandard"
	System_Contract_GetCallFlags InteropService = "System.Contract.GetCallFlags"
	/// Calculate corresponding account scripthash for given public key
	/// Warning: check first that input public key is valid, before creating the script.
	System_Contract_CreateStandardAccount InteropService = "System.Contract.CreateStandardAccount"
	System_Contract_CreateMultisigAccount InteropService = "System.Contract.CreateMultisigAccount"
	System_Contract_NativeOnPersist       InteropService = "System.Contract.NativeOnPersist"
	System_Contract_NativePostPersist     InteropService = "System.Contract.NativePostPersist"

	// -----Crypto-----
	System_Crypto_CheckSig      InteropService = "System.Crypto.CheckSig"
	System_Crypto_CheckMultisig InteropService = "System.Crypto.CheckMultisig"
)

// ToInteropMethodHash converts a method name to 32 bytes hash
func (p InteropService) ToInteropMethodHash() uint {
	temp := crypto.Sha256([]byte(p))
	u := binary.LittleEndian.Uint32(temp)
	return uint(u)
}
