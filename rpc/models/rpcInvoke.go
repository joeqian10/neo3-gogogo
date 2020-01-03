package models

type InvokeResult struct {
	Script      string              `json:"script"`
	State       string              `json:"state"`
	GasConsumed string              `json:"gas_consumed"`
	Stack       []InvokeStackResult `json:"stack"`
}

type InvokeStackResult struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type InvokeFunctionStackArg struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func NewInvokeFunctionStackArg(t string, v string) InvokeFunctionStackArg {
	return InvokeFunctionStackArg{Type: t, Value: v}
}

func NewInvokeFunctionStackByteArray(value string) InvokeFunctionStackArg {
	return InvokeFunctionStackArg{Type: "ByteArray", Value: value}
}
