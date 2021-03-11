package mpt

import (
	"bytes"
	"errors"
	"github.com/joeqian10/neo3-gogogo/blockchain"
	"github.com/joeqian10/neo3-gogogo/helper"
	nio "github.com/joeqian10/neo3-gogogo/io"
)

const prefix byte = 0xf0

// ToNibbles ..
func ToNibbles(data []byte) []byte {
	r := make([]byte, len(data)*2)
	for i := 0; i < len(data); i++ {
		r[i*2] = data[i] >> 4
		r[i*2+1] = data[i] & 0x0f
	}
	return r
}

func FromNibbles(data []byte) []byte {
	if len(data)%2 != 0 {
		return nil
	}
	key := make([]byte, len(data)/2)
	for i := 0; i < len(key); i++ {
		key[i] = data[i*2] << 4
		key[i] |= data[i*2+1]
	}
	return key
}

//Trie mpt tree
type Trie struct {
	db   *trieDb
	root *Node
}

//NewTrie new a trie instance
func NewTrie(root *helper.UInt256, db IKVReadOnlyDb) (*Trie, error) {
	if db == nil {
		return nil, errors.New("failed initialize Trie, invalid db")
	}
	t := &Trie{
		db: newTrieDb(db),
	}
	if root == nil {
		t.root = NewNode()
	} else {
		r, err := t.resolve(root)
		if err != nil {
			return nil, err
		}
		t.root = r
	}
	return t, nil
}

func (t *Trie) resolve(hash *helper.UInt256) (*Node, error) {
	return t.db.node(hash)
}

//Get try get value
func (t *Trie) Get(key []byte) ([]byte, error) {
	path := ToNibbles(key)
	vn, err := t.get(t.root, path)
	if err != nil {
		return nil, err
	}
	return vn.Value, err
}

func (t *Trie) get(n *Node, path []byte) (*Node, error) {
	switch n.nodeType {
	case LeafNode:
		if len(path) == 0 {
			return n, nil
		}
		break
	case Empty:
		break
	case HashNode:
		nn, err := t.resolve(n.hash)
		if err != nil {
			return nil, err
		}
		return t.get(nn, path)
	case BranchNode:
		if len(path) == 0 {
			return t.get(&n.Children[16], path)
		}
		return t.get(&n.Children[path[0]], path[1:])
	case ExtensionNode:
		if bytes.HasPrefix(path, n.Key) {
			return t.get(n.Next, bytes.TrimPrefix(path, n.Key))
		}
		break
	}
	return nil, errors.New("invalid node or path for the trie")
}

//VerifyProof directly verify proof
func VerifyProof(root *helper.UInt256, id int, key []byte, proof [][]byte) ([]byte, error) {
	sKey := blockchain.StorageKey{
		Id:  id,
		Key: key,
	}
	vkey, err := nio.ToArray(&sKey)
	if err != nil {
		return nil, err
	}
	proofdb := NewProofDb(proof)
	trie, err := NewTrie(root, proofdb)
	if err != nil {
		return nil, err
	}
	value, err := trie.Get(vkey)
	if err != nil {
		return nil, err
	}
	return resolveValue(value)
}

//ResolveProof get key and proofs from proofdata
func ResolveProof(proofBytes []byte) (id int, key []byte, proof [][]byte, err error) {
	br := nio.NewBinaryReaderFromBuf(proofBytes)
	key = br.ReadVarBytes()
	if br.Err != nil {
		err = br.Err
		return id, key, proof, err
	}
	count := br.ReadVarUInt()
	proof = make([][]byte, count)
	for i := uint64(0); i < count; i++ {
		proof[i] = br.ReadVarBytes()
	}
	id, key, err = resolveKey(key)
	return id, key, proof, err
}

func resolveValue(value []byte) ([]byte, error) {
	item := blockchain.StorageItem{}
	err := nio.AsSerializable(&item, value)
	return item.Value, err
}

func resolveKey(key []byte) (id int, kk []byte, err error) {
	sKey := blockchain.StorageKey{}
	err = nio.AsSerializable(&sKey, key)
	return sKey.Id, sKey.Key, err
}
