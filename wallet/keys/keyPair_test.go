package keys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExportWIF(t *testing.T) {
	for _, testCase := range KeyCases {
		keyPair, err := NewKeyPairFromWIF(testCase.Wif)
		assert.Nil(t, err)

		privateKey := keyPair.String()
		publicKey := keyPair.PublicKey.String()
		wif := keyPair.ExportWIF()
		nep2, err := keyPair.ExportNep2(testCase.Passphrase)

		assert.Nil(t, err)
		assert.Equal(t, testCase.PrivateKey, privateKey)
		assert.Equal(t, testCase.PublicKey, publicKey)
		assert.Equal(t, testCase.Wif, wif)
		assert.Equal(t, testCase.Nep2key, nep2)
	}
}

func TestExportNEP2(t *testing.T) {
	for _, testCase := range KeyCases {
		keyPair, err := NewKeyPairFromNEP2(testCase.Nep2key, testCase.Passphrase)
		assert.Nil(t, err)

		privateKey := keyPair.String()
		publicKey := keyPair.PublicKey.String()
		wif := keyPair.ExportWIF()
		nep2, err := keyPair.ExportNep2(testCase.Passphrase)

		assert.Nil(t, err)
		assert.Equal(t, testCase.PrivateKey, privateKey)
		assert.Equal(t, testCase.PublicKey, publicKey)
		assert.Equal(t, testCase.Wif, wif)
		assert.Equal(t, testCase.Nep2key, nep2)
	}
}

func TestPubKeyVerify(t *testing.T) {
	var data = []byte("sample")
	keyPair, err := GenerateKeyPair()
	assert.Nil(t, err)
	signedData, err := keyPair.Sign(data)
	assert.Nil(t, err)
	pubKey := keyPair.PublicKey
	result := VerifySignature(data, signedData, pubKey)
	assert.Equal(t, true, result)
}

func TestWrongPubKey(t *testing.T) {
	keyPair, _ := GenerateKeyPair()
	sample := []byte("sample")
	signedData, _ := keyPair.Sign(sample)

	secondKeyPair, _ := GenerateKeyPair()
	wrongPubKey := secondKeyPair.PublicKey

	actual := VerifySignature(sample, signedData, wrongPubKey)
	assert.Equal(t, false, actual)
}
