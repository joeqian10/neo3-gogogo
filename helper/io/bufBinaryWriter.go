package io

import (
	"bytes"
)

// BufBinaryWriter is an additional layer on top of BinaryWriter that
// automatically creates buffer to write into that you can get after all
// writes via Bytes().
type BufBinaryWriter struct {
	*BinaryWriter
	buf *bytes.Buffer
}

// NewBufBinaryWriter makes a BufBinaryWriter with an empty byte buffer.
func NewBufBinaryWriter() *BufBinaryWriter {
	b := new(bytes.Buffer)
	return &BufBinaryWriter{BinaryWriter: NewBinaryWriterFromIO(b), buf: b}
}

// Bytes returns resulting buffer and reset to prevent future writes
func (bw *BufBinaryWriter) Bytes() []byte {
	if bw.Err != nil {
		return nil
	}
	b := bw.buf.Bytes()
	bw.Reset()
	return b
}

// Reset resets the state of the buffer, making it usable again. It can
// make buffer usage somewhat more efficient, because you don't need to
// create it again, but beware that the buffer is gonna be the same as the one
// returned by Bytes(), so if you need that data after Reset() you have to copy
// it yourself.
func (bw *BufBinaryWriter) Reset() {
	bw.Err = nil
	bw.buf.Reset()
}
