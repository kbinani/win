package win

//ref DWORD
//ref LARGE_INTEGER
//ref ULONGLONG
func (this *MIB_UDPROW_OWNER_MODULE) DwFlags() *int32 {
	return &this.dwFlags
}
func (this *MIB_UDPROW_OWNER_MODULE) SpecificPortBind() int32 {
	return this.dwFlags & 0x1
}
