package keys

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateKeyPair(t *testing.T) {
	pair, err := GenerateKeyPair()
	assert.Nil(t, err)
	assert.NotNil(t, pair)
}

func TestNewKeyPair(t *testing.T) {
	for _, testCase := range KeyCases {
		keyPair, err := NewKeyPair(helper.HexToBytes(testCase.PrivateKey))
		assert.Nil(t, err)

		publicKey := keyPair.PublicKey.String()
		assert.Equal(t, testCase.PublicKey, publicKey)
	}
}

func TestNewKeyPairFromNEP2(t *testing.T) {
	for _, testCase := range KeyCases {
		keyPair, err := NewKeyPairFromNEP2(testCase.Nep2key, testCase.Passphrase, N, R, P)
		assert.Nil(t, err)

		privateKey := keyPair.String()
		publicKey := keyPair.PublicKey.String()

		assert.Nil(t, err)
		assert.Equal(t, testCase.PrivateKey, privateKey)
		assert.Equal(t, testCase.PublicKey, publicKey)
	}
}

func TestNewKeyPairFromNEP2_2(t *testing.T) {
	keyPair, err := NewKeyPairFromNEP2("6PYKfoRbX69ttWsoe1gM5GARbD1G6tBSurQ84AkjkaKR4TGwbRyZ3ZC5ie", "t", 16384, 8, 8)
	assert.Nil(t, err)

	publicKey := keyPair.PublicKey.String()

	assert.Nil(t, err)
	assert.Equal(t, "031f8ea9859b304a90128fcd85164dcf5de9abb7c950d761a25b43dbae453b0275", publicKey)

}

func TestNewKeyPairFromWIF(t *testing.T) {
	for _, testCase := range KeyCases {
		keyPair, err := NewKeyPairFromWIF(testCase.Wif)
		assert.Nil(t, err)

		privateKey := keyPair.String()
		publicKey := keyPair.PublicKey.String()

		assert.Equal(t, testCase.PrivateKey, privateKey)
		assert.Equal(t, testCase.PublicKey, publicKey)
	}
}

func TestKeyPair_CompareTo(t *testing.T) {
	privateKey1 := helper.HexToBytes(KeyCases[0].PrivateKey)
	pair1, err := NewKeyPair(privateKey1)
	assert.Nil(t, err)
	privateKey2 := helper.HexToBytes(KeyCases[1].PrivateKey)
	pair2, err := NewKeyPair(privateKey2)
	assert.Nil(t, err)
	assert.Equal(t, 1, pair1.CompareTo(pair2))
}

func TestKeyPair_ExistsIn(t *testing.T) {
	privateKey1 := helper.HexToBytes(KeyCases[0].PrivateKey)
	pair1, err := NewKeyPair(privateKey1)
	assert.Nil(t, err)
	privateKey2 := helper.HexToBytes(KeyCases[1].PrivateKey)
	pair2, err := NewKeyPair(privateKey2)
	assert.Nil(t, err)
	a := []KeyPair{*pair1, *pair2}
	assert.Equal(t, true, pair1.ExistsIn(a))
}

func TestKeyPair_Export(t *testing.T) {
	for _, testCase := range KeyCases {
		keyPair, err := NewKeyPair(helper.HexToBytes(testCase.PrivateKey))
		assert.Nil(t, err)

		wif := keyPair.Export()
		assert.Equal(t, testCase.Wif, wif)
	}
}

func TestKeyPair_ExportWithPassword(t *testing.T) {
	for _, testCase := range KeyCases {
		keyPair, err := NewKeyPair(helper.HexToBytes(testCase.PrivateKey))
		assert.Nil(t, err)

		nep2, err := keyPair.ExportWithPassword(testCase.Passphrase, N, R, P)
		assert.Nil(t, err)
		assert.Equal(t, testCase.Nep2key, nep2)
	}
}

func TestKeyPair_Sign(t *testing.T) {
	var data = []byte("sample")
	keyPair, err := GenerateKeyPair()
	assert.Nil(t, err)
	signedData, err := keyPair.Sign(data)
	assert.Nil(t, err)
	pubKey := keyPair.PublicKey
	result := VerifySignature(data, signedData, pubKey)
	assert.Equal(t, true, result)
}

func TestKeyPair_String(t *testing.T) {
	for _, testCase := range KeyCases {
		keyPair, err := NewKeyPair(helper.HexToBytes(testCase.PrivateKey))
		assert.Nil(t, err)

		assert.Equal(t, testCase.PrivateKey, keyPair.String())
	}
}

func TestVerifySignature(t *testing.T) {
	var data = []byte("sample")
	keyPair, err := GenerateKeyPair()
	assert.Nil(t, err)
	signedData, err := keyPair.Sign(data)
	assert.Nil(t, err)
	pubKey := keyPair.PublicKey
	result := VerifySignature(data, signedData, pubKey)
	assert.Equal(t, true, result)
}

func TestVerifySignature2(t *testing.T) {
	keyPair, _ := GenerateKeyPair()
	sample := []byte("sample")
	signedData, _ := keyPair.Sign(sample)

	secondKeyPair, _ := GenerateKeyPair()
	wrongPubKey := secondKeyPair.PublicKey

	actual := VerifySignature(sample, signedData, wrongPubKey)
	assert.Equal(t, false, actual)
}
