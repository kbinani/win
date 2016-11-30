package win

import "unsafe"

//ref ADDRESS_FAMILY
//ref USHORT
//ref ULONG
//ref IN6_ADDR
//ref SCOPE_ID

type SOCKADDR_IN6_LH struct {
	sin6_family   ADDRESS_FAMILY
	sin6_port     USHORT
	sin6_flowinfo ULONG
	sin6_addr     IN6_ADDR
	union1        ULONG
}

func (this *SOCKADDR_IN6_LH) Sin6_scope_id() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union1))
}

func (this *SOCKADDR_IN6_LH) Sin6_scope_struct() *SCOPE_ID {
	return (*SCOPE_ID)(unsafe.Pointer(&this.union1))
}
