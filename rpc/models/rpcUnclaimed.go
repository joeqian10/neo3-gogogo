package models

type UnclaimedGasInWallet struct {
	Available   string `json:"available"`
	Unavailable string `json:"unavailable"`
}

type UnclaimedGasInAddress struct {
	Available   uint64 `json:"available"`
	Unavailable uint64 `json:"unavailable"`
	Unclaimed   uint64 `json:"unclaimed"`
}

type UnclaimedGas struct {
	Unclaimed uint64 `json:"unclaimed"`
	Address   string `json:"address"`
}
