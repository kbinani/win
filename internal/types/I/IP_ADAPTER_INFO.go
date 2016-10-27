package win

//ref DWORD
//ref CHAR
//ref UINT
//ref BYTE
//ref PIP_ADDR_STRING
//ref IP_ADDR_STRING
//ref BOOL
//ref Time_t
type IP_ADAPTER_INFO struct {
	Next                *IP_ADAPTER_INFO
	ComboIndex          DWORD
	AdapterName         [MAX_ADAPTER_NAME_LENGTH + 4]CHAR
	Description         [MAX_ADAPTER_DESCRIPTION_LENGTH + 4]CHAR
	AddressLength       UINT
	Address             [MAX_ADAPTER_ADDRESS_LENGTH]BYTE
	Index               DWORD
	Type                UINT
	DhcpEnabled         UINT
	CurrentIpAddress    PIP_ADDR_STRING
	IpAddressList       IP_ADDR_STRING
	GatewayList         IP_ADDR_STRING
	DhcpServer          IP_ADDR_STRING
	HaveWins            BOOL
	PrimaryWinsServer   IP_ADDR_STRING
	SecondaryWinsServer IP_ADDR_STRING
	LeaseObtained       Time_t
	LeaseExpires        Time_t
}
