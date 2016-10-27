package win

//ref ULONGLONG
//ref SOCKET_ADDRESS
//ref ULONG
//ref DWORD
func (this *IP_ADAPTER_GATEWAY_ADDRESS_LH) Alignment() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_GATEWAY_ADDRESS_LH) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_GATEWAY_ADDRESS_LH) Reserved() *DWORD {
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}
