package win

//ref ULONG
//ref ULONG64
//ref GUID
//ref HANDLE
//ref LARGE_INTEGER

type WNODE_HEADER struct {
	BufferSize    ULONG
	ProviderId    ULONG
	union1        ULONG64
	union2        uint64
	Guid          GUID
	ClientContext ULONG
	Flags         ULONG
}

func (this *WNODE_HEADER) GetHistoricalContext() ULONG64 {
	return this.union1
}

func (this *WNODE_HEADER) SetHistoricalContext(v ULONG64) {
	this.union1 = v
}

func (this *WNODE_HEADER) GetVersion() ULONG {
	return *(*ULONG)(unsafe.Pointer(&this.union1))
}

func (this *WNODE_HEADER) SetVersion(v ULONG) {
	*(*ULONG)(unsafe.Pointer(&this.union1)) = v
}

func (this *WNODE_HEADER) GetLinkage() ULONG {
	return *(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}

func (this *WNODE_HEADER) SetLinkage(v ULONG) {
	*(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4))) = v
}

func (this *WNODE_HEADER) GetCountLost() ULONG {
	return *(*ULONG)(unsafe.Pointer(&this.union2))
}

func (this *WNODE_HEADER) SetCountLost(v ULONG) {
	*(*ULONG)(unsafe.Pointer(&this.union2)) = v
}

func (this *WNODE_HEADER) GetKernelHandle() HANDLE {
	return *(*HANDLE)(unsafe.Pointer(&this.union2))
}

func (this *WNODE_HEADER) SetKernelHandle(v HANDLE) {
	*(*HANDLE)(unsafe.Pointer(&this.union2)) = v
}

func (this *WNODE_HEADER) GetTimeStamp() LARGE_INTEGER {
	return *(*LARGE_INTEGER)(unsafe.Pointer(&this.union2))
}

func (this *WNODE_HEADER) SetTimeStamp(v LARGE_INTEGER) {
	*(*LARGE_INTEGER)(unsafe.Pointer(&this.union2)) = v
}
