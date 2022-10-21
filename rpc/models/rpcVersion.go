package models

type RpcVersion struct {
	TcpPort   int         `json:"tcpPort"`
	WsPort    int         `json:"wsPort"`
	Nonce     string      `json:"nonce"`
	UserAgent string      `json:"useragent"`
	Protocol  RpcProtocol `json:"protocol"`
}

type RpcProtocol struct {
	AddressVersion              byte   `json:"addressversion"`
	Network                     uint32 `json:"network"`
	ValidatorsCount             int32  `json:"validatorscount"`
	MillisecondsPerBlock        uint32 `json:"msperblock"`
	MaxTraceableBlocks          uint32 `json:"maxtraceableblocks"`
	MaxValidUntilBlockIncrement uint32 `json:"maxvaliduntilblockincrement"`
	MaxTransactionsPerBlock     uint32 `json:"maxtransactionsperblock"`
	MemoryPoolMaxTransactions   int32  `json:"memorypoolmaxtransactions"`
	InitialGasDistribution      uint64 `json:"initialgasdistribution"`
}
