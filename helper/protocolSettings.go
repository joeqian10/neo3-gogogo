package helper

const (
	Neo3Magic_MainNet     uint32 = 5195086 // 0x4F454Eu
	Neo3Magic_TestNet     uint32 = 1951352142
	DefaultAddressVersion byte   = 0x35
)

type ProtocolSettings struct {
	Magic uint32
	AddressVersion byte
}

var DefaultProtocolSettings = ProtocolSettings{
	Magic:          Neo3Magic_MainNet,
	AddressVersion: DefaultAddressVersion,
}
