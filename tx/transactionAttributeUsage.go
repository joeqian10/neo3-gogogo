package tx

// Transaction attribute usages
type TransactionAttributeUsage uint8

const (
	Url TransactionAttributeUsage = 0x81 // 129
)

func (u TransactionAttributeUsage) String() string {
	if uint8(u) == 0x81 {return "Url"}
	return "Not Defined"
}

func (u TransactionAttributeUsage) IsDefined() bool {
	if u == Url {return true}
	return false
}