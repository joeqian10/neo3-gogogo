package io

import (
	"encoding/binary"
	"io"
)

// BinaryWriter is a convenient wrapper around a io.Writer and err object.
// Used to simplify error handling when writing into a io.Writer
// from a struct with many fields.
type BinaryWriter struct {
	w   io.Writer
	Err error
}

// NewBinaryWriterFromIO makes a BinaryWriter from io.Writer.
func NewBinaryWriterFromIO(iow io.Writer) *BinaryWriter {
	return &BinaryWriter{w: iow}
}

// WriteLE writes into the underlying io.Writer from an object v in little-endian format.
func (w *BinaryWriter) WriteLE(v interface{}) {
	if w.Err != nil {
		return
	}
	w.Err = binary.Write(w.w, binary.LittleEndian, v)
}

// WriteBE writes into the underlying io.Writer from an object v in big-endian format.
func (w *BinaryWriter) WriteBE(v interface{}) {
	if w.Err != nil {
		return
	}
	w.Err = binary.Write(w.w, binary.BigEndian, v)
}

// WriteVarUint writes a uint64 into the underlying writer using variable-length encoding.
func (w *BinaryWriter) WriteVarUInt(val uint64) {
	if w.Err != nil {
		return
	}

	if val < 0xfd {
		w.Err = binary.Write(w.w, binary.LittleEndian, uint8(val))
		return
	}
	if val < 0xffff {
		w.Err = binary.Write(w.w, binary.LittleEndian, byte(0xfd))
		w.Err = binary.Write(w.w, binary.LittleEndian, uint16(val))
		return
	}
	if val < 0xFFFFFFFF {
		w.Err = binary.Write(w.w, binary.LittleEndian, byte(0xfe))
		w.Err = binary.Write(w.w, binary.LittleEndian, uint32(val))
		return
	}
	w.Err = binary.Write(w.w, binary.LittleEndian, byte(0xff))
	w.Err = binary.Write(w.w, binary.LittleEndian, val)
}

// WriteBytes writes a variable length byte array into the underlying io.Writer.
func (w *BinaryWriter) WriteVarBytes(b []byte) {
	w.WriteVarUInt(uint64(len(b)))
	w.WriteLE(b)
}

// WriteVarString writes a variable length string into the underlying io.Writer.
func (w *BinaryWriter) WriteVarString(s string) {
	w.WriteVarBytes([]byte(s))
}

//WriteBytesWithGrouping ...
func (writer *BinaryWriter) WriteBytesWithGrouping(value []byte) {
	index := 0
	remain := len(value)
	for remain >= 16 {
		writer.WriteLE(value[index : index+16])
		writer.WriteLE(byte(0))
		index += 16
		remain -= 16
	}
	if remain > 0 {
		writer.WriteLE(value[index:])
	}
	padding := 16 - remain
	for i := 0; i < padding; i++ {
		writer.WriteLE(byte(0))
	}
	writer.WriteLE(byte(padding))
}