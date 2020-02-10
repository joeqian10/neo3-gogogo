package sc

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteSlice_GetVarSize(t *testing.T) {
	b := helper.HexTobytes("deadbeef")
	size := ByteSlice(b).GetVarSize()
	assert.Equal(t, 5, size)
}

func TestByteSlice_IsMultiSigContract(t *testing.T) {
	m := 4
	sb := NewScriptBuilder()
	_ = sb.EmitPushInt(m)
	_ = sb.EmitPushBytes(helper.HexTobytes("025e5eb8e89ab16cda6e5f646de54a8e9e9e8ce0a64e44db6b6ffeff8a6369f268"))
	_ = sb.EmitPushBytes(helper.HexTobytes("025e5eb8e89ab16cda6e5f646de54a8e9e9e8ce0a64e44db6b6ffeff8a6369f268"))
	_ = sb.EmitPushBytes(helper.HexTobytes("025e5eb8e89ab16cda6e5f646de54a8e9e9e8ce0a64e44db6b6ffeff8a6369f268"))
	_ = sb.EmitPushBytes(helper.HexTobytes("025e5eb8e89ab16cda6e5f646de54a8e9e9e8ce0a64e44db6b6ffeff8a6369f268"))
	_ = sb.EmitPushInt(m)
	_ = sb.Emit(PUSHNULL)
	_ = sb.EmitSysCall(ECDsaCheckMultiSig.ToInteropMethodHash())
	script := sb.ToArray()
	b, mt, nt := ByteSlice(script).IsMultiSigContract()
	assert.Equal(t, true, b)
	assert.Equal(t, 4, mt)
	assert.Equal(t, 4, nt)
}

func TestByteSlice_IsSignatureContract(t *testing.T) {
	//p, err := keys.NewPublicKeyFromString("025e5eb8e89ab16cda6e5f646de54a8e9e9e8ce0a64e44db6b6ffeff8a6369f268")
	//assert.Nil(t, err)
	//script := keys.CreateSignatureRedeemScript(p)
	//b := ByteSlice(script).IsSignatureContract()
	sb := NewScriptBuilder()
	_ = sb.EmitPushBytes(helper.HexTobytes("025e5eb8e89ab16cda6e5f646de54a8e9e9e8ce0a64e44db6b6ffeff8a6369f268"))
	_ = sb.Emit(PUSHNULL)
	_ = sb.EmitSysCall(ECDsaVerify.ToInteropMethodHash())
	script := sb.ToArray()
	b := ByteSlice(script).IsSignatureContract()
	assert.Equal(t, true, b)
}