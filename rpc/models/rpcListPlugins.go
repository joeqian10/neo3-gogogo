package models

type RpcListPlugin struct {
	Name       string   `json:"name"`
	Version    string   `json:"version"`
	Interfaces []string `json:"interfaces"`
}
