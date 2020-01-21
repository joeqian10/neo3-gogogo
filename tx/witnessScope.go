package tx

type WitnessScope uint8

const (
	/// Global allows this witness in all contexts (default Neo2 behavior).
	/// This cannot be combined with other flags.
	Global WitnessScope = 0x00

	/// CalledByEntry means that this condition must hold: EntryScriptHash == CallingScriptHash
	/// No params is needed, as the witness/permission/signature given on first invocation will automatically expire if entering deeper internal invokes
	/// This can be default safe choice for native NEO/GAS (previously used on Neo 2 as "attach" mode)
	CalledByEntry WitnessScope = 0x01

	/// Custom hash for contract-specific
	CustomContracts WitnessScope = 0x10

	/// Custom pubkey for group members
	CustomGroups WitnessScope = 0x20
)
