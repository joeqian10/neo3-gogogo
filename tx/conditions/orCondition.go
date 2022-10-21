package conditions

import "github.com/joeqian10/neo3-gogogo/helper"

// OrCondition Represents the condition that any of the conditions musts.
type OrCondition struct {
	expressions []IWitnessCondition
}

func (this OrCondition) GetType() WitnessConditionType {
	return Or
}

func (this OrCondition) GetSize() int {
	sum := 0
	for _, expr := range this.expressions {
		sum += expr.GetSize()
	}
	return Or.GetSize() + helper.GetVarSize(len(this.expressions)) + sum
}
