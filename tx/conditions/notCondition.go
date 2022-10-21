package conditions

// NotCondition Reverses another condition.
type NotCondition struct {
	expression IWitnessCondition
}

func (this NotCondition) GetType() WitnessConditionType {
	return Not
}

func (this NotCondition) GetSize() int {
	return Not.GetSize() + this.expression.GetSize()
}
