package win

//ref IN6_ADDR
//ref DWORD
//ref MIB_TCP_STATE
//ref TCP_CONNECTION_OFFLOAD_STATE

type MIB_TCP6ROW2 struct {
	LocalAddr       IN6_ADDR
	DwLocalScopeId  DWORD
	DwLocalPort     DWORD
	RemoteAddr      IN6_ADDR
	DwRemoteScopeId DWORD
	DwRemotePort    DWORD
	State           MIB_TCP_STATE
	DwOwningPid     DWORD
	DwOffloadState  TCP_CONNECTION_OFFLOAD_STATE
}
