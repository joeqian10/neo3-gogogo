package tx

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/wallet/keys"
)

type WitnessSlice []*Witness
func (ws WitnessSlice) Len() int           { return len(ws) }
func (ws WitnessSlice) Less(i, j int) bool { return ws[i]._scriptHash.Less(ws[j]._scriptHash) }
func (ws WitnessSlice) Swap(i, j int)      { ws[i], ws[j] = ws[j], ws[i] }

// Witness
type Witness struct {
	InvocationScript   []byte         // signature
	VerificationScript []byte         // pub key
	_scriptHash        helper.UInt160 // script hash
}

func (w *Witness) Size() int {
	return sc.ByteSlice(w.InvocationScript).GetVarSize() + sc.ByteSlice(w.VerificationScript).GetVarSize()
}

func (w *Witness) GetScriptHash() helper.UInt160 {
	w._scriptHash, _ = helper.UInt160FromBytes(crypto.Hash160(w.VerificationScript))
	return w._scriptHash
}

// Deserialize implements Serializable interface.
func (w *Witness) Deserialize(br *io.BinaryReader) {
	// This is designed to allow a MultiSig 10/10 (around 1003 bytes) ~1024 bytes
	// Invocation = 10 * 64 + 10 = 650 ~ 664  (exact is 653)
	w.InvocationScript = br.ReadVarBytes(663)
	// Verification = 10 * 33 + 10 = 340 ~ 360   (exact is 351)
	w.VerificationScript = br.ReadVarBytes(361)
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
func CreateWitness(invocationScript []byte, verificationScript []byte) (witness *Witness, err error) {
	if len(verificationScript) == 0 {
		return nil, fmt.Errorf("verificationScript should not be empty")
	}
	witness = &Witness{InvocationScript: invocationScript, VerificationScript: verificationScript}
	witness._scriptHash, err = helper.UInt160FromBytes(crypto.Hash160(witness.VerificationScript))
	return
}

// this is only used for empty VerificationScript, neo block chain will search the contract script with scriptHash
func CreateWitnessWithScriptHash(scriptHash helper.UInt160, invocationScript []byte) (witness *Witness) {
	witness = &Witness{InvocationScript: invocationScript, VerificationScript: []byte{}, _scriptHash: scriptHash}
	return
}

// CreateContractWitness
func CreateContractWitness(msg []byte, pairs []*keys.KeyPair, contract *sc.Contract) (witness *Witness, err error) {
	invocationScript, err := CreateSignatureInvocation(msg, pairs)
	if err != nil {
		return witness, err
	}

	return CreateWitness(invocationScript, contract.Script)
}

// CreateSignatureInvocation pushes signature
func CreateSignatureInvocation(msg []byte, pairs []*keys.KeyPair) (invocationScript []byte, err error) {
	// invocationScript: push signature
	sort.Sort(sort.Reverse(keys.KeyPairSlice(pairs))) // sort in descending order

	builder := sc.NewScriptBuilder()
	for _, pair := range pairs {
		signature, err := pair.Sign(msg)
		if err != nil {
			return invocationScript, err
		}
		err = builder.EmitPushBytes(signature)
		if err != nil {
			return invocationScript, err
		}
	}
	return builder.ToArray(), nil
}
