package sc

import (
	"encoding/binary"

	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
)

type Contract struct {
	Script        []byte
	ParameterList []ContractParameterType
	_scriptHash   helper.UInt160
	_address      string
}

// GetAddress is the getter of _address
func (c *Contract) GetAddress() string {
	if c._address == "" {
		c._address = helper.ScriptHashToAddress(c._scriptHash)
	}
	return c._address
}

// GetScriptHash is the getter of _scriptHash
func (c *Contract) GetScriptHash() helper.UInt160 {
	c._scriptHash, _ = helper.UInt160FromBytes(crypto.Hash160(c.Script))
	return c._scriptHash
}

type ByteSlice []byte

func (bs ByteSlice) GetVarSize() int {
	return helper.GetVarSize(len(bs)) + len(bs)*1
}

func (bs ByteSlice) IsSignatureContract() bool {
	if len(bs) != 41 {
		return false
	}
	if bs[0] != byte(PUSHDATA1) ||
		bs[1] != 33 ||
		bs[35] != byte(PUSHNULL) ||
		bs[36] != byte(SYSCALL) ||
		uint(binary.LittleEndian.Uint32(bs[37:])) != ECDsaVerify.ToInteropMethodHash() {
		return false
	}
	return true
}

func (bs ByteSlice) IsMultiSigContract() (bool, int, int) {
	var m, n int = 0, 0
	var i int = 0
	if len(bs) < 43 {
		return false, m, n
	}
	switch bs[i] {
	case byte(PUSHINT8):
		i++
		m = int(bs[i])
		i++
	case byte(PUSHINT16):
		i++
		m = int(binary.LittleEndian.Uint16(bs[i : i+2]))
		i += 2
	case byte(PUSH1), byte(PUSH2), byte(PUSH3), byte(PUSH4),
		byte(PUSH5), byte(PUSH6), byte(PUSH7), byte(PUSH8),
		byte(PUSH9), byte(PUSH10), byte(PUSH11), byte(PUSH12),
		byte(PUSH13), byte(PUSH14), byte(PUSH15), byte(PUSH16):
		m = int(bs[i] - byte(PUSH0))
		i++
	default:
		return false, 0, 0
	}
	if m < 1 || m > 1024 {
		return false, 0, 0
	}
	for bs[i] == byte(PUSHDATA1) {
		if len(bs) <= i+35 {
			return false, 0, 0
		}
		i++
		if bs[i] != 33 {
			return false, 0, 0
		}
		i += 34
		n++
	}
	if n < m || n > 1024 {
		return false, 0, 0
	}
	switch bs[i] {
	case byte(PUSHINT8):
		i++
		if n != int(bs[i]) {
			return false, 0, 0
		}
		i++
	case byte(PUSHINT16):
		if len(bs) < i+3 {
			return false, 0, 0
		} else if i++; n != int(binary.LittleEndian.Uint16(bs[i:i+2])) {
			return false, 0, 0
		}
		i += 2
	case byte(PUSH1), byte(PUSH2), byte(PUSH3), byte(PUSH4),
		byte(PUSH5), byte(PUSH6), byte(PUSH7), byte(PUSH8),
		byte(PUSH9), byte(PUSH10), byte(PUSH11), byte(PUSH12),
		byte(PUSH13), byte(PUSH14), byte(PUSH15), byte(PUSH16):
		if n != int(bs[i]-byte(PUSH0)) {
			return false, 0, 0
		}
		i++
	default:
		return false, 0, 0
	}
	if bs[i] != byte(PUSHNULL) {
		return false, 0, 0
	}
	i++
	if bs[i] != byte(SYSCALL) {
		return false, 0, 0
	}
	i++
	if len(bs) != i+4 {
		return false, 0, 0
	}
	if uint(binary.LittleEndian.Uint32(bs[i:])) != ECDsaCheckMultiSig.ToInteropMethodHash() {
		return false, 0, 0
	}
	return true, m, n
}
