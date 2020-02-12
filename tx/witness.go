package tx

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/wallet/keys"
	"sort"
)

// Witness
type Witness struct {
	InvocationScript   []byte         // signature
	VerificationScript []byte         // pub key
	_scriptHash         helper.UInt160 // script hash
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

// create single signature witness
func CreateSignatureWitness(msg []byte, pair *keys.KeyPair) (witness *Witness, err error) {
	// 	invocationScript: push signature
	signature, err := pair.Sign(msg)
	if err != nil {
		return
	}
	builder := sc.NewScriptBuilder()
	_ = builder.EmitPushBytes(signature)
	invocationScript := builder.ToArray()

	// verificationScript: SignatureRedeemScript
	verificationScript := keys.CreateSignatureRedeemScript(pair.PublicKey)
	return CreateWitness(invocationScript, verificationScript)
}

// create multi-signature witness
func CreateMultiSignatureWitness(msg []byte, pairs []*keys.KeyPair, least int, publicKeys []*keys.PublicKey) (witness *Witness, err error) {
	// TODO ensure the pairs match with publicKeys
	if len(pairs) == least {
		return witness, fmt.Errorf("the multi-signature contract needs least %v signatures", least)
	}
	// invocationScript: push signature
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].PublicKey.Compare(pairs[j].PublicKey) == 1
	})
	builder := sc.NewScriptBuilder()
	for _, pair := range pairs {
		signature, err := pair.Sign(msg)
		if err != nil {
			return witness, err
		}
		err = builder.EmitPushBytes(signature)
		if err != nil {
			return witness, err
		}
	}
	invocationScript := builder.ToArray()

	// verificationScript: CreateMultiSigRedeemScript
	verificationScript, _ := keys.CreateMultiSigRedeemScript(least, publicKeys...)
	return CreateWitness(invocationScript, verificationScript)
}
