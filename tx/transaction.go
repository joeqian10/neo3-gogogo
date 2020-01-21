package tx

import (
	"bytes"
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/helper/io"
	"sort"
)

const (
	TransactionVersion          uint8 = 0 // neo-2.x
	MaxTransactionSize                = 102400
	MaxValidUntilBlockIncrement       = 2102400
	MaxTransactionAttributes          = 16 // Maximum number of attributes that can be contained within a transaction
	MaxCosigners                      = 16 // Maximum number of cosigners that can be contained within a transaction
)

// base class
type Transaction struct {
	version         uint8
	nonce           uint
	sender          helper.UInt160
	sysfee          int64
	netfee          int64
	validUntilBlock uint
	attributes      []*TransactionAttribute
	cosigners       []*Cosigner
	script          []byte
	witnesses       []*Witness

	_hash *helper.UInt256
	_size int
}

func (tx *Transaction) HeaderSize() int {
	buf := bytes.Buffer{}
	buf.WriteByte(byte(tx.version))
	buf.Write(helper.UInt32ToBytes(uint32(tx.nonce)))
	buf.Write(tx.sender.Bytes())
	buf.Write(helper.UInt64ToBytes(uint64(tx.sysfee)))
	buf.Write(helper.UInt64ToBytes(uint64(tx.netfee)))
	buf.Write(helper.UInt32ToBytes(uint32(tx.validUntilBlock)))
	return len(buf.Bytes())
}

// GetAttributes is the getter of tx.attributes
func (tx *Transaction) GetAttributes() []*TransactionAttribute {
	return tx.attributes
}

// SetAttributes is the setter of tx.attributes
func (tx *Transaction) SetAttributes(value []*TransactionAttribute) {
	tx.attributes = value
	tx._hash = nil
	tx._size = 0
}

// GetCosigners is the getter of tx.cosigners
func (tx *Transaction) GetCosigners() []*Cosigner {
	return tx.cosigners
}

// SetCosigners is the setter of tx.cosigners
func (tx *Transaction) SetCosigners(value []*Cosigner) {
	tx.cosigners = value
	tx._hash = nil
	tx._size = 0
}

// GetHash is the getter of tx._hash
func (tx *Transaction) GetHash() helper.UInt256 {
	if tx._hash == nil {
		hash, _ := helper.UInt256FromBytes(crypto.Hash256(tx.GetHashData()))
		tx._hash = &hash
	}
	return *tx._hash
}

// TODO UT
func (tx *Transaction) GetHashData() []byte {
	buf := io.NewBufBinaryWriter()
	tx.SerializeUnsigned(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

// GetNetworkFee is the getter of tx.netfee
func (tx *Transaction) GetNetworkFee() int64 {
	return tx.netfee
}

// SetNetworkFee is the setter of tx.netfee
func (tx *Transaction) SetNetworkFee(value int64) {
	tx.netfee = value
	tx._hash = nil
}

// GetNonce is the getter of tx.nonce
func (tx *Transaction) GetNonce() uint {
	return tx.nonce
}

// SetNonce is the setter of tx.nonce
func (tx *Transaction) SetNonce(value uint) {
	tx.nonce = value
	tx._hash = nil
}

// GetScript is the getter of tx.script
func (tx *Transaction) GetScript() []byte {
	return tx.script
}

// SetScript is the setter of tx.script
func (tx *Transaction) SetScript(value []byte) {
	tx.script = value
	tx._hash = nil
	tx._size = 0
}

// GetSender is the getter of tx.sender
func (tx *Transaction) GetSender() helper.UInt160 {
	return tx.sender
}

// SetSender is the setter of tx.sender
func (tx *Transaction) SetSender(value helper.UInt160) {
	tx.sender = value
	tx._hash = nil
}

func (tx *Transaction) GetSize() int {
	if tx._size == 0 {
		tx._size = len(tx.RawTransaction())
	}
	return tx._size
}

func (tx *Transaction) RawTransaction() []byte {
	buf := io.NewBufBinaryWriter()
	tx.Serialize(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
}

// GetSystemFee is the getter of tx.sysfee
func (tx *Transaction) GetSystemFee() int64 {
	return tx.sysfee
}

// SetSystemFee is the setter of tx.sysfee
func (tx *Transaction) SetSystemFee(value int64) {
	tx.sysfee = value
	tx._hash = nil
}

// GetValidUntilBlock is the getter of tx.validUntilBlock
func (tx *Transaction) GetValidUntilBlock() uint {
	return tx.validUntilBlock
}

// SetValidUntilBlock is the setter of tx.validUntilBlock
func (tx *Transaction) SetValidUntilBlock(value uint) {
	tx.validUntilBlock = value
	tx._hash = nil
}

// GetVersion is the getter of tx.version
func (tx *Transaction) GetVersion() uint8 {
	return tx.version
}

// SetVersion is the setter of tx.version
func (tx *Transaction) SetVersion(value uint8) {
	tx.version = value
	tx._hash = nil
}

// GetWitnesses is the getter of tx.witnesses
func (tx *Transaction) GetWitnesses() []*Witness {
	return tx.witnesses
}

// SetWitnesses is the setter of tx.witnesses
func (tx *Transaction) SetWitnesses(value []*Witness) {
	tx.witnesses = value
	tx._hash = nil
}

// Deserialize implements Serializable interface.
func (tx *Transaction) Deserialize(br *io.BinaryReader) {
	tx.DeserializeUnsigned(br)
	tx.DeserializeWitnesses(br)
}

func (tx *Transaction) DeserializeUnsigned(br *io.BinaryReader) {
	// version
	br.ReadLE(&tx.version)
	if tx.version > 0 {
		br.Err = fmt.Errorf("format error: version > 0")
	}
	// nonce
	br.ReadLE(&tx.nonce)
	// sender
	br.ReadLE(&tx.sender)
	// sysfee
	br.ReadLE(&tx.sysfee)
	if tx.sysfee < 0 {
		br.Err = fmt.Errorf("format error: sysfee < 0")
	}
	if tx.sysfee%100000000 != 0 {
		br.Err = fmt.Errorf("format error: sysfee is not an integer")
	}
	// netfee
	br.ReadLE(&tx.netfee)
	if tx.netfee < 0 {
		br.Err = fmt.Errorf("format error: netfee < 0")
	}
	// validUntilBlock
	br.ReadLE(&tx.validUntilBlock)
	// attributes
	lenAttributes := br.ReadVarUInt(MaxTransactionAttributes)
	tx.attributes = make([]*TransactionAttribute, lenAttributes)
	for i := 0; i < int(lenAttributes); i++ {
		tx.attributes[i] = &TransactionAttribute{} // may not be needed
		tx.attributes[i].Deserialize(br)
	}
	// cosigners
	lenCosigners := br.ReadVarUInt(MaxCosigners)
	tx.cosigners = make([]*Cosigner, lenCosigners)
	for i := 0; i < int(lenCosigners); i++ {
		tx.cosigners[i] = &Cosigner{} // may not be needed
		tx.cosigners[i].Deserialize(br)
	}
	// script
	tx.script = br.ReadVarBytes(65535)
	if len(tx.script) == 0 {
		br.Err = fmt.Errorf("format error: script is empty")
	}
}

func (tx *Transaction) DeserializeWitnesses(br *io.BinaryReader) {
	lenWitnesses := br.ReadVarUInt(16777216)
	tx.witnesses = make([]*Witness, lenWitnesses)
	for i := 0; i < int(lenWitnesses); i++ {
		tx.witnesses[i] = &Witness{}
		tx.witnesses[i].Deserialize(br)
	}
}

// Serialize implements Serializable interface.
func (tx *Transaction) Serialize(bw *io.BinaryWriter) {
	tx.SerializeUnsigned(bw)
	tx.SerializeWitnesses(bw)
}

func (tx *Transaction) SerializeUnsigned(bw *io.BinaryWriter) {
	// version
	bw.WriteLE(tx.version)
	// nonce
	bw.WriteLE(tx.nonce)
	// sender
	bw.WriteLE(tx.sender)
	// sysfee
	bw.WriteLE(tx.sysfee)
	// netfee
	bw.WriteLE(tx.netfee)
	// validUntilBlock
	bw.WriteLE(tx.validUntilBlock)
	// attributes
	bw.WriteVarUInt(uint64(len(tx.attributes)))
	for _, attr := range tx.attributes {
		attr.Serialize(bw)
	}
	// cosigners
	bw.WriteVarUInt(uint64(len(tx.cosigners)))
	for _, cosigner := range tx.cosigners {
		cosigner.Serialize(bw)
	}
	// script
	bw.WriteVarBytes(tx.script)
}

func (tx *Transaction) SerializeWitnesses(bw *io.BinaryWriter) {
	bw.WriteVarUInt(uint64(len(tx.witnesses)))
	for _, s := range tx.witnesses {
		s.Serialize(bw)
	}
}

func (tx *Transaction) GetScriptHashesForVerifying() []helper.UInt160 {
	hashmaps := map[helper.UInt160]bool{tx.sender:true}
	for _, cosigner := range tx.cosigners {
		if !hashmaps[cosigner.Account] {
			hashmaps[cosigner.Account] = true
		}
	}
	hashes := []helper.UInt160{}
	for key, _ := range hashmaps {
		hashes = append(hashes, key)
	}
	sort.Sort(helper.UInt160Slice(hashes))
	return hashes
}

