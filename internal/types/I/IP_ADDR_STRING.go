package win

//ref IP_ADDRESS_STRING
//ref IP_MASK_STRING
//ref DWORD
type IP_ADDR_STRING struct {
	Next      *IP_ADDR_STRING
	IpAddress IP_ADDRESS_STRING
	IpMask    IP_MASK_STRING
	Context   DWORD
}
