package win

//ref UCHAR
//ref DWORD
//ref LARGE_INTEGER
//ref ULONGLONG

type MIB_UDP6ROW_OWNER_MODULE struct {
	UcLocalAddr       [16]UCHAR
	DwLocalScopeId    DWORD
	DwLocalPort       DWORD
	DwOwningPid       DWORD
	padding1          [pad0for64_4for32]byte
	LiCreateTimestamp LARGE_INTEGER
	dwFlags           int32
	padding2          [pad0for64_4for32]byte
	OwningModuleInfo  [TCPIP_OWNING_MODULE_SIZE]ULONGLONG
}

func (this *MIB_UDP6ROW_OWNER_MODULE) DwFlags() *int32 {
	return &this.dwFlags
}

func (this *MIB_UDP6ROW_OWNER_MODULE) SpecificPortBind() int32 {
	return this.dwFlags & 0x1
}
