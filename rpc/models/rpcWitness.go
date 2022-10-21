package models

import (
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/tx/conditions"
)

type RpcWitness struct {
	Invocation   string `json:"invocation"`
	Verification string `json:"verification"`
}

type RpcWitnessRule struct {
	Action    string              `json:"action"`
	Condition RpcWitnessCondition `json:"condition"`
}

// RpcWitnessCondition combines all types of conditions into one struct
type RpcWitnessCondition struct {
	Type        string                `json:"type"`                  // type of this condition
	Expression  interface{}           `json:"expression,omitempty"`  // BooleanCondition: "true" || "false" | NotCondition: !RpcWitnessCondition
	Expressions []RpcWitnessCondition `json:"expressions,omitempty"` // AndCondition: RpcWitnessCondition && RpcWitnessCondition | OrCondition: RpcWitnessCondition || RpcWitnessCondition
	Hash        string                `json:"hash,omitempty"`        // ScriptHashCondition | CalledByContractCondition: UInt160.ToString()
	Group       string                `json:"group,omitempty"`       // GroupCondition | CalledByGroupCondition: ECPoint.ToString()
}

func CreateRpcWitnessCondition(c *conditions.WitnessCondition) RpcWitnessCondition {
	r := RpcWitnessCondition{Type: c.Type.String()}
	switch c.Type {
	case conditions.Boolean:
		b := c.GetCondition().(*bool)
		if *b {
			r.Expression = "true"
		} else {
			r.Expression = "false"
		}
		break
	case conditions.Not:
		inner := CreateRpcWitnessCondition(c.GetCondition().(*conditions.WitnessCondition))
		r.Expression = inner // get the value from the pointer
		break
	case conditions.And, conditions.Or:
		cs := c.GetCondition().([]*conditions.WitnessCondition)
		inners := make([]RpcWitnessCondition, len(cs))
		for i, _ := range cs {
			inners[i] = CreateRpcWitnessCondition(cs[i])
		}
		r.Expressions = inners
		break
	case conditions.ScriptHash, conditions.CalledByContract:
		hash := c.GetCondition().(*helper.UInt160)
		r.Hash = hash.String()
		break
	case conditions.Group, conditions.CalledByGroup:
		group := c.GetCondition().(*crypto.ECPoint)
		r.Group = group.String()
		break
	case conditions.CalledByEntryType:
		break
	default:
		break
	}
	return r
}
