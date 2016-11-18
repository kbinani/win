package win

import "unsafe"

type CY struct {
	union1 [8]byte
}

func (this CY) Lo() uint32 {
	return *(*uint32)(unsafe.Pointer(&this.union1[0]))
}
func (this CY) Hi() int32 {
	return *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + uintptr(4)))
}
func (this CY) Int64() int64 {
	return *(*int64)(unsafe.Pointer(&this.union1[0]))
}
