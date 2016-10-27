package win

//ref USHORT
//ref BYTE
//ref USHORT
//ref ULONG
//ref ULONGLONG
//import unsafe
type DECIMAL struct {
	WReserved USHORT
	union1    [2]byte
	Hi32      ULONG
	union2    [8]byte
}

func (this *DECIMAL) Scale() *BYTE {
	return (*BYTE)(unsafe.Pointer(&this.union1[0]))
}
func (this *DECIMAL) Sign() *BYTE {
	return (*BYTE)(unsafe.Pointer(&this.union1[1]))
}
func (this *DECIMAL) Signscale() *USHORT {
	return (*USHORT)(unsafe.Pointer(&this.union1[0]))
}
func (this *DECIMAL) Lo32() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union2[0]))
}
func (this *DECIMAL) Mid32() *ULONG {
	return (*ULONG)(unsafe.Pointer(&this.union2[4]))
}
func (this *DECIMAL) Lo64() *ULONGLONG {
	return (*ULONGLONG)(unsafe.Pointer(&this.union2[0]))
}
