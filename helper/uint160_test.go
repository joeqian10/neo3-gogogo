package helper

import (
	"github.com/joeqian10/neo3-gogogo/io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUInt160(t *testing.T) {
	u := NewUInt160()
	assert.NotNil(t, u)
	assert.Equal(t, uint64(0), u.Value1)
}

func TestUInt160FromBytes(t *testing.T) {
	b := make([]byte, UINT160SIZE)
	u := UInt160FromBytes(b)
	assert.NotNil(t, u)
	assert.Equal(t, "0000000000000000000000000000000000000000", u.String())
}

func TestUInt160FromString(t *testing.T) {
	s := "2d3b96ae1bcc5a585e075e3b81920210dec16302"
	u, err := UInt160FromString(s)
	assert.Nil(t, err)
	assert.Equal(t, s, u.String())
}

func TestUInt160_CompareTo(t *testing.T) {
	tmp := make([]byte, UINT160SIZE)
	tmp[UINT160SIZE-1] = 0x01
	u := UInt160FromBytes(tmp)
	assert.Equal(t, 0, UInt160Zero.CompareTo(UInt160Zero))
	assert.Equal(t, 1, u.CompareTo(UInt160Zero))
	assert.Equal(t, -1, UInt160Zero.CompareTo(u))
}

func TestUInt160_Equals(t *testing.T) {
	tmp := make([]byte, UINT160SIZE)
	tmp[UINT160SIZE-1] = 0x01
	u := UInt160FromBytes(tmp)
	assert.Equal(t, true, UInt160Zero.Equals(UInt160Zero))
	assert.Equal(t, false, u.Equals(UInt160Zero))
	assert.Equal(t, false, u.Equals(nil))
}

func TestUInt160_Equals2(t *testing.T)  {
	tmp := make([]byte, UINT160SIZE)
	u := UInt160FromBytes(tmp)
	assert.Equal(t, true, u.Equals(UInt160Zero))
	tmp1 := make([]byte, UINT160SIZE)
	tmp1[UINT160SIZE-1] = 0x01
	u1 := UInt160FromBytes(tmp1)

	if *u == *UInt160Zero && *u1 != *UInt160Zero { // value equals
		log.Println("can use ==")
	} else {
		log.Println("can not use ==")
	}
}

func TestUInt160_Less(t *testing.T) {
	a := "2d3b96ae1bcc5a585e075e3b81920210dec16302"
	b := "2d3b96ae1bcc5a585e075e3b81920210dec16303"

	ua, err := UInt160FromString(a)
	assert.Nil(t, err)
	ua2, err := UInt160FromString(a)
	assert.Nil(t, err)
	ub, err := UInt160FromString(b)
	assert.Nil(t, err)
	assert.Equal(t, true, ua.Less(ub))
	assert.Equal(t, false, ua.Less(ua2))
	assert.Equal(t, false, ub.Less(ua))
}

func TestUInt160_String(t *testing.T) {
	s := "b28427088a3729b2536d10122960394e8be6721f"
	u, err := UInt160FromString(s)
	assert.Nil(t, err)
	assert.Equal(t, s, u.String())
}

func TestUInt160_Size(t *testing.T) {
	s := UInt160Zero.Size()
	assert.Equal(t, UINT160SIZE, s)
}

func TestUInt160_ExistsIn(t *testing.T) {
	a := []UInt160{*UInt160Zero}
	b := UInt160Zero.ExistsIn(a)
	assert.Equal(t, true, b)
}

func TestUInt160_ToByteArray(t *testing.T) {
	s := "2d3b96ae1bcc5a585e075e3b81920210dec16302"
	u, err := UInt160FromString(s)
	assert.Nil(t, err)
	assert.Equal(t, HexToBytes(s), ReverseBytes(u.ToByteArray()))
}

func TestUInt160_Deserialize(t *testing.T) {
	s := "2d3b96ae1bcc5a585e075e3b81920210dec16302"
	b := ReverseBytes(HexToBytes(s))
	br := io.NewBinaryReaderFromBuf(b)
	u := NewUInt160()
	u.Deserialize(br)
	assert.Equal(t, s, u.String())
}

func TestUInt160_Serialize(t *testing.T) {
	s := "2d3b96ae1bcc5a585e075e3b81920210dec16302"
	u, err := UInt160FromString(s)
	assert.Nil(t, err)
	b := io.NewBufBinaryWriter()
	u.Serialize(b.BinaryWriter)
	assert.Equal(t, HexToBytes(s), ReverseBytes(b.Bytes()))
}

func TestUInt160UnmarshalAndMarshal(t *testing.T) {
	s := "2d3b96ae1bcc5a585e075e3b81920210dec16302"
	expected, err := UInt160FromString(s)
	assert.Nil(t, err)

	// UnmarshalJSON decodes hex-strings
	u1 := NewUInt160()
	err = u1.UnmarshalJSON([]byte(`"` + s + `"`))
	assert.Nil(t, err)
	assert.Equal(t, true, expected.Equals(u1))

	b, err := expected.MarshalJSON()
	assert.Nil(t, err)

	// UnmarshalJSON decodes hex-strings prefixed by 0x
	u2 := NewUInt160()
	err = u2.UnmarshalJSON(b)
	assert.Nil(t, err)
	assert.Equal(t, true, expected.Equals(u2))
}
