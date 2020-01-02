package io

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// BinaryReader is a convenient wrapper around a io.Reader and err object.
// Used to simplify error handling when reading into a struct with many fields.
type BinaryReader struct {
	r   io.Reader
	Err error
}

// NewBinaryReaderFromIO makes a BinaryReader from io.Reader.
func NewBinaryReaderFromIO(ior io.Reader) *BinaryReader {
	return &BinaryReader{r: ior}
}

// NewBinaryReaderFromBuf makes a BinaryReader from byte buffer.
func NewBinaryReaderFromBuf(b []byte) *BinaryReader {
	r := bytes.NewReader(b)
	return NewBinaryReaderFromIO(r)
}

// ReadLE reads from the underlying io.Reader
// into the interface v in little-endian format.
func (br *BinaryReader) ReadLE(v interface{}) {
	if br.Err != nil {
		return
	}
	br.Err = binary.Read(br.r, binary.LittleEndian, v)
}

// ReadBE reads from the underlying io.Reader
// into the interface v in big-endian format.
func (br *BinaryReader) ReadBE(v interface{}) {
	if br.Err != nil {
		return
	}
	br.Err = binary.Read(br.r, binary.BigEndian, v)
}

// ReadUInt64Bytes reads from the underlying io.Reader
// into the interface v in little-endian format
func (br *BinaryReader) ReadUInt64Bytes() []byte {
	b := make([]byte, 8)
	br.ReadLE(b)
	if br.Err != nil {
		return nil
	}
	return b;
}

// ReadVarUInt reads a variable-length-encoded integer from the underlying reader.
// The result should not exceed the max value of uint64
func (br *BinaryReader) ReadVarUInt(max uint64) uint64 {
	if br.Err != nil {
		return 0
	}

	var b uint8
	br.Err = binary.Read(br.r, binary.LittleEndian, &b)
	var result uint64

	if b == 0xfd {
		var v uint16
		br.Err = binary.Read(br.r, binary.LittleEndian, &v)
		result = uint64(v)
	} else if b == 0xfe {
		var v uint32
		br.Err = binary.Read(br.r, binary.LittleEndian, &v)
		result = uint64(v)
	} else if b == 0xff {
		var v uint64
		br.Err = binary.Read(br.r, binary.LittleEndian, &v)
		result = v
	} else {
		result = uint64(b)
	}
	if result > max {
		br.Err = fmt.Errorf("max value exceeded")}
	return result
}

// ReadVarBytes reads the next set of bytes from the underlying reader.
// ReadVarUInt() is used to determine how large that slice is
// length should not exceed 0x1000000 = 16,777,216‬
func (br *BinaryReader) ReadVarBytes(max int) []byte {
	if max > 0x1000000 {
		br.Err = fmt.Errorf("max length exceeded")
		return nil
	}
	n := br.ReadVarUInt(uint64(max))
	b := make([]byte, n)
	br.ReadLE(b)
	return b
}

// ReadVarString calls ReadVarBytes and casts the results as a string.
// "max" should not exceed 0x1000000 = 16,777,216‬
func (br *BinaryReader) ReadVarString(max int) string {
	b := br.ReadVarBytes(max)
	return string(b)
}
