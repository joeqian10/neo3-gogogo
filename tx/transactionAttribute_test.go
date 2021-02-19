package tx

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTransactionAttribute(t *testing.T) {
	attr := CreateTransactionAttribute(HighPriority)
	assert.Equal(t, HighPriority, attr.GetAttributeType())
}

func TestDeserializeFrom(t *testing.T) {
	b := helper.HexToBytes("110100000000000000000401020304")
	br := io.NewBinaryReaderFromBuf(b)
	a := DeserializeFrom(br)
	assert.Equal(t, OracleResponse, a.GetAttributeType())
}
