package block

import (
	//"github.com/joeqian10/neo-gogogo/rpc/models"
	"github.com/joeqian10/neo3-gogogo/tx"
)

const (
	MaxContentsPerBlock = 65535
	MaxTransactionsPerBlock = MaxContentsPerBlock - 1
)

type Block struct {
	BlockHeader
	ConsensusData *ConsensusData
	Tx []tx.Transaction
}

//func NewBlockFromRPC(rpcBlock *models.RpcBlock) (*Block, error) {
//	var block = new &Block{
//		BlockHeader: NewBlockHeaderFromRPC(&rpcBlock.RpcBlockHeader),
//		Tx:          nil,
//	}
//}
