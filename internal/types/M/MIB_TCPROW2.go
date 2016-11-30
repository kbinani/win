package win

//ref DWORD
//ref TCP_CONNECTION_OFFLOAD_STATE

type MIB_TCPROW2 struct {
	DwState        DWORD
	DwLocalAddr    DWORD
	DwLocalPort    DWORD
	DwRemoteAddr   DWORD
	DwRemotePort   DWORD
	DwOwningPid    DWORD
	DwOffloadState TCP_CONNECTION_OFFLOAD_STATE
}
