package win

//ref ULONGLONG
//ref SOCKET_ADDRESS
//ref ULONG
//ref DWORD
import "unsafe"

type IP_ADAPTER_PREFIX_XP struct {
	union1       ULONGLONG
	Next         *IP_ADAPTER_PREFIX_XP
	Address      SOCKET_ADDRESS
	PrefixLength ULONG
}

func (this *IP_ADAPTER_PREFIX_XP) Alignment() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_PREFIX_XP) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_PREFIX_XP) Flags() *DWORD {
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}
