package tx

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/io"
)

// only committee can use this attribute
type HighPriorityAttribute struct {
	
}

func (h *HighPriorityAttribute) GetAttributeType() TransactionAttributeType {
	return HighPriority
}

func (h *HighPriorityAttribute) AllowMultiple() bool {
	return false
}

func (h *HighPriorityAttribute) GetAttributeSize() int {
	return 1
}

func (h *HighPriorityAttribute) Deserialize(br *io.BinaryReader) {
	if br.ReadByte() != byte(HighPriority) {
		br.Err = fmt.Errorf("format error: not HighPriority")
	}
	h.DeserializeWithoutType(br)
}

func (h *HighPriorityAttribute) Serialize(bw *io.BinaryWriter)  {
	bw.WriteLE(byte(HighPriority))
	h.SerializeWithoutType(bw)
}

func (h *HighPriorityAttribute) DeserializeWithoutType(br *io.BinaryReader) {

}

func (h *HighPriorityAttribute) SerializeWithoutType(bw *io.BinaryWriter)  {

}
