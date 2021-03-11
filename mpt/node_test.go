package mpt

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/stretchr/testify/assert"
	"testing"
)

func nodeToArrayAsChild(n *Node) []byte {
	bbw := io.NewBufBinaryWriter()
	n.SerializeAsChild(bbw.BinaryWriter)
	return bbw.Bytes()
}

func TestNode_Hash_Serialize(t *testing.T) {
	n := NewHashNode(helper.UInt256Zero)
	expected := "030000000000000000000000000000000000000000000000000000000000000000"
	bs, err := io.ToArray(n)
	assert.Nil(t, err)
	assert.Equal(t, expected, helper.BytesToHex(bs))
}

func TestNode_Empty_Serialize(t *testing.T) {
	n := NewNode()
	expected := "04"
	bs, err := io.ToArray(n)
	assert.Nil(t, err)
	assert.Equal(t, expected, helper.BytesToHex(bs))
}

func TestNode_Leaf_Serialize(t *testing.T) {
	n := NewLeafNode([]byte("leaf"))
	expected := "02" + "04" + helper.BytesToHex([]byte("leaf"))
	assert.Equal(t, expected, helper.BytesToHex(n.ToArrayWithoutReference()))
	expected += "01"
	bs, err := io.ToArray(n)
	assert.Nil(t, err)
	assert.Equal(t, expected, helper.BytesToHex(bs))
}

func TestNode_Extension_Serialize(t *testing.T) {
	n := NewExtensionNode(helper.HexToBytes("010a"), NewNode())
	expected := "01" + "02" + "010a" + "04"
	assert.Equal(t, expected, helper.BytesToHex(n.ToArrayWithoutReference()))
	expected += "01"
	bs, err := io.ToArray(n)
	assert.Nil(t, err)
	assert.Equal(t, expected, helper.BytesToHex(bs))
	assert.Equal(t, 6, n.Size())
}

func TestNode_Branch_Serialize(t *testing.T) {
	n := NewBranchNode()
	n.Children[1] = *NewLeafNode([]byte("leaf1"))
	n.Children[10] = *NewLeafNode([]byte("leafa"))
	expected := "00"
	for i:= 0; i<BranchChildCount; i++ {
		if i == 1 {
			expected += "03" + helper.BytesToHex(crypto.Hash256(append([]byte{0x02, 0x05}, []byte("leaf1")...)))
		} else if i == 10 {
			expected += "03" + helper.BytesToHex(crypto.Hash256(append([]byte{0x02, 0x05}, []byte("leafa")...)))
		} else {
			expected += "04"
		}
	}
	expected += "01"
	bs, err := io.ToArray(n)
	assert.Nil(t, err)
	assert.Equal(t, expected, helper.BytesToHex(bs))
	assert.Equal(t, 83, n.Size())
}

func TestNode_Leaf_SerializeAsChild(t *testing.T) {
	n := NewLeafNode([]byte("leaf"))
	expected := "03" + helper.BytesToHex(crypto.Hash256(append([]byte{0x02, 0x04}, []byte("leaf")...)))
	assert.Equal(t, expected, helper.BytesToHex(nodeToArrayAsChild(n)))
}

func TestNode_Extension_SerializeAsChild(t *testing.T) {
	n := NewExtensionNode(helper.HexToBytes("010a"), NewNode())
	expected := "03" + helper.BytesToHex(crypto.Hash256([]byte{0x01, 0x02, 0x01, 0x0a, 0x04}))
	assert.Equal(t, expected, helper.BytesToHex(nodeToArrayAsChild(n)))
}

func TestNode_Branch_SerializeAsChild(t *testing.T) {
	n := NewBranchNode()
	data := []byte{0x00}
	for i:= 0; i<BranchChildCount; i++ {
		data = append(data, 0x04)
	}
	expected := "03" + helper.BytesToHex(crypto.Hash256(data))
	assert.Equal(t, expected, helper.BytesToHex(nodeToArrayAsChild(n)))
}

func TestNode_Size(t *testing.T) {
	n := NewNode()
	assert.Equal(t, 1, n.Size())
	n = NewBranchNode()
	assert.Equal(t, 19, n.Size())
	n = NewExtensionNode([]byte{0x00}, NewNode())
	assert.Equal(t, 5, n.Size())
	n = NewLeafNode([]byte{0x00})
	assert.Equal(t, 4, n.Size())
	n = NewHashNode(helper.UInt256Zero)
	assert.Equal(t, 33, n.Size())
}
