package helper

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

const UINT160SIZE = 20

// UInt160 is a 20 byte long unsigned integer. little endian
type UInt160 [UINT160SIZE]uint8

// UInt160FromString attempts to decode the given big endian string into an UInt160.
func UInt160FromString(s string) (UInt160, error) {
	var u UInt160
	s = strings.TrimPrefix(s, "0x")
	if len(s) != UINT160SIZE*2 {
		return u, fmt.Errorf("expected string size of %d got %d", UINT160SIZE*2, len(s))
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return u, err
	}
	return UInt160FromBytes(ReverseBytes(b))
}

// UInt160FromBytes attempts to decode the given bytes into an UInt160.
func UInt160FromBytes(b []byte) (u UInt160, err error) {
	if len(b) != UINT160SIZE {
		return u, fmt.Errorf("expected byte size of %d got %d", UINT160SIZE, len(b))
	}
	copy(u[:], b)
	return
}

// Bytes returns the little endian byte slice representation of u.
func (u UInt160) Bytes() []byte {
	return u[:]
}

// String implements the stringer interface. Return big endian hex string.
func (u UInt160) String() string {
	return hex.EncodeToString(ReverseBytes(u.Bytes()))
}

// Equals returns true if both UInt256 values are the same.
func (u UInt160) Equals(other UInt160) bool {
	return u == other
}

// Less returns true if this value is less than given UInt160 value. It's
// primarily intended to be used for sorting purposes.
func (u UInt160) Less(other UInt160) bool {
	for k := len(u.Bytes()) - 1; k >= 0; k-- {
		if u[k] == other[k] {
			continue
		}
		return u[k] < other[k]
	}
	return false
}

// UnmarshalJSON implements the json unmarshaller interface.
func (u *UInt160) UnmarshalJSON(data []byte) (err error) {
	var js string
	if err = json.Unmarshal(data, &js); err != nil {
		return err
	}
	js = strings.TrimPrefix(js, "0x")
	*u, err = UInt160FromString(js)
	return err
}

// MarshalJSON implements the json marshaller interface.
func (u UInt160) MarshalJSON() ([]byte, error) {
	return []byte(`"0x` + u.String() + `"`), nil
}

type UInt160Slice []UInt160

func (us UInt160Slice) Len() int {
	return len(us)
}

func (us UInt160Slice) Less(i int, j int) bool {
	return us[i].Less(us[j])
}

func (us UInt160Slice) Swap(i, j int) {
	t := us[i]
	us[i] = us[j]
	us[j] = t
}

func (us UInt160Slice) GetVarSize() int {
	return GetVarSize(len(us)) + len(us)*UINT160SIZE
}
