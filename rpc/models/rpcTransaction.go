package models

import (
	"github.com/joeqian10/neo3-gogogo/tx"
)

type RpcTransaction struct {
	Hash            string                    `json:"hash"`
	Size            int                       `json:"size"`
	Version         int                       `json:"version"`
	Nonce           int                       `json:"nonce"`
	Sender          string                    `json:"sender"`
	SysFee          string                    `json:"sysfee"`
	NetFee          string                    `json:"netfee"`
	ValidUntilBlock int                       `json:"validuntilblock"`
	Signers         []RpcSigner               `json:"signers"`
	Attributes      []RpcTransactionAttribute `json:"attributes"`
	Script          string                    `json:"script"`
	Witnesses       []RpcWitness              `json:"witnesses"`
	BlockHash       string                    `json:"blockhash"`
	Confirmations   int                       `json:"confirmations"`
	Blocktime       int                       `json:"blocktime"`
}

type RpcTransactionAttribute struct {
	Usage string `json:"usage"`
	Data  string `json:"data"`
}

type RpcSigner struct {
	Account          string   `json:"account"` // script hash
	Scopes           string   `json:"scopes"`
	AllowedContracts []string `json:"allowedcontracts"`
	AllowedGroups    []string `json:"allowedgroups"`
}

func CreateRpcSigners(signers []tx.Signer) []RpcSigner {
	rpcSigners := make([]RpcSigner, len(signers))
	for i, _ := range signers {
		rpcSigners[i] = CreateRpcSigner(signers[i])
	}
	return rpcSigners
}

func CreateRpcSigner(signer tx.Signer) RpcSigner {
	allowedContracts := make([]string, len(signer.AllowedContracts))
	allowedGroups := make([]string, len(signer.AllowedGroups))
	for i, _ := range signer.AllowedContracts {
		allowedContracts[i] = signer.AllowedContracts[i].String()
	}
	for i, _ := range signer.AllowedGroups {
		allowedGroups[i] = signer.AllowedGroups[i].String()
	}
	return RpcSigner{
		Account:          signer.Account.String(),
		Scopes:           signer.Scopes.String(),
		AllowedContracts: allowedContracts,
		AllowedGroups:    allowedGroups,
	}
}

type RpcWitness struct {
	Invocation   string `json:"invocation"`
	Verification string `json:"verification"`
}
