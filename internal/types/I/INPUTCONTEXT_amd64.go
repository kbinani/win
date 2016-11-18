package win

import "unsafe"

type INPUTCONTEXT struct {
	storage [352]byte
}

func (this *INPUTCONTEXT) HWnd() *HWND { // 8
	return (*HWND)(unsafe.Pointer(&this.storage[0]))
}
func (this *INPUTCONTEXT) FOpen() *BOOL { // 4
	return (*BOOL)(unsafe.Pointer(&this.storage[8]))
}
func (this *INPUTCONTEXT) PtStatusWndPos() *POINT { // 8
	return (*POINT)(unsafe.Pointer(&this.storage[12]))
}
func (this *INPUTCONTEXT) PtSoftKbdPos() *POINT { // 8
	return (*POINT)(unsafe.Pointer(&this.storage[20]))
}
func (this *INPUTCONTEXT) FdwConversion() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[28]))
}
func (this *INPUTCONTEXT) FdwSentence() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[32]))
}
func (this *INPUTCONTEXT) LfFont() *LOGFONT { // 92
	return (*LOGFONT)(unsafe.Pointer(&this.storage[36]))
}
func (this *INPUTCONTEXT) CfCompForm() *COMPOSITIONFORM { // 28
	return (*COMPOSITIONFORM)(unsafe.Pointer(&this.storage[128]))
}
func (this *INPUTCONTEXT) CfCandForm() *[4]CANDIDATEFORM { // 128
	return (*[4]CANDIDATEFORM)(unsafe.Pointer(&this.storage[156]))
}

//padding 4
func (this *INPUTCONTEXT) HCompStr() *HIMCC { // 8
	return (*HIMCC)(unsafe.Pointer(&this.storage[288]))
}
func (this *INPUTCONTEXT) HCandInfo() *HIMCC { // 8
	return (*HIMCC)(unsafe.Pointer(&this.storage[296]))
}
func (this *INPUTCONTEXT) HGuideLine() *HIMCC { // 8
	return (*HIMCC)(unsafe.Pointer(&this.storage[304]))
}
func (this *INPUTCONTEXT) HPrivate() *HIMCC { // 8
	return (*HIMCC)(unsafe.Pointer(&this.storage[312]))
}
func (this *INPUTCONTEXT) DwNumMsgBuf() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[320]))
}

//padding 4
func (this *INPUTCONTEXT) HMsgBuf() *HIMCC { // 8
	return (*HIMCC)(unsafe.Pointer(&this.storage[328]))
}
func (this *INPUTCONTEXT) FdwInit() *DWORD { // 4
	return (*DWORD)(unsafe.Pointer(&this.storage[336]))
}
func (this *INPUTCONTEXT) DwReserve() *[3]DWORD { // 12
	return (*[3]DWORD)(unsafe.Pointer(&this.storage[340]))
}
