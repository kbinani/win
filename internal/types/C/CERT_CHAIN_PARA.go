package win

//ref DWORD
//ref CERT_USAGE_MATCH

type CERT_CHAIN_PARA struct {
	CbSize         DWORD
	RequestedUsage CERT_USAGE_MATCH
}
