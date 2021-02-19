package io

import (
	"bytes"
	"io"
)

// ISerializable defines the binary encoding/decoding interface. Errors are
// returned via BinaryReader/BinaryWriter Err field. These functions must have safe
// behavior when passed BinaryReader/BinaryWriter with Err already set. Invocations
// to these functions tend to be nested, with this mechanism only the top-level
// caller should handle the error once and all the other code should just not
// panic in presence of error.
type ISerializable interface {
	Deserialize(*BinaryReader)
	Serialize(*BinaryWriter)
}

func ToArray(p ISerializable) ([]byte, error) {
	buf := NewBufBinaryWriter()
	p.Serialize(buf.BinaryWriter)
	return buf.Bytes(), buf.Err
}

func AsSerializable(se ISerializable, data []byte) error {
	buffer := bytes.NewBuffer(data)
	reader := NewBinaryReaderFromIO(io.Reader(buffer))
	se.Deserialize(reader)
	return reader.Err
}
