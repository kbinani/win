package win

import "unsafe"

//ref DWORD
//ref DWORD_PTR
//ref WCHAR
//ref MMVERSION
//ref WORD

type MIXERLINE struct {
	CbStruct      DWORD
	DwDestination DWORD
	DwSource      DWORD
	DwLineID      DWORD
	FdwLine       DWORD
	//	DwUser          DWORD_PTR
	storage1        [4]byte
	storage2        [pad4for64_0for32]byte
	DwComponentType DWORD
	CChannels       DWORD
	CConnections    DWORD
	CControls       DWORD
	SzShortName     [MIXER_SHORT_NAME_CHARS]WCHAR
	SzName          [MIXER_LONG_NAME_CHARS]WCHAR
	Target          MIXERLINE_Target
}

type MIXERLINE_Target struct {
	DwType         DWORD
	DwDeviceID     DWORD
	WMid           WORD
	WPid           WORD
	VDriverVersion MMVERSION
	SzPname        [MAXPNAMELEN]WCHAR
}

func (this *MIXERLINE) DwUser() *DWORD_PTR {
	return (*DWORD_PTR)(unsafe.Pointer(&this.storage1[0]))
}
