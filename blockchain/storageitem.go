package blockchain

import "github.com/joeqian10/neo3-gogogo/helper/io"

//StorageItem value to store on blockchain
type StorageItem struct {
	Version    byte
	Value      []byte
	IsConstant bool
}

//Deserialize deserialize from byte array
func (si *StorageItem) Deserialize(reader *io.BinaryReader) {
	reader.ReadLE(&si.Version)
	si.Value = reader.ReadVarBytes()
	reader.ReadLE(&si.IsConstant)
}

//Serialize serialize to byte array
func (si *StorageItem) Serialize(writer *io.BinaryWriter) {
	writer.WriteLE(si.Version)
	writer.WriteVarBytes(si.Value)
	writer.WriteLE(si.IsConstant)
}
