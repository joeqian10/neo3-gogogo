package models

type RpcApplicationLog struct {
	TxId        string `json:"txid"`
	Trigger     string `json:"trigger"`
	VMState     string `json:"vmstate"`
	GasConsumed string `json:"gas_consumed"`
	Stack       []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"stack"`
	Notifications []struct {
		Contract string `json:"contract"`
		State    struct {
			Type  string      `json:"type"`
			Value interface{} `json:"value"`
		} `json:"state"`
	} `json:"notifications"`
}

type RpcContractParameter struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
