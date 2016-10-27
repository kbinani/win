package win

//ref DWORD
type WSAPROTOCOLCHAIN struct {
	ChainLen     int32
	ChainEntries [MAX_PROTOCOL_CHAIN]DWORD
}
