package win

import "unsafe"

//ref NTSTATUS
//ref PVOID
//ref ULONG_PTR

type IO_STATUS_BLOCK struct {
	union1      uintptr
	Information ULONG_PTR
}

func (this *IO_STATUS_BLOCK) Status() *NTSTATUS {
	return (*NTSTATUS)(unsafe.Pointer(&this.union1))
}

func (this *IO_STATUS_BLOCK) Pointer() *PVOID {
	return (*PVOID)(unsafe.Pointer(&this.union1))
}
