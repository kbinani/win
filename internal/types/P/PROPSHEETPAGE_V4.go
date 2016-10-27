package win

//import unsafe
//ref DWORD
//ref LPCWSTR
//ref UINT
//ref HANDLE
//ref PROPSHEETPAGE_RESOURCE
//ref HICON
//ref HANDLE
//ref HINSTANCE
//ref DLGPROC
//ref PSPCALLBACK
type PROPSHEETPAGE_V4 struct {
	DwSize            DWORD
	DwFlags           DWORD
	HInstance         HINSTANCE
	union1            uintptr
	union2            uintptr
	PszTitle          LPCWSTR
	PfnDlgProc        uintptr // DLGPROC
	LParam            uintptr
	PfnCallback       uintptr // PSPCALLBACK
	PcRefParent       *UINT
	PszHeaderTitle    LPCWSTR
	PszHeaderSubTitle LPCWSTR
	HActCtx           HANDLE
	union3            uintptr
}

func (this *PROPSHEETPAGE_V4) PszTemplate() *LPCWSTR {
	return (*LPCWSTR)(unsafe.Pointer(&this.union1))
}
func (this *PROPSHEETPAGE_V4) PResource() *PROPSHEETPAGE_RESOURCE {
	return (*PROPSHEETPAGE_RESOURCE)(unsafe.Pointer(&this.union1))
}
func (this *PROPSHEETPAGE_V4) HIcon() *HICON {
	return (*HICON)(unsafe.Pointer(&this.union2))
}
func (this *PROPSHEETPAGE_V4) PszIcon() *LPCWSTR {
	return (*LPCWSTR)(unsafe.Pointer(&this.union2))
}
func (this *PROPSHEETPAGE_V4) HbmHeader() *HBITMAP {
	return (*HBITMAP)(unsafe.Pointer(&this.union3))
}
func (this *PROPSHEETPAGE_V4) PszbmHeader() *LPCWSTR {
	return (*LPCWSTR)(unsafe.Pointer(&this.union3))
}
