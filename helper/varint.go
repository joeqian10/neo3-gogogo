package helper

import (
	"encoding/binary"
	"fmt"
)

// VarInt
type VarInt struct {
	Value uint64
}

// Length returns the serialized bytes length of the var int
func (v VarInt) Length() int {
	if v.Value < 0xfd {
		return 1
	}
	if v.Value <= 0xffff {
		return 3
	}
	if v.Value <= 0xffffffff {
		return 5
	}
	return 9
}

// Bytes returns the serialized bytes of the var int
func (v VarInt) Bytes() []byte {
	if v.Value < 0xfd {
		ret := make([]byte, 1)
		ret[0] = byte(v.Value)
		return ret
	}
	if v.Value <= 0xffff {
		ret := make([]byte, 3)
		ret[0] = 0xfd
		binary.LittleEndian.PutUint16(ret[1:], uint16(v.Value))
		return ret
	}
	if v.Value <= 0xffffffff {
		ret := make([]byte, 5)
		ret[0] = 0xfe
		binary.LittleEndian.PutUint32(ret[1:], uint32(v.Value))
		return ret
	}
	ret := make([]byte, 9)
	ret[0] = 0xff
	binary.LittleEndian.PutUint64(ret[1:], uint64(v.Value))
	return ret
}

// ParseVarInt parse the serialized bytes of the var int and return VarInt
func ParseVarInt(bytes []byte) (VarInt, error) {
	ret := VarInt{}
	if len(bytes) < 1 {
		return ret, fmt.Errorf("ParseVarInt: input bytes length 0")
	}
	if bytes[0] < 0xfd {
		ret.Value = uint64(bytes[0])
		return ret, nil
	}
	if bytes[0] == 0xfd {
		if len(bytes) < 3 {
			return ret, fmt.Errorf("ParseVarInt: input bytes starts with 0xfd but length %d", len(bytes))
		}
		ret.Value = uint64(binary.LittleEndian.Uint16(bytes[1:]))
		return ret, nil
	}
	if bytes[0] == 0xfe {
		if len(bytes) < 5 {
			return ret, fmt.Errorf("ParseVarInt: input bytes starts with 0xfe but length %d", len(bytes))
		}
		ret.Value = uint64(binary.LittleEndian.Uint32(bytes[1:]))
		return ret, nil
	}

	if bytes[0] == 0xff {
		if len(bytes) < 9 {
			return ret, fmt.Errorf("ParseVarInt: input bytes starts with 0xfe but length %d", len(bytes))
		}
		ret.Value = binary.LittleEndian.Uint64(bytes[1:])
		return ret, nil
	}
	return ret, fmt.Errorf("invalid input")
}

func VarIntFromUInt64(input uint64) VarInt {
	var data VarInt
	data.Value = input
	return data
}

func VarIntFromInt(input int) VarInt {
	return VarIntFromUInt64(uint64(input))
}

func VarIntFromInt16(input int16) VarInt {
	return VarIntFromUInt64(uint64(input))
}
