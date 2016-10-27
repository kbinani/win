package win

//ref ULONGLONG
//ref PIP_ADAPTER_UNICAST_ADDRESS_LH
//ref PIP_ADAPTER_ANYCAST_ADDRESS_XP
//ref PIP_ADAPTER_MULTICAST_ADDRESS_XP
//ref PIP_ADAPTER_DNS_SERVER_ADDRESS_XP
//ref PWCHAR
//ref BYTE
//ref ULONG
//ref IFTYPE
//ref PIP_ADAPTER_PREFIX_XP
//ref PIP_ADAPTER_WINS_SERVER_ADDRESS_LH
//ref PIP_ADAPTER_GATEWAY_ADDRESS_LH
//ref IF_LUID
//ref NET_IF_COMPARTMENT_ID
//ref NET_IF_NETWORK_GUID
//ref NET_IF_CONNECTION_TYPE
//ref TUNNEL_TYPE
//ref PIP_ADAPTER_DNS_SUFFIX
//import unsafe
func (this *IP_ADAPTER_ADDRESSES_LH) Alignment() *ULONGLONG {
	return &this.union1
}
func (this *IP_ADAPTER_ADDRESSES_LH) Length() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}
func (this *IP_ADAPTER_ADDRESSES_LH) IfIndex() *IF_INDEX {
	return (*IF_INDEX)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}
func (this *IP_ADAPTER_ADDRESSES_LH) Flags() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union2))
}
func (this *IP_ADAPTER_ADDRESSES_LH) DdnsEnabled() bool {
	return this.union2 == 0x80000000
}
func (this *IP_ADAPTER_ADDRESSES_LH) RegisterAdapterSuffix() bool {
	return this.union2 == 0x20000000
}
func (this *IP_ADAPTER_ADDRESSES_LH) Dhcpv4Enabled() bool {
	return this.union2 == 0x10000000
}
func (this *IP_ADAPTER_ADDRESSES_LH) ReceiveOnly() bool {
	return this.union2 == 0x8000000
}
func (this *IP_ADAPTER_ADDRESSES_LH) NoMulticast() bool {
	return this.union2 == 0x2000000
}
func (this *IP_ADAPTER_ADDRESSES_LH) Ipv6OtherStatefulConfig() bool {
	return this.union2 == 0x1000000
}
func (this *IP_ADAPTER_ADDRESSES_LH) NetbiosOverTcpipEnabled() bool {
	return this.union2 == 0x800000
}
func (this *IP_ADAPTER_ADDRESSES_LH) Ipv4Enabled() bool {
	return this.union2 == 0x200000
}
func (this *IP_ADAPTER_ADDRESSES_LH) Ipv6Enabled() bool {
	return this.union2 == 0x100000
}
func (this *IP_ADAPTER_ADDRESSES_LH) Ipv6ManagedAddressConfigurationSupported() bool {
	return this.union2 == 0x80000
}
