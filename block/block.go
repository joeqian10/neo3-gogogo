package block

import (
	//"github.com/joeqian10/neo-gogogo/rpc/models"
	"github.com/joeqian10/neo3-gogogo/tx"
)

type Block struct {
	Header
	Transactions []tx.Transaction
}

func (b *Block) GetSize() int {
	sz := 0
	for _, tx := range b.Transactions {
		sz += tx.GetSize()
	}
	return b.Header.GetSize() + sz
}
