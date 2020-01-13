package sc

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/joeqian10/neo3-gogogo/helper"
)

type ScriptBuilder struct {
	buff *bytes.Buffer
}

func NewScriptBuilder() ScriptBuilder {
	return ScriptBuilder{buff: new(bytes.Buffer)}
}

// ToArray converts ScriptBuilder to byte array
func (sb *ScriptBuilder) ToArray() []byte {
	return sb.buff.Bytes()
}

//func (sb *ScriptBuilder) MakeInvocationScript(scriptHash []byte, operation string, args []ContractParameter) {
//	if len(operation) == 0 { // Neo.VM.Helper.cs: Line 28
//		l := len(args)
//		for i := l - 1; i >= 0; i-- {
//			sb.EmitPushParameter(args[i])
//		}
//		sb.EmitAppCall(scriptHash, false)
//	} else {
//		if args != nil { // Neo.VM.Helper.cs: Line 43
//			l := len(args)
//			for i := l - 1; i >= 0; i-- {
//				sb.EmitPushParameter(args[i])
//			}
//			sb.EmitPushInt(l)
//			sb.Emit(PACK)
//			sb.EmitPushString(operation)
//			sb.EmitAppCall(scriptHash, false)
//		} else { // Neo.VM.Helper.cs: Line 35
//			sb.EmitPushBool(false)
//			sb.EmitPushString(operation)
//			sb.EmitAppCall(scriptHash, false)
//		}
//	}
//}

func (sb *ScriptBuilder) Emit(op OpCode, arg ...byte) error {
	err := sb.buff.WriteByte(byte(op))
	if err != nil {
		return err
	}

	if arg != nil {
		_, err = sb.buff.Write(arg)
	}
	return err
}

func (sb *ScriptBuilder) EmitCall(offset int) error {
	if offset < -128 || offset > 127 {
		return sb.Emit(CALL_L, helper.IntToBytes(offset)...)
	} else {
		return sb.Emit(CALL, byte(offset))
	}
}

func (sb *ScriptBuilder) EmitJump(op OpCode, offset int) error {
	if op < JMP || op > JMPLE_L {
		return fmt.Errorf("invalid OpCode")
	}
	if int(op)%2 == 0 && offset < -128 || offset > 127 {
		op += 1
	}
	if int(op)%2 == 0 {
		return sb.Emit(op, byte(offset))
	} else {
		return sb.Emit(op, helper.IntToBytes(offset)...)
	}
}

func (sb *ScriptBuilder) EmitPushBigInt(number big.Int) error {
	if number.Cmp(big.NewInt(-1)) >= 0 && number.Cmp(big.NewInt(16)) <= 0 {
		var b = byte(number.Int64())
		return sb.Emit(PUSH0 + OpCode(b))
	}
	// need little endian
	data := helper.ReverseBytes(number.Bytes()) // Bytes() returns big-endian
	if len(data) == 1 {
		if number.Cmp(big.NewInt(128)) >= 0 && number.Cmp(big.NewInt(0xff)) <= 0 {
			return sb.Emit(PUSHINT16, PadRight(data, 2)...)
		} // 0b_00000000_10000000 ~ 0b_00000000_11111111, 8 -> 16 bits
		return sb.Emit(PUSHINT8, data...)
	}
	if len(data) == 2 {
		if number.Cmp(big.NewInt(0x8000)) >= 0 && number.Cmp(big.NewInt(0xffff)) <= 0 {
			return sb.Emit(PUSHINT32, PadRight(data, 4)...)
		} // 0b_00000000_10000000_00000000 ~ 0b_00000000_11111111_11111111, 16 -> 32 bits
		return sb.Emit(PUSHINT16, data...)
	}
	if len(data) <= 4 {
		if number.Cmp(big.NewInt(0x80000000)) >= 0 && number.Cmp(big.NewInt(0xffffffff)) <= 0 {
			return sb.Emit(PUSHINT64, PadRight(data, 8)...)
		} // 0b_00000000_10000000_00000000_00000000_00000000 ~ 0b_00000000_11111111_11111111_11111111_11111111, 32 -> 64 bits
		return sb.Emit(PUSHINT32, PadRight(data, 4)...)
	}
	if len(data) <= 8 {
		if number.Cmp(new(big.Int).SetUint64(0x8000000000000000)) >= 0 && number.Cmp(new(big.Int).SetUint64(0xffffffffffffffff)) <= 0 {
			return sb.Emit(PUSHINT128, PadRight(data, 16)...)
		}
		return sb.Emit(PUSHINT64, PadRight(data, 8)...)
	}
	if len(data) <= 16 {
		return sb.Emit(PUSHINT128, PadRight(data, 16)...)
	}
	if len(data) <= 32 {
		return sb.Emit(PUSHINT256, PadRight(data, 32)...)
	}
	return fmt.Errorf("argument out of range")
}

func PadRight(data []byte, length int) []byte {
	if len(data) >= length {
		return data
	}
	newData := data
	for len(newData) < length {
		newData = append(newData, byte(0))
	}
	return newData
}

func (sb *ScriptBuilder) EmitPushInt(number int) error {
	return sb.EmitPushBigInt(*big.NewInt(int64(number)))
}

func (sb *ScriptBuilder) EmitPushBool(data bool) error {
	if data {
		return sb.Emit(PUSH1)
	} else {
		return sb.Emit(PUSH0)
	}
}

func (sb *ScriptBuilder) EmitPushBytes(data []byte) error {
	if data == nil {
		return fmt.Errorf("data is empty")
	}
	le := len(data)
	var err error
	if le < int(0x100) {
		err = sb.Emit(PUSHDATA1)
		sb.buff.WriteByte(byte(le))
		sb.buff.Write(data)
	} else if le < int(0x10000) {
		err = sb.Emit(PUSHDATA2)
		sb.buff.Write(helper.UInt16ToBytes(uint16(le)))
		sb.buff.Write(data)
	} else {
		err = sb.Emit(PUSHDATA4)
		sb.buff.Write(helper.UInt32ToBytes(uint32(le)))
		sb.buff.Write(data)
	}
	if err != nil {
		return err
	}
	return nil
}

// Convert the string to UTF8 format encoded byte array
func (sb *ScriptBuilder) EmitPushString(data string) error {
	return sb.EmitPushBytes([]byte(data))
}

// EmitRaw pushes a raw byte array
func (sb *ScriptBuilder) EmitRaw(arg []byte) {
	if arg != nil {
		sb.buff.Write(arg)
	}
}

func (sb *ScriptBuilder) EmitPushParameter(data ContractParameter) error {
	var err error
	switch data.Type {
	case Signature:
	case ByteArray:
		err = sb.EmitPushBytes(data.Value.([]byte))
	case Boolean:
		err = sb.EmitPushBool(data.Value.(bool))
	case Integer:
		num := data.Value.(int64)
		err = sb.EmitPushBigInt(*big.NewInt(num))
	case Hash160:
		u, e := helper.UInt160FromBytes(data.Value.([]byte))
		if e != nil {
			return e
		}
		err = sb.EmitPushBytes(u.Bytes())
	case Hash256:
		u, e := helper.UInt256FromBytes(data.Value.([]byte))
		if e != nil {
			return e
		}
		err = sb.EmitPushBytes(u.Bytes())
	case PublicKey:
		err = sb.EmitPushBytes(data.Value.([]byte))
	case String:
		s := string(data.Value.(string))
		err = sb.EmitPushString(s)
	case Array:
		a := data.Value.([]ContractParameter)
		for i := len(a) - 1; i >= 0; i-- {
			e := sb.EmitPushParameter(a[i])
			if e != nil {
				return e
			}
		}
		err = sb.EmitPushInt(len(a))
		if err != nil {
			return err
		}
		err = sb.Emit(PACK)
	}
	if err != nil {
		return err
	}
	return nil
}

func (sb *ScriptBuilder) EmitSysCall(api uint) error {
	return sb.Emit(SYSCALL, helper.UInt32ToBytes(uint32(api))...)
}
