package win

import "unsafe"

//ref RAWINPUTHEADER
//ref RAWMOUSE
//ref RAWKEYBOARD
//ref RAWHID
type RAWINPUT struct {
	Header RAWINPUTHEADER
	Data   RAWINPUT_data
}
type RAWINPUT_data struct {
	storage [24]byte
}

func (this *RAWINPUT_data) Mouse() *RAWMOUSE {
	return (*RAWMOUSE)(unsafe.Pointer(this))
}

func (this *RAWINPUT_data) Keyboard() *RAWKEYBOARD {
	return (*RAWKEYBOARD)(unsafe.Pointer(this))
}

func (this *RAWINPUT_data) HID() *RAWHID {
	return (*RAWHID)(unsafe.Pointer(this))
}
