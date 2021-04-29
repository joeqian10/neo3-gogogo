package sc

import (
	"fmt"
	"strings"
)

type ContractParameterType byte

const (
	Any ContractParameterType = 0x00

	Boolean ContractParameterType = 0x10
	Integer ContractParameterType = 0x11
	ByteArray ContractParameterType = 0x12
	String    ContractParameterType = 0x13
	Hash160   ContractParameterType = 0x14
	Hash256   ContractParameterType = 0x15
	PublicKey ContractParameterType = 0x16
	Signature ContractParameterType = 0x17

	Array ContractParameterType = 0x20
	Map   ContractParameterType = 0x22

	InteropInterface ContractParameterType = 0x30

	Void ContractParameterType = 0xff
)

type ContractParameter struct {
	Type  ContractParameterType
	Value interface{}
}

func NewContractParameterTypeFromString(s string) (ContractParameterType, error) {
	t := strings.ToLower(s)
	switch t {
	case "any":
		return Any, nil
	case "boolean":
		return Boolean, nil
	case "integer":
		return Integer, nil
	case "bytearray":
		return ByteArray, nil
	case "string":
		return String, nil
	case "hash160":
		return Hash160, nil
	case "hash256":
		return Hash256, nil
	case "publickey":
		return PublicKey, nil
	case "signature":
		return Signature, nil
	case "array":
		return Array, nil
	case "map":
		return Map, nil
	case "interopinterface":
		return InteropInterface, nil
	case "void":
		return Void, nil
	default:
		return Void, fmt.Errorf("not supported string")
	}
}

func (cpt ContractParameterType) String() string {
	var s string
	switch byte(cpt) {
	case 0x00:
		s = "Any"
	case 0x10:
		s = "Boolean"
	case 0x11:
		s = "Integer"
	case 0x12:
		s = "ByteArray"
	case 0x13:
		s = "String"
	case 0x14:
		s = "Hash160"
	case 0x15:
		s = "Hash256"
	case 0x16:
		s = "PublicKey"
	case 0x17:
		s = "Signature"
	case 0x20:
		s = "Array"
	case 0x21:
		s = "Map"
	case 0x30:
		s = "InteropInterface"
	case 0xff:
		s = "Void"
	default:
		s = ""
	}
	return s
}
