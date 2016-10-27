package win

//ref DWORD
//ref GUID
//ref WSAPROTOCOLCHAIN
//ref WCHAR
type WSAPROTOCOL_INFO struct {
	DwServiceFlags1    DWORD
	DwServiceFlags2    DWORD
	DwServiceFlags3    DWORD
	DwServiceFlags4    DWORD
	DwProviderFlags    DWORD
	ProviderId         GUID
	DwCatalogEntryId   DWORD
	ProtocolChain      WSAPROTOCOLCHAIN
	IVersion           int32
	IAddressFamily     int32
	IMaxSockAddr       int32
	IMinSockAddr       int32
	ISocketType        int32
	IProtocol          int32
	IProtocolMaxOffset int32
	INetworkByteOrder  int32
	ISecurityScheme    int32
	DwMessageSize      DWORD
	DwProviderReserved DWORD
	SzProtocol         [WSAPROTOCOL_LEN + 1]WCHAR
}
