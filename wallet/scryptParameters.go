package wallet

// ScryptParameters is a json-serializable container for scrypt KDF parameters.
type ScryptParameters struct {
	N int `json:"n"`
	R int `json:"r"`
	P int `json:"p"`
}

var DefaultScryptParameters = NewScryptParameters(16384, 8, 8)

func NewScryptParameters(n int, r int, p int) *ScryptParameters {
	return &ScryptParameters{
		N: n,
		R: r,
		P: p,
	}
}
