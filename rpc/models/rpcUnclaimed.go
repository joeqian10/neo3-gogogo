package models

type UnclaimedGasInWallet struct {
	Available   string `json:"available"`
	Unavailable string `json:"unavailable"`
}

type UnclaimedGasInAddress struct {
	Available   float64 `json:"available"`
	Unavailable float64 `json:"unavailable"`
	Unclaimed   float64 `json:"unclaimed"`
}
