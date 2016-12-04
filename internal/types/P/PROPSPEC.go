package win

import "unsafe"

//ref ULONG
//ref PROPID
//ref LPOLESTR

type PROPSPEC struct {
	ulKind ULONG
	union1 uintptr
}

func (this *PROPSPEC) Propid() *PROPID {
	return (*PROPID)(unsafe.Pointer(&this.union1))
}

func (this *PROPSPEC) Lpwstr() *LPOLESTR {
	return (*LPOLESTR)(unsafe.Pointer(&this.union1))
}
