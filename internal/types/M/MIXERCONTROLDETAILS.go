package win

import "unsafe"

//ref DWORD
//ref LPVOID
//ref HWND

type MIXERCONTROLDETAILS struct {
	CbStruct    DWORD
	DwControlID DWORD
	CChannels   DWORD
	storage1    [4]byte
	pad1        [pad4for64_0for32]byte
	CbDetails   DWORD
	PaDetails   LPVOID
}

func (this *MIXERCONTROLDETAILS) HwndOwner() *HWND {
	return (*HWND)(unsafe.Pointer(&this.storage1[0]))
}

func (this *MIXERCONTROLDETAILS) CMultipleItems() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.storage1[0]))
}
