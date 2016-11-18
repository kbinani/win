package win

import "unsafe"

type INPUTCONTEXT struct {
	storage [320]byte
}

func (this *INPUTCONTEXT) HWnd() *HWND { // 4
	return (*HWND)(unsafe.Pointer(&this.storage[0]))
}
func (this *INPUTCONTEXT) FOpen() *BOOL { // 4
	return (*BOOL)(unsafe.Pointer(&this.storage[4]))
}
func (this *INPUTCONTEXT) PtStatusWndPos() *POINT { // 8
	return (*POINT)(unsafe.Pointer(&this.storage[8]))
}
func (this *INPUTCONTEXT) PtSoftKbdPos() *POINT { // 8
	return (*POINT)(unsafe.Pointer(&this.storage[16]))
}
func (this *INPUTCONTEXT) FdwConversion() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[24]))
}
func (this *INPUTCONTEXT) FdwSentence() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[28]))
}
func (this *INPUTCONTEXT) LfFont() *LOGFONT { // 92
	return (*LOGFONT)(unsafe.Pointer(&this.storage[32]))
}
func (this *INPUTCONTEXT) CfCompForm() *COMPOSITIONFORM { // 28
	return (*COMPOSITIONFORM)(unsafe.Pointer(&this.storage[124]))
}
func (this *INPUTCONTEXT) CfCandForm() *[4]CANDIDATEFORM { // 128
	return (*[4]CANDIDATEFORM)(unsafe.Pointer(&this.storage[152]))
}
func (this *INPUTCONTEXT) HCompStr() *HIMCC { // 4
	return (*HIMCC)(unsafe.Pointer(&this.storage[280]))
}
func (this *INPUTCONTEXT) HCandInfo() *HIMCC { // 4
	return (*HIMCC)(unsafe.Pointer(&this.storage[284]))
}
func (this *INPUTCONTEXT) HGuideLine() *HIMCC { // 4
	return (*HIMCC)(unsafe.Pointer(&this.storage[288]))
}
func (this *INPUTCONTEXT) HPrivate() *HIMCC { // 4
	return (*HIMCC)(unsafe.Pointer(&this.storage[292]))
}
func (this *INPUTCONTEXT) DwNumMsgBuf() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[296]))
}
func (this *INPUTCONTEXT) HMsgBuf() *HIMCC { // 4
	return (*HIMCC)(unsafe.Pointer(&this.storage[300]))
}
func (this *INPUTCONTEXT) FdwInit() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[304]))
}
func (this *INPUTCONTEXT) DwReserve() *[3]DWORD { // 12
	return (*[3]DWORD)(unsafe.Pointer(&this.storage[308]))
}
