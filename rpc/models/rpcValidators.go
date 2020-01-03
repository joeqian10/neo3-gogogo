package models

type RpcValidator struct {
	PublicKey string `json:"publickey"`
	Votes     string `json:"votes"`
	Active    bool   `json:"active"`
}
