package win

import "unsafe"

type MIDIHDR struct {
	storage [64]byte
}

func (this *MIDIHDR) LpData() *LPSTR { // 4
	return (*LPSTR)(unsafe.Pointer(&this.storage[0]))
}
func (this *MIDIHDR) DwBufferLength() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[4]))
}
func (this *MIDIHDR) DwBytesRecorded() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[8]))
}
func (this *MIDIHDR) DwUser() *DWORD_PTR { // 4
	return (*DWORD_PTR)(unsafe.Pointer(&this.storage[12]))
}
func (this *MIDIHDR) DwFlags() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[16]))
}
func (this *MIDIHDR) LpNext() **MIDIHDR { // 4
	return (**MIDIHDR)(unsafe.Pointer(&this.storage[20]))
}
func (this *MIDIHDR) Reserved() *DWORD_PTR { // 4
	return (*DWORD_PTR)(unsafe.Pointer(&this.storage[24]))
}
func (this *MIDIHDR) DwOffset() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[28]))
}
func (this *MIDIHDR) DwReserved() *[8]DWORD_PTR { // 32
	return (*[8]DWORD_PTR)(unsafe.Pointer(&this.storage[32]))
}
