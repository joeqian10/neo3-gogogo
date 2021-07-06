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
		nonce:         0,
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
		0, 0, 0, 0, 0, 0, 0, 0, // Nonce
		0, 0, 0, 0, // Index
		0,                                                          // PrimaryIndex
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
		0, 0, 0, 0, 0, 0, 0, 0, // Nonce
		0, 0, 0, 0, // Index
		0,                                                          // PrimaryIndex
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
		0, 0, 0, 0, 0, 0, 0, 0, // Nonce
		0, 0, 0, 0, // Index
		0,                                                          // PrimaryIndex
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
		0, 0, 0, 0, 0, 0, 0, 0, // Nonce
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
		Hash:              "0x159359b9a57c1d94f946dfb510409ac4a32008c31f6742b21d3d8e165cdd660d",
		Size:              401,
		Version:           0,
		PreviousBlockHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
		MerkleRoot:        "0x803ff4abe3ea6533bcc0be574efa02f83ae8fdc651c879056b0d9be336c01bf4",
		Time:              1468595301,
		Nonce:             "00",
		Index:             0,
		PrimaryIndex:      0x00,
		NextConsensus:     "NNU67Fvdy3LEQTM374EJ9iMbCRxVExgM8Y",
		Witnesses: []models.RpcWitness{{
			Invocation:   "DEB9U2OHJaj1fYGFA+yQbVEQA7oncKTpcTdxeA/r8+GP2zSYyfuM5jOwz4ggVmUEApevohEm3pMXQEWX2iJRoPWXDED87joySzn/NE8wBclEiJhPTqriDAQB7B9Vqz1XYlRpgk8hcQoO4dFDagV+wVgS/ihj5vyPUeOAIEJ5OL52o6RMDECtqJv94+JKu59rmwB8Oiagf8hkNrgxcLCVuwfbW4IGjTd5/fhUItyu7davOcFQKKr0SJayMQMwD38yqS4BBhWRDEC2z1+QgD6L/denhGfo0FpGvpcTEHFiGjJlXqGdC1Sm73KXLmiA4jnEmkvZQhfBVji6NMyI57Fm2LwziEQpGh4HDECifyH8SPERHo+Z+P3/dgzuBc16inO8aF9eOPexkJ9RkHH+A5sYOtzROiTcZ4LQVdthpYxl191ccl70L5rDNDC2",
			Verification: "FQwhAwCbdUDhDyVi5f2PrJ6uwlFmpYsm5BI0j/WoaSe/rCKiDCEDAgXpzvrqWh38WAryDI1aokaLsBSPGl5GBfxiLIDmBLoMIQIUuvDO6jpm8X5+HoOeol/YvtbNgua7bmglAYkGX0T/AQwhAj6bMuqJuU0GbmSbEk/VDjlu6RNp6OKmrhsRwXDQIiVtDCEDQI3NQWOW9keDrFh+oeFZPFfZ/qiAyKahkg6SollHeAYMIQKng0vpsy4pgdFXy1u9OstCz9EepcOxAiTXpE6YxZEPGwwhAroscPWZbzV6QxmHBYWfriz+oT4RcpYoAHcrPViKnUq9F0Ge0Nw6",
		}},
		Confirmations: 5276880,
		NextBlockHash: "0xd782db8a38b0eea0d7394e0f007c61c71798867578c77c387c08113903946cc9",
	}

	header, err := NewBlockHeaderFromRPC(&rpcHeader)
	assert.Nil(t, err)
	assert.Equal(t, 252, len(header.Witness.VerificationScript))
}

//func TestNewBlockHeaderFromRPC2(t *testing.T) {
//	cli := rpc.NewClient("http://seed1t.neo.org:21332")
//	res := cli.GetBlockHeader("1")
//	h, err := NewBlockHeaderFromRPC(&res.Result)
//	assert.Nil(t, err)
//	log.Println(h.GetHash().String())
//}
