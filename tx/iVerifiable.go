package tx

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/helper/io"
)

type IVerifiable interface {
	GetSize() int

	Deserialize(br *io.BinaryReader)
	Serialize(bw *io.BinaryWriter)
	DeserializeUnsigned(br *io.BinaryReader)
	SerializeUnsigned(bw *io.BinaryWriter)

	GetWitnesses() []Witness
	SetWitnesses(data []Witness)
	GetScriptHashesForVerifying() []helper.UInt160
}

func GetHashData(verifiable IVerifiable) []byte {
	return GetHashDataWithMagic(verifiable, Neo3Magic)
}

func GetHashDataWithMagic(verifiable IVerifiable, magic uint32) []byte {
	buf := io.NewBufBinaryWriter()
	buf.BinaryWriter.WriteLE(magic)
	verifiable.Serialize(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func CalculateHash(verifiable IVerifiable) *helper.UInt256 {
	return helper.UInt256FromBytes(crypto.Hash256(GetHashData(verifiable)))
}

func CalculateHashWithMagic(verifiable IVerifiable, magic uint32) *helper.UInt256 {
	return helper.UInt256FromBytes(crypto.Hash256(GetHashDataWithMagic(verifiable, magic)))
}
