package conditions

type CalledByEntryCondition struct {
}

func (this CalledByEntryCondition) GetType() WitnessConditionType {
	return CalledByEntryType
}

func (this CalledByEntryCondition) GetSize() int {
	return CalledByEntryType.GetSize()
}
