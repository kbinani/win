package win

//ref FLOATL
//ref LONG
import "unsafe"

type FLOAT_LONG struct {
	storage [4]byte
}

func (this *FLOAT_LONG) E() *FLOATL {
	return (*FLOATL)(unsafe.Pointer(&this.storage[0]))
}
func (this *FLOAT_LONG) L() *LONG {
	return (*LONG)(unsafe.Pointer(&this.storage[0]))
}
