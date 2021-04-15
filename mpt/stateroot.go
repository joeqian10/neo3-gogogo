package mpt

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
)

//StateRoot truct of StateRoot message
type StateRoot struct {
	Version  byte              `json:"version"`
	Index    uint32            `json:"index"`
	RootHash string            `json:"roothash"`
	Witness  models.RpcWitness `json:"witness"`
}

func (sr *StateRoot) Deserialize(br *io.BinaryReader) {
	sr.DeserializeUnsigned(br)
	l := br.ReadVarUInt()
	if l != 1 {
		return
	}
	sr.Witness.Invocation = crypto.Base64Encode(br.ReadVarBytes())
	sr.Witness.Verification = crypto.Base64Encode(br.ReadVarBytes())
}

func (sr *StateRoot) Serialize(bw *io.BinaryWriter) {
	sr.SerializeUnsigned(bw)
	bw.WriteVarUInt(1)
	is, err := crypto.Base64Decode(sr.Witness.Invocation)
	if err != nil {
		bw.Err = err
		return
	}
	vs, err := crypto.Base64Decode(sr.Witness.Verification)
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
	var rootHash, stateRoot helper.UInt256
	br.ReadLE(&rootHash)
	br.ReadLE(&stateRoot)
	sr.RootHash = "0x" + rootHash.String()
}

func (sr *StateRoot) SerializeUnsigned(bw *io.BinaryWriter) {
	bw.WriteLE(sr.Version)
	bw.WriteLE(sr.Index)
	rootHash, _ := helper.UInt256FromString(sr.RootHash)
	bw.WriteLE(rootHash)
}
