package win

import "unsafe"

//ref VARTYPE
//ref WORD
//ref CHAR
//ref UCHAR
//ref SHORT
//ref USHORT
//ref LONG
//ref ULONG
//ref INT
//ref UINT
//ref LARGE_INTEGER
//ref ULARGE_INTEGER
//ref FLOAT
//ref DOUBLE
//ref VARIANT_BOOL
//ref SCODE
//ref CY
//ref DATE
//ref FILETIME
//ref CLSID
//ref CLIPDATA
//ref BSTR
//ref BSTRBLOB
//ref BLOB
//ref LPSTR
//ref LPWSTR
//ref IUnknown
//ref IDispatch
//ref IStream
//ref IStorage
//ref LPVERSIONEDSTREAM
//ref LPSAFEARRAY
//ref CAC
//ref CAUB
//ref CAI
//ref CAUI
//ref CAL
//ref CAUL
//ref CAH
//ref CAUH
//ref CAFLT
//ref CADBL
//ref CABOOL
//ref CASCODE
//ref CACY
//ref CADATE
//ref CAFILETIME
//ref CACLSID
//ref CACLIPDATA
//ref CABSTR
//ref CALPSTR
//ref CALPWSTR
//ref CAPROPVARIANT
//ref DECIMAL
//ref CABSTRBLOB

type PROPVARIANT struct {
	Vt         VARTYPE
	WReserved1 WORD
	WReserved2 WORD
	WReserved3 WORD
	storage1   [ptrsize * 2]byte
}

func (this *PROPVARIANT) CVal() *CHAR    { return (*CHAR)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) BVal() *UCHAR   { return (*UCHAR)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) IVal() *SHORT   { return (*SHORT)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) UiVal() *USHORT { return (*USHORT)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) LVal() *LONG    { return (*LONG)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) UlVal() *ULONG  { return (*ULONG)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) IntVal() *INT   { return (*INT)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) UintVal() *UINT { return (*UINT)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) HVal() *LARGE_INTEGER {
	return (*LARGE_INTEGER)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) UhVal() *ULARGE_INTEGER {
	return (*ULARGE_INTEGER)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) FltVal() *FLOAT  { return (*FLOAT)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) DblVal() *DOUBLE { return (*DOUBLE)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) BoolVal() *VARIANT_BOOL {
	return (*VARIANT_BOOL)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) Scode() *SCODE       { return (*SCODE)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) CyVal() *CY          { return (*CY)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Date() *DATE         { return (*DATE)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Filetime() *FILETIME { return (*FILETIME)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Puuid() **CLSID      { return (**CLSID)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Pclipdata() **CLIPDATA {
	return (**CLIPDATA)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) BstrVal() *BSTR { return (*BSTR)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) BstrblobVal() *BSTRBLOB {
	return (*BSTRBLOB)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) Blob() *BLOB         { return (*BLOB)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PszVal() *LPSTR      { return (*LPSTR)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PwszVal() *LPWSTR    { return (*LPWSTR)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PunkVal() **IUnknown { return (**IUnknown)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PdispVal() **IDispatch {
	return (**IDispatch)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) PStream() **IStream   { return (**IStream)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PStorage() **IStorage { return (**IStorage)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PVersionedStream() *LPVERSIONEDSTREAM {
	return (*LPVERSIONEDSTREAM)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) Parray() *LPSAFEARRAY {
	return (*LPSAFEARRAY)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) Cac() *CAC         { return (*CAC)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Caub() *CAUB       { return (*CAUB)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Cai() *CAI         { return (*CAI)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Caui() *CAUI       { return (*CAUI)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Cal() *CAL         { return (*CAL)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Caul() *CAUL       { return (*CAUL)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Cah() *CAH         { return (*CAH)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Cauh() *CAUH       { return (*CAUH)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Caflt() *CAFLT     { return (*CAFLT)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Cadbl() *CADBL     { return (*CADBL)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Cabool() *CABOOL   { return (*CABOOL)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Cascode() *CASCODE { return (*CASCODE)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Cacy() *CACY       { return (*CACY)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Cadate() *CADATE   { return (*CADATE)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Cafiletime() *CAFILETIME {
	return (*CAFILETIME)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) Cauuid() *CACLSID { return (*CACLSID)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Caclipdata() *CACLIPDATA {
	return (*CACLIPDATA)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) Cabstr() *CABSTR { return (*CABSTR)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Cabstrblob() *CABSTRBLOB {
	return (*CABSTRBLOB)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) Calpstr() *CALPSTR   { return (*CALPSTR)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Calpwstr() *CALPWSTR { return (*CALPWSTR)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Capropvar() *CAPROPVARIANT {
	return (*CAPROPVARIANT)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) PcVal() **CHAR     { return (**CHAR)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PbVal() **UCHAR    { return (**UCHAR)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PiVal() **SHORT    { return (**SHORT)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PuiVal() **USHORT  { return (**USHORT)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PlVal() **LONG     { return (**LONG)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PulVal() **ULONG   { return (**ULONG)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PintVal() **INT    { return (**INT)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PuintVal() **UINT  { return (**UINT)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PfltVal() **FLOAT  { return (**FLOAT)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PdblVal() **DOUBLE { return (**DOUBLE)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PboolVal() **VARIANT_BOOL {
	return (**VARIANT_BOOL)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) PdecVal() **DECIMAL { return (**DECIMAL)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Pscode() **SCODE    { return (**SCODE)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PcyVal() **CY       { return (**CY)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) Pdate() **DATE      { return (**DATE)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PbstrVal() **BSTR   { return (**BSTR)(unsafe.Pointer(&this.storage1[0])) }
func (this *PROPVARIANT) PpunkVal() ***IUnknown {
	return (***IUnknown)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) PpdispVal() ***IDispatch {
	return (***IDispatch)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) Pparray() **LPSAFEARRAY {
	return (**LPSAFEARRAY)(unsafe.Pointer(&this.storage1[0]))
}
func (this *PROPVARIANT) PvarVal() **PROPVARIANT {
	return (**PROPVARIANT)(unsafe.Pointer(&this.storage1[0]))
}
