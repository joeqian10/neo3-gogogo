package tx

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOracleResponseAttribute_AllowMultiple(t *testing.T) {
	oracleRes, err := NewOracleResponseAttribute()
	assert.Nil(t, err)
	assert.Equal(t, false, oracleRes.AllowMultiple())
}

func TestOracleResponseAttribute_Deserialize(t *testing.T) {
	oracleRes, err := NewOracleResponseAttribute()
	assert.Nil(t, err)
	b := helper.HexToBytes("110100000000000000000401020304")
	br := io.NewBinaryReaderFromBuf(b)
	oracleRes.Deserialize(br)
	assert.Equal(t, "01020304", helper.BytesToHex(oracleRes.Result))
}

func TestOracleResponseAttribute_GetAttributeSize(t *testing.T) {
	oracleRes, err := NewOracleResponseAttribute()
	assert.Nil(t, err)
	oracleRes.Id = 1
	oracleRes.Code = Success
	oracleRes.Result = helper.HexToBytes("01020304")
	assert.Equal(t, 15, oracleRes.GetAttributeSize())
}

func TestOracleResponseAttribute_GetAttributeType(t *testing.T) {
	oracleRes, err := NewOracleResponseAttribute()
	assert.Nil(t, err)
	oracleRes.Id = 1
	oracleRes.Code = Success
	oracleRes.Result = helper.HexToBytes("01020304")
	assert.Equal(t, OracleResponse, oracleRes.GetAttributeType())
}

func TestOracleResponseAttribute_Serialize(t *testing.T) {
	oracleRes, err := NewOracleResponseAttribute()
	assert.Nil(t, err)
	oracleRes.Id = 1
	oracleRes.Code = Success
	oracleRes.Result = helper.HexToBytes("01020304")

	bbw := io.NewBufBinaryWriter()
	oracleRes.Serialize(bbw.BinaryWriter)
	assert.Equal(t, "110100000000000000000401020304", helper.BytesToHex(bbw.Bytes()))
}
