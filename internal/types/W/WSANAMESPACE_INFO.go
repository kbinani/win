package win

//ref GUID
//ref DWORD
//ref BOOL
//ref LPWSTR
type WSANAMESPACE_INFO struct {
	NSProviderId   GUID
	DwNameSpace    DWORD
	FActive        BOOL
	DwVersion      DWORD
	LpszIdentifier LPWSTR
}
