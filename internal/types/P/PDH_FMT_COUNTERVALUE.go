package win

//ref DWORD
//ref LONG
//ref LONGLONG
//ref LPCSTR
//ref LPCWSTR
//import unsafe
type PDH_FMT_COUNTERVALUE struct {
	storage [16]byte
}

func (this *PDH_FMT_COUNTERVALUE) CStatus() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}
func (this *PDH_FMT_COUNTERVALUE) LongValue() *LONG {
	return (*LONG)(unsafe.Pointer(&this.storage[4]))
}
func (this *PDH_FMT_COUNTERVALUE) DoubleValue() *float64 {
	return (*float64)(unsafe.Pointer(&this.storage[4]))
}
func (this *PDH_FMT_COUNTERVALUE) LargeValue() *LONGLONG {
	return (*LONGLONG)(unsafe.Pointer(&this.storage[4]))
}
func (this *PDH_FMT_COUNTERVALUE) AnsiStringValue() *LPCSTR {
	return (*LPCSTR)(unsafe.Pointer(&this.storage[4]))
}
func (this *PDH_FMT_COUNTERVALUE) WideStringValue() *LPCWSTR {
	return (*LPCWSTR)(unsafe.Pointer(&this.storage[4]))
}
