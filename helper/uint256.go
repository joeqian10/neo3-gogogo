package helper

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

const uint256Size = 32

// UInt256 is a 32 byte long unsigned integer.
type UInt256 [uint256Size]uint8

// UInt256FromString attempts to decode the given string (in BE representation) into an UInt256.
func UInt256FromString(s string) (u UInt256, err error) {
	s = strings.TrimPrefix(s, "0x")
	if len(s) != uint256Size*2 {
		return u, fmt.Errorf("expected string size of %d got %d", uint256Size*2, len(s))
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return u, err
	}
	return UInt256FromBytes(ReverseBytes(b))
}

// UInt256FromBytes attempts to decode the given bytes (in LE representation) into an UInt256.
func UInt256FromBytes(b []byte) (u UInt256, err error) {
	if len(b) != uint256Size {
		return u, fmt.Errorf("expected []byte of size %d got %d", uint256Size, len(b))
	}
	copy(u[:], b)
	return u, nil
}

// Bytes returns a byte slice representation of u.
func (u UInt256) Bytes() []byte {
	return u[:]
}

// Equals returns true if both UInt256 values are the same.
func (u UInt256) Equals(other UInt256) bool {
	return u == other
}

// String implements the stringer interface.
func (u UInt256) String() string {
	return hex.EncodeToString(ReverseBytes(u.Bytes()))
}

// UnmarshalJSON implements the json unmarshaller interface.
func (u *UInt256) UnmarshalJSON(data []byte) (err error) {
	var js string
	if err = json.Unmarshal(data, &js); err != nil {
		return err
	}
	js = strings.TrimPrefix(js, "0x")
	*u, err = UInt256FromString(js)
	return err
}

// MarshalJSON implements the json marshaller interface.
func (u UInt256) MarshalJSON() ([]byte, error) {
	return []byte(`"0x` + u.String() + `"`), nil
}

// CompareTo compares two UInt256 with each other. Possible output: 1, -1, 0
//  1 implies u > other.
// -1 implies u < other.
//  0 implies  u = other.
func (u UInt256) CompareTo(other UInt256) int {
	for k := len(u.Bytes()) - 1; k >= 0; k-- {
		if u[k] < other[k] {
			return -1
		} else {
			return 1
		}
	}
	return 0
}
