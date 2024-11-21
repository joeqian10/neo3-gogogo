package sc

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/joeqian10/neo3-gogogo/crypto"

	"github.com/stretchr/testify/assert"

	"github.com/joeqian10/neo3-gogogo/helper"
)

func TestScriptBuilder_Emit(t *testing.T) {
	sb := NewScriptBuilder()
	sb.Emit(NOP)
	b, err := sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, []byte{0x21}, b)

	sb = NewScriptBuilder()
	sb.Emit(NOP, 0x66)
	b, err = sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, []byte{0x21, 0x66}, b)
}

func TestScriptBuilder_EmitSysCall(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitSysCall(0xE393C875)
	b, err := sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, []byte{byte(SYSCALL), 0x75, 0xc8, 0x93, 0xe3}, b)
}

func TestScriptBuilder_EmitCall(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitCall(0)
	b, err := sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, []byte{byte(CALL), byte(0)}, b)

	sb = NewScriptBuilder()
	sb.EmitCall(12345) // 0x3039
	b, err = sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, append([]byte{byte(CALL_L)}, helper.IntToBytes(12345)...), b)

	sb = NewScriptBuilder()
	sb.EmitCall(-12345)
	b, err = sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, append([]byte{byte(CALL_L)}, helper.IntToBytes(-12345)...), b)
}

func TestScriptBuilder_EmitJump(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitJump(NOP, 127)
	_, err := sb.ToArray()
	assert.NotNil(t, err)

	offset_i8 := 127
	offset_i32 := 2147483647
	var i byte
	for i = 0x22; i < 0x34; i++ {
		sb = NewScriptBuilder()
		sb.EmitJump(OpCode(i), offset_i8)
		sb.EmitJump(OpCode(i), offset_i32)
		b, err := sb.ToArray()
		assert.Nil(t, err)
		if i%2 == 0 {
			expected := append([]byte{i, byte(offset_i8), i + 1}, helper.IntToBytes(offset_i32)...)
			assert.Equal(t, true, bytes.Equal(expected, b))
		} else {
			expected := append(append([]byte{i}, helper.IntToBytes(offset_i8)...), append([]byte{i}, helper.IntToBytes(offset_i32)...)...)
			assert.Equal(t, true, bytes.Equal(expected, b))
		}
	}

	offset_i8 = -128
	offset_i32 = -2147483648
	for i = 0x22; i < 0x34; i++ {
		sb = NewScriptBuilder()
		sb.EmitJump(OpCode(i), offset_i8)
		sb.EmitJump(OpCode(i), offset_i32)
		b, err := sb.ToArray()
		assert.Nil(t, err)
		if i%2 == 0 {
			expected := append([]byte{i, byte(offset_i8), i + 1}, helper.IntToBytes(offset_i32)...)
			assert.Equal(t, true, bytes.Equal(expected, b))
		} else {
			expected := append(append([]byte{i}, helper.IntToBytes(offset_i8)...), append([]byte{i}, helper.IntToBytes(offset_i32)...)...)
			assert.Equal(t, true, bytes.Equal(expected, b))
		}
	}
}

func TestScriptBuilder_EmitPushBigInt(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(big.NewInt(-1))
	b, err := sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, true, bytes.Equal([]byte{0x0f}, b))
	var x byte
	for x = 0; x <= 16; x++ {
		sb = NewScriptBuilder()
		sb.EmitPushBigInt(big.NewInt(int64(x)))
		b, err = sb.ToArray()
		assert.Nil(t, err)
		assert.Equal(t, true, bytes.Equal([]byte{byte(PUSH0) + x}, b))
	}

	data := make(map[string]*big.Int)
	data["0080"] = big.NewInt(-128)
	data["007f"] = big.NewInt(127)
	data["01ff00"] = big.NewInt(255)
	data["010080"] = big.NewInt(-32768)
	data["01ff7f"] = big.NewInt(32767)
	data["02ffff0000"] = big.NewInt(65535)
	data["0200000080"] = big.NewInt(-2147483648)
	data["02ffffff7f"] = big.NewInt(2147483647)
	data["03ffffffff00000000"] = big.NewInt(4294967295)
	data["030000000000000080"] = big.NewInt(-9223372036854775808)
	data["03ffffffffffffff7f"] = big.NewInt(9223372036854775807)
	data["04ffffffffffffffff0000000000000000"] = new(big.Int).SetUint64(18446744073709551615)
	data["050100000000000000feffffffffffffff00000000000000000000000000000000"] = new(big.Int).Mul(new(big.Int).SetUint64(18446744073709551615), new(big.Int).SetUint64(18446744073709551615))

	for k, v := range data {
		sb = NewScriptBuilder()
		sb.EmitPushBigInt(v)
		b, err = sb.ToArray()
		assert.Nil(t, err)
		assert.Equal(t, k, helper.BytesToHex(b))
	}

	y := new(big.Int).SetBytes(helper.ReverseBytes(helper.HexToBytes("0100000000000000feffffffffffffff0100000000000000feffffffffffffff00000000000000000000000000000000")))
	sb = NewScriptBuilder()
	sb.EmitPushBigInt(y)
	_, err = sb.ToArray()
	assert.NotNil(t, err)
}

func TestScriptBuilder_EmitPushBool(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBool(true)
	b, err := sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, []byte{byte(PUSH1)}, b)

	sb = NewScriptBuilder()
	sb.EmitPushBool(false)
	b, err = sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, []byte{byte(PUSH0)}, b)
}

func TestScriptBuilder_EmitPushBytes(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBytes(nil)
	_, err := sb.ToArray()
	assert.NotNil(t, err)

	// length = 0x4c
	bs, err := helper.GenerateRandomBytes(0x4c)
	assert.Nil(t, err)
	sb = NewScriptBuilder()
	sb.EmitPushBytes(bs)
	b, err := sb.ToArray()
	expected := append([]byte{byte(PUSHDATA1), byte(len(bs))}, bs...)
	assert.Equal(t, true, bytes.Equal(expected, b))

	// length = 0x100
	bs, err = helper.GenerateRandomBytes(0x100)
	for i, _ := range bs {
		bs[i] = byte(i)
	}
	assert.Nil(t, err)
	sb = NewScriptBuilder()
	sb.EmitPushBytes(bs)
	b, err = sb.ToArray()
	expected = append([]byte{byte(PUSHDATA2)}, helper.Int16ToBytes(int16(len(bs)))...)
	expected = append(expected, bs...)
	s1 := helper.BytesToHex(expected)
	s2 := helper.BytesToHex(b)
	assert.Equal(t, s1, s2)
	assert.Equal(t, true, bytes.Equal(expected, b))

	//  length = 0x10000
	bs, err = helper.GenerateRandomBytes(0x10000)
	assert.Nil(t, err)
	sb = NewScriptBuilder()
	sb.EmitPushBytes(bs)
	b, err = sb.ToArray()
	expected = append(append([]byte{byte(PUSHDATA4)}, helper.IntToBytes(len(bs))...), bs...)
	assert.Equal(t, true, bytes.Equal(expected, b))
}

func TestScriptBuilder_EmitPushString(t *testing.T) {
	// length = 0x4c
	bs, err := helper.GenerateRandomBytes(0x4c)
	assert.Nil(t, err)
	sb := NewScriptBuilder()
	sb.EmitPushString(string(bs))
	b, err := sb.ToArray()
	expected := append([]byte{byte(PUSHDATA1), byte(len(bs))}, bs...)
	assert.Equal(t, true, bytes.Equal(expected, b))

	// length = 0x100
	bs, err = helper.GenerateRandomBytes(0x100)
	assert.Nil(t, err)
	sb = NewScriptBuilder()
	sb.EmitPushString(string(bs))
	b, err = sb.ToArray()
	expected = append(append([]byte{byte(PUSHDATA2)}, helper.Int16ToBytes(int16(len(bs)))...), bs...)
	assert.Equal(t, true, bytes.Equal(expected, b))

	//  length = 0x10000
	bs, err = helper.GenerateRandomBytes(0x10000)
	assert.Nil(t, err)
	sb = NewScriptBuilder()
	sb.EmitPushString(string(bs))
	b, err = sb.ToArray()
	expected = append(append([]byte{byte(PUSHDATA4)}, helper.IntToBytes(len(bs))...), bs...)
	assert.Equal(t, true, bytes.Equal(expected, b))
}

func TestScriptBuilder_EmitOpCodes(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitOpCodes([]OpCode{PUSH0}...)
	b, err := sb.ToArray()
	assert.Nil(t, err)
	expected := []byte{byte(PUSH0)}
	assert.Equal(t, true, bytes.Equal(expected, b))
}

func TestScriptBuilder_EmitDynamicCall(t *testing.T) {
	//format:(byte)0x10+(byte)OpCode.NEWARRAY+(string)operation+(Uint160)scriptHash+(uint)InteropService.System_Contract_Call
	sb := NewScriptBuilder()
	sb.EmitDynamicCall(helper.UInt160Zero, "AAAAA", nil)

	tmp := make([]byte, 36)
	tmp[0] = byte(NEWARRAY0)
	tmp[1] = byte(PUSH15)
	tmp[2] = byte(PUSHDATA1)
	tmp[3] = 5            // operation.Length
	bs := []byte("AAAAA") // operation.data
	copy(tmp[4:9], bs)
	tmp[9] = byte(PUSHDATA1)
	tmp[10] = 0x14                                     // scripthash.Length
	copy(tmp[11:31], helper.UInt160Zero.ToByteArray()) // scripthash.data
	api := System_Contract_Call.ToInteropMethodHash()
	tmp[31] = byte(SYSCALL)
	copy(tmp[32:], helper.UInt32ToBytes(uint32(api)))

	b, err := sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, true, bytes.Equal(tmp, b))
}

func TestScriptBuilder_EmitDynamicCall2(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitDynamicCall(helper.UInt160Zero, "AAAAA", []interface{}{ContractParameter{Type: Integer, Value: 0}})

	tmp := make([]byte, 38)
	tmp[0] = byte(PUSH0) // arg
	tmp[1] = byte(PUSH1) // arg.Length
	tmp[2] = byte(PACK)
	tmp[3] = byte(PUSH15) // CallFlags.All
	tmp[4] = byte(PUSHDATA1)
	tmp[5] = 0x05         // operation.Length
	bs := []byte("AAAAA") // operation.data
	copy(tmp[6:11], bs)
	tmp[11] = byte(PUSHDATA1)
	tmp[12] = 0x14                                     // scripthash.Length
	copy(tmp[13:33], helper.UInt160Zero.ToByteArray()) // scripthash.data
	api := System_Contract_Call.ToInteropMethodHash()
	tmp[33] = byte(SYSCALL)
	copy(tmp[34:], helper.UInt32ToBytes(uint32(api)))

	b, err := sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, true, bytes.Equal(tmp, b))
}

func TestScriptBuilder_EmitDynamicCall3(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitDynamicCall(helper.UInt160Zero, "AAAAA", []interface{}{true})

	tmp := make([]byte, 38)
	tmp[0] = byte(PUSH1) // arg
	tmp[1] = byte(PUSH1) // arg.Length
	tmp[2] = byte(PACK)
	tmp[3] = byte(PUSH15) // CallFlags.All
	tmp[4] = byte(PUSHDATA1)
	tmp[5] = 0x05         // operation.Length
	bs := []byte("AAAAA") // operation.data
	copy(tmp[6:11], bs)
	tmp[11] = byte(PUSHDATA1)
	tmp[12] = 0x14                                     // scripthash.Length
	copy(tmp[13:33], helper.UInt160Zero.ToByteArray()) // scripthash.data
	api := System_Contract_Call.ToInteropMethodHash()
	tmp[33] = byte(SYSCALL)
	copy(tmp[34:], helper.UInt32ToBytes(uint32(api)))

	b, err := sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, true, bytes.Equal(tmp, b))
}

func TestScriptBuilder_CreateArray(t *testing.T) {
	a := []interface{}{*big.NewInt(1), *big.NewInt(2), *big.NewInt(3)}
	sb := NewScriptBuilder()
	sb.CreateArray(a)
	expected := []byte{byte(PUSH3), byte(PUSH2), byte(PUSH1), byte(PUSH3), byte(PACK)}
	b, err := sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, true, bytes.Equal(expected, b))
}

func TestScriptBuilder_CreateArray2(t *testing.T) {
	a := []interface{}{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
	sb := NewScriptBuilder()
	sb.CreateArray(a)
	expected := []byte{byte(PUSH3), byte(PUSH2), byte(PUSH1), byte(PUSH3), byte(PACK)}
	b, err := sb.ToArray()
	assert.Nil(t, err)
	assert.Equal(t, true, bytes.Equal(expected, b))
}

func TestScriptBuilder_CreateMap(t *testing.T) {
	a := map[interface{}]interface{}{}
	a[big.NewInt(1)] = big.NewInt(2)
	a[big.NewInt(3)] = big.NewInt(4)
	sb := NewScriptBuilder()
	sb.CreateMap(a)
	b, err := sb.ToArray()
	assert.Nil(t, err)
	s1 := "c84a1112d04a1314d0"
	s2 := "c84a1314d04a1112d0"
	r := helper.BytesToHex(b)
	assert.Contains(t, []string{s1, s2}, r)
}

func TestMakeScript(t *testing.T) {
	b, err := MakeScript(helper.UInt160FromBytes(helper.HexToBytes("28b3adab7269f9c2181db3cb741ebf551930e270")), "balanceOf", []interface{}{helper.UInt160Zero})
	assert.Nil(t, err)
	expected := "0c14000000000000000000000000000000000000000011c01f0c0962616c616e63654f660c1428b3adab7269f9c2181db3cb741ebf551930e27041627d5b52"
	actual := helper.BytesToHex(b)
	assert.Equal(t, expected, actual)
}

func TestMakeScript2(t *testing.T) {
	scriptHash, _ := helper.UInt160FromString("f8328398c4c8e77b6c5843f5e404be0170d5012e")
	gas, _ := helper.UInt160FromString("0xd2a4cff31913016155e38e474a2c06d08be276cf")
	cp1 := ContractParameter{
		Type:  Hash160,
		Value: gas,
	}
	b, err := MakeScript(scriptHash, "extractFee", []interface{}{cp1})
	assert.Nil(t, err)
	bb := crypto.Base64Encode(b)
	fmt.Println(bb)
}
