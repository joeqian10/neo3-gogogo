package helper

const (
	Neo3Magic_MainNet     uint32 = 860833102 // NEO3
	Neo3Magic_TestNet     uint32 = 894710606 // N3T5
	DefaultAddressVersion byte   = 0x35      // 53
)

type ProtocolSettings struct {
	Magic          uint32
	AddressVersion byte
}

var DefaultProtocolSettings = ProtocolSettings{
	Magic:          Neo3Magic_MainNet,
	AddressVersion: DefaultAddressVersion,
}
