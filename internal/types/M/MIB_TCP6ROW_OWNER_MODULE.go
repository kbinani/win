package win

//ref UCHAR
//ref DWORD
//ref LARGE_INTEGER
//ref ULONGLONG
type MIB_TCP6ROW_OWNER_MODULE struct {
	UcLocalAddr       [16]UCHAR
	DwLocalScopeId    DWORD
	DwLocalPort       DWORD
	UcRemoteAddr      [16]UCHAR
	DwRemoteScopeId   DWORD
	DwRemotePort      DWORD
	DwState           DWORD
	DwOwningPid       DWORD
	LiCreateTimestamp LARGE_INTEGER
	OwningModuleInfo  [TCPIP_OWNING_MODULE_SIZE]ULONGLONG
}
