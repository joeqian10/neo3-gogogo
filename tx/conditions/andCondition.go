package conditions

import "github.com/joeqian10/neo3-gogogo/helper"

// AndCondition Represents the condition that all conditions must be met.
type AndCondition struct {
	expressions []IWitnessCondition
}

func (this AndCondition) GetType() WitnessConditionType {
	return And
}

func (this AndCondition) GetSize() int {
	sum := 0
	for _, expr := range this.expressions {
		sum += expr.GetSize()
	}
	return And.GetSize() + helper.GetVarSize(len(this.expressions)) + sum
}
