package win

type MIB_UDPROW_OWNER_MODULE struct {
	DwLocalAddr       DWORD
	DwLocalPort       DWORD
	DwOwningPid       DWORD
	padding1          [4]byte
	LiCreateTimestamp LARGE_INTEGER
	dwFlags           int32
	padding2          [4]byte
	OwningModuleInfo  [TCPIP_OWNING_MODULE_SIZE]ULONGLONG
}
