package tx

type WitnessRuleAction byte

const (
	Deny  WitnessRuleAction = 0
	Allow WitnessRuleAction = 1
)

func (this WitnessRuleAction) String() string {
	b := byte(this)
	switch b {
	case 0x00:
		return "Deny"
	case 0x01:
		return "Allow"
	default:
		return ""
	}
}

func (this WitnessRuleAction) GetSize() int {
	return 1
}
