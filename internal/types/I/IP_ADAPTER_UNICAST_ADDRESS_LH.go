package win

//ref ULONGLONG
//ref SOCKET_ADDRESS
//ref IP_PREFIX_ORIGIN
//ref IP_SUFFIX_ORIGIN
//ref IP_DAD_STATE
//ref ULONG
//ref DWORD
//ref UINT8
type IP_ADAPTER_UNICAST_ADDRESS_LH struct {
	union1             ULONGLONG
	Next               *IP_ADAPTER_UNICAST_ADDRESS_LH
	Address            SOCKET_ADDRESS
	PrefixOrigin       IP_PREFIX_ORIGIN
	SuffixOrigin       IP_SUFFIX_ORIGIN
	DadState           IP_DAD_STATE
	ValidLifetime      ULONG
	PreferredLifetime  ULONG
	LeaseLifetime      ULONG
	OnLinkPrefixLength UINT8
}

func (this *IP_ADAPTER_UNICAST_ADDRESS_LH) Alignment() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_UNICAST_ADDRESS_LH) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_UNICAST_ADDRESS_LH) Flags() *DWORD {
	return (*DWORD)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}
