package conditions

type WitnessConditionType byte

const (
	Boolean           WitnessConditionType = 0x00
	Not               WitnessConditionType = 0x01
	And               WitnessConditionType = 0x02
	Or                WitnessConditionType = 0x03
	ScriptHash        WitnessConditionType = 0x18 // 0b_0001_1000
	Group             WitnessConditionType = 0x19 // 0b_0001_1001
	CalledByEntryType WitnessConditionType = 0x20 // 0b_0010_0000
	CalledByContract  WitnessConditionType = 0x28 // 0b_0010_1000
	CalledByGroup     WitnessConditionType = 0x29 // 0b_0010_1001
)

func (w WitnessConditionType) String() string {
	b := byte(w)
	switch b {
	case 0x00:
		return "Boolean"
	case 0x01:
		return "Not"
	case 0x02:
		return "And"
	case 0x03:
		return "Or"
	case 0x18:
		return "ScriptHash"
	case 0x19:
		return "Group"
	case 0x20:
		return "CalledByEntry"
	case 0x28:
		return "CalledByContract"
	case 0x29:
		return "CalledByGroup"
	default:
		return ""
	}
}

func (w WitnessConditionType) GetSize() int {
	return 1
}
