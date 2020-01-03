package models

type ContractState struct {
	Hash     string           `json:"hash"`
	Script   string           `json:"script"`
	Manifest ContractManifest `json:"manifest"`
}

type ContractManifest struct {
	Groups   []ContractGroup `json:"groups"`
	Features struct {
		HasStorage bool `json:"storage"`
		Payable    bool `json:"payable"`
	} `json:"features"`
	Abi         ContractAbi          `json:"abi"`
	Permissions []ContractPermission `json:"permissions"`
	Trusts      []string             `json:"trusts"`
	SafeMethods []string             `json:"safeMethods"`
	Extra       interface{}          `json:"extra"`
}

type ContractGroup struct {
	PubKey    string `json:"pubKey"`
	Signature string `json:"signature"`
}

type ContractAbi struct {
	Hash       string                     `json:"hash"`
	EntryPoint ContractMethodDescriptor   `json:"entryPoint"`
	Methods    []ContractMethodDescriptor `json:"methods"`
	Events     []ContractMethodDescriptor `json:"events"`
}

type ContractMethodDescriptor struct {
	Name       string `json:"name"`
	Parameters struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"parameters"`
	ReturnType string `json:"returnType"`
}

type ContractPermission struct {
	Contract string   `json:"contract"`
	Methods  []string `json:"methods"`
}
