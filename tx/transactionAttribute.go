package tx

import (
	"encoding/hex"
	"encoding/json"
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"github.com/joeqian10/neo3-gogogo/sc"
)

type TransactionAttribute struct {
	Usage TransactionAttributeUsage
	Data  []byte
}

func (attr *TransactionAttribute) Size() int {
	size := 1 + sc.ByteSlice(attr.Data).GetVarSize()
	return size
}

// Deserialize implements Serializable interface.
func (attr *TransactionAttribute) Deserialize(br *io.BinaryReader) {
	br.ReadLE(&attr.Usage)
	if attr.Usage.IsDefined() == false {attr.Usage = Url}
	attr.Data = br.ReadVarBytes(252)
}

// Serialize implements Serializable interface.
func (attr *TransactionAttribute) Serialize(bw *io.BinaryWriter) {
	bw.WriteLE(attr.Usage)
	bw.WriteVarBytes(attr.Data)
}

// MarshalJSON implements the json Marshaller interface.
func (attr *TransactionAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"usage": attr.Usage.String(),
		"data":  hex.EncodeToString(attr.Data),
	})
}

type TransactionAttributeSlice []*TransactionAttribute

func (ts TransactionAttributeSlice) GetVarSize() int {
	var size int = 0
	for _, t := range ts {
		size += t.Size()
	}
	return size
}