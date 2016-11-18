package win

//ref UINT
//ref LPWSTR
import "unsafe"

func (this *STRRET) POleStr() *LPWSTR {
	return (*LPWSTR)(unsafe.Pointer(&this.cStr[0]))
}
func (this *STRRET) UOffset() *UINT {
	return (*UINT)(unsafe.Pointer(&this.cStr[0]))
}
func (this *STRRET) CStr() **byte {
	return (**byte)(unsafe.Pointer(&this.cStr[0]))
}
