package win

import "unsafe"

//ref LPMIXERCONTROL
//ref DWORD

func (this *MIXERLINECONTROLS) DwControlID() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.union1))
}

func (this *MIXERLINECONTROLS) DwControlType() *DWORD {
	return (*DWORD)(unsafe.Pointer(&this.union1))
}

func (this *MIXERLINECONTROLS) Pamxctrl() *LPMIXERCONTROL {
	return (*LPMIXERCONTROL)(unsafe.Pointer(&this.storage1[0]))
}
