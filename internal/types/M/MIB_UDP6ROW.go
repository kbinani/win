package win

//ref IN6_ADDR
//ref DWORD

type MIB_UDP6ROW struct {
	DwLocalAddr    IN6_ADDR
	DwLocalScopeId DWORD
	DwLocalPort    DWORD
}
