package win

//ref MIB_TCP_STATE
//ref IN6_ADDR
//ref DWORD

type MIB_TCP6ROW struct {
	State           MIB_TCP_STATE
	LocalAddr       IN6_ADDR
	DwLocalScopeId  DWORD
	DwLocalPort     DWORD
	RemoteAddr      IN6_ADDR
	DwRemoteScopeId DWORD
	DwRemotePort    DWORD
}
