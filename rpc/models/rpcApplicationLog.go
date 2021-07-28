package models

type RpcApplicationLog struct {
	TxId       string         `json:"txid"`
	Executions []RpcExecution `json:"executions"`
}

type RpcExecution struct {
	Trigger       string            `json:"trigger"`
	VMState       string            `json:"vmstate"`
	Exception	  string			`json:"exception"`
	GasConsumed   string            `json:"gasconsumed"`
	Stack         []InvokeStack     `json:"stack"`
	Notifications []RpcNotification `json:"notifications"`
}

type RpcNotification struct {
	Contract  string      `json:"contract"`
	EventName string      `json:"eventname"`
	State     InvokeStack `json:"state"`
}
