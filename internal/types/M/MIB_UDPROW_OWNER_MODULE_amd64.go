package win

type MIB_UDPROW_OWNER_MODULE struct {
	DwLocalAddr       DWORD
	DwLocalPort       DWORD
	DwOwningPid       DWORD
	LiCreateTimestamp LARGE_INTEGER
	dwFlags           int32
	OwningModuleInfo  [TCPIP_OWNING_MODULE_SIZE]ULONGLONG
}
