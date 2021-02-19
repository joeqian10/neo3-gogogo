package block

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
)

const ValidatorsCount = 7

type ConsensusData struct {
	PrimaryIndex byte
	Nonce uint64

	hash *helper.UInt256
}

func (c *ConsensusData) GetHash() *helper.UInt256 {
	if c.hash == nil {
		c.hash = helper.UInt256FromBytes(crypto.Hash256(c.ToArray()))
	}
	return c.hash
}

func (c *ConsensusData) GetSize() int {
	return 1 + 8
}

func (c *ConsensusData) Deserialize(br *io.BinaryReader)  {
	primaryIndex := br.ReadByte()
	if primaryIndex >= ValidatorsCount {
		br.Err = fmt.Errorf("format error: index out of range")
	}
	c.PrimaryIndex = primaryIndex
	br.ReadLE(&c.Nonce)
}

func (c *ConsensusData) Serialize(bw *io.BinaryWriter)  {
	bw.WriteLE(c.PrimaryIndex)
	bw.WriteLE(c.Nonce)
}

func (c *ConsensusData) ToArray() []byte {
	buf := io.NewBufBinaryWriter()
	bw := buf.BinaryWriter
	c.Serialize(bw)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}
