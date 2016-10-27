package win

//ref DWORD
//ref LPWSTR
//ref LPGUID
//ref LPWSAVERSION
//ref LPAFPROTOCOLS
//ref LPCSADDR_INFO
//ref LPBLOB
type WSAQUERYSET struct {
	DwSize                  DWORD
	LpszServiceInstanceName LPWSTR
	LpServiceClassId        LPGUID
	LpVersion               LPWSAVERSION
	LpszComment             LPWSTR
	DwNameSpace             DWORD
	LpNSProviderId          LPGUID
	LpszContext             LPWSTR
	DwNumberOfProtocols     DWORD
	LpafpProtocols          LPAFPROTOCOLS
	LpszQueryString         LPWSTR
	DwNumberOfCsAddrs       DWORD
	LpcsaBuffer             LPCSADDR_INFO
	DwOutputFlags           DWORD
	LpBlob                  LPBLOB
}
