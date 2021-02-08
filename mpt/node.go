package mpt

import (
	"errors"

	"github.com/joeqian10/neo3-gogogo/helper/io"
)

const (
	fullNodeType  byte = 0x00
	shortNodeType byte = 0x01
	hashNodeType  byte = 0x02
	valueNodeType byte = 0x03
)

type node interface {
}

func decodeNode(data []byte) (node, error) {
	var ntype byte
	reader := io.NewBinaryReaderFromBuf(data)
	reader.ReadLE(&ntype)
	switch ntype {
	case fullNodeType:
		return decodeFullNode(reader)
	case shortNodeType:
		return decodeShortNode(reader)
	case valueNodeType:
		return valueNode(reader.ReadVarBytes()), reader.Err
	}
	return nil, errors.New("invalid node type to decode")
}

type fullNode struct {
	children [17]node
}

func decodeFullNode(reader *io.BinaryReader) (fullNode, error) {
	f := fullNode{}
	for i := range f.children {
		f.children[i] = hashNode(reader.ReadVarBytes())
	}
	return f, reader.Err
}

type shortNode struct {
	next node
	key  []byte
}

func decodeShortNode(reader *io.BinaryReader) (shortNode, error) {
	s := new(shortNode)
	s.key = reader.ReadVarBytes()
	s.next = hashNode(reader.ReadVarBytes())
	return *s, reader.Err
}

type hashNode []byte
type valueNode []byte
