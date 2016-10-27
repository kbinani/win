package win

//import unsafe
type DLGTEMPLATE struct {
	storage [18]byte
}

func (this *DLGTEMPLATE) Style() *uint32 {
	return (*uint32)(unsafe.Pointer(&this.storage[0]))
}
func (this *DLGTEMPLATE) DwExtendedStyle() *uint32 {
	return (*uint32)(unsafe.Pointer(&this.storage[4]))
}
func (this *DLGTEMPLATE) Cdit() *uint16 {
	return (*uint16)(unsafe.Pointer(&this.storage[8]))
}
func (this *DLGTEMPLATE) X() *int16 {
	return (*int16)(unsafe.Pointer(&this.storage[10]))
}
func (this *DLGTEMPLATE) Y() *int16 {
	return (*int16)(unsafe.Pointer(&this.storage[12]))
}
func (this *DLGTEMPLATE) Cx() *int16 {
	return (*int16)(unsafe.Pointer(&this.storage[14]))
}
func (this *DLGTEMPLATE) Cy() *int16 {
	return (*int16)(unsafe.Pointer(&this.storage[16]))
}
