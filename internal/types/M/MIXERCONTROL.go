package win

import "unsafe"

//ref DWORD
//ref WCHAR

type MIXERCONTROL struct {
	cbStruct       DWORD
	dwControlID    DWORD
	dwControlType  DWORD
	fdwControl     DWORD
	cMultipleItems DWORD
	szShortName    [MIXER_SHORT_NAME_CHARS]WCHAR
	szName         [MIXER_LONG_NAME_CHARS]WCHAR
	Bounds         MIXERLINECONTROL_Bounds
	Metrics        MIXERLINECONTROL_Metrics
}

type MIXERLINECONTROL_Bounds struct {
	storage [24]byte
}

func (this *MIXERLINECONTROL_Bounds) DwReserved() *[6]DWORD {
	return (*[6]DWORD)(unsafe.Pointer(&this.storage[0]))
}

func (this *MIXERLINECONTROL_Bounds) LMinimum() *LONG {
	return (*LONG)(unsafe.Pointer(&this.storage[0]))
}

func (this *MIXERLINECONTROL_Bounds) LMaximum() *LONG {
	return (*LONG)(unsafe.Pointer(&this.storage[4]))
}

func (this *MIXERLINECONTROL_Bounds) DwMinimum() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}

func (this *MIXERLINECONTROL_Bounds) DwMaximum() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[4]))
}

type MIXERLINECONTROL_Metrics struct {
	storage [24]byte
}

func (this *MIXERLINECONTROL_Metrics) CSteps() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}

func (this *MIXERLINECONTROL_Metrics) CbCustomData() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage[0]))
}

func (this *MIXERLINECONTROL_Metrics) DwReserved() *[6]DWORD {
	return (*[6]DWORD)(unsafe.Pointer(&this.storage[0]))
}
