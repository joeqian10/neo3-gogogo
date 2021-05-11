package crypto

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAddressToScriptHash(t *testing.T) {
	r, err := AddressToScriptHash("NdtB8RXRmJ7Nhw1FPTm7E6HoDZGnDw37nf", helper.DefaultAddressVersion)
	assert.Nil(t, err)
	u := helper.UInt160FromBytes(Hash160([]byte{0x01}))
	assert.Equal(t, u.String(), r.String())
}

func TestAddressToScriptHash2(t *testing.T) {
	r, err := AddressToScriptHash("NbG6HCirXABhtAakkJPsFhzsVFVgC3xuCT", helper.DefaultAddressVersion)
	assert.Nil(t, err)
	log.Println(r.String())
}

func TestScriptHashToAddress(t *testing.T) {
	u := helper.UInt160FromBytes(Hash160([]byte{0x01}))
	a := ScriptHashToAddress(u, helper.DefaultAddressVersion)
	assert.Equal(t, "NdtB8RXRmJ7Nhw1FPTm7E6HoDZGnDw37nf", a)
}
