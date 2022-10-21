package models

import (
	"github.com/joeqian10/neo3-gogogo/tx"
)

type RpcSigner struct {
	Account          string           `json:"account"` // script hash
	Scopes           string           `json:"scopes"`
	AllowedContracts []string         `json:"allowedcontracts,omitempty"`
	AllowedGroups    []string         `json:"allowedgroups,omitempty"`
	Rules            []RpcWitnessRule `json:"rules,omitempty"`
}

func CreateRpcSigners(signers []*tx.Signer) []RpcSigner {
	rpcSigners := make([]RpcSigner, len(signers))
	for i, _ := range signers {
		rpcSigners[i] = CreateRpcSigner(signers[i])
	}
	return rpcSigners
}

func CreateRpcSigner(signer *tx.Signer) RpcSigner {
	allowedContracts := make([]string, len(signer.AllowedContracts))
	allowedGroups := make([]string, len(signer.AllowedGroups))
	rules := make([]RpcWitnessRule, len(signer.Rules))
	for i, _ := range signer.AllowedContracts {
		allowedContracts[i] = signer.AllowedContracts[i].String()
	}
	for i, _ := range signer.AllowedGroups {
		allowedGroups[i] = signer.AllowedGroups[i].String()
	}
	for i, _ := range signer.Rules {
		rule := signer.Rules[i]
		rules[i] = RpcWitnessRule{
			Action:    rule.Action.String(),
			Condition: CreateRpcWitnessCondition(rule.Condition),
		}
	}
	return RpcSigner{
		Account:          signer.Account.String(),
		Scopes:           signer.Scopes.String(),
		AllowedContracts: allowedContracts,
		AllowedGroups:    allowedGroups,
		Rules:            rules,
	}
}
