package io

// ISerializable defines the binary encoding/decoding interface. Errors are
// returned via BinaryReader/BinaryWriter Err field. These functions must have safe
// behavior when passed BinaryReader/BinaryWriter with Err already set. Invocations
// to these functions tend to be nested, with this mechanism only the top-level
// caller should handle the error once and all the other code should just not
// panic in presence of error.
type ISerializable interface {
	Deserialize(*BinaryReader)
	Serialize(*BinaryWriter)
}
