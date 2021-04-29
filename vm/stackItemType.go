package vm

import (
	"fmt"
	"strings"
)

type StackItemType byte

const (
	Any              StackItemType = 0x00 // Represents any type.
	Pointer          StackItemType = 0x10 // Represents a code pointer.
	Boolean          StackItemType = 0x20 // Represents the boolean type.
	Integer          StackItemType = 0x21 // Represents an integer.
	ByteString       StackItemType = 0x28 // Represents an immutable memory block.
	Buffer           StackItemType = 0x30 // Represents a memory block that can be used for reading and writing.
	Array            StackItemType = 0x40 // Represents an array of a complex object.
	Struct           StackItemType = 0x41 // Represents a structure.
	Map              StackItemType = 0x48 // Represents an ordered collection of key-value pairs.
	InteropInterface StackItemType = 0x60 // Represents an interface used to interoperate with the outside of the VM.
)

func NewStackItemTypeFromString(s string) (StackItemType, error) {
	t := strings.ToLower(s)
	switch t {
	case "any":
		return Any, nil
	case "pointer":
		return Pointer, nil
	case "boolean":
		return Boolean, nil
	case "integer":
		return Integer, nil
	case "bytestring":
		return ByteString, nil
	case "buffer":
		return Buffer, nil
	case "array":
		return Array, nil
	case "struct":
		return Struct, nil
	case "map":
		return Map, nil
	case "interopinterface":
		return InteropInterface, nil
	default:
		return Any, fmt.Errorf("not supported string")
	}
}

func (sit StackItemType) String() string {
	var s string
	switch byte(sit) {
	case 0x00:
		s = "Any"
	case 0x10:
		s = "Pointer"
	case 0x20:
		s = "Boolean"
	case 0x21:
		s = "Integer"
	case 0x28:
		s = "ByteString"
	case 0x30:
		s = "Buffer"
	case 0x40:
		s = "Array"
	case 0x41:
		s = "Struct"
	case 0x48:
		s = "Map"
	case 0x60:
		s = "InteropInterface"
	default:
		s = ""
	}
	return s
}
