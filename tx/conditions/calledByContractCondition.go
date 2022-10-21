package conditions

import "github.com/joeqian10/neo3-gogogo/helper"

// CalledByContractCondition represents the contract script hash to be checked.
type CalledByContractCondition struct {
	hash helper.UInt160
}

func (this CalledByContractCondition) GetType() WitnessConditionType {
	return CalledByContract
}

func (this CalledByContractCondition) GetSize() int {
	return CalledByContract.GetSize() + this.hash.GetSize()
}
