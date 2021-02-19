package helper

import (
	"encoding/hex"
	"encoding/json"
	"github.com/joeqian10/neo3-gogogo/io"
	"strings"
)

const UINT256SIZE = 32

var UInt256Zero = NewUInt256()

/// This class stores a 256 bit unsigned int, represented as a 32-byte little-endian byte array
/// Composed by ulong(64) + ulong(64) + ulong(64) + ulong(64) = UInt256(256)
type UInt256 struct {
	Value1 uint64
	Value2 uint64
	Value3 uint64
	Value4 uint64
}

func NewUInt256() *UInt256 {
	return &UInt256{}
}

// UInt256FromBytes attempts to decode the given bytes (in LE representation) into an UInt256.
func UInt256FromBytes(b []byte) *UInt256 {
	var r []byte
	if b == nil {
		r = make([]byte, UINT256SIZE)
	} else if len(b) < UINT256SIZE {
		r = PadRight(b, UINT256SIZE)
	} else {
		r = b[:UINT256SIZE]
	}

	return &UInt256{
		Value1: BytesToUInt64(r[:UINT64SIZE]),
		Value2: BytesToUInt64(r[UINT64SIZE:UINT64SIZE*2]),
		Value3: BytesToUInt64(r[UINT64SIZE*2:UINT64SIZE*3]),
		Value4: BytesToUInt64(r[UINT64SIZE*3:]),
	}
}

// UInt256FromString attempts to decode the given string (in BE representation) into an UInt256.
func UInt256FromString(s string) (u *UInt256, err error) {
	s = strings.TrimPrefix(s, "0x")
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return UInt256FromBytes(ReverseBytes(b)), nil
}

/// Method CompareTo returns 1 if this UInt256 is bigger than other UInt256; -1 if it's smaller; 0 if it's equals
/// Example: assume this is 01ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00a4, this.CompareTo(02ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00ff00a3) returns 1
func (u *UInt256) CompareTo(other *UInt256) int {
	r := CompareTo(u.Value4, other.Value4)
	if r != 0 {return r}
	r = CompareTo(u.Value3, other.Value3)
	if r != 0 {
		return r
	}
	r = CompareTo(u.Value2, other.Value2)
	if r != 0 {
		return r
	}
	r = CompareTo(u.Value1, other.Value1)
	return r
}

// Equals returns true if both UInt256 values are the same.
func (u *UInt256) Equals(other *UInt256) bool {
	if other == nil {
		return false
	}
	return u.CompareTo(other) == 0
}

func (u *UInt256) Less(other *UInt256) bool {
	return u.CompareTo(other) == -1
}

func (u *UInt256) Deserialize(br *io.BinaryReader) {
	br.ReadLE(&u.Value1)
	br.ReadLE(&u.Value2)
	br.ReadLE(&u.Value3)
	br.ReadLE(&u.Value4)
}

func (u *UInt256) Serialize(bw *io.BinaryWriter) {
	bw.WriteLE(u.Value1)
	bw.WriteLE(u.Value2)
	bw.WriteLE(u.Value3)
	bw.WriteLE(u.Value4)
}

// String implements the stringer interface. Return big endian hex string.
func (u UInt256) String() string {
	return hex.EncodeToString(ReverseBytes(u.ToByteArray()))
}

// ToByteArray returns a byte slice representation of u.
func (u *UInt256) ToByteArray() []byte {
	b, e := io.ToArray(u)
	if e != nil {
		return nil
	}
	return b
}

// Size returns the size of a UInt256 object in byte
func (u *UInt256) Size() int {
	return UINT256SIZE
}

// UnmarshalJSON implements the json unmarshaller interface.
func (u *UInt256) UnmarshalJSON(data []byte) (err error) {
	var js string
	if err = json.Unmarshal(data, &js); err != nil {
		return err
	}
	js = strings.TrimPrefix(js, "0x")
	v, err := UInt256FromString(js)
	*u = *v
	return err
}

// MarshalJSON implements the json marshaller interface.
func (u UInt256) MarshalJSON() ([]byte, error) {
	return []byte(`"0x` + u.String() + `"`), nil
}

// ExistsIn checks if u exists in list
func (u UInt256) ExistsIn(list []UInt256) bool {
	for _, a := range list {
		if (&u).Equals(&a) {
			return true
		}
	}
	return false
}

type UInt256Slice []UInt256

func (us UInt256Slice) Len() int {
	return len(us)
}

func (us UInt256Slice) Less(i int, j int) bool {
	return (&us[i]).Less(&us[j])
}

func (us UInt256Slice) Swap(i, j int) {
	t := us[i]
	us[i] = us[j]
	us[j] = t
}

func (us UInt256Slice) GetVarSize() int {
	return GetVarSize(len(us)) + len(us)*UINT256SIZE
}
