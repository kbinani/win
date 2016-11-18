package win

import "unsafe"

type INPUT struct {
	Type uint32
	data [24]byte
}

func (this *INPUT) Mi() *MOUSEINPUT {
	return (*MOUSEINPUT)(unsafe.Pointer(&this.data[0]))
}
func (this *INPUT) Ki() *KEYBDINPUT {
	return (*KEYBDINPUT)(unsafe.Pointer(&this.data[0]))
}
func (this *INPUT) Hi() *HARDWAREINPUT {
	return (*HARDWAREINPUT)(unsafe.Pointer(&this.data[0]))
}
