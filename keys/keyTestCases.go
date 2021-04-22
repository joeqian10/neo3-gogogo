package keys

// Ktype represents key test case values (different encodings of the key).
type Ktype struct {
	Address,
	PrivateKey,
	PublicKey,
	Wif,
	Passphrase,
	Nep2key,
	ScriptHash string
}

// Arr contains a set of known keys in Ktype format.
var KeyCases = []Ktype{
	{
		Passphrase: "neo3-gogogo",
		PrivateKey: "831cb932167332a768f1c898d2cf4586a14aa606b7f078eba028c849c306cce6",
		PublicKey:  "03b7a7f933199f28cc1c48d22a21c78ac3992cf7fceb038a9c670fe55444426619",
		Wif:        "L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z",
		Nep2key:    "6PYXStTdYPTC3XzscAyL3MZikD4hTWgeV1aaDqTSxiWKWfk9VjZcUbm6Uz",
		ScriptHash: "0685cf3f7b1ea94f8f61be3fb8e06d081af34cc9", // big endian
		Address:    "NeGMBsEZ44B52dEvnm73ZU2QqLoRmTzMfr",
	},
	{
		Passphrase: "",
		PrivateKey: "82a4ff38f5de304c2fae62d1a736c343816412f7e4fe3badaf5e1e940c8f07c3",
		PublicKey:  "027d73c8b02e446340caceee7a517cddff72440e60c28cbb84884f307760ecad5b",
		Wif:        "L1bfdDaFQErh7gGMz32zgBXZCN65AKeexKzxeEvS7d4Cq6zf2Rpf",
		Nep2key:    "6PYL8gFU1rZZ5qbKqJJAK7fPTvM8RJfEiF8k7ranqjKKgzB3RxxbH1uySg",
		ScriptHash: "97c96fe6fcb153de1293cc5537337b62ab2c84cb",
		Address:    "NeU4hLo5Lgkp4RkZnMbQjbNSLtYYxJ6eZG",
	},
	{
		Passphrase: "密码@👍",
		PrivateKey: "31ab808b68c25377b2068500e264f164d1d75eda748a8e0a98db4c74db181b66",
		PublicKey:  "038a2151948a908cdf2d680eead6512217769e34b9db196574572cb98e273516a1",
		Wif:        "KxtGC6QHFKSiHVLY1ANkwS78ebfhworv6LnkJH2MUxE8AbbgAHVW",
		Nep2key:    "6PYS6KAwfyD26moSEE3L7DvQ4LSjdyy9HA35oC82Yaiu8523HHzYxzxoFc",
		ScriptHash: "8febbbd7e397c2c784b4bdef2992e2bbe6edca65",
		Address:    "NVCCXuVjRKt1ac7ri4WuVC8sDrYeZ6NeCx",
	},
}
