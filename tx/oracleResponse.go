package tx

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/joeqian10/neo3-gogogo/sc"
)

// only oracle can use this attribute
type OracleResponseAttribute struct {
	FixedScript []byte
	Id          uint64
	Code        OracleResponseCode
	Result      []byte
}

func NewOracleResponseAttribute() (*OracleResponseAttribute, error) {
	orcleContractHash, _ := helper.UInt160FromString("0x8dc0e742cbdfdeda51ff8a8b78d46829144c80ee")
	sb := sc.NewScriptBuilder()
	sb.EmitDynamicCall(orcleContractHash, "finish") // 0x8dc0e742cbdfdeda51ff8a8b78d46829144c80ee
	b, err := sb.ToArray()
	if err != nil {
		return nil, err
	}
	return &OracleResponseAttribute{
		FixedScript: b,
	}, nil
}

func (o *OracleResponseAttribute) GetMaxResultSize() int {
	return 65535
}

func (o *OracleResponseAttribute) GetAttributeType() TransactionAttributeType {
	return OracleResponse
}

func (o *OracleResponseAttribute) AllowMultiple() bool {
	return false
}

func (o *OracleResponseAttribute) GetAttributeSize() int {
	return 1 + // base size
		8 + // Id
		1 + // ResponseCode
		sc.ByteSlice(o.Result).GetVarSize() // Result
}

func (o *OracleResponseAttribute) Deserialize(br *io.BinaryReader) {
	if br.ReadByte() != byte(OracleResponse) {
		br.Err = fmt.Errorf("format error: not HighPriority")
	}
	o.DeserializeWithoutType(br)
}

func (o *OracleResponseAttribute) Serialize(bw *io.BinaryWriter) {
	bw.WriteLE(byte(OracleResponse))
	o.SerializeWithoutType(bw)
}

func (o *OracleResponseAttribute) DeserializeWithoutType(br *io.BinaryReader) {
	br.ReadLE(&o.Id)
	code := br.ReadByte()
	if OracleResponseCode(code).IsDefined() {
		o.Code = OracleResponseCode(code)
	} else {
		br.Err = fmt.Errorf("format error: oracle response code is not defined")
		return
	}
	result := br.ReadVarBytesWithMaxLimit(o.GetMaxResultSize())
	if code != byte(Success) && len(result) > 0 {
		br.Err = fmt.Errorf("format error: wrong result")
	}
	o.Result = result
}

func (o *OracleResponseAttribute) SerializeWithoutType(bw *io.BinaryWriter) {
	bw.WriteLE(o.Id)
	bw.WriteLE(byte(o.Code))
	bw.WriteVarBytes(o.Result)
}
