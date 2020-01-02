package sc

type ContractParameterType byte

const (
	Signature ContractParameterType = 0x00
	Boolean   ContractParameterType = 0x01
	Integer   ContractParameterType = 0x02
	Hash160   ContractParameterType = 0x03
	Hash256   ContractParameterType = 0x04
	ByteArray ContractParameterType = 0x05
	PublicKey ContractParameterType = 0x06
	String    ContractParameterType = 0x07

	Array ContractParameterType = 0x10
	Map   ContractParameterType = 0x12

	InteropInterface ContractParameterType = 0xf0

	Any  ContractParameterType = 0xfe
	Void ContractParameterType = 0xff
)

type ContractParameter struct {
	Type  ContractParameterType
	Value interface{}
}
