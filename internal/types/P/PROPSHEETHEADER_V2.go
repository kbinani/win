package win

//ref DWORD
//ref HWND
//ref HINSTANCE
//ref HICON
//ref LPCWSTR
//ref UINT
//ref PROPSHEETPAGE
//ref HPROPSHEETPAGE
//ref PROPSHEETCALLBACK
//ref HBITMAP
//ref HPALETTE
type PROPSHEETHEADER_V2 struct {
	dwSize       DWORD
	dwFlags      DWORD
	hwndParent   HWND
	hInstance    HINSTANCE
	union1       uintptr
	PszCaption   LPCWSTR
	NPages       UINT
	union2       uintptr
	union3       uintptr
	PfnCallback  uintptr // PFNPROPSHEETCALLBACK
	union4       uintptr
	HplWatermark HPALETTE
	union5       uintptr
}

func (this *PROPSHEETHEADER_V2) HIcon() HICON {
	return HICON(this.union1)
}

func (this *PROPSHEETHEADER_V2) PszIcon() string {
	return stringFromUnicode16((*uint16)(unsafe.Pointer(this.union1)))
}

func (this *PROPSHEETHEADER_V2) NStartPage() UINT {
	return *(*UINT)(unsafe.Pointer(&this.union2))
}

func (this *PROPSHEETHEADER_V2) PStartPage() string {
	return stringFromUnicode16((*uint16)(unsafe.Pointer(this.union2)))
}

func (this *PROPSHEETHEADER_V2) Ppsp() /*const*/ **PROPSHEETPAGE {
	return (**PROPSHEETPAGE)(unsafe.Pointer(&this.union3))
}

func (this *PROPSHEETHEADER_V2) Phpage() *HPROPSHEETPAGE {
	return (*HPROPSHEETPAGE)(unsafe.Pointer(&this.union3))
}

func (this *PROPSHEETHEADER_V2) HbmWatermark() HBITMAP {
	return HBITMAP(this.union4)
}

func (this *PROPSHEETHEADER_V2) PszbmWatermark() string {
	return stringFromUnicode16((*uint16)(unsafe.Pointer(this.union4)))
}

func (this *PROPSHEETHEADER_V2) HbmHeader() HBITMAP {
	return HBITMAP(this.union5)
}

func (this *PROPSHEETHEADER_V2) PszbmHeader() string {
	return stringFromUnicode16((*uint16)(unsafe.Pointer(this.union5)))
}
