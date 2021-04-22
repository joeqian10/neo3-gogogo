package sc

import (
	"encoding/binary"
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"sort"
)

type Contract struct {
	Script        []byte
	ParameterList []ContractParameterType
	scriptHash    *helper.UInt160
}

func CreateContract(parameterList []ContractParameterType, redeemScript []byte) *Contract {
	return &Contract{
		Script:        redeemScript,
		ParameterList: parameterList,
	}
}

/// Construct special Contract with empty Script, will get the Script with scriptHash from blockchain when doing the Verify
/// verification = snapshot.Contracts.TryGet(hashes[i])?.Script;
func CreateContractWithScriptHash(scriptHash *helper.UInt160, parameterList []ContractParameterType) *Contract {
	return &Contract{
		Script:        []byte{},
		ParameterList: parameterList,
		scriptHash:    scriptHash,
	}
}

// GetScriptHash is the getter of _scriptHash
func (c *Contract) GetScriptHash() *helper.UInt160 {
	if c.scriptHash == nil {
		c.scriptHash = helper.UInt160FromBytes(crypto.Hash160(c.Script))
	}
	return c.scriptHash
}

// create signature check script
func CreateSignatureRedeemScript(p *crypto.ECPoint) ([]byte, error) {
	sb := NewScriptBuilder()
	sb.EmitPushBytes(p.EncodePoint(true))
	sb.EmitSysCall(System_Crypto_CheckSig.ToInteropMethodHash())
	b, err := sb.ToArray()
	if err != nil {
		return nil, err
	}
	return b, nil
}

// CreateSignatureContract
func CreateSignatureContract(publicKey *crypto.ECPoint) (*Contract, error) {
	script, err := CreateSignatureRedeemScript(publicKey)
	if err != nil {
		return nil, err
	}
	return &Contract{
		Script:        script,
		ParameterList: []ContractParameterType{Signature},
	}, nil
}

// create multi-signature check script
func CreateMultiSigRedeemScript(m int, ps []crypto.ECPoint) ([]byte, error) {
	if !(m >= 1 && m <= len(ps) && len(ps) <= 1024) {
		return nil, fmt.Errorf("argument exception: %v, %v", m, len(ps))
	}
	sb := NewScriptBuilder()
	sb.EmitPushInteger(m)
	// sort public keys
	pubKeys := crypto.PublicKeySlice(ps)
	sort.Sort(pubKeys)

	for _, p := range pubKeys {
		sb.EmitPushBytes(p.EncodePoint(true))
	}
	sb.EmitPushInteger(pubKeys.Len())
	sb.EmitSysCall(System_Crypto_CheckMultisig.ToInteropMethodHash())
	b, err := sb.ToArray()
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Create Multi-Signature Contract
func CreateMultiSigContract(m int, publicKeys []crypto.ECPoint) (*Contract, error) {
	script, err := CreateMultiSigRedeemScript(m, publicKeys)
	if err != nil {
		return nil, err
	}
	parameters := make([]ContractParameterType, m)
	for i := 0; i < m; i++ {
		parameters[i] = Signature
	}

	return &Contract{
		Script:        script,
		ParameterList: parameters,
	}, nil
}

type ByteSlice []byte

func (bs ByteSlice) GetVarSize() int {
	return helper.GetVarSize(len(bs)) + len(bs)*1
}

func (bs ByteSlice) IsStandardContract() bool {
	return bs.IsSignatureContract() || bs.IsMultiSigContract()
}

func (bs ByteSlice) IsSignatureContract() bool {
	return IsSignatureContract(bs)
}

func (bs ByteSlice) IsMultiSigContract() bool {
	b, _, _, _ := IsMultiSigContract(bs)
	return b
}

func (bs ByteSlice) IsMultiSigContractWithCounts() (bool, int, int) {
	b, m, n, _ := IsMultiSigContract(bs)
	return b, m, n
}

func (bs ByteSlice) IsMultiSigContractWithPoints() (bool, int, []crypto.ECPoint) {
	b, m, _, points := IsMultiSigContract(bs)
	return b, m, points
}

func IsSignatureContract(script []byte) bool {
	if len(script) != 40 {
		return false
	}
	if script[0] != byte(PUSHDATA1) ||
		script[1] != 33 ||
		script[35] != byte(SYSCALL) ||
		uint(binary.LittleEndian.Uint32(script[36:])) != System_Crypto_CheckSig.ToInteropMethodHash() {
		return false
	}
	return true
}

func IsMultiSigContract(script []byte) (bool, int, int, []crypto.ECPoint) {
	var m, n int = 0, 0
	var i int = 0
	if len(script) < 42 {
		return false, m, n, nil
	}
	switch script[i] {
	case byte(PUSHINT8):
		i++
		m = int(script[i])
		i++
		break
	case byte(PUSHINT16):
		i++
		m = int(binary.LittleEndian.Uint16(script[i : i+2]))
		i += 2
		break
	case byte(PUSH1), byte(PUSH2), byte(PUSH3), byte(PUSH4),
		byte(PUSH5), byte(PUSH6), byte(PUSH7), byte(PUSH8),
		byte(PUSH9), byte(PUSH10), byte(PUSH11), byte(PUSH12),
		byte(PUSH13), byte(PUSH14), byte(PUSH15), byte(PUSH16):
		m = int(script[i] - byte(PUSH0))
		i++
		break
	default:
		return false, 0, 0, nil
	}
	if m < 1 || m > 1024 {
		return false, 0, 0, nil
	}

	points := make([]crypto.ECPoint, 0)
	for script[i] == byte(PUSHDATA1) {
		if len(script) <= i+35 {
			return false, 0, 0, nil
		}
		i++
		if script[i] != 33 {
			return false, 0, 0, nil
		}
		// add point
		point, _ := crypto.DecodePoint(script[i+1:i+34], &crypto.P256)
		points = append(points, *point)
		i += 34
		n++
	}
	if n < m || n > 1024 {
		return false, 0, 0, nil
	}
	switch script[i] {
	case byte(PUSHINT8):
		if len(script) <= i+1 {
			return false, 0, 0, nil
		}
		i++
		if n != int(script[i]) {
			return false, 0, 0, nil
		}
		i++
		break
	case byte(PUSHINT16):
		if len(script) < i+3 {
			return false, 0, 0, nil
		} else if i++; n != int(binary.LittleEndian.Uint16(script[i:i+2])) {
			return false, 0, 0, nil
		}
		i += 2
		break
	case byte(PUSH1), byte(PUSH2), byte(PUSH3), byte(PUSH4),
		byte(PUSH5), byte(PUSH6), byte(PUSH7), byte(PUSH8),
		byte(PUSH9), byte(PUSH10), byte(PUSH11), byte(PUSH12),
		byte(PUSH13), byte(PUSH14), byte(PUSH15), byte(PUSH16):
		if n != int(script[i]-byte(PUSH0)) {
			return false, 0, 0, nil
		}
		i++
		break
	default:
		return false, 0, 0, nil
	}
	if len(script) != i+5 {
		return false, 0, 0, nil
	}
	if script[i] != byte(SYSCALL) {
		return false, 0, 0, nil
	}
	i++
	if uint(binary.LittleEndian.Uint32(script[i:])) != System_Crypto_CheckMultisig.ToInteropMethodHash() {
		return false, 0, 0, nil
	}
	return true, m, n, points
}
