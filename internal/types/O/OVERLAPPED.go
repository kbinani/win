package win

//ref ULONG_PTR
//ref HANDLE
//ref DWORD
//ref PVOID
//import unsafe
type OVERLAPPED struct {
	Internal     ULONG_PTR
	InternalHigh ULONG_PTR
	union1       [8]byte
	HEvent       HANDLE
}

func (this *OVERLAPPED) Offset() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.union1[0]))
}
func (this *OVERLAPPED) OffsetHigh() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.union1[4]))
}
func (this *OVERLAPPED) Pointer() *PVOID {
	return (*PVOID)(unsafe.Pointer(&this.union1[0]))
}
