package keys

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	for _, testCase := range KeyCases {
		p, err := NewPublicKeyFromString(testCase.PublicKey)
		assert.Nil(t, err)

		address := p.Address()
		assert.Equal(t, testCase.Address, address)

		scripthash := p.ScriptHash()
		assert.Equal(t, testCase.ScriptHash, scripthash.String())
	}
}

func TestCreateMultiSigRedeemScript(t *testing.T) {
	privateKey1, _ := hex.DecodeString(KeyCases[0].PrivateKey)
	privateKey2, _ := hex.DecodeString(KeyCases[1].PrivateKey)
	privateKey3, _ := hex.DecodeString(KeyCases[2].PrivateKey)

	keyPair1, _ := NewKeyPair(privateKey1)
	keyPair2, _ := NewKeyPair(privateKey2)
	keyPair3, _ := NewKeyPair(privateKey3)

	multiSignature, _ := CreateMultiSigRedeemScript(2, keyPair1.PublicKey, keyPair2.PublicKey, keyPair3.PublicKey)

	assert.Equal(t, "120c21027d73c8b02e446340caceee7a517cddff72440e60c28cbb84884f307760ecad5b0c21038a2151948a908cdf2d680eead6512217769e34b9db196574572cb98e273516a10c2103b7a7f933199f28cc1c48d22a21c78ac3992cf7fceb038a9c670fe55444426619130b413073b3bb", hex.EncodeToString(multiSignature))
}
