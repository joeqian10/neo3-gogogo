package models

type RpcTransaction struct {
	Hash    string `json:"hash"`
	Size    int    `json:"size"`
	Version int    `json:"version"`
	Nonce   int    `json:"nonce"`
	// address
	Sender          string                    `json:"sender"`
	SysFee          string                    `json:"sys_fee"`
	NetFee          string                    `json:"net_fee"`
	ValidUntilBlock int                       `json:"valid_until_block"`
	Attributes      []RpcTransactionAttribute `json:"attributes"`
	Cosigners       []RpcCosigner             `json:"cosigners"`
	Script          string                    `json:"script"`
	Witnesses       []RpcWitness              `json:"witnesses"`
	BlockHash       string                    `json:"blockhash"`
	Confirmations   int                       `json:"confirmations"`
	Blocktime       int                       `json:"blocktime"`
	VMState         string                    `json:"vmState"`
}

type RpcTransactionAttribute struct {
	Usage string `json:"usage"`
	Data  string `json:"data"`
}

type RpcCosigner struct {
	Account          string   `json:"account"`
	Scopes           string   `json:"scopes"`
	AllowedContracts []string `json:"allowedContracts"`
	AllowedGroups    []string `json:"allowedGroups"`
}

type RpcWitness struct {
	Invocation   string `json:"invocation"`
	Verification string `json:"verification"`
}
