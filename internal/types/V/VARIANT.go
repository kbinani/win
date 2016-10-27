package win

//ref BSTR
//ref BYTE
//ref CHAR
//ref CY
//ref DATE
//ref DECIMAL
//ref DOUBLE
//ref FLOAT
//ref IDispatch
//ref INT
//ref IUnknown
//ref LONG
//ref LONGLONG
//ref PVOID
//ref SAFEARRAY
//ref SCODE
//ref SHORT
//ref UINT
//ref ULONG
//ref ULONGLONG
//ref USHORT
//ref VARIANT
//ref VARIANT_BOOL
//ref VARTYPE
//ref WORD
//ref IRecordInfo
//import unsafe
func (this *VARIANT) Vt() *VARTYPE {
	return (*VARTYPE)(unsafe.Pointer(&this.union1[0]))
}
func (this *VARIANT) WReserved1() *WORD {
	return (*WORD)(unsafe.Pointer(&this.union1[2]))
}
func (this *VARIANT) WReserved2() *WORD {
	return (*WORD)(unsafe.Pointer(&this.union1[4]))
}
func (this *VARIANT) WReserved3() *WORD {
	return (*WORD)(unsafe.Pointer(&this.union1[6]))
}
func (this *VARIANT) LlVal() *LONGLONG        { return (*LONGLONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) LVal() *LONG             { return (*LONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) BVal() *BYTE             { return (*BYTE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) IVal() *SHORT            { return (*SHORT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) FltVal() *FLOAT          { return (*FLOAT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) DblVal() *DOUBLE         { return (*DOUBLE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) BoolVal() *VARIANT_BOOL  { return (*VARIANT_BOOL)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Bool() *VARIANT_BOOL     { return (*VARIANT_BOOL)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Scode() *SCODE           { return (*SCODE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) CyVal() *CY              { return (*CY)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Date() *DATE             { return (*DATE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) BstrVal() *BSTR          { return (*BSTR)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PunkVal() *IUnknown      { return (*IUnknown)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PdispVal() *IDispatch    { return (*IDispatch)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Parray() *SAFEARRAY      { return (*SAFEARRAY)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PbVal() *BYTE            { return (*BYTE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PiVal() *SHORT           { return (*SHORT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PlVal() *LONG            { return (*LONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PllVal() *LONGLONG       { return (*LONGLONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PfltVal() *FLOAT         { return (*FLOAT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PdblVal() *DOUBLE        { return (*DOUBLE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PboolVal() *VARIANT_BOOL { return (*VARIANT_BOOL)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Pbool() *VARIANT_BOOL    { return (*VARIANT_BOOL)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Pscode() *SCODE          { return (*SCODE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PcyVal() *CY             { return (*CY)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Pdate() *DATE            { return (*DATE)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PbstrVal() *BSTR         { return (*BSTR)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PpunkVal() **IUnknown    { return (**IUnknown)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PpdispVal() **IDispatch  { return (**IDispatch)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Pparray() **SAFEARRAY    { return (**SAFEARRAY)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PvarVal() *VARIANT       { return (*VARIANT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) Byref() PVOID            { return (PVOID)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) CVal() CHAR              { return *(*CHAR)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) UiVal() USHORT           { return *(*USHORT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) UlVal() ULONG            { return *(*ULONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) UllVal() ULONGLONG       { return *(*ULONGLONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) IntVal() INT             { return *(*INT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) UintVal() UINT           { return *(*UINT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PdecVal() *DECIMAL       { return (*DECIMAL)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PcVal() *CHAR            { return (*CHAR)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PuiVal() *USHORT         { return (*USHORT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PulVal() *ULONG          { return (*ULONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PullVal() *ULONGLONG     { return (*ULONGLONG)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PintVal() *INT           { return (*INT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PuintVal() *UINT         { return (*UINT)(unsafe.Pointer(&this.union1[8])) }
func (this *VARIANT) PvRecord() PVOID {
	return (PVOID)(unsafe.Pointer(&this.union1[8]))
}
func (this *VARIANT) DecVal() DECIMAL {
	return *(*DECIMAL)(unsafe.Pointer(&this.union1[0]))
}
func unpackVARIANT(v VARIANT) []uintptr {
	size := int(unsafe.Sizeof(v))
	size += size % 4
	step := 4
	n := size / step
	ret := []uintptr{}
	ptr := uintptr(unsafe.Pointer(&v))
	for i := 0; i < n; i++ {
		ret = append(ret, *(*uintptr)(unsafe.Pointer(ptr + uintptr(step*i))))
	}
	return ret
}
