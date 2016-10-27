package win

//ref UCHAR
//ref DWORD
//ref LARGE_INTEGER
//ref ULONGLONG
func (this *MIB_UDP6ROW_OWNER_MODULE) DwFlags() *int32 {
	return &this.dwFlags
}
func (this *MIB_UDP6ROW_OWNER_MODULE) SpecificPortBind() int32 {
	return this.dwFlags & 0x1
}
