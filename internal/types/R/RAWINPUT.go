package win

//import unsafe
//ref RAWINPUTHEADER
//ref RAWMOUSE
//ref RAWKEYBOARD
//ref RAWHID
type RAWINPUT struct {
	Header RAWINPUTHEADER
	data   [24]byte
}

func (r *RAWINPUT) Mouse() *RAWMOUSE {
	return (*RAWMOUSE)(unsafe.Pointer(&r.data[0]))
}

func (r *RAWINPUT) Keyboard() *RAWKEYBOARD {
	return (*RAWKEYBOARD)(unsafe.Pointer(&r.data[0]))
}

func (r *RAWINPUT) HID() *RAWHID {
	return (*RAWHID)(unsafe.Pointer(&r.data[0]))
}
