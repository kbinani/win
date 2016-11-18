package win

import "unsafe"

type PRINTDLG struct {
	storage [66]byte
}

func (this *PRINTDLG) LStructSize() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}
func (this *PRINTDLG) HwndOwner() *HWND { // 4
	return (*HWND)(unsafe.Pointer(&this.storage[4]))
}
func (this *PRINTDLG) HDevMode() *HGLOBAL { // 4
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[8]))
}
func (this *PRINTDLG) HDevNames() *HGLOBAL { // 4
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[12]))
}
func (this *PRINTDLG) HDC() *HDC { // 4
	return (*HDC)(unsafe.Pointer(&this.storage[16]))
}
func (this *PRINTDLG) Flags() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[20]))
}
func (this *PRINTDLG) NFromPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[24]))
}
func (this *PRINTDLG) NToPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[26]))
}
func (this *PRINTDLG) NMinPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[28]))
}
func (this *PRINTDLG) NMaxPage() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[30]))
}
func (this *PRINTDLG) NCopies() *WORD { // 2
	return (*WORD)(unsafe.Pointer(&this.storage[32]))
}
func (this *PRINTDLG) HInstance() *HINSTANCE { // 4
	return (*HINSTANCE)(unsafe.Pointer(&this.storage[34]))
}
func (this *PRINTDLG) LCustData() *LPARAM { // 4
	return (*LPARAM)(unsafe.Pointer(&this.storage[38]))
}

// LPPRINTHOOKPROC
func (this *PRINTDLG) LpfnPrintHook() *uintptr { // 4
	return (*uintptr)(unsafe.Pointer(&this.storage[42]))
}

// LPSETUPHOOKPROC
func (this *PRINTDLG) LpfnSetupHook() *uintptr { // 4
	return (*uintptr)(unsafe.Pointer(&this.storage[46]))
}
func (this *PRINTDLG) LpPrintTemplateName() *LPCWSTR { // 4
	return (*LPCWSTR)(unsafe.Pointer(&this.storage[50]))
}
func (this *PRINTDLG) LpSetupTemplateName() *LPCWSTR { // 4
	return (*LPCWSTR)(unsafe.Pointer(&this.storage[54]))
}
func (this *PRINTDLG) HPrintTemplate() *HGLOBAL { // 4
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[58]))
}
func (this *PRINTDLG) HSetupTemplate() *HGLOBAL { // 4
	return (*HGLOBAL)(unsafe.Pointer(&this.storage[62]))
}
