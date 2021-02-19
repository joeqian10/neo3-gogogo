package blockchain

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
)

//Storagekey key use to store StorageItem on blockchain
type Storagekey struct {
	ScriptHash helper.UInt160
	Key        []byte
}

//Deserialize deserialize from byte array
func (sk *Storagekey) Deserialize(reader *io.BinaryReader) {
	reader.ReadLE(&sk.ScriptHash)
	sk.Key, _ = reader.ReadBytesWithGrouping()
}

//Serialize serialize to byte array
func (sk *Storagekey) Serialize(writer *io.BinaryWriter) {
	writer.WriteLE(sk.ScriptHash)
	writer.WriteBytesWithGrouping(sk.Key)
}
