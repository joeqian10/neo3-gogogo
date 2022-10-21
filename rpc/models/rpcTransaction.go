package models

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
