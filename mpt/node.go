package mpt

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/sc"
	io2 "io"

	"github.com/joeqian10/neo3-gogogo/io"
)

const BranchChildCount int = 17
const MaxKeyLength = (64+4)*2
const MaxValueLength = 3 + 65535 + 1

type Node struct {
	nodeType  NodeType
	hash      *helper.UInt256
	Reference int

	// BranchNode
	Children []Node

	// ExtensionNode
	Key []byte
	Next *Node

	// LeafNode
	Value []byte
}

func (n *Node) GetHash() *helper.UInt256 {
	if n.hash == nil {
		n.hash = helper.UInt256FromBytes(crypto.Hash256(n.ToArrayWithoutReference()))
	}
	return n.hash
}

func (n *Node) GetNodeType() NodeType {
	return n.nodeType
}

func (n *Node) IsEmpty() bool {
	return n.nodeType == Empty
}

func NewNode() *Node {
	return &Node{
		nodeType: Empty,
	}
}

func NewBranchNode() *Node {
	n := &Node{
		nodeType:  BranchNode,
		Reference: 1,
		Children:  make([]Node, BranchChildCount),
	}
	for i := 0; i < BranchChildCount; i++ {
		n.Children[i] = *NewNode()
	}
	return n
}

func NewExtensionNode(key []byte, next *Node) *Node {
	if len(key) == 0 || next == nil {
		return nil
	}
	return &Node{
		nodeType:  ExtensionNode,
		Key: key,
		Next: next,
		Reference: 1,
	}
}

func NewHashNode(hash *helper.UInt256) *Node {
	if hash == nil {
		return nil
	}
	return &Node{
		nodeType:  HashNode,
		hash:      hash,
	}
}

func NewLeafNode(value []byte) *Node {
	if len(value) == 0 {
		return nil
	}
	return &Node{
		nodeType: LeafNode,
		Value: value,
		Reference: 1,
	}
}

// ----------BranchNode functions----------
func (n *Node) getBranchSize() int {
	size := 0
	for i:=0;i < BranchChildCount; i++ {
		size += n.Children[i].SizeAsChild()
	}
	return size
}

func (n *Node) serializeBranch(bw *io.BinaryWriter)  {
	for i:=0;i < BranchChildCount; i++ {
		n.Children[i].SerializeAsChild(bw)
	}
}

func (n *Node) deserializeBranch(br *io.BinaryReader)  {
	n.Children = make([]Node, BranchChildCount)
	for i:=0;i < BranchChildCount; i++ {
		nn := NewNode()
		nn.Deserialize(br)
		n.Children[i] = *nn
	}
}
// ----------------------------------------

// ----------ExtensionNode functions----------
func (n *Node) getExtensionSize() int {
	return sc.ByteSlice(n.Key).GetVarSize() + n.Next.SizeAsChild()
}

func (n *Node) serializeExtension(bw *io.BinaryWriter)  {
	bw.WriteVarBytes(n.Key)
	n.Next.SerializeAsChild(bw)
}

func (n *Node) deserializeExtension(br *io.BinaryReader)  {
	n.Key = br.ReadVarBytes()
	nn := NewNode()
	nn.Deserialize(br)
	n.Next = nn
}
// -------------------------------------------

// ----------HashNode functions----------
func (n *Node) getHashSize() int {
	return n.hash.Size()
}

func (n *Node) serializeHash(bw *io.BinaryWriter)  {
	bw.WriteLE(n.hash)
}

func (n *Node) deserializeHash(br *io.BinaryReader)  {
	h := helper.NewUInt256()
	br.ReadLE(h)
	n.hash = h
}
// --------------------------------------

// ----------LeafNode functions----------
func (n *Node) getLeafSize() int {
	return sc.ByteSlice(n.Value).GetVarSize()
}

func (n *Node) serializeLeaf(bw *io.BinaryWriter) {
	bw.WriteVarBytes(n.Value)
}

func (n *Node) deserializeLeaf(br *io.BinaryReader) {
	n.Value = br.ReadVarBytes()
}
// --------------------------------------

func (n *Node) Size() int {
	size := 1 // emptyNode
	switch n.nodeType {
	case BranchNode:
		return size + n.getBranchSize() + helper.GetVarSize(n.Reference)
	case ExtensionNode:
		return size + n.getExtensionSize()+ helper.GetVarSize(n.Reference)
	case LeafNode:
		return size + n.getLeafSize() + helper.GetVarSize(n.Reference)
	case HashNode:
		return size + n.getHashSize()
	case Empty:
		return size
	default:
		panic(n.nodeType)
	}
}

func (n *Node) SetDirty()  {
	n.hash = nil
}

func (n *Node) SizeAsChild() int {
	switch n.nodeType {
	case BranchNode, ExtensionNode, LeafNode:
		return NewHashNode(n.GetHash()).Size()
	case HashNode, Empty:
		return n.Size()
	default:
		panic(n.nodeType)
	}
}

func (n *Node) SerializeAsChild(bw *io.BinaryWriter)  {
	switch n.nodeType {
	case BranchNode, ExtensionNode, LeafNode:
		nn := NewHashNode(n.GetHash())
		nn.Serialize(bw)
		break
	case HashNode, Empty:
		n.Serialize(bw)
		break
	default:
		panic(n.nodeType)
	}
}

func (n *Node) serializeWithoutReference(bw *io.BinaryWriter)  {
	bw.WriteLE(n.nodeType)
	switch n.nodeType {
	case BranchNode:
		n.serializeBranch(bw)
		break
	case ExtensionNode:
		n.serializeExtension(bw)
		break
	case LeafNode:
		n.serializeLeaf(bw)
		break
	case HashNode:
		n.serializeHash(bw)
		break
	case Empty:
		break
	default:
		panic(n.nodeType)
	}
}

func (n *Node) Serialize(bw *io.BinaryWriter)  {
	n.serializeWithoutReference(bw)
	if n.nodeType == BranchNode || n.nodeType == ExtensionNode || n.nodeType == LeafNode {
		bw.WriteVarUInt(uint64(n.Reference))
	}
}

func (n *Node) ToArrayWithoutReference() []byte {
	bbw := io.NewBufBinaryWriter()
	n.serializeWithoutReference(bbw.BinaryWriter)
	return bbw.Bytes()
}

func (n *Node) Deserialize(br *io.BinaryReader)  {
	var nodeType byte
	br.ReadLE(&nodeType)
	n.nodeType = NodeType(nodeType)
	switch n.nodeType {
	case BranchNode:
		n.deserializeBranch(br)
		n.Reference = int(br.ReadVarUInt())
		break
	case ExtensionNode:
		n.deserializeExtension(br)
		n.Reference = int(br.ReadVarUInt())
		break
	case LeafNode:
		n.deserializeLeaf(br)
		n.Reference = int(br.ReadVarUInt())
		break
	case HashNode:
		n.deserializeHash(br)
		break
	case Empty:
		break
	default:
		panic(n.nodeType)
	}
}

func decodeNode(data []byte) (*Node, error) {
	br := io.NewBinaryReaderFromBuf(data)
	n := NewNode()
	n.Deserialize(br)
	if br.Err != nil && br.Err != io2.EOF {
		return nil, br.Err
	}
	return n, nil
}
