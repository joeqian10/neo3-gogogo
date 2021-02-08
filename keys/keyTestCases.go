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
		Address:    "NTY5R9UBKDUmLCGvQVba5JqnrGKctZkEig",
		PrivateKey: "831cb932167332a768f1c898d2cf4586a14aa606b7f078eba028c849c306cce6",
		PublicKey:  "03b7a7f933199f28cc1c48d22a21c78ac3992cf7fceb038a9c670fe55444426619",
		Wif:        "L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z",
		Passphrase: "neo3-gogogo",
		Nep2key:    "6PYX7SaH7EdDgeHo1V8yXfhrgrGXjHaMzsgWLMpaxLkDR9u6HHnEDMmjhh",
		ScriptHash: "dc1a4b4ea1681d872b426cad34e88d426f0e9d53",
	},
	{
		Address:    "NfVhCCpDsnW4Cb6hfKUFdkJVBLnnuBwe2S",
		PrivateKey: "82a4ff38f5de304c2fae62d1a736c343816412f7e4fe3badaf5e1e940c8f07c3",
		PublicKey:  "027d73c8b02e446340caceee7a517cddff72440e60c28cbb84884f307760ecad5b",
		Wif:        "L1bfdDaFQErh7gGMz32zgBXZCN65AKeexKzxeEvS7d4Cq6zf2Rpf",
		Passphrase: "",
		Nep2key:    "6PYUHgyMMWL4kfUm7NVCPjJ1rZCkzPmZ5VpmQQ4N3ymaZSqN2xYrBqkGNn",
		ScriptHash: "6e7d9f206f03d3c4d60657908f3afbe9d32ecbd6",
	},
	{
		Address:    "NRMsJdbogVVmsfFsjDHf6wTywpSRCCba3o",
		PrivateKey: "31ab808b68c25377b2068500e264f164d1d75eda748a8e0a98db4c74db181b66",
		PublicKey:  "038a2151948a908cdf2d680eead6512217769e34b9db196574572cb98e273516a1",
		Wif:        "KxtGC6QHFKSiHVLY1ANkwS78ebfhworv6LnkJH2MUxE8AbbgAHVW",
		Passphrase: "密码@👍",
		Nep2key:    "6PYRECeQ1Fpjs8YBfCk4qnwJHJsEMpCKk17ZNFVu8mzR4WuH7dfyvqDq8A",
		ScriptHash: "cc930ef8ca95a3f88a7ed10c5975db63ca88be3b",
	},
}
