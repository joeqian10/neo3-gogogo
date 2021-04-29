package mpt

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
)

//StateRoot truct of StateRoot message
type StateRoot struct {
	Version   byte                `json:"version"`
	Index     uint32              `json:"index"`
	RootHash  string              `json:"roothash"`
	Witnesses []models.RpcWitness `json:"witnesses"`
}

func (sr *StateRoot) Deserialize(br *io.BinaryReader) {
	sr.DeserializeUnsigned(br)
	l := br.ReadVarUInt()
	if l != 1 {
		return
	}
	inv := crypto.Base64Encode(br.ReadVarBytes())
	ver := crypto.Base64Encode(br.ReadVarBytes())
	sr.Witnesses = []models.RpcWitness{
		{
			Invocation: inv,
			Verification: ver,
		},
	}
}

func (sr *StateRoot) Serialize(bw *io.BinaryWriter) {
	sr.SerializeUnsigned(bw)
	bw.WriteVarUInt(1)
	is, err := crypto.Base64Decode(sr.Witnesses[0].Invocation)
	if err != nil {
		bw.Err = err
		return
	}
	vs, err := crypto.Base64Decode(sr.Witnesses[0].Verification)
	if err != nil {
		bw.Err = err
		return
	}
	bw.WriteVarBytes(is)
	bw.WriteVarBytes(vs)
}

func (sr *StateRoot) DeserializeUnsigned(br *io.BinaryReader) {
	br.ReadLE(&sr.Version)
	br.ReadLE(&sr.Index)
	var rootHash helper.UInt256
	br.ReadLE(&rootHash)
	sr.RootHash = "0x" + rootHash.String()
}

func (sr *StateRoot) SerializeUnsigned(bw *io.BinaryWriter) {
	bw.WriteLE(sr.Version)
	bw.WriteLE(sr.Index)
	rootHash, _ := helper.UInt256FromString(sr.RootHash)
	bw.WriteLE(rootHash)
}
