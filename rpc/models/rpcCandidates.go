package models

type RpcCandidates struct {
	PublicKey string `json:"publickey"`
	Votes     string `json:"votes"`
	Active    bool   `json:"active"`
}
