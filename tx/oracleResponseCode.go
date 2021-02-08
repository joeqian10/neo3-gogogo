package tx

type OracleResponseCode byte

const (
	Success OracleResponseCode = 0x00

	ProtocolNotSupported = 0x10
	ConsensusUnreachable = 0x12
	NotFound = 0x14
	Timeout = 0x16
	Forbidden = 0x18
	ResponseTooLarge = 0x1a
	InsufficientFunds = 0x1c

	Error = 0xff
)

func (code OracleResponseCode) IsDefined() bool {
	b := byte(code)
	switch b {
	case 0x00, 0x10, 0x12, 0x14,
		0x16, 0x18, 0x1a, 0x1c,
		0xff:
		return true
	default:
		return false
	}
}
