package tx

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
)

type ITransactionAttribute interface {
	GetAttributeType() TransactionAttributeType
	AllowMultiple() bool
	GetAttributeSize() int

	Deserialize(br *io.BinaryReader)
	Serialize(bw *io.BinaryWriter)
	DeserializeWithoutType(br *io.BinaryReader)
	SerializeWithoutType(bw *io.BinaryWriter)
}

func CreateTransactionAttribute(attributeType TransactionAttributeType) ITransactionAttribute {
	b := byte(attributeType)
	switch b {
	case 0x01:
		return &HighPriorityAttribute{}
	case 0x11:
		a, _ := NewOracleResponseAttribute()
		return a
	default:
		return nil
	}
}

func DeserializeFrom(br *io.BinaryReader) ITransactionAttribute {
	t := TransactionAttributeType(br.ReadOneByte())
	a := CreateTransactionAttribute(t)
	if a == nil {
		br.Err = fmt.Errorf("format error: invalid attribute type")
		return nil
	}
	a.DeserializeWithoutType(br)
	return a
}

type TransactionAttributeSlice []ITransactionAttribute

func (ts TransactionAttributeSlice) GetVarSize() int {
	var size int = 0
	for _, t := range ts {
		size += t.GetAttributeSize()
	}
	return helper.GetVarSize(len(ts)) + size
}
