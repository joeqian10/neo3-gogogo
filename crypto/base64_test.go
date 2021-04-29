package crypto

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	bs := []byte{0x01, 0x02, 0x05, 0x00}
	encoded := Base64Encode(bs)
	log.Println(encoded)
}

func TestBase64Decode(t *testing.T) {
	s := "AoFykYVAsrUS6uGHKiouOijZicYNldq4gprafX3XBtZY"
	decoded, err := Base64Decode(s)
	assert.Nil(t, err)
	ss := helper.BytesToHex(decoded)
	log.Println(ss)
}

func TestBase64Decode2(t *testing.T) {
	s := "A4uK9iEOz9y8qyJVLvjYz0HG+G+c+atT2GV0HP24M/Br"
	decoded, err := Base64Decode(s)
	assert.Nil(t, err)
	ss := helper.BytesToHex(decoded)
	log.Println(ss)
}
