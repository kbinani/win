package win

//ref DWORD
//ref LARGE_INTEGER
//ref ULONGLONG
type MIB_TCPROW_OWNER_MODULE struct {
	DwState           DWORD
	DwLocalAddr       DWORD
	DwLocalPort       DWORD
	DwRemoteAddr      DWORD
	DwRemotePort      DWORD
	DwOwningPid       DWORD
	LiCreateTimestamp LARGE_INTEGER
	OwningModuleInfo  [TCPIP_OWNING_MODULE_SIZE]ULONGLONG
}
