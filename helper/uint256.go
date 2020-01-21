package helper

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

const UINT256SIZE = 32

// UInt256 is a 32 byte long unsigned integer.
type UInt256 [UINT256SIZE]uint8

// UInt256FromString attempts to decode the given string (in BE representation) into an UInt256.
func UInt256FromString(s string) (u UInt256, err error) {
	s = strings.TrimPrefix(s, "0x")
	if len(s) != UINT256SIZE*2 {
		return u, fmt.Errorf("expected string size of %d got %d", UINT256SIZE*2, len(s))
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return u, err
	}
	return UInt256FromBytes(ReverseBytes(b))
}

// UInt256FromBytes attempts to decode the given bytes (in LE representation) into an UInt256.
func UInt256FromBytes(b []byte) (u UInt256, err error) {
	if len(b) != UINT256SIZE {
		return u, fmt.Errorf("expected []byte of size %d got %d", UINT256SIZE, len(b))
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

func (u UInt256) Less(other UInt256) bool {
	for k := len(u.Bytes()) - 1; k >= 0; k-- {
		if u[k] == other[k] {
			continue
		}
		return u[k] < other[k]
	}
	return false
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

type UInt256Slice []UInt256

func (us UInt256Slice) Len() int {
	return len(us)
}

func (us UInt256Slice) Less(i int, j int) bool {
	return us[i].Less(us[j])
}

func (us UInt256Slice) Swap(i, j int) {
	t := us[i]
	us[i] = us[j]
	us[j] = t
}

func (us UInt256Slice) GetVarSize() int {
	return GetVarSize(len(us)) + len(us)*UINT256SIZE
}
