package keys

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
)

func TestPublicKeyToScriptHash(t *testing.T) {
	for _, testCase := range KeyCases {
		pubKey, err := crypto.NewECPointFromString(testCase.PublicKey)
		assert.Nil(t, err)
		scriptHash := PublicKeyToScriptHash(pubKey)
		s := scriptHash.String()
		assert.Equal(t, testCase.ScriptHash, s)
	}
}

func TestPublicKeyToAddress(t *testing.T) {
	for _, testCase := range KeyCases {
		pubKey, err := crypto.NewECPointFromString(testCase.PublicKey)
		assert.Nil(t, err)
		address := PublicKeyToAddress(pubKey, helper.DefaultAddressVersion)
		assert.Equal(t, testCase.Address, address)
	}
}

func TestRecoverPubKeyFromSig(t *testing.T) {
	// here is my message
	msg := []byte("Hello World")

	// here is my private key
	privateKey := []byte{
		0xde, 0xed, 0xbe, 0xef, 0xde, 0xed, 0xbe, 0xef,
		0xde, 0xed, 0xbe, 0xef, 0xde, 0xed, 0xbe, 0xef,
		0xde, 0xed, 0xbe, 0xef, 0xde, 0xed, 0xbe, 0xef,
		0xde, 0xed, 0xbe, 0xef, 0xde, 0xed, 0xbe, 0xef,
	}

	pair, err := NewKeyPair(privateKey)
	assert.Nil(t, err)
	fmt.Println(pair.PublicKey.String())
	fmt.Println()

	count := 0
	for i := 0; i < 100; i++ {
		sig, _ := pair.Sign(msg)
		pubKeys, _ := RecoverPubKeyFromSigOnSecp256r1(msg, sig)
		match := false
		for _, pubKey := range pubKeys {
			if pair.PublicKey.String() == pubKey.String() {
				match = true
				break
			}
		}
		if !match {
			//fmt.Println(helper.BytesToHex(sig))
			//fmt.Println(pubKey.String())
			count++
		}
	}

	fmt.Println(count)
	//sig, err := pair.Sign(msg)
	//assert.Nil(t, err)
	//
	//fmt.Println(helper.BytesToHex(sig))
	//
	//pubKey, err := RecoverPubKeyFromSigOnSecp256r1(msg, sig)
	//assert.Nil(t, err)

	//assert.Equal(t, pair.PublicKey.String(), pubKey.String())
}

func TestVerifySignature3(t *testing.T) {
	msg := []byte("Hello World")
	sig := helper.HexToBytes("4c2c6fd93dd5ae2feb4857747383b6ebb82bcf8db84cb02a62dc8cc5999958e34add7302209ee9905cf49b365d475967bfcc3ff8a220b5914b83275e76648676")
	pubKey, err := crypto.NewECPointFromString("03f6838d75c7bdc8cbe3f79b2ffee75012abb79ef7c9f96dad713903cd420ebab8")
	assert.Nil(t, err)

	b := VerifySignature(msg, sig, pubKey)
	fmt.Println(b)
}

//func TestRecoverPubKeyFromSigOnSecp256r1(t *testing.T) {
//	msg := []byte("Hello World")
//	sig := helper.HexToBytes("4c2c6fd93dd5ae2feb4857747383b6ebb82bcf8db84cb02a62dc8cc5999958e34add7302209ee9905cf49b365d475967bfcc3ff8a220b5914b83275e76648676")
//	pubKey, _ := RecoverPubKeyFromSigOnSecp256r1(msg, sig)
//	fmt.Println(pubKey.String())
//}
