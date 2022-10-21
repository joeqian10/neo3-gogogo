package conditions

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/io"
)

const (
	MaxSubItems     = 16
	MaxNestingDepth = 2
)

// IWitnessCondition is not used, use WitnessCondition instead
type IWitnessCondition interface {
	GetType() WitnessConditionType
	GetSize() int
}

type WitnessCondition struct {
	Type                      WitnessConditionType
	booleanCondition          *bool
	notCondition              *WitnessCondition
	andConditions             []*WitnessCondition
	orConditions              []*WitnessCondition
	scriptHashCondition       *helper.UInt160
	groupCondition            *crypto.ECPoint
	calledByContractCondition *helper.UInt160
	calledByGroupCondition    *crypto.ECPoint
}

func NewWitnessCondition(t WitnessConditionType, v interface{}) *WitnessCondition {
	wc, err := NewWitnessConditionWithError(t, v)
	if err != nil {
		return nil
	}
	return wc
}

func NewWitnessConditionWithError(t WitnessConditionType, v interface{}) (*WitnessCondition, error) {
	switch t {
	case Boolean:
		x, ok := v.(*bool)
		if !ok {
			return nil, fmt.Errorf("invalid value for BooleanCondition")
		}
		return &WitnessCondition{Type: Boolean, booleanCondition: x}, nil
	case Not:
		x, ok := v.(*WitnessCondition)
		if !ok {
			return nil, fmt.Errorf("invalid value for NotCondition")
		}
		return &WitnessCondition{Type: Not, notCondition: x}, nil
	case And:
		x, ok := v.([]*WitnessCondition)
		if !ok {
			return nil, fmt.Errorf("invalid value for AndCondition")
		}
		return &WitnessCondition{Type: And, andConditions: x}, nil
	case Or:
		x, ok := v.([]*WitnessCondition)
		if !ok {
			return nil, fmt.Errorf("invalid value for OrCondition")
		}
		return &WitnessCondition{Type: Or, orConditions: x}, nil
	case ScriptHash:
		x, ok := v.(*helper.UInt160)
		if !ok {
			return nil, fmt.Errorf("invalid value for ScriptHashCondition")
		}
		return &WitnessCondition{Type: ScriptHash, scriptHashCondition: x}, nil
	case Group:
		x, ok := v.(*crypto.ECPoint)
		if !ok {
			return nil, fmt.Errorf("invalid value for GroupCondition")
		}
		return &WitnessCondition{Type: Group, groupCondition: x}, nil
	case CalledByEntryType:
		return &WitnessCondition{Type: CalledByEntryType}, nil
	case CalledByContract:
		x, ok := v.(*helper.UInt160)
		if !ok {
			return nil, fmt.Errorf("invalid value for CalledByContractCondition")
		}
		return &WitnessCondition{Type: CalledByContract, calledByContractCondition: x}, nil
	case CalledByGroup:
		x, ok := v.(*crypto.ECPoint)
		if !ok {
			return nil, fmt.Errorf("invalid value for CalledByGroupCondition")
		}
		return &WitnessCondition{Type: CalledByGroup, calledByGroupCondition: x}, nil
	default:
		return nil, fmt.Errorf("not supported witness condition type")
	}
}

func (this *WitnessCondition) GetCondition() interface{} {
	switch this.Type {
	case Boolean:
		return this.booleanCondition
	case Not:
		return this.notCondition
	case And:
		return this.andConditions
	case Or:
		return this.orConditions
	case ScriptHash:
		return this.scriptHashCondition
	case Group:
		return this.groupCondition
	case CalledByEntryType:
		return nil
	case CalledByContract:
		return this.calledByContractCondition
	case CalledByGroup:
		return this.calledByGroupCondition
	default:
		return nil
	}
}

func (this *WitnessCondition) GetSize() int {
	size := this.Type.GetSize()
	switch this.Type {
	case Boolean:
		size += 1 // bool
	case Not:
		size += this.notCondition.GetSize()
	case And:
		size += helper.GetVarSize(len(this.andConditions))
		for _, expr := range this.andConditions {
			size += expr.GetSize()
		}
	case Or:
		size += helper.GetVarSize(len(this.orConditions))
		for _, expr := range this.orConditions {
			size += expr.GetSize()
		}
	case ScriptHash:
		size += this.scriptHashCondition.GetSize()
	case Group:
		size += this.groupCondition.GetSize()
	case CalledByEntryType:
		return size
	case CalledByContract:
		size += this.calledByContractCondition.GetSize()
	case CalledByGroup:
		size += this.calledByGroupCondition.GetSize()
	default:
		return size
	}
	return size
}

func (this *WitnessCondition) Deserialize(br *io.BinaryReader) {
	wc := DeserializeConditionFrom(br, MaxNestingDepth)
	if br.Err != nil {
		return
	}
	if this.Type != wc.Type {
		br.Err = fmt.Errorf("witness condition type does not match when deserializing")
		return
	}
	switch this.Type {
	case Boolean:
		this.booleanCondition = wc.booleanCondition
	case Not:
		this.notCondition = wc.notCondition
	case And:
		this.andConditions = wc.andConditions
	case Or:
		this.orConditions = wc.orConditions
	case ScriptHash:
		this.scriptHashCondition = wc.scriptHashCondition
	case Group:
		this.groupCondition = wc.groupCondition
	case CalledByEntryType:
		return
	case CalledByContract:
		this.calledByContractCondition = wc.calledByContractCondition
	case CalledByGroup:
		this.calledByGroupCondition = wc.calledByGroupCondition
	default:
		br.Err = fmt.Errorf("not supported witness condition type")
		return
	}
}

func DeserializeConditions(br *io.BinaryReader, maxNestDepth int) []*WitnessCondition {
	length := br.ReadVarUIntWithMaxLimit(MaxSubItems)
	conditions := make([]*WitnessCondition, length)
	for i, _ := range conditions {
		conditions[i] = DeserializeConditionFrom(br, maxNestDepth)
		if br.Err != nil {
			return nil
		}
	}
	return conditions
}

func DeserializeConditionFrom(br *io.BinaryReader, maxNestDepth int) *WitnessCondition {
	t := br.ReadOneByte()
	if br.Err != nil {
		return nil
	}
	condition := &WitnessCondition{Type: WitnessConditionType(t)}
	switch condition.Type {
	case Boolean:
		br.ReadLE(condition.booleanCondition)
	case Not:
		if maxNestDepth <= 0 {
			br.Err = fmt.Errorf("max nest depth exceeded")
		}
		condition.notCondition = DeserializeConditionFrom(br, maxNestDepth-1)
	case And:
		if maxNestDepth <= 0 {
			br.Err = fmt.Errorf("max nest depth exceeded")
		}
		condition.andConditions = DeserializeConditions(br, maxNestDepth-1)
	case Or:
		if maxNestDepth <= 0 {
			br.Err = fmt.Errorf("max nest depth exceeded")
		}
		condition.andConditions = DeserializeConditions(br, maxNestDepth-1)
	case ScriptHash:
		condition.scriptHashCondition = new(helper.UInt160)
		condition.scriptHashCondition.Deserialize(br)
	case Group:
		condition.groupCondition = new(crypto.ECPoint)
		condition.groupCondition.Deserialize(br)
	case CalledByEntryType:
		return condition
	case CalledByContract:
		condition.calledByContractCondition = new(helper.UInt160)
		condition.calledByContractCondition.Deserialize(br)
	case CalledByGroup:
		condition.calledByGroupCondition = new(crypto.ECPoint)
		condition.calledByGroupCondition.Deserialize(br)
	default:
		br.Err = fmt.Errorf("not supported witness condition type")
	}
	if br.Err != nil {
		return nil
	}
	return condition
}

func (this *WitnessCondition) Serialize(bw *io.BinaryWriter) {
	bw.WriteLE(this.Type)
	switch this.Type {
	case Boolean:
		bw.WriteLE(this.booleanCondition)
	case Not:
		this.notCondition.Serialize(bw)
	case And:
		bw.WriteVarUInt(uint64(len(this.andConditions)))
		for _, x := range this.andConditions {
			x.Serialize(bw)
		}
	case Or:
		bw.WriteVarUInt(uint64(len(this.orConditions)))
		for _, x := range this.orConditions {
			x.Serialize(bw)
		}
	case ScriptHash:
		this.scriptHashCondition.Serialize(bw)
	case Group:
		this.groupCondition.Serialize(bw)
	case CalledByEntryType:
		return
	case CalledByContract:
		this.calledByContractCondition.Serialize(bw)
	case CalledByGroup:
		this.calledByGroupCondition.Serialize(bw)
	default:
		bw.Err = fmt.Errorf("not supported witness condition type")
	}
}

func CreateWitnessCondition(t WitnessConditionType, v interface{}) (IWitnessCondition, error) {
	switch t {
	case Boolean:
		x, ok := v.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid value for BooleanCondition")
		}
		return BooleanCondition{expression: x}, nil
	case Not:
		x, ok := v.(IWitnessCondition)
		if !ok {
			return nil, fmt.Errorf("invalid value for NotCondition")
		}
		return NotCondition{expression: x}, nil
	case And:
		x, ok := v.([]IWitnessCondition)
		if !ok {
			return nil, fmt.Errorf("invalid value for AndCondition")
		}
		return AndCondition{expressions: x}, nil
	case Or:
		x, ok := v.([]IWitnessCondition)
		if !ok {
			return nil, fmt.Errorf("invalid value for OrCondition")
		}
		return OrCondition{expressions: x}, nil
	case ScriptHash:
		x, ok := v.(helper.UInt160)
		if !ok {
			return nil, fmt.Errorf("invalid value for ScriptHashCondition")
		}
		return ScriptHashCondition{hash: x}, nil
	case Group:
		x, ok := v.(crypto.ECPoint)
		if !ok {
			return nil, fmt.Errorf("invalid value for GroupCondition")
		}
		return GroupCondition{group: x}, nil
	case CalledByEntryType:
		return CalledByEntryCondition{}, nil
	case CalledByContract:
		x, ok := v.(helper.UInt160)
		if !ok {
			return nil, fmt.Errorf("invalid value for CalledByContractCondition")
		}
		return CalledByContractCondition{hash: x}, nil
	case CalledByGroup:
		x, ok := v.(crypto.ECPoint)
		if !ok {
			return nil, fmt.Errorf("invalid value for CalledByGroupCondition")
		}
		return CalledByGroupCondition{group: x}, nil
	default:
		return nil, fmt.Errorf("not supported witness condition type")
	}
}
