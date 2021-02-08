package mpt

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

func (t *trieDb) node(hash hashNode) (node, error) {
	data, err := t.db.Get(hash)
	if err != nil {
		return nil, err
	}
	node, err := decodeNode(data)
	return node, err
}
