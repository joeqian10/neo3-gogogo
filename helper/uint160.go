package helper

import (
	"encoding/hex"
	"encoding/json"
	"github.com/joeqian10/neo3-gogogo/io"
	"strings"
)

const UINT160SIZE = 20
const UINT64SIZE = 8
const UINT32SIZE = 4

var UInt160Zero = NewUInt160()

/// This class stores a 160 bit unsigned int, represented as a 20-byte little-endian byte array
/// It is composed by ulong(64) + ulong(64) + uint(32) = UInt160(160)
type UInt160 struct {
	Value1 uint64
	Value2 uint64
	Value3 uint32
}

func NewUInt160() *UInt160 {
	return &UInt160{}
}

// UInt160FromBytes attempts to decode the given little endian bytes into an UInt160.
func UInt160FromBytes(b []byte) *UInt160 {
	var r []byte
	if b == nil {
		r = make([]byte, UINT160SIZE)
	} else if len(b) < UINT160SIZE {
		r = PadRight(b, UINT160SIZE)
	} else {
		r = b[:UINT160SIZE]
	}

	return &UInt160{
		Value1: BytesToUInt64(r[:UINT64SIZE]),
		Value2: BytesToUInt64(r[UINT64SIZE:UINT64SIZE*2]),
		Value3: BytesToUInt32(r[UINT64SIZE*2:]),
	}
}

// UInt160FromString attempts to decode the given big endian string into an UInt160.
func UInt160FromString(s string) (*UInt160, error) {
	s = strings.TrimPrefix(s, "0x")
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return UInt160FromBytes(ReverseBytes(b)), nil
}

/// Method CompareTo returns 1 if this UInt160 is bigger than other UInt160; -1 if it's smaller; 0 if it's equals
/// Example: assume this is 01ff00ff00ff00ff00ff00ff00ff00ff00ff00a4, this.CompareTo(02ff00ff00ff00ff00ff00ff00ff00ff00ff00a3) returns 1
func (u *UInt160) CompareTo(other *UInt160) int {
	r := CompareTo(uint64(u.Value3), uint64(other.Value3))
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

// Equals returns true if both UInt160 values are the same.
func (u *UInt160) Equals(other *UInt160) bool {
	if other == nil {
		return false
	}
	return u.CompareTo(other) == 0
}

// Less returns true if this value is less than given UInt160 value. It's
// primarily intended to be used for sorting purposes.
func (u *UInt160) Less(other *UInt160) bool {
	return u.CompareTo(other) == -1
}

func (u *UInt160) Deserialize(br *io.BinaryReader) {
	br.ReadLE(&u.Value1)
	br.ReadLE(&u.Value2)
	br.ReadLE(&u.Value3)
}

func (u *UInt160) Serialize(bw *io.BinaryWriter) {
	bw.WriteLE(u.Value1)
	bw.WriteLE(u.Value2)
	bw.WriteLE(u.Value3)
}

// String implements the stringer interface. Return big endian hex string.
func (u UInt160) String() string {
	return hex.EncodeToString(ReverseBytes(u.ToByteArray()))
}

// ToByteArray returns the little endian byte slice representation of u.
func (u *UInt160) ToByteArray() []byte {
	b, e := io.ToArray(u)
	if e != nil {
		return nil
	}
	return b
}

// Size returns the size of a UInt160 object in byte
func (u *UInt160) Size() int {
	return UINT160SIZE
}

// UnmarshalJSON implements the json unmarshaller interface.
func (u *UInt160) UnmarshalJSON(data []byte) (err error) {
	var js string
	if err = json.Unmarshal(data, &js); err != nil {
		return err
	}
	js = strings.TrimPrefix(js, "0x")
	v, err := UInt160FromString(js)
	*u = *v
	return err
}

// MarshalJSON implements the json marshaller interface.
func (u *UInt160) MarshalJSON() ([]byte, error) {
	return []byte(`"0x` + u.String() + `"`), nil
}

// ExistsIn checks if u exists in list
func (u *UInt160) ExistsIn(list []UInt160) bool {
	for _, a := range list {
		if (u).Equals(&a) {
			return true
		}
	}
	return false
}

type UInt160Slice []UInt160

func (us UInt160Slice) Len() int {
	return len(us)
}

func (us UInt160Slice) Less(i int, j int) bool {
	return us[i].Less(&us[j])
}

func (us UInt160Slice) Swap(i, j int) {
	t := us[i]
	us[i] = us[j]
	us[j] = t
}

func (us UInt160Slice) GetVarSize() int {
	return GetVarSize(len(us)) + len(us)*UINT160SIZE
}
