package tx

import (
	"github.com/joeqian10/neo3-gogogo/io"
	"github.com/joeqian10/neo3-gogogo/tx/conditions"
)

type WitnessRule struct {
	Action    WitnessRuleAction
	Condition *conditions.WitnessCondition
}

func (this *WitnessRule) GetSize() int {
	return this.Action.GetSize() + this.Condition.GetSize()
}

func (this *WitnessRule) Deserialize(br *io.BinaryReader) {
	a := br.ReadOneByte()
	if br.Err != nil {
		return
	}
	this.Action = WitnessRuleAction(a)
	this.Condition = new(conditions.WitnessCondition)
	this.Condition.Deserialize(br)
}

func (this *WitnessRule) Serialize(bw *io.BinaryWriter) {
	bw.WriteLE(this.Action)
	this.Condition.Serialize(bw)
}
