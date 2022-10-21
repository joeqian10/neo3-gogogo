package conditions

// BooleanCondition Indicates that the condition will always be met or not met.
type BooleanCondition struct {
	expression bool
}

func (this BooleanCondition) GetType() WitnessConditionType {
	return Boolean
}

func (this BooleanCondition) GetSize() int {
	return Boolean.GetSize() + 1
}
