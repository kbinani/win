package win

import "unsafe"

//ref NET_ADDRESS_FORMAT
//ref WCHAR
//ref SOCKADDR_IN
//ref SOCKADDR_IN6
//ref SOCKADDR

type NET_ADDRESS_INFO struct {
	Format  NET_ADDRESS_FORMAT
	storage [524]byte
}

type NET_ADDRESS_INFO_NamedAddress struct {
	Address [DNS_MAX_NAME_BUFFER_LENGTH]WCHAR
	Port    [6]WCHAR
}

func (this *NET_ADDRESS_INFO) NamedAddress() *NET_ADDRESS_INFO_NamedAddress {
	return (*NET_ADDRESS_INFO_NamedAddress)(unsafe.Pointer(&this.storage[0]))
}

func (this *NET_ADDRESS_INFO) Ipv4Address() *SOCKADDR_IN {
	return (*SOCKADDR_IN)(unsafe.Pointer(&this.storage[0]))
}

func (this *NET_ADDRESS_INFO) Ipv6Address() *SOCKADDR_IN6 {
	return (*SOCKADDR_IN6)(unsafe.Pointer(&this.storage[0]))
}

func (this *NET_ADDRESS_INFO) IpAddress() *SOCKADDR {
	return (*SOCKADDR)(unsafe.Pointer(&this.storage[0]))
}
