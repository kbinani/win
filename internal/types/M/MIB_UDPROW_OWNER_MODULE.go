package win

//ref DWORD
//ref LARGE_INTEGER
//ref ULONGLONG

type MIB_UDPROW_OWNER_MODULE struct {
	DwLocalAddr       DWORD
	DwLocalPort       DWORD
	DwOwningPid       DWORD
	padding1          [pad0for64_4for32]byte
	LiCreateTimestamp LARGE_INTEGER
	dwFlags           int32
	padding2          [pad0for64_4for32]byte
	OwningModuleInfo  [TCPIP_OWNING_MODULE_SIZE]ULONGLONG
}

func (this *MIB_UDPROW_OWNER_MODULE) DwFlags() *int32 {
	return &this.dwFlags
}

func (this *MIB_UDPROW_OWNER_MODULE) SpecificPortBind() int32 {
	return this.dwFlags & 0x1
}
