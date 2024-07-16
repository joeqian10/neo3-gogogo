package crypto

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joeqian10/neo3-gogogo/helper"
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
	s := "U2FsdGVkX1+iX5Ey7GqLND5UFUoV0b7rUJ2eEvHkYqA="
	decoded, err := Base64Decode(s)
	assert.Nil(t, err)
	ss := helper.BytesToHex(decoded)
	log.Println(ss)
}

func TestBase64Decode3(t *testing.T) {
	s := "VHJhbnNhY3Rpb24gaGFzIGJlZW4gZXhlY3V0ZWQ="
	decoded, err := Base64Decode(s)
	assert.Nil(t, err)
	ss := string(decoded)
	assert.Equal(t, "Transaction has been executed", ss)
}

func TestBase64Decode4(t *testing.T) {
	s := "FPHOAtHWr86lTMPb00238i9KzaWhFK17JowKjp28yE2ZYrJNkaYtDc1e05dZaMvZ1w0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD9IQMEbQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfuhtBlhDF/iAa1Gfzmowxg8BULN+6G0GWEMX+IBrUZ/OajDGDwFQs37obQZYQxf4gGtRn85qMMYPAVCzfug="
	decoded, err := Base64Decode(s)
	assert.Nil(t, err)
	log.Println(helper.BytesToHex(decoded))
}

func TestBase64Decode5(t *testing.T) {
	s := "TW/Nz+UkOwvhddhFmcl3fCXb3Uw="
	decoded, err := Base64Decode(s)
	assert.Nil(t, err)
	ss := helper.BytesToHex(decoded)
	log.Println(ss)
}
