package win

//ref ULONGLONG
//ref ULONG
//ref DWORD
//ref SOCKET_ADDRESS
//import unsafe
func (this *IP_ADAPTER_WINS_SERVER_ADDRESS_LH) Alignment() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_WINS_SERVER_ADDRESS_LH) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_WINS_SERVER_ADDRESS_LH) Reserved() *DWORD {
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}
