package win

import "unsafe"

//ref MOUSEINPUT
//ref KEYBDINPUT
//ref HARDWAREINPUT

type INPUT struct {
	Type     uint32
	padding1 [pad4for64_0for32]byte
	data     [8 * pad4for64_3for32]byte
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
