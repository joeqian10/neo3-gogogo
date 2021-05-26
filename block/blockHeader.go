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

type Header struct {
	version       uint32
	prevHash      *helper.UInt256
	merkleRoot    *helper.UInt256
	timestamp     uint64
	index         uint32
	primaryIndex  byte
	nextConsensus *helper.UInt160

	Witness *tx.Witness

	_hash *helper.UInt256
	_size int
}

func NewBlockHeader() *Header {
	return &Header{
		version:       0,
		prevHash:      helper.NewUInt256(),
		merkleRoot:    helper.NewUInt256(),
		timestamp:     0,
		index:         0,
		primaryIndex:  0,
		nextConsensus: helper.NewUInt160(),
		Witness:       &tx.Witness{
			InvocationScript:   []byte{},
			VerificationScript: []byte{},
		},
		_hash:         nil,
		_size:         0,
	}
}

func NewBlockHeaderFromRPC(header *models.RpcBlockHeader) (*Header, error) {
	version := uint32(header.Version)
	prevHash, err := helper.UInt256FromString(header.PreviousBlockHash)
	if err != nil {
		return nil, err
	}
	merkleRoot, err := helper.UInt256FromString(header.MerkleRoot)
	if err != nil {
		return nil, err
	}
	timeStamp := uint64(header.Time)
	index := uint32(header.Index)
	primaryIndex := header.PrimaryIndex
	nextConsensus, err := crypto.AddressToScriptHash(header.NextConsensus, helper.DefaultAddressVersion)
	if err != nil {
		return nil, err
	}
	var witness *tx.Witness
	if len(header.Witnesses) != 0 {
		inv, err := crypto.Base64Decode(header.Witnesses[0].Invocation)
		if err != nil {
			return nil, err
		}
		ver, err := crypto.Base64Decode(header.Witnesses[0].Verification)
		if err != nil {
			return nil, err
		}
		witness = &tx.Witness{
			InvocationScript:   inv,
			VerificationScript: ver,
		}
	}
	hash, err := helper.UInt256FromString(header.Hash)
	if err != nil {
		return nil, err
	}
	bh := Header{
		version:       version,
		prevHash:      prevHash,
		merkleRoot:    merkleRoot,
		timestamp:     timeStamp,
		index:         index,
		primaryIndex:  primaryIndex,
		nextConsensus: nextConsensus,
		Witness:       witness,
	}
	fmt.Println(bh.GetHash().String())
	if !bh.GetHash().Equals(hash) {
		return nil, fmt.Errorf("wrong block hash")
	}
	return &bh, nil
}

func (h *Header) GetVersion() uint32 {
	return h.version
}
func (h *Header) SetVersion(value uint32) {
	h.version = value
	h._hash = nil
}

func (h *Header) GetPrevHash() *helper.UInt256 {
	return h.prevHash
}
func (h *Header) SetPrevHash(value *helper.UInt256) {
	h.prevHash = value
	h._hash = nil
}

func (h *Header) GetMerkleRoot() *helper.UInt256 {
	return h.merkleRoot
}
func (h *Header) SetMerkleRoot(value *helper.UInt256) {
	h.merkleRoot = value
	h._hash = nil
}

func (h *Header) GetTimeStamp() uint64 {
	return h.timestamp
}
func (h *Header) SetTimeStamp(value uint64) {
	h.timestamp = value
	h._hash = nil
}

func (h *Header) GetIndex() uint32 {
	return h.index
}
func (h *Header) SetIndex(value uint32) {
	h.index = value
	h._hash = nil
}

func (h *Header) GetPrimaryIndex() byte {
	return h.primaryIndex
}
func (h *Header) SetPrimaryIndex(value byte) {
	h.primaryIndex = value
	h._hash = nil
}

func (h *Header) GetNextConsensus() *helper.UInt160 {
	return h.nextConsensus
}
func (h *Header) SetNextConsensus(value *helper.UInt160) {
	h.nextConsensus = value
	h._hash = nil
}

func (h *Header) GetHash() *helper.UInt256 {
	if h._hash == nil {
		h._hash = tx.CalculateHash(h)
	}
	return h._hash
}

func (h *Header) GetSize() int {
	return 4 + // version
		32 + // prevHash
		32 + // merkleRoot
		8 + // timestamp
		4 + // index
		1 + // primaryIndex
		20 + // nextConsensus
		1 + h.Witness.Size()
}

func (h *Header) GetWitnesses() []tx.Witness {
	return []tx.Witness{*h.Witness}
}
func (h *Header) SetWitnesses(value []tx.Witness)  {
	if len(value) != 1 {
		return
	}
	h.Witness = &value[0]
}

func (h *Header) GetScriptHashesForVerifying() []helper.UInt160 {
	if h.prevHash.Equals(helper.UInt256Zero) {
		return []helper.UInt160 {*h.Witness.GetScriptHash()}
	}
	// todo, get prev block header
	return nil
}

func (h *Header) Deserialize(br *io.BinaryReader) {
	h.DeserializeUnsigned(br)
	var b byte
	br.ReadLE(&b)
	if b != byte(1) {
		br.Err = fmt.Errorf("format error: padding must equal 1 got %d", b)
	}
	if h.Witness == nil {
		h.Witness = &tx.Witness{}
	}
	h.Witness.Deserialize(br)
}

//DeserializeUnsigned deserialize blockheader without witness
func (h *Header) DeserializeUnsigned(br *io.BinaryReader) {
	br.ReadLE(&h.version)
	br.ReadLE(h.prevHash)
	br.ReadLE(h.merkleRoot)
	br.ReadLE(&h.timestamp)
	br.ReadLE(&h.index)
	h.primaryIndex = br.ReadByte()
	br.ReadLE(h.nextConsensus)
}

func (h *Header) Serialize(bw *io.BinaryWriter) {
	h.SerializeUnsigned(bw)
	bw.WriteVarUInt(uint64(len([]tx.Witness{*h.Witness})))
	h.Witness.Serialize(bw)
}

//SerializeUnsigned serialize blockheader without witness
func (h *Header) SerializeUnsigned(bw *io.BinaryWriter) {
	bw.WriteLE(h.version)
	bw.WriteLE(h.prevHash)
	bw.WriteLE(h.merkleRoot)
	bw.WriteLE(h.timestamp)
	bw.WriteLE(h.index)
	bw.WriteLE(h.primaryIndex)
	bw.WriteLE(h.nextConsensus)
}

func (h *Header) GetHashString() string {
	return hex.EncodeToString(helper.ReverseBytes(h.GetHash().ToByteArray())) // reverse to big endian
}
