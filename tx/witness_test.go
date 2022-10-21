package tx

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/joeqian10/neo3-gogogo/keys"
	"github.com/stretchr/testify/assert"
	"testing"
)

var caseLen int = len(keys.KeyCases)

func TestWitness_Deserialize(t *testing.T) {
	s := "41" + "40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58" +
		"23" + "2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac"
	br := io.NewBinaryReaderFromBuf(helper.HexToBytes(s))
	w := Witness{}
	w.Deserialize(br)
	assert.Equal(t, "40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58", helper.BytesToHex(w.InvocationScript))
	assert.Equal(t, "2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac", helper.BytesToHex(w.VerificationScript))
}

func TestWitness_GetScriptHash(t *testing.T) {
	w := Witness{
		InvocationScript:   helper.HexToBytes("40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58"), //65
		VerificationScript: helper.HexToBytes("2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac"),                                                             //35
	}
	scriptHash := w.GetScriptHash()
	assert.Equal(t, "71cb588c8291c18fa87fa07ce16c3fd92ab5aa30", scriptHash.String())
}

func TestWitness_Serialize(t *testing.T) {
	w := Witness{
		InvocationScript:   helper.HexToBytes("40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58"), //65
		VerificationScript: helper.HexToBytes("2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac"),                                                             //35
	}
	bbw := io.NewBufBinaryWriter()
	w.Serialize(bbw.BinaryWriter)
	b := bbw.Bytes()
	assert.Equal(t, "41"+"40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58"+
		"23"+"2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac", helper.BytesToHex(b))
}

func TestWitness_Size(t *testing.T) {
	w := Witness{
		InvocationScript:   helper.HexToBytes("40915467ecd359684b2dc358024ca750609591aa731a0b309c7fb3cab5cd0836ad3992aa0a24da431f43b68883ea5651d548feb6bd3c8e16376e6e426f91f84c58"), //65
		VerificationScript: helper.HexToBytes("2103322f35c7819267e721335948d385fae5be66e7ba8c748ac15467dcca0693692dac"),                                                             //35
	}
	size := w.GetSize()
	assert.Equal(t, 1+65+1+35, size)
}

func TestCreateSignatureWitness(t *testing.T) {
	msg := []byte("sample")
	pair, _ := keys.NewKeyPairFromWIF(keys.KeyCases[0].Wif)
	witness, err := CreateSignatureWitness(msg, pair)
	assert.Nil(t, err)
	assert.Equal(t, 66, len(witness.InvocationScript))
	assert.Equal(t, 40, len(witness.VerificationScript))
}

func TestCreateMultiSignatureWitness(t *testing.T) {
	msg := []byte("sample")
	pairs := make([]*keys.KeyPair, caseLen)
	pubKeys := make([]*crypto.ECPoint, caseLen)
	for i := 0; i < caseLen; i++ {
		pair, _ := keys.NewKeyPairFromWIF(keys.KeyCases[i].Wif)
		pairs[i] = pair
		pubKeys[i] = pair.PublicKey
	}

	witness, err := CreateMultiSignatureWitness(msg, pairs[:caseLen-1], caseLen-1, pubKeys)
	assert.Nil(t, err)
	assert.Equal(t, 66*(caseLen-1), len(witness.InvocationScript))
	assert.Equal(t, 1+35*caseLen+1+5, len(witness.VerificationScript))
}

func TestVerifySignatureWitness(t *testing.T) {
	msg := []byte("sample")
	pair, _ := keys.NewKeyPairFromWIF(keys.KeyCases[0].Wif)
	witness, err := CreateSignatureWitness(msg, pair)
	b := VerifySignatureWitness(msg, witness)
	assert.Nil(t, err)
	assert.Equal(t, true, b)
}

func TestVerifyMultiSignatureWitness(t *testing.T) {
	msg := []byte("sample")

	pairs := make([]*keys.KeyPair, caseLen)
	pubKeys := make([]*crypto.ECPoint, caseLen)
	for i := 0; i < caseLen; i++ {
		pair, _ := keys.NewKeyPairFromWIF(keys.KeyCases[i].Wif)
		pairs[i] = pair
		pubKeys[i] = pair.PublicKey
	}
	witness, err := CreateMultiSignatureWitness(msg, pairs[:caseLen-1], caseLen-1, pubKeys)
	assert.Nil(t, err)

	b := VerifyMultiSignatureWitness(msg, witness)
	assert.Nil(t, err)
	assert.Equal(t, true, b)
}

func TestVerifyMultiSignatureWitness2(t *testing.T) {
	msg := helper.HexToBytes("3a454f0062365ed35c4ba19853e2aba246c8b42609da75503b8139f3b96dc9af359b1488")
	inv, err := crypto.Base64Decode("DEDZ6kN4T3okrkTcD+jmOyuvghXhr7U7r6Ttkgyrd1o5mxw1jKtAlj/6ZTrC42E/zERrNbjHp+G2v6njnrSJ7osBDEBPjxq+5rsquclbpIogtlFnc/tDQGC3QeWI443/4sgrJ7ssZoXHZxqNVm/q1lEshG7S3qv6tpDumxMnOSAQ8F4QDEACncLL35Wo+kURJiUv4rqzbUeuwciLI+8+rdvbR3sPOmeYh2hviF2FKpwaxgpaJ+rZxwTrew1IWJLsnKKQ1l8I")
	assert.Nil(t, err)
	ver, err := crypto.Base64Decode("EwwhAxPsv6UkXoRiWuh70MHrB5dHOlH2oXJ89glc1+9SaX0GDCEDVTSosWpgZS1C4uMwXFoy6EnvOlRxn1CahCfJwrDYsesMIQLoSLZZEYoSlcy7zwoCIkrGjb25fvDOrKF9Rcmvbv+HZAwhA/7G8sahf7WFskw8v+nWpjyHRUDFC0RdOk0I9WWyhPGaFEGe0Nw6")
	assert.Nil(t, err)
	witness, err := CreateWitness(inv, ver)
	assert.Nil(t, err)
	b := VerifyMultiSignatureWitness(msg, witness)
	assert.Equal(t, true, b)
}
