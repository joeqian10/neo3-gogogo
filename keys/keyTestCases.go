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
		Nep2key:    "6PYN7P7VnqHXEtsmn98gU9Vi65zg1rhLVdk4m8Uj9LChnVyZ7Cdq3rBLJK",
		ScriptHash: "ddf8645739c8fc40ef728237c060d15b6cdba70e",
		Address:    "NMFTjMmZH2HvGHQ7auPivhnV6ksXjQxSJF",
	},
	{
		Passphrase: "",
		PrivateKey: "82a4ff38f5de304c2fae62d1a736c343816412f7e4fe3badaf5e1e940c8f07c3",
		PublicKey:  "027d73c8b02e446340caceee7a517cddff72440e60c28cbb84884f307760ecad5b",
		Wif:        "L1bfdDaFQErh7gGMz32zgBXZCN65AKeexKzxeEvS7d4Cq6zf2Rpf",
		Nep2key:    "6PYRaBycimfkXpwMdg1Uxni8JaumKxNRpJfzvgQEFmq5ZkER51Ls77cAhJ",
		ScriptHash: "e80c0ea8ab8e92f7ca52e5d9515aac09648875c2",
		Address:    "NdeB5Hrwh5yPeZ32K8HPcnQVbq91ihYgAK",
	},
	{
		Passphrase: "密码@👍",
		PrivateKey: "31ab808b68c25377b2068500e264f164d1d75eda748a8e0a98db4c74db181b66",
		PublicKey:  "038a2151948a908cdf2d680eead6512217769e34b9db196574572cb98e273516a1",
		Wif:        "KxtGC6QHFKSiHVLY1ANkwS78ebfhworv6LnkJH2MUxE8AbbgAHVW",
		Nep2key:    "6PYSAoG1fxCvBm7FBLj5Yasa98NwpGuZLSBi5tAdHnrLHwXyiHhx83VB7F",
		ScriptHash: "b6da874681fc2280a5d400418d072277153d0bfb",
		Address:    "NioNFn49Qng91bLaoi123ieGr7pAgcZnxf",
	},
}
