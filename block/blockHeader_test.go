package block

import (
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/joeqian10/neo3-gogogo/sc"
	"github.com/joeqian10/neo3-gogogo/tx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func SetupBlockHeaderWithValues() *BlockHeader {
	mr := helper.UInt256FromBytes([]byte{214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251})
	//ts := time.Date(1968, 06, 01, 0, 0, 0, 0, time.UTC)

	bh := BlockHeader{
		Version:       0,
		PrevHash:      helper.NewUInt256(),
		MerkleRoot:    mr,
		Timestamp:     4244941696,
		Index:         0,
		NextConsensus: helper.NewUInt160(),
		Witness: &tx.Witness{
			InvocationScript:   []byte{},
			VerificationScript: []byte{byte(sc.PUSH1)},
		},
	}
	return &bh
}

func TestBlockHeader_Deserialize(t *testing.T) {
	rawBlock := []byte{0, 0, 0, 0, // Version
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // PrevHash
		214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251, // MerkleRoot
		128, 171, 4, 253, // Timestamp
		0, 0, 0, 0, // Index
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // NextConsensus
		1, 0, 1, 17, // Witness
		0} // check bit

	br := io.NewBinaryReaderFromBuf(rawBlock)
	bh := NewBlockHeader()
	bh.Deserialize(br)
	assert.Nil(t, br.Err)
	assert.Equal(t, uint32(0), bh.Version)
	assert.Equal(t, uint32(4244941696), bh.Timestamp)
	assert.Equal(t, byte(sc.PUSH1), bh.Witness.VerificationScript[0])
}

func TestBlockHeader_DeserializeUnsigned(t *testing.T) {
	rawBlock := []byte{0, 0, 0, 0, // Version
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // PrevHash
		214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251, // MerkleRoot
		128, 171, 4, 253, // Timestamp
		0, 0, 0, 0, // Index
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // NextConsensus
    }

	br := io.NewBinaryReaderFromBuf(rawBlock)
	bh := NewBlockHeader()
	bh.DeserializeUnsigned(br)
	assert.Equal(t, uint32(0), bh.Version)
	assert.Equal(t, uint32(4244941696), bh.Timestamp)
}

func TestBlockHeader_GetHashData(t *testing.T) {
	bh := SetupBlockHeaderWithValues()
	bs := bh.GetHashData()
	requiredData := []byte{0, 0, 0, 0, // Version
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // PrevHash
		214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251, // MerkleRoot
		128, 171, 4, 253, // Timestamp
		0, 0, 0, 0, // Index
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // NextConsensus

	assert.Equal(t, requiredData, bs)
}

func TestBlockHeader_Serialize(t *testing.T) {
	bh := SetupBlockHeaderWithValues()
	buf := io.NewBufBinaryWriter()
	bh.Serialize(buf.BinaryWriter)
	requiredData := []byte{0, 0, 0, 0, // Version
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // PrevHash
		214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251, // MerkleRoot
		128, 171, 4, 253, // Timestamp
		0, 0, 0, 0, // Index
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // NextConsensus
		1, 0, 1, 17, // Witness
		0} // check bit

	assert.Nil(t, buf.Err)
	assert.Equal(t, helper.BytesToHex(requiredData), helper.BytesToHex(buf.Bytes()))
}

func TestBlockHeader_SerializeUnsigned(t *testing.T) {
	bh := SetupBlockHeaderWithValues()
	buf := io.NewBufBinaryWriter()
	bh.SerializeUnsigned(buf.BinaryWriter)
	requiredData := []byte{0, 0, 0, 0, // Version
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // PrevHash
		214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251, // MerkleRoot
		128, 171, 4, 253, // Timestamp
		0, 0, 0, 0, // Index
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // NextConsensus

	assert.Nil(t, buf.Err)
	assert.Equal(t, helper.BytesToHex(requiredData), helper.BytesToHex(buf.Bytes()))
}

func TestNewBlockHeaderFromRPC(t *testing.T) {
	//consensusData := binary.BigEndian.Uint64(helper.HexToBytes("000000007c2bac1d"))
	//assert.Equal(t, uint64(2083236893), consensusData)
	rpcHeader := models.RpcBlockHeader{
		Hash:              "0xd42561e3d30e15be6400b6df2f328e02d2bf6354c41dce433bc57687c82144bf",
		Size:              401,
		Version:           0,
		PreviousBlockHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
		MerkleRoot:        "0x803ff4abe3ea6533bcc0be574efa02f83ae8fdc651c879056b0d9be336c01bf4",
		Time:              1468595301,
		Index:             0,
		NextConsensus:     "NNU67Fvdy3LEQTM374EJ9iMbCRxVExgM8Y",
		Witnesses: []models.RpcWitness{{
			Invocation:   "",
			Verification: "11",
		}},
		Confirmations: 5276880,
		NextBlockHash: "0xd782db8a38b0eea0d7394e0f007c61c71798867578c77c387c08113903946cc9",
	}

	header, err := NewBlockHeaderFromRPC(&rpcHeader)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(header.Witness.VerificationScript))
}
