package mpt

type NodeType byte

const (
	BranchNode NodeType = 0x00
	ExtensionNode NodeType = 0x01
	LeafNode NodeType = 0x02
	HashNode NodeType = 0x03
	Empty NodeType = 0x04

	fullNodeType  byte = 0x00
	shortNodeType byte = 0x01
	hashNodeType  byte = 0x02
	valueNodeType byte = 0x03
)

