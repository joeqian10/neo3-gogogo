package tx

import (
	"bytes"
	"fmt"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
)

const (
	TransactionVersion          uint8  = 0
	MaxTransactionSize                 = 102400
	MaxValidUntilBlockIncrement uint32 = 5760 // 24 hours
	MaxTransactionAttributes           = 16   // Maximum number of attributes that can be contained within a transaction
	MaxSigners                         = 16   // Maximum number of cosigners that can be contained within a transaction
)

const NeoTokenId = "0xef4073a0f2b305a38ec4050e4d3d28bc40ea63f5"
const GasTokenId = "0xd2a4cff31913016155e38e474a2c06d08be276cf"

const GasFactor = 100000000
const ExecFeeFactor = 30
const FeePerByte = 1000
const ECDsaVerifyPrice = 1 << 15

var NeoToken, _ = helper.UInt160FromString(NeoTokenId)
var GasToken, _ = helper.UInt160FromString(GasTokenId)

type Transaction struct {
	version         uint8
	nonce           uint32
	sysfee          int64
	netfee          int64
	validUntilBlock uint32
	signers         []Signer
	attributes      []ITransactionAttribute
	script          []byte
	witnesses       []Witness

	_hash *helper.UInt256
	_size int
}

func NewTransaction() *Transaction {
	return &Transaction{
		version:         uint8(0),
		nonce:           0,
		sysfee:          0,
		netfee:          0,
		validUntilBlock: 0,
		signers:         []Signer{},
		attributes:      []ITransactionAttribute{},
		script:          []byte{},
		witnesses:       []Witness{},
	}
}

func (tx *Transaction) HeaderSize() int {
	buf := bytes.Buffer{}
	buf.WriteByte(tx.version)                           // 1
	buf.Write(helper.UInt32ToBytes(tx.nonce))           // 4
	buf.Write(helper.UInt64ToBytes(uint64(tx.sysfee)))  // 8
	buf.Write(helper.UInt64ToBytes(uint64(tx.netfee)))  // 8
	buf.Write(helper.UInt32ToBytes(tx.validUntilBlock)) // 4
	return len(buf.Bytes())                             // total 25
}

// GetAttributes is the getter of tx.attributes
func (tx *Transaction) GetAttributes() []ITransactionAttribute {
	return tx.attributes
}

// SetAttributes is the setter of tx.attributes
func (tx *Transaction) SetAttributes(value []ITransactionAttribute) {
	tx.attributes = value
	tx._hash = nil
	tx._size = 0
}

/// The <c>NetworkFee</c> for the transaction divided by its <c>Size</c>.
/// <para>Note that this property must be used with care. Getting the value of this property multiple times will return the same result. The value of this property can only be obtained after the transaction has been completely built (no longer modified).</para>
func (tx *Transaction) FeePerByte() int64 {
	return tx.netfee / int64(tx._size)
}

// GetHash is the getter of tx._hash, using default magic
func (tx *Transaction) GetHash() *helper.UInt256 {
	if tx._hash == nil {
		tx._hash = CalculateHash(tx)
	}
	return tx._hash
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
func (tx *Transaction) GetNonce() uint32 {
	return tx.nonce
}

// SetNonce is the setter of tx.nonce
func (tx *Transaction) SetNonce(value uint32) {
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

// GetSender
func (tx *Transaction) GetSender() *helper.UInt160 {
	return tx.signers[0].Account
}

// GetSigners is the getter of tx.signers
func (tx *Transaction) GetSigners() []Signer {
	return tx.signers
}

// SetSigners is the setter of tx.signers
func (tx *Transaction) SetSigners(value []Signer) {
	tx.signers = value
	tx._hash = nil
	tx._size = 0
}

// GetSize is the getter of tx._size
func (tx *Transaction) GetSize() int {
	if tx._size == 0 {
		tx._size = len(tx.ToByteArray())
	}
	return tx._size
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
func (tx *Transaction) GetValidUntilBlock() uint32 {
	return tx.validUntilBlock
}

// SetValidUntilBlock is the setter of tx.validUntilBlock
func (tx *Transaction) SetValidUntilBlock(value uint32) {
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
func (tx *Transaction) GetWitnesses() []Witness {
	return tx.witnesses
}

// SetWitnesses is the setter of tx.witnesses
func (tx *Transaction) SetWitnesses(value []Witness) {
	tx.witnesses = value
	tx._hash = nil
}

// ToByteArray returns signed tx data
func (tx *Transaction) ToByteArray() []byte {
	buf := io.NewBufBinaryWriter()
	tx.Serialize(buf.BinaryWriter)
	if buf.Err != nil {
		return nil
	}
	return buf.Bytes()
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
		return
	}
	// nonce
	br.ReadLE(&tx.nonce)
	// sysfee
	br.ReadLE(&tx.sysfee)
	if tx.sysfee < 0 {
		br.Err = fmt.Errorf("format error: sysfee < 0")
		return
	}
	// netfee
	br.ReadLE(&tx.netfee)
	if tx.netfee < 0 {
		br.Err = fmt.Errorf("format error: netfee < 0")
		return
	}
	if tx.sysfee+tx.netfee < tx.sysfee {
		br.Err = fmt.Errorf("format error: overflow")
		return
	}
	// validUntilBlock
	br.ReadLE(&tx.validUntilBlock)
	// signers
	tx.signers = deserializeSigners(br, MaxTransactionAttributes)
	// attributes
	tx.attributes = deserializeAttributes(br, MaxTransactionAttributes-len(tx.signers))
	// script
	tx.script = br.ReadVarBytesWithMaxLimit(65535)
	if len(tx.script) == 0 {
		br.Err = fmt.Errorf("format error: script is empty")
	}
}

func deserializeAttributes(br *io.BinaryReader, maxCount int) []ITransactionAttribute {
	count := int(br.ReadVarUIntWithMaxLimit(uint64(maxCount)))
	result := make([]ITransactionAttribute, count)
	m := make(map[TransactionAttributeType]ITransactionAttribute)
	for i := 0; i < count; i++ {
		attribute := DeserializeFrom(br)
		if attribute == nil {
			return nil
		}
		if !attribute.AllowMultiple() && m[attribute.GetAttributeType()] == attribute {
			br.Err = fmt.Errorf("format error: duplicate attribute")
			return nil
		}
		result[i] = attribute
	}
	return result
}

func deserializeSigners(br *io.BinaryReader, maxCount int) []Signer {
	count := int(br.ReadVarUIntWithMaxLimit(uint64(maxCount)))
	if count == 0 {
		br.Err = fmt.Errorf("format error: signer count is zero")
		return nil
	}
	result := make([]Signer, count)
	m := make(map[helper.UInt160]Signer)
	for i := 0; i < count; i++ {
		signer := NewDefaultSigner()
		signer.Deserialize(br)
		if t, ok := m[*signer.Account]; ok && (&t).CompareTo(signer) == 0 {
			br.Err = fmt.Errorf("format error: duplicate signer")
			return nil
		}
		result[i] = *signer
	}
	return result
}

func (tx *Transaction) DeserializeWitnesses(br *io.BinaryReader) {
	lenWitnesses := br.ReadVarUInt()
	tx.witnesses = make([]Witness, lenWitnesses)
	for i := 0; i < int(lenWitnesses); i++ {
		tx.witnesses[i] = Witness{}
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
	// sysfee
	bw.WriteLE(tx.sysfee)
	// netfee
	bw.WriteLE(tx.netfee)
	// validUntilBlock
	bw.WriteLE(tx.validUntilBlock)
	// signers
	bw.WriteVarUInt(uint64(len(tx.signers)))
	for _, signer := range tx.signers {
		signer.Serialize(bw)
	}
	// attributes
	bw.WriteVarUInt(uint64(len(tx.attributes)))
	for _, attr := range tx.attributes {
		attr.Serialize(bw)
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
	result := make([]helper.UInt160, len(tx.signers))
	for i, s := range tx.signers {
		result[i] = *s.Account
	}
	return result
}
