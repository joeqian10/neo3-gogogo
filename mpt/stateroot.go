package mpt

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/joeqian10/neo3-gogogo/tx"
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
			Invocation:   inv,
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

func (sr *StateRoot) GetHash() *helper.UInt256 {
	return tx.CalculateHash(sr)
}

func (sr *StateRoot) GetSize() int {
	size := 1 + //Version
		4 + // Index
		32 // RootHash
	if len(sr.Witnesses) == 0 {
		size += 1
	} else {
		w0 := sr.Witnesses[0]
		inv, _ := crypto.Base64Decode(w0.Invocation)
		ver, _ := crypto.Base64Decode(w0.Verification)
		witness, _ := tx.CreateWitness(inv, ver)
		size += 1 + witness.GetSize()
	}
	return size
}

func (sr *StateRoot) GetWitnesses() []*tx.Witness {
	ws := make([]*tx.Witness, len(sr.Witnesses))
	for i, v := range sr.Witnesses {
		inv, _ := crypto.Base64Decode(v.Invocation)
		ver, _ := crypto.Base64Decode(v.Verification)
		w, _ := tx.CreateWitness(inv, ver)
		ws[i] = w
	}
	return ws
}

func (sr *StateRoot) SetWitnesses(data []*tx.Witness) {
	rws := make([]models.RpcWitness, len(data))
	for i, v := range data {
		rws[i] = models.RpcWitness{
			Invocation:   crypto.Base64Encode(v.InvocationScript),
			Verification: crypto.Base64Encode(v.VerificationScript),
		}
	}
	sr.Witnesses = rws
}

func (sr *StateRoot) GetScriptHashesForVerifying() []*helper.UInt160 {
	if len(sr.Witnesses) == 0 {
		return []*helper.UInt160{}
	}
	verificationScriptBs, _ := crypto.Base64Decode(sr.Witnesses[0].Verification) // base64
	if len(verificationScriptBs) == 0 {
		return []*helper.UInt160{}
	}
	scriptHash := helper.UInt160FromBytes(crypto.Hash160(verificationScriptBs))
	return []*helper.UInt160{scriptHash}
}
