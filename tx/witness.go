package tx

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/joeqian10/neo3-gogogo/keys"
	"github.com/joeqian10/neo3-gogogo/sc"
)

type WitnessSlice []Witness

func (ws WitnessSlice) Len() int           { return len(ws) }
func (ws WitnessSlice) Less(i, j int) bool { return ws[i]._scriptHash.Less(ws[j]._scriptHash) }
func (ws WitnessSlice) Swap(i, j int)      { ws[i], ws[j] = ws[j], ws[i] }

/// <summary>
/// This is designed to allow a MultiSig 21/11 (committee)
/// Invocation = 11 * (64 + 2) = 726
/// </summary>
const MaxInvocationScript = 1024
/// <summary>
/// Verification = m + (PUSH_PubKey * 21) + length + null + syscall = 1 + ((2 + 33) * 21) + 2 + 1 + 5 = 744
/// </summary>
const MaxVerificationScript = 1024

// Witness
type Witness struct {
	InvocationScript   []byte         // signature
	VerificationScript []byte         // pub key
	_scriptHash        *helper.UInt160 // script hash
}

func (w *Witness) GetScriptHash() *helper.UInt160 {
	w._scriptHash = crypto.BytesToScriptHash(w.VerificationScript)
	return w._scriptHash
}

func (w *Witness) Size() int {
	return sc.ByteSlice(w.InvocationScript).GetVarSize() + sc.ByteSlice(w.VerificationScript).GetVarSize()
}

// Deserialize implements Serializable interface.
func (w *Witness) Deserialize(br *io.BinaryReader) {
	w.InvocationScript = br.ReadVarBytesWithMaxLimit(MaxInvocationScript)
	w.VerificationScript = br.ReadVarBytesWithMaxLimit(MaxVerificationScript)
}

// Serialize implements Serializable interface.
func (w *Witness) Serialize(bw *io.BinaryWriter) {
	bw.WriteVarBytes(w.InvocationScript)
	bw.WriteVarBytes(w.VerificationScript)
}

// MarshalJSON implements the json marshaller interface.
func (w *Witness) MarshalJSON() ([]byte, error) {
	data := map[string]string{
		"invocation":   hex.EncodeToString(w.InvocationScript),
		"verification": hex.EncodeToString(w.VerificationScript),
	}

	return json.Marshal(data)
}

// Create Witness with invocationScript and verificationScript
func CreateWitness(invocationScript []byte, verificationScript []byte) (*Witness, error) {
	if len(verificationScript) == 0 {
		return nil, fmt.Errorf("verificationScript should not be empty")
	}
	witness := &Witness{InvocationScript: invocationScript, VerificationScript: verificationScript}
	witness._scriptHash = helper.UInt160FromBytes(crypto.Hash160(witness.VerificationScript))
	return witness, nil
}

// this is only used for empty VerificationScript, neo block chain will search the contract script with scriptHash
func CreateWitnessWithScriptHash(scriptHash *helper.UInt160, invocationScript []byte) (witness *Witness) {
	witness = &Witness{InvocationScript: invocationScript, VerificationScript: []byte{}, _scriptHash: scriptHash}
	return
}

// CreateContractWitness
func CreateContractWitness(msg []byte, pairs []keys.KeyPair, contract *sc.Contract) (*Witness, error) {
	invocationScript, err := CreateSignatureInvocation(msg, pairs)
	if err != nil {
		return nil, err
	}

	return CreateWitness(invocationScript, contract.Script)
}

// CreateSignatureInvocation pushes signature
func CreateSignatureInvocation(msg []byte, pairs []keys.KeyPair) ([]byte, error) {
	// invocationScript: push signature
	sort.Sort(keys.KeyPairSlice(pairs)) // sort in ascending order

	builder := sc.NewScriptBuilder()
	for _, pair := range pairs {
		signature, err := pair.Sign(msg)
		if err != nil {
			return nil, err
		}
		builder.EmitPushBytes(signature)
	}
	return builder.ToArray()
}
