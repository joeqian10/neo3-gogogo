package tx

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
)

type IVerifiable interface {
	GetHash() *helper.UInt256
	GetSize() int

	Deserialize(br *io.BinaryReader)
	Serialize(bw *io.BinaryWriter)
	DeserializeUnsigned(br *io.BinaryReader)
	SerializeUnsigned(bw *io.BinaryWriter)

	GetWitnesses() []Witness
	SetWitnesses(data []Witness)
	GetScriptHashesForVerifying() []helper.UInt160
}

func GetSignData(verifiable IVerifiable, magic uint32) []byte {
	buf := io.NewBufBinaryWriter()
	buf.BinaryWriter.WriteLE(magic)
	buf.BinaryWriter.WriteLE(verifiable.GetHash())
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func CalculateHash(verifiable IVerifiable) *helper.UInt256 {
	buf := io.NewBufBinaryWriter()
	verifiable.SerializeUnsigned(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}

	return helper.UInt256FromBytes(crypto.Sha256(buf.Bytes()))
}
