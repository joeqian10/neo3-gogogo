package tx

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"github.com/stretchr/testify/assert"
	"testing"
)

var highPriority = CreateTransactionAttribute(HighPriority)

func TestHighPriorityAttribute_AllowMultiple(t *testing.T) {
	assert.Equal(t, false, highPriority.AllowMultiple())
}

func TestHighPriorityAttribute_Deserialize(t *testing.T) {
	b := helper.HexToBytes("01")
	br := io.NewBinaryReaderFromBuf(b)
	hp := HighPriorityAttribute{}
	hp.Deserialize(br)
	assert.Equal(t, highPriority.GetAttributeType(), hp.GetAttributeType())
}

func TestHighPriorityAttribute_GetAttributeSize(t *testing.T) {
	assert.Equal(t, 1, highPriority.GetAttributeSize())
}

func TestHighPriorityAttribute_GetAttributeType(t *testing.T) {
	assert.Equal(t, HighPriority, highPriority.GetAttributeType())
}

func TestHighPriorityAttribute_Serialize(t *testing.T) {
	bbw := io.NewBufBinaryWriter()
	highPriority.Serialize(bbw.BinaryWriter)
	assert.Equal(t, "01", helper.BytesToHex(bbw.Bytes()))
}
