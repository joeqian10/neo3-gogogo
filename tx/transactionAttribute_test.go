package tx

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionAttribute_Size(t *testing.T) {
	ta := TransactionAttribute{
		Usage: 0x81,
		Data:  helper.HexTobytes("676f6f676c65"),
	}
	size := ta.Size()
	assert.Equal(t, 1+1+6, size)
}

func TestTransactionAttribute_Deserialize(t *testing.T) {
	s := "8106676f6f676c65"
	br := io.NewBinaryReaderFromBuf(helper.HexTobytes(s))
	ta := TransactionAttribute{}
	ta.Deserialize(br)
	assert.Equal(t, uint8(0x81), uint8(ta.Usage))
}

func TestTransactionAttribute_Serialize(t *testing.T) {
	ta := TransactionAttribute{
		Usage: 0x81,
		Data:  helper.HexTobytes("676f6f676c65"),
	}
	bbw := io.NewBufBinaryWriter()
	ta.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	assert.Equal(t, "8106676f6f676c65", helper.BytesToHex(b))
}

func TestTransactionAttributeSlice_GetVarSize(t *testing.T) {
	ta1 := TransactionAttribute{
		Usage: 0x81,
		Data:  helper.HexTobytes("777777"),
	}
	ta2 := TransactionAttribute{
		Usage: 0x81,
		Data:  helper.HexTobytes("676f6f676c65"),
	}
	ta3 := TransactionAttribute{
		Usage: 0x81,
		Data:  helper.HexTobytes("636f6d"),
	}
	tas := []*TransactionAttribute{&ta1, &ta2, &ta3}
	size := TransactionAttributeSlice(tas).GetVarSize()
	assert.Equal(t, 3, len(TransactionAttributeSlice(tas)))
	assert.Equal(t, 1+1+1+3+1+1+6+1+1+3, size)
}
