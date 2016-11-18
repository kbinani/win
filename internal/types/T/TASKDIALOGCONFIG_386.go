package win

import "unsafe"

type TASKDIALOGCONFIG struct {
	storage [96]byte
}

func (this *TASKDIALOGCONFIG) CbSize() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[0]))
}
func (this *TASKDIALOGCONFIG) HwndParent() *HWND {
	return (*HWND)(unsafe.Pointer(&this.storage[4]))
}
func (this *TASKDIALOGCONFIG) HInstance() *HINSTANCE {
	return (*HINSTANCE)(unsafe.Pointer(&this.storage[8]))
}
func (this *TASKDIALOGCONFIG) DwFlags() *TASKDIALOG_FLAGS {
	return (*TASKDIALOG_FLAGS)(unsafe.Pointer(&this.storage[12]))
}
func (this *TASKDIALOGCONFIG) DwCommonButtons() *TASKDIALOG_COMMON_BUTTON_FLAGS {
	return (*TASKDIALOG_COMMON_BUTTON_FLAGS)(unsafe.Pointer(&this.storage[16]))
}
func (this *TASKDIALOGCONFIG) PszWindowTitle() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[20]))
}
func (this *TASKDIALOGCONFIG) HMainIcon() *HICON {
	return (*HICON)(unsafe.Pointer(&this.storage[24]))
}
func (this *TASKDIALOGCONFIG) PszMainIcon() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[24]))
}
func (this *TASKDIALOGCONFIG) PszMainInstruction() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[28]))
}
func (this *TASKDIALOGCONFIG) PszContent() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[32]))
}
func (this *TASKDIALOGCONFIG) CButtons() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[36]))
}
func (this *TASKDIALOGCONFIG) PButtons() **TASKDIALOG_BUTTON {
	return (**TASKDIALOG_BUTTON)(unsafe.Pointer(&this.storage[40]))
}
func (this *TASKDIALOGCONFIG) NDefaultButton() *int32 {
	return (*int32)(unsafe.Pointer(&this.storage[44]))
}
func (this *TASKDIALOGCONFIG) CRadioButtons() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[48]))
}
func (this *TASKDIALOGCONFIG) PRadioButtons() **TASKDIALOG_BUTTON {
	return (**TASKDIALOG_BUTTON)(unsafe.Pointer(&this.storage[52]))
}
func (this *TASKDIALOGCONFIG) NDefaultRadioButton() *int32 {
	return (*int32)(unsafe.Pointer(&this.storage[56]))
}
func (this *TASKDIALOGCONFIG) PszVerificationText() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[60]))
}
func (this *TASKDIALOGCONFIG) PszExpandedInformation() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[64]))
}
func (this *TASKDIALOGCONFIG) PszExpandedControlText() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[68]))
}
func (this *TASKDIALOGCONFIG) PszCollapsedControlText() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[72]))
}
func (this *TASKDIALOGCONFIG) HFooterIcon() *HICON {
	return (*HICON)(unsafe.Pointer(&this.storage[76]))
}
func (this *TASKDIALOGCONFIG) PszFooterIcon() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[76]))
}
func (this *TASKDIALOGCONFIG) PszFooter() *PCWSTR {
	return (*PCWSTR)(unsafe.Pointer(&this.storage[80]))
}
func (this *TASKDIALOGCONFIG) PfCallback() *uintptr {
	return (*uintptr)(unsafe.Pointer(&this.storage[84]))
}
func (this *TASKDIALOGCONFIG) LpCallbackData() *LONG_PTR {
	return (*LONG_PTR)(unsafe.Pointer(&this.storage[88]))
}
func (this *TASKDIALOGCONFIG) CxWidth() *UINT {
	return (*UINT)(unsafe.Pointer(&this.storage[92]))
}
