package helper

import (
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUInt256(t *testing.T) {
	u := NewUInt256()
	assert.NotNil(t, u)
	assert.Equal(t, uint64(0), u.Value1)
}

func TestUInt256FromBytes(t *testing.T) {
	b := make([]byte, UINT256SIZE)
	u := UInt256FromBytes(b)
	assert.NotNil(t, u)
	assert.Equal(t, "0000000000000000000000000000000000000000000000000000000000000000", u.String())
}

func TestUInt256FromString(t *testing.T) {
	s := "f037308fa0ab18155bccfc08485468c112409ea5064595699e98c545f245f32d"
	val, err := UInt256FromString(s)
	assert.Nil(t, err)
	assert.Equal(t, s, val.String())
}

func TestUInt256_CompareTo(t *testing.T) {
	tmp := make([]byte, UINT256SIZE)
	tmp[UINT256SIZE-1] = 0x01
	u := UInt256FromBytes(tmp)
	assert.Equal(t, 0, UInt256Zero.CompareTo(UInt256Zero))
	assert.Equal(t, 1, u.CompareTo(UInt256Zero))
	assert.Equal(t, -1, UInt256Zero.CompareTo(u))
}

func TestUInt256_Equals(t *testing.T) {
	tmp := make([]byte, UINT256SIZE)
	tmp[UINT256SIZE-1] = 0x01
	u := UInt256FromBytes(tmp)
	assert.Equal(t, true, UInt256Zero.Equals(UInt256Zero))
	assert.Equal(t, false, u.Equals(UInt256Zero))
	assert.Equal(t, false, u.Equals(nil))
}

func TestUInt256_Less(t *testing.T) {
	a := "f037308fa0ab18155bccfc08485468c112409ea5064595699e98c545f245f322"
	b := "f037308fa0ab18155bccfc08485468c112409ea5064595699e98c545f245f323"

	ua, err := UInt256FromString(a)
	assert.Nil(t, err)
	ua2, err := UInt256FromString(a)
	assert.Nil(t, err)
	ub, err := UInt256FromString(b)
	assert.Nil(t, err)
	assert.Equal(t, true, ua.Less(ub))
	assert.Equal(t, false, ua.Less(ua2))
	assert.Equal(t, false, ub.Less(ua))
}

func TestUInt256_String(t *testing.T) {
	s := "f037308fa0ab18155bccfc08485468c112409ea5064595699e98c545f245f322"
	u, err := UInt256FromString(s)
	assert.Nil(t, err)
	assert.Equal(t, s, u.String())
}

func TestUInt256_Size(t *testing.T) {
	s := UInt256Zero.Size()
	assert.Equal(t, UINT256SIZE, s)
}

func TestUInt256_ExistsIn(t *testing.T) {
	a := []UInt256{*UInt256Zero}
	b := UInt256Zero.ExistsIn(a)
	assert.Equal(t, true, b)
}

func TestUInt256_ToByteArray(t *testing.T) {
	s := "f037308fa0ab18155bccfc08485468c112409ea5064595699e98c545f245f322"
	u, err := UInt256FromString(s)
	assert.Nil(t, err)
	assert.Equal(t, HexToBytes(s), ReverseBytes(u.ToByteArray()))
}

func TestUInt256_Deserialize(t *testing.T) {
	s := "f037308fa0ab18155bccfc08485468c112409ea5064595699e98c545f245f322"
	b := ReverseBytes(HexToBytes(s))
	br := io.NewBinaryReaderFromBuf(b)
	u := NewUInt256()
	u.Deserialize(br)
	assert.Equal(t, s, u.String())
}

func TestUInt256_Serialize(t *testing.T) {
	s := "f037308fa0ab18155bccfc08485468c112409ea5064595699e98c545f245f322"
	u, err := UInt256FromString(s)
	assert.Nil(t, err)
	b := io.NewBufBinaryWriter()
	u.Serialize(b.BinaryWriter)
	assert.Equal(t, HexToBytes(s), ReverseBytes(b.Bytes()))
}

func TestUInt256UnmarshalJSON(t *testing.T) {
	str := "f037308fa0ab18155bccfc08485468c112409ea5064595699e98c545f245f32d"
	expected, err := UInt256FromString(str)
	assert.Nil(t, err)

	// UnmarshalJSON decodes hex-strings

	u1 := NewUInt256()
	err = u1.UnmarshalJSON([]byte(`"` + str + `"`))
	assert.Nil(t, err)
	assert.True(t, expected.Equals(u1))

	s, err := expected.MarshalJSON()
	assert.Nil(t, err)

	// UnmarshalJSON decodes hex-strings prefixed by 0x
	u2 := NewUInt256()
	if err = u2.UnmarshalJSON(s); err != nil {
		t.Fatal(err)
	}
	assert.True(t, expected.Equals(u1))
}
