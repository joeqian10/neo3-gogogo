package keys

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNEP2Encrypt(t *testing.T) {
	for _, testCase := range KeyCases {
		b, _ := hex.DecodeString(testCase.PrivateKey)

		keyPair, err := NewKeyPair(b)
		assert.Nil(t, err)

		nep2Key, err := NEP2Encrypt(keyPair, testCase.Passphrase)
		assert.Nil(t, err)

		assert.Equal(t, testCase.Nep2key, nep2Key)
	}
}

func TestNEP2Decrypt(t *testing.T) {
	for _, testCase := range KeyCases {

		keyPair, err := NEP2Decrypt(testCase.Nep2key, testCase.Passphrase)
		assert.Nil(t, err)

		assert.Equal(t, testCase.PrivateKey, keyPair.String())

		wif := keyPair.ExportWIF()
		assert.Equal(t, testCase.Wif, wif)

		address := keyPair.PublicKey.Address()
		assert.Equal(t, testCase.Address, address)
	}
}
