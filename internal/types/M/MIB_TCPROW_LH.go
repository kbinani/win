package win

//ref MIB_TCP_STATE
//ref DWORD
type MIB_TCPROW_LH struct {
	State        MIB_TCP_STATE
	DwLocalAddr  DWORD
	DwLocalPort  DWORD
	DwRemoteAddr DWORD
	DwRemotePort DWORD
}
