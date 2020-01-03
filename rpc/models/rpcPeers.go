package models

type RpcPeers struct {
	Unconnected []Peer `json:"unconnected"`
	Bad         []Peer `json:"bad"`
	Connected   []Peer `json:"connected"`
}

type Peer struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}
