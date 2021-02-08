package models

type RpcApplicationLog struct {
	TxId       string         `json:"txid"`
	BlockHash  string         `json:"blockhash"`
	Executions []RpcExecution `json:"executions"`
}

type RpcExecution struct {
	Trigger       string            `json:"trigger"`
	VMState       string            `json:"vmstate"`
	GasConsumed   string            `json:"gasconsumed"`
	Stack         []InvokeStack     `json:"stack"`
	Notifications []RpcNotification `json:"notifications"`
}

type RpcNotification struct {
	Contract  string      `json:"contract"`
	EventName string      `json:"eventname"`
	State     InvokeStack `json:"state"`
}
