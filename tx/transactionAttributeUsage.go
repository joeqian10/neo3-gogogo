package tx

// Transaction attribute usages
type TransactionAttributeUsage uint8

const (
	Url TransactionAttributeUsage = 0x81 // 129
)

func (u TransactionAttributeUsage) String() string {
	return "Url"
}