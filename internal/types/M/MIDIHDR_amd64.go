package win

//import unsafe
type MIDIHDR struct {
	storage [112]byte
}

func (this *MIDIHDR) LpData() *LPSTR {
	return (*LPSTR)(unsafe.Pointer(&this.storage[0]))
}
func (this *MIDIHDR) DwBufferLength() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[8]))
}
func (this *MIDIHDR) DwBytesRecorded() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[12]))
}
func (this *MIDIHDR) DwUser() *DWORD_PTR {
	return (*DWORD_PTR)(unsafe.Pointer(&this.storage[16]))
}
func (this *MIDIHDR) DwFlags() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[24]))
}
func (this *MIDIHDR) LpNext() **MIDIHDR {
	return (**MIDIHDR)(unsafe.Pointer(&this.storage[28]))
}
func (this *MIDIHDR) Reserved() *DWORD_PTR {
	return (*DWORD_PTR)(unsafe.Pointer(&this.storage[36]))
}
func (this *MIDIHDR) DwOffset() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[44]))
}
func (this *MIDIHDR) DwReserved() *[8]DWORD_PTR {
	return (*[8]DWORD_PTR)(unsafe.Pointer(&this.storage[48]))
}
