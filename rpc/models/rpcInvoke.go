package models

type InvokeResult struct {
	Script           string                      `json:"script"`
	State            string                      `json:"state"`
	GasConsumed      string                      `json:"gasconsumed"`
	Exception        string                      `json:"exception"`
	Notifications    []RpcNotification           `json:"notifications"`
	Diagnostics      RpcDiagnostic               `json:"diagnostics,omitempty"`
	Stack            []InvokeStack               `json:"stack"` // "error: invalid operation" | InvokeStack[]
	Session          string                      `json:"session,omitempty"`
	Tx               string                      `json:"tx,omitempty"`
	PendingSignature RpcContractParameterContext `json:"pendingsignature,omitempty"`
}

type RpcDiagnostic struct {
	InvokedContracts []RpcInvocationTreeNode `json:"invokedcontracts"`
}

type RpcInvocationTreeNode struct {
	Hash string                  `json:"hash"`
	Call []RpcInvocationTreeNode `json:"call,omitempty"`
}

type RpcStorageChange struct {
	State string `json:"state"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type InvokeStack struct {
	Type      string      `json:"type"`
	Value     interface{} `json:"value"`
	Interface string      `json:"interface,omitempty"`
	Id        string      `json:"id,omitempty"`
}
