package blockchain

import (
	"errors"

	"github.com/joeqian10/neo3-gogogo/io"
)

func writeBytesWithGrouping(writer *io.BinaryWriter, value []byte) {
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

func readBytesWithGrouping(reader *io.BinaryReader) (key []byte, err error) {
	padding := byte(0)
	for padding == 0 {
		group := [16]byte{}
		reader.ReadLE(&group)
		reader.ReadLE(&padding)
		if 16 < padding {
			return key, errors.New("padding error")
		}
		count := 16 - padding
		if count > 0 {
			key = append(key, group[:count]...)
		}
	}
	return key, nil
}
