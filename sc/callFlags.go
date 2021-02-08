package sc

type CallFlags byte

const (
	None CallFlags = 0

	ReadStates  CallFlags = 0b00000001
	WriteStates CallFlags = 0b00000010
	AllowCall   CallFlags = 0b00000100
	AllowNotify CallFlags = 0b00001000

	States CallFlags = ReadStates | WriteStates
	ReadOnly CallFlags = ReadStates | AllowCall
	All CallFlags = States | AllowCall |AllowNotify
)
