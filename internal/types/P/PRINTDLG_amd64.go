package win

type PRINTDLG struct {
	storage [120]byte
}

func (this *PRINTDLG) LStructSize() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}
func (this *PRINTDLG) HwndOwner() *HWND { // 8
	return (*HWND)(unsafe.Pointer(&this.storage[4]))
}
func (this *PRINTDLG) HDevMode() *HGLOBAL { // 8
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[12]))
}
func (this *PRINTDLG) HDevNames() *HGLOBAL { // 8
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[20]))
}
func (this *PRINTDLG) HDC() *HDC { // 8
	return (*HDC)(unsafe.Pointer(&this.storage[28]))
}
func (this *PRINTDLG) Flags() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[36]))
}
func (this *PRINTDLG) NFromPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[40]))
}
func (this *PRINTDLG) NToPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[42]))
}
func (this *PRINTDLG) NMinPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[44]))
}
func (this *PRINTDLG) NMaxPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[46]))
}
func (this *PRINTDLG) NCopies() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[48]))
}
func (this *PRINTDLG) HInstance() *HINSTANCE { // 8
	return (*HINSTANCE)(unsafe.Pointer(&this.storage[56]))
}
func (this *PRINTDLG) LCustData() *LPARAM { // 8
	return (*LPARAM)(unsafe.Pointer(&this.storage[64]))
}

// LPPRINTHOOKPROC
func (this *PRINTDLG) LpfnPrintHook() *uintptr { // 8
	return (*uintptr)(unsafe.Pointer(&this.storage[72]))
}

// LPSETUPHOOKPROC
func (this *PRINTDLG) LpfnSetupHook() *uintptr { // 8
	return (*uintptr)(unsafe.Pointer(&this.storage[80]))
}
func (this *PRINTDLG) LpPrintTemplateName() *LPCWSTR { // 8
	return (*LPCWSTR)(unsafe.Pointer(&this.storage[88]))
}
func (this *PRINTDLG) LpSetupTemplateName() *LPCWSTR { // 8
	return (*LPCWSTR)(unsafe.Pointer(&this.storage[96]))
}
func (this *PRINTDLG) HPrintTemplate() *HGLOBAL { // 8
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[104]))
}
func (this *PRINTDLG) HSetupTemplate() *HGLOBAL { // 8
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[112]))
}
