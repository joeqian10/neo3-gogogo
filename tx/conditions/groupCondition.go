package conditions

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
)

// GroupCondition represents the group to be checked.
type GroupCondition struct {
	group crypto.ECPoint
}

func (this GroupCondition) GetType() WitnessConditionType {
	return Group
}

func (this GroupCondition) GetSize() int {
	return Group.GetSize() + this.group.GetSize()
}
