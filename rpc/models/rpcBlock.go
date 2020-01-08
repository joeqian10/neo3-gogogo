package models

type RpcBlockHeader struct {
	Hash              string       `json:"hash"`
	Size              int          `json:"size"`
	Version           int          `json:"version"`
	PreviousBlockHash string       `json:"previousblockhash"`
	MerkleRoot        string       `json:"merkleroot"`
	Time              int          `json:"time"`
	Index             int          `json:"index"`
	NextConsensus     string       `json:"nextconsensus"` //address
	Witnesses         []RpcWitness `json:"witnesses"`
	Confirmations     int          `json:"confirmations"`
	NextBlockHash     string       `json:"nextblockhash"`
	CrossStatesRoot   string       `json:"crossstatesroot"`
	ChainID           string       `json:"chainid"` //ulong = uint64
}

type RpcBlock struct {
	RpcBlockHeader
	ConsensusData struct {
		PrimaryIndex int    `json:"primary"`
		Nonce        string `json:"nonce"`
	} `json:"consensus_data"`
	Tx []RpcTransaction `json:"tx"`
}

//func (bh *RpcBlockHeader) Deserialize(br *io.BinaryReader) {
//	bh.DeserializeUnsigned(br)
//	bh.DeserializeWitness(br)
//}
//
////DeserializeUnsigned deserialize blockheader without witness
//func (bh *RpcBlockHeader) DeserializeUnsigned(br *io.BinaryReader) {
//	var h, ph, mr, cr helper.UInt256
//	br.ReadLE(&h)
//	bh.Hash = h.String()
//
//	br.ReadLE(&bh.Version)
//
//	br.ReadLE(&ph)
//	bh.PreviousBlockHash = ph.String()
//
//	br.ReadLE(&mr)
//	bh.MerkleRoot = mr.String()
//	br.ReadLE(&bh.Time)
//	br.ReadLE(&bh.Index)
//
//	var nextConsensus helper.UInt160
//	br.ReadLE(&nextConsensus)
//	bh.NextBlockHash = helper.ScriptHashToAddress(nextConsensus)
//
//	bh.ChainID = helper.BytesToHex(helper.ReverseBytes(br.ReadUInt64Bytes()))
//	br.ReadLE(&cr)
//	bh.CrossStatesRoot = cr.String()
//}
//
////DeserializeWitness deserialize witness
//func (bh *RpcBlockHeader) DeserializeWitness(br *io.BinaryReader) {
//	var padding uint8
//	br.ReadLE(&padding)
//	if padding != 1 {
//		br.Err = fmt.Errorf("format error: padding must equal 1 got %d", padding)
//		return
//	}
//	bh.Witnesses[0].Invocation = helper.BytesToHex(br.ReadVarBytes(663))
//	bh.Witnesses[0].Verification = helper.BytesToHex(br.ReadVarBytes(361))
//}
//
//func (bh *RpcBlockHeader) Serialize(bw *io.BufBinaryWriter) {
//	bh.SerializeUnsigned(bw)
//	bh.SerializeWitness(bw)
//}
//
////SerializeUnsigned serialize blockheader without witness
//func (bh *RpcBlockHeader) SerializeUnsigned(bw *io.BufBinaryWriter) {
//	var h, ph, mr, cr helper.UInt256
//	h, _ = helper.UInt256FromString(bh.Hash)
//	bw.WriteLE(h)
//	bw.WriteLE(bh.Version)
//	ph, _ = helper.UInt256FromString(bh.PreviousBlockHash)
//	bw.WriteLE(ph)
//	mr, _ = helper.UInt256FromString(bh.MerkleRoot)
//	bw.WriteLE(mr)
//	bw.WriteLE(bh.Time)
//	bw.WriteLE(bh.Index)
//	var nc helper.UInt160
//	nc, _ = helper.AddressToScriptHash(bh.NextConsensus)
//	bw.WriteLE(nc)
//	bw.WriteLE(helper.ReverseBytes(helper.HexTobytes(bh.ChainID)))
//	cr, _ = helper.UInt256FromString(bh.CrossStatesRoot)
//	bw.WriteLE(cr)
//}
//
////SerializeWitness serialize witness
//func (bh *RpcBlockHeader) SerializeWitness(bw *io.BufBinaryWriter) {
//	bw.WriteLE(uint8(1))
//	bw.WriteVarBytes(helper.HexTobytes(bh.Witnesses[0].Invocation))
//	bw.WriteVarBytes(helper.HexTobytes(bh.Witnesses[0].Verification))
//}
