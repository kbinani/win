package win

type MIB_UDP6ROW_OWNER_MODULE struct {
	UcLocalAddr       [16]UCHAR
	DwLocalScopeId    DWORD
	DwLocalPort       DWORD
	DwOwningPid       DWORD
	LiCreateTimestamp LARGE_INTEGER
	dwFlags           int32
	OwningModuleInfo  [TCPIP_OWNING_MODULE_SIZE]ULONGLONG
}
