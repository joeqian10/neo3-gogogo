package tx

// Transaction attribute type
type TransactionAttributeType byte

const (
	HighPriority   TransactionAttributeType = 0x01 //
	OracleResponse TransactionAttributeType = 0x11
)

func (u TransactionAttributeType) String() string {
	b := byte(u)
	switch b {
	case 0x01:
		return "HighPriority"
	case 0x11:
		return "OracleResponse"
	default:
		return "Not Defined"
	}
}

func (u TransactionAttributeType) IsDefined() bool {
	if u.String() != "Not Defined" {
		return true
	}
	return false
}
