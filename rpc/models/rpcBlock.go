package models

type RpcBlockHeader struct {
	Hash              string       `json:"hash"`
	Size              int          `json:"size"`
	Version           int          `json:"version"`
	PreviousBlockHash string       `json:"previousblockhash"`
	MerkleRoot        string       `json:"merkleroot"`
	Time              int          `json:"time"`
	Nonce             string       `json:"nonce"`
	Index             int          `json:"index"`
	PrimaryIndex      byte         `json:"primary"`
	NextConsensus     string       `json:"nextconsensus"` //address
	Witnesses         []RpcWitness `json:"witnesses"`
	Confirmations     int          `json:"confirmations"`
	NextBlockHash     string       `json:"nextblockhash"`
}

type RpcBlock struct {
	RpcBlockHeader
	Tx []RpcTransaction `json:"tx"`
}
