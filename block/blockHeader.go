package block

import (
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/joeqian10/neo3-gogogo/tx"
)

type BlockHeader struct {
	Version       uint32
	PrevHash      *helper.UInt256
	MerkleRoot    *helper.UInt256
	Timestamp     uint32
	Index         uint32
	NextConsensus *helper.UInt160

	Witness       *tx.Witness

	_hash *helper.UInt256
	_size int
}

func NewBlockHeader() *BlockHeader {
	return &BlockHeader{
		Version:       0,
		PrevHash:      helper.NewUInt256(),
		MerkleRoot:    helper.NewUInt256(),
		Timestamp:     0,
		Index:         0,
		NextConsensus: helper.NewUInt160(),
		Witness:       nil,
		_hash:         nil,
		_size:         0,
	}
}

func NewBlockHeaderFromRPC(header *models.RpcBlockHeader) (*BlockHeader, error) {
	version := uint32(header.Version)
	prevHash, err := helper.UInt256FromString(header.PreviousBlockHash)
	if err != nil {
		return nil, err
	}
	merkleRoot, err := helper.UInt256FromString(header.MerkleRoot)
	if err != nil {
		return nil, err
	}
	timeStamp := uint32(header.Time)
	index := uint32(header.Index)
	nextConsensus, err := crypto.AddressToScriptHash(header.NextConsensus)
	if err != nil {
		return nil, err
	}
	var witness *tx.Witness
	if header.Witnesses != nil && len(header.Witnesses) != 0 {
		witness = &tx.Witness{
			InvocationScript:   helper.HexToBytes(header.Witnesses[0].Invocation),
			VerificationScript: helper.HexToBytes(header.Witnesses[0].Verification),
		}
	}
	hash, err := helper.UInt256FromString(header.Hash)
	if err != nil {
		return nil, err
	}
	bh := BlockHeader{
		Version:       version,
		PrevHash:      prevHash,
		MerkleRoot:    merkleRoot,
		Timestamp:     timeStamp,
		Index:         index,
		NextConsensus: nextConsensus,
		Witness:       witness,
		_hash:         hash,
	}
	return &bh, nil
}

func (bh *BlockHeader) Deserialize(br *io.BinaryReader) {
	bh.DeserializeUnsigned(br)
	var b byte
	br.ReadLE(&b)
	if b != byte(1) {
		br.Err = fmt.Errorf("format error: padding must equal 1 got %d", b)
	}
	if bh.Witness == nil {
		bh.Witness = &tx.Witness{}
	}
	bh.Witness.Deserialize(br)
	br.ReadLE(&b)
	if b != byte(0) {
		br.Err = fmt.Errorf("format error: check byte must equal 0 got %d", b)
	}
}

//DeserializeUnsigned deserialize blockheader without witness
func (bh *BlockHeader) DeserializeUnsigned(br *io.BinaryReader) {
	br.ReadLE(&bh.Version)
	br.ReadLE(bh.PrevHash)
	br.ReadLE(bh.MerkleRoot)
	br.ReadLE(&bh.Timestamp)
	br.ReadLE(&bh.Index)
	br.ReadLE(bh.NextConsensus)
}

func (bh *BlockHeader) Serialize(bw *io.BinaryWriter) {
	bh.SerializeUnsigned(bw)
	bw.WriteVarUInt(uint64(len([]tx.Witness{*bh.Witness})))
	bh.Witness.Serialize(bw)
	bw.WriteLE(byte(0))
}

//SerializeUnsigned serialize blockheader without witness
func (bh *BlockHeader) SerializeUnsigned(bw *io.BinaryWriter) {
	bw.WriteLE(bh.Version)
	bw.WriteLE(bh.PrevHash)
	bw.WriteLE(bh.MerkleRoot)
	bw.WriteLE(bh.Timestamp)
	bw.WriteLE(bh.Index)
	bw.WriteLE(bh.NextConsensus)
}

func (bh *BlockHeader) GetHashData() []byte {
	buf := io.NewBufBinaryWriter()
	bh.SerializeUnsigned(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

func (bh *BlockHeader) HashString() string {
	hash := crypto.Hash256(bh.GetHashData())
	bh._hash = helper.UInt256FromBytes(hash)
	return hex.EncodeToString(helper.ReverseBytes(hash)) // reverse to big endian
}

func (bh *BlockHeader) Hash() *helper.UInt256 {
	hash := crypto.Hash256(bh.GetHashData())
	bh._hash = helper.UInt256FromBytes(hash)
	return bh._hash
}
