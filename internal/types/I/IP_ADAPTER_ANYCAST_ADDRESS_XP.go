package win

//ref ULONGLONG
//ref SOCKET_ADDRESS
import "unsafe"

func (this *IP_ADAPTER_ANYCAST_ADDRESS_XP) Alignment() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_ANYCAST_ADDRESS_XP) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_ANYCAST_ADDRESS_XP) Flags() *DWORD {
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}
