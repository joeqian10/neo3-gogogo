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

func SetupBlockHeaderWithValues() *Header {
	mr := helper.UInt256FromBytes([]byte{214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251})
	//ts := time.Date(1968, 06, 01, 0, 0, 0, 0, time.UTC)

	bh := Header{
		version:       0,
		prevHash:      helper.NewUInt256(),
		merkleRoot:    mr,
		timestamp:     4244941696,
		index:         0,
		nextConsensus: helper.NewUInt160(),
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
		128, 171, 4, 253, 0, 0, 0, 0, // Timestamp
		0, 0, 0, 0, // Index
		0, // PrimaryIndex
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // NextConsensus
		1, 0, 1, 17, // Witness
		}

	br := io.NewBinaryReaderFromBuf(rawBlock)
	bh := NewBlockHeader()
	bh.Deserialize(br)
	assert.Nil(t, br.Err)
	assert.Equal(t, uint32(0), bh.version)
	assert.Equal(t, uint64(4244941696), bh.timestamp)
	assert.Equal(t, byte(sc.PUSH1), bh.Witness.VerificationScript[0])
}

func TestBlockHeader_DeserializeUnsigned(t *testing.T) {
	rawBlock := []byte{0, 0, 0, 0, // Version
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // PrevHash
		214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251, // MerkleRoot
		128, 171, 4, 253, // Timestamp
		0, 0, 0, 0, // Index
		0, // PrimaryIndex
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // NextConsensus
    }

	br := io.NewBinaryReaderFromBuf(rawBlock)
	bh := NewBlockHeader()
	bh.DeserializeUnsigned(br)
	assert.Equal(t, uint32(0), bh.version)
	assert.Equal(t, uint64(4244941696), bh.timestamp)
}

func TestBlockHeader_Serialize(t *testing.T) {
	bh := SetupBlockHeaderWithValues()
	buf := io.NewBufBinaryWriter()
	bh.Serialize(buf.BinaryWriter)
	requiredData := []byte{0, 0, 0, 0, // Version
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // PrevHash
		214, 87, 42, 69, 155, 149, 217, 19, 107, 122, 113, 60, 84, 133, 202, 112, 159, 158, 250, 79, 8, 241, 194, 93, 215, 146, 103, 45, 43, 215, 91, 251, // MerkleRoot
		128, 171, 4, 253, 0, 0, 0, 0, // Timestamp
		0, 0, 0, 0, // Index
		0, // PrimaryIndex
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // NextConsensus
		1, 0, 1, 17, // Witness
		}

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
		128, 171, 4, 253, 0, 0, 0, 0, // Timestamp
		0, 0, 0, 0, // Index
		0, // PrimaryIndex
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // NextConsensus

	assert.Nil(t, buf.Err)
	assert.Equal(t, helper.BytesToHex(requiredData), helper.BytesToHex(buf.Bytes()))
}

func TestNewBlockHeaderFromRPC(t *testing.T) {
	//consensusData := binary.BigEndian.Uint64(helper.HexToBytes("000000007c2bac1d"))
	//assert.Equal(t, uint64(2083236893), consensusData)
	rpcHeader := models.RpcBlockHeader{
		Hash:              "0x1a3bdcad1cdaa90ecf731a412a89cb60ec49cbf3c605317f3f6911cd5def5761",
		Size:              401,
		Version:           0,
		PreviousBlockHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
		MerkleRoot:        "0x803ff4abe3ea6533bcc0be574efa02f83ae8fdc651c879056b0d9be336c01bf4",
		Time:              1468595301,
		Index:             0,
		PrimaryIndex:      0x00,
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
