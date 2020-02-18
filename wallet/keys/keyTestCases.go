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
		Address:    "NNG9H8bHqUTEphsLVW4X2aGB6W9LnrWZv2",
		PrivateKey: "831cb932167332a768f1c898d2cf4586a14aa606b7f078eba028c849c306cce6",
		PublicKey:  "03b7a7f933199f28cc1c48d22a21c78ac3992cf7fceb038a9c670fe55444426619",
		Wif:        "L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z",
		Passphrase: "neo3-gogogo",
		Nep2key:    "6PYMDM4ZwugPFbFqtDwen33yEYtnSSc3fdR8z3iseTH3FeuGxSc4GDcSMj",
		ScriptHash: "9ab8400ee456ceb9eb81678ed9f076a7e9fec019",
	},
	{
		Address:    "NZw4dfRsrWkrJYJRx6Gk6B3Pms4tsfgN8x",
		PrivateKey: "82a4ff38f5de304c2fae62d1a736c343816412f7e4fe3badaf5e1e940c8f07c3",
		PublicKey:  "027d73c8b02e446340caceee7a517cddff72440e60c28cbb84884f307760ecad5b",
		Wif:        "L1bfdDaFQErh7gGMz32zgBXZCN65AKeexKzxeEvS7d4Cq6zf2Rpf",
		Passphrase: "",
		Nep2key:    "6PYNGvyRHwNNRfNFZZAp9ZJ4cUk7n29HsyWJoLw5AshfVqfvbwYaDrk9ne",
		ScriptHash: "9792dad81144defd6fea088e32fb37763db9c699",
	},
	{
		Address:    "NPU5MwmwKLGberKaUhnuBUMCqFGYFtVvhF",
		PrivateKey: "31ab808b68c25377b2068500e264f164d1d75eda748a8e0a98db4c74db181b66",
		PublicKey:  "038a2151948a908cdf2d680eead6512217769e34b9db196574572cb98e273516a1",
		Wif:        "KxtGC6QHFKSiHVLY1ANkwS78ebfhworv6LnkJH2MUxE8AbbgAHVW",
		Passphrase: "密码@👍",
		Nep2key:    "6PYK1i6kKYStFfraDN6hGENYnbtzmXWJwrr2u4dJpFWtDvsZo5NKe2Nqsb",
		ScriptHash: "7842a4dbdc88ca887328443fa4f27615e4d7fa26",
	},
}
