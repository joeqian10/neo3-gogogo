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


// MaxInvocationScript This is designed to allow a MultiSig 21/11 (committee)
// Invocation = 11 * (64 + 2) = 726
const MaxInvocationScript = 1024

// MaxVerificationScript Verification = m + (PUSH_PubKey * 21) + length + null + syscall = 1 + ((2 + 33) * 21) + 2 + 1 + 5 = 744
const MaxVerificationScript = 1024

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

// CreateWitness with invocationScript and verificationScript
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

// create single signature witness
func CreateContractWitness(msg []byte, pairs []keys.KeyPair, contract *sc.Contract) (*Witness, error) {
	invocationScript, err := CreateInvocationScript(msg, pairs)
	if err != nil {
		return nil, err
	}

	return CreateWitness(invocationScript, contract.Script)
}

// CreateInvocationScript pushes signature
func CreateInvocationScript(msg []byte, pairs []keys.KeyPair) ([]byte, error) {
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

func CreateSignatureWitness(msg []byte, pair *keys.KeyPair) (*Witness, error) {
	// 	invocationScript: push signature
	signature, err := pair.Sign(msg)

	if err != nil {
		return nil, err
	}
	sb := sc.NewScriptBuilder()
	sb.EmitPushBytes(signature)
	invocationScript, err := sb.ToArray() // length 66
	if err != nil {
		return nil, err
	}

	// verificationScript: SignatureRedeemScript
	verificationScript, err := sc.CreateSignatureRedeemScript(pair.PublicKey)
	if err != nil {
		return nil, err
	}
	return CreateWitness(invocationScript, verificationScript)
}

// create multi-signature witness
func CreateMultiSignatureWitness(msg []byte, pairs []keys.KeyPair, least int, publicKeys []crypto.ECPoint) (*Witness, error) {
	if len(pairs) < least {
		return nil, fmt.Errorf("the multi-signature contract needs least %v signatures", least)
	}
	// invocationScript: push signature
	keyPairs := keys.KeyPairSlice(pairs)
	sort.Sort(keyPairs) // ascending

	sb := sc.NewScriptBuilder()
	for _, pair := range keyPairs {
		signature, err := pair.Sign(msg)
		if err != nil {
			return nil, err
		}
		sb.EmitPushBytes(signature)
	}
	invocationScript, err := sb.ToArray()
	if err != nil {
		return nil, err
	}

	// verificationScript: CreateMultiSigRedeemScript
	verificationScript, _ := sc.CreateMultiSigRedeemScript(least, publicKeys)
	return CreateWitness(invocationScript, verificationScript)
}

func VerifySignatureWitness(msg []byte, witness *Witness) bool {
	invocationScript := witness.InvocationScript
	if len(invocationScript) != 66 {
		return false
	}
	if invocationScript[0] != 0x0c || invocationScript[1] != 0x40 {
		return false
	}
	signature := invocationScript[2:] // length 64

	verificationScript := witness.VerificationScript
	if len(verificationScript) != 40 {
		return false
	}
	data := verificationScript[:35] // length 35
	publicKey, _ := crypto.NewECPointFromBytes(data[2:]) // length 33
	return keys.VerifySignature(msg, signature, publicKey)
}

func VerifyMultiSignatureWitness(msg []byte, witness *Witness) bool {
	invocationScript := witness.InvocationScript
	lenInvoScript := len(invocationScript)
	if lenInvoScript%66 != 0 {
		return false
	}
	m := lenInvoScript / 66 // m signatures

	verificationScript := witness.VerificationScript
	least := verificationScript[0] - byte(sc.PUSH1) + 1 // least required signatures, limited to 16 here

	if m < int(least) {
		return false
	} // not enough signatures
	var signatures = make([][]byte, m)
	for i := 0; i < m; i++ {
		signatures[i] = invocationScript[i*66+2 : i*66+66] // signature length is 64
	}

	lenVeriScript := len(verificationScript)
	n := verificationScript[lenVeriScript-6] - byte(sc.PUSH1) + 1 // public keys, limited to 16 here

	if m > int(n) {
		return false
	} // too many signatures

	var pubKeys = make([]crypto.ECPoint, n)
	for i := 0; i < int(n); i++ {
		data := verificationScript[i*35+1 : i*35+36] // length 35
		publicKey, _ := crypto.NewECPointFromBytes(data[2:]) // length 33
		pubKeys[i] = *publicKey
	}
	return keys.VerifyMultiSig(msg, signatures, pubKeys)
}
