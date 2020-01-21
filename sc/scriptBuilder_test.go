package sc

import (
	"math/big"
	"testing"

	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/stretchr/testify/assert"
)

func TestScriptBuilder_Emit(t *testing.T) {
	sb := NewScriptBuilder()
	scriptHash, _ := helper.UInt160FromString("14df5d02f9a52d3e92ab8cdcce5fc76c743a9b26")
	sb.Emit(SYSCALL, scriptHash.Bytes()...)
	b := sb.ToArray()
	assert.Equal(t, "41269b3a746cc75fcedc8cab923e2da5f9025ddf14", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitSysCall(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitSysCall(0xE393C875)
	b := sb.ToArray()
	assert.Equal(t, "4175c893e3", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitCall(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitCall(12345) // 0x3039
	b := sb.ToArray()
	assert.Equal(t, "3539300000", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitJump(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitJump(JMP, 127)
	sb.EmitJump(JMP, 2147483647)
	b := sb.ToArray()
	assert.Equal(t, "227f23ffffff7f", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt1(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(-1))
	b := sb.ToArray()
	assert.Equal(t, "0f", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt2(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(0))
	b := sb.ToArray()
	assert.Equal(t, "10", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt3(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(-128))
	b := sb.ToArray()
	assert.Equal(t, "0080", helper.BytesToHex(b)) //0b_1000_0000
}

func TestScriptBuilder_EmitPushBigInt4(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(127)) //0b_0111_1111, 0b_1111_1111 = 0d_-1
	b := sb.ToArray()
	assert.Equal(t, "007f", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt5(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(255)) //0b_00000000_11111111
	b := sb.ToArray()
	assert.Equal(t, "01ff00", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt6(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(-32768)) //0b_1000_0000_0000_0000
	b := sb.ToArray()
	assert.Equal(t, "010080", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt7(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(32767)) //0b_1111_1111_1111_1111
	b := sb.ToArray()
	assert.Equal(t, "01ff7f", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt8(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(65535)) //0b_1111_1111_1111_1111
	b := sb.ToArray()
	assert.Equal(t, "02ffff0000", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt9(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(-2147483648)) //0b_1111_1111_1111_1111
	b := sb.ToArray()
	assert.Equal(t, "0200000080", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt10(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(2147483647)) //0b_1111_1111_1111_1111
	b := sb.ToArray()
	assert.Equal(t, "02ffffff7f", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt11(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(4294967295)) //0b_1111_1111_1111_1111
	b := sb.ToArray()
	assert.Equal(t, "03ffffffff00000000", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt12(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(-9223372036854775808)) //0b_1111_1111_1111_1111
	b := sb.ToArray()
	assert.Equal(t, "030000000000000080", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt13(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(9223372036854775807)) //0b_1111_1111_1111_1111
	b := sb.ToArray()
	assert.Equal(t, "03ffffffffffffff7f", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt14(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*new(big.Int).SetUint64(uint64(18446744073709551615))) //0b_1111_1111_1111_1111
	b := sb.ToArray()
	assert.Equal(t, "04ffffffffffffffff0000000000000000", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBool(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBool(true)
	b := sb.ToArray()
	assert.Equal(t, "11", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBytes(t *testing.T) {
	sb := NewScriptBuilder()
	n := *big.NewInt(7777777777)
	bytes := helper.ReverseBytes(n.Bytes())
	sb.EmitPushBytes(bytes)
	b := sb.ToArray()
	assert.Equal(t, "0c05717897cf01", helper.BytesToHex(b)) // PUSHDATA1 + length + data
}

func TestScriptBuilder_EmitPushInt(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushInt(-1)
	sb.EmitPushInt(0)
	sb.EmitPushInt(8)
	sb.EmitPushInt(100)
	sb.EmitPushInt(1000)
	sb.EmitPushInt(10000)
	sb.EmitPushInt(0x20000)
	bytes := sb.ToArray()
	assert.Equal(t, "0f1018006401e8030110270200000200", helper.BytesToHex(bytes)) // 0f + 10 + 18 + 0064 + 01e803 + 011027 + 0200000200, pad right
}

func TestScriptBuilder_EmitPushParameter(t *testing.T) {
	u, _ := helper.UInt256FromString("c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b")
	cp := ContractParameter{
		Type:  Hash256,
		Value: u.Bytes(),
	}
	sb := NewScriptBuilder()
	sb.EmitPushParameter(cp)
	b := sb.ToArray()
	assert.Equal(t, "0c209b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc5", helper.BytesToHex(b)) // 0c + length + data
}

func TestScriptBuilder_EmitPushString(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushString("Hello World!")
	b := sb.ToArray()
	assert.Equal(t, "0c0c48656c6c6f20576f726c6421", helper.BytesToHex(b))
}

func TestScriptBuilder_ToArray(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushString("Hello World!")
	b := sb.ToArray()
	assert.Equal(t, "0c0c48656c6c6f20576f726c6421", helper.BytesToHex(b))
}
