package conditions

import "github.com/joeqian10/neo3-gogogo/helper"

// ScriptHashCondition represents the script hash to be checked.
type ScriptHashCondition struct {
	hash helper.UInt160
}

func (this ScriptHashCondition) GetType() WitnessConditionType {
	return ScriptHash
}

func (this ScriptHashCondition) GetSize() int {
	return ScriptHash.GetSize() + this.hash.GetSize()
}
