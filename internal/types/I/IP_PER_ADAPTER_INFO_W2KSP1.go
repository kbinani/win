package win

//ref UINT
//ref PIP_ADDR_STRING
//ref IP_ADDR_STRING
type IP_PER_ADAPTER_INFO_W2KSP1 struct {
	AutoconfigEnabled UINT
	AutoconfigActive  UINT
	CurrentDnsServer  PIP_ADDR_STRING
	DnsServerList     IP_ADDR_STRING
}
