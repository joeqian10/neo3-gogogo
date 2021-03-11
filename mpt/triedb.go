package mpt

import "github.com/joeqian10/neo3-gogogo/helper"

//IKVReadOnlyDb to store data
type IKVReadOnlyDb interface {
	Get(key []byte) ([]byte, error)
}

type trieDb struct {
	db IKVReadOnlyDb
}

func newTrieDb(kvdb IKVReadOnlyDb) *trieDb {
	return &trieDb{
		db: kvdb,
	}
}

func (t *trieDb) node(hash *helper.UInt256) (*Node, error) {
	data, err := t.db.Get(hash.ToByteArray())
	if err != nil {
		return nil, err
	}
	node, err := decodeNode(data)
	return node, err
}
