package blockchain

import "github.com/joeqian10/neo3-gogogo/io"

//StorageItem value to store on blockchain
type StorageItem struct {
	Value      []byte
}

//Deserialize deserialize from byte array
func (si *StorageItem) Deserialize(reader *io.BinaryReader) {
	si.Value = reader.ReadVarBytes()
}

//Serialize serialize to byte array
func (si *StorageItem) Serialize(writer *io.BinaryWriter) {
	writer.WriteVarBytes(si.Value)
}
