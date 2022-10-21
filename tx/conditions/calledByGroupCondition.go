package conditions

import "github.com/joeqian10/neo3-gogogo/crypto"

// CalledByGroupCondition represents the group to be checked.
type CalledByGroupCondition struct {
	group crypto.ECPoint
}

func (this CalledByGroupCondition) GetType() WitnessConditionType {
	return CalledByGroup
}

func (this CalledByGroupCondition) GetSize() int {
	return CalledByGroup.GetSize() + this.group.GetSize()
}
