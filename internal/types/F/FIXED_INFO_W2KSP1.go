package win

//ref CHAR
//ref PIP_ADDR_STRING
//ref UINT
//ref IP_ADDR_STRING
//ref PIP_ADDR_STRING
type FIXED_INFO_W2KSP1 struct {
	HostName         [MAX_HOSTNAME_LEN + 4]CHAR
	DomainName       [MAX_DOMAIN_NAME_LEN + 4]CHAR
	CurrentDnsServer PIP_ADDR_STRING
	DnsServerList    IP_ADDR_STRING
	NodeType         UINT
	ScopeId          [MAX_SCOPE_ID_LEN + 4]CHAR
	EnableRouting    UINT
	EnableProxy      UINT
	EnableDns        UINT
}
