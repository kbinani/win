package win

import "unsafe"

type CY struct {
	union1 [8]byte
}

func (this *CY) GetLo() uint32 {
	return *(*uint32)(unsafe.Pointer(&this.union1[0]))
}

func (this *CY) SetLo(v uint32) {
	*(*uint32)(unsafe.Pointer(&this.union1[0])) = v
}

func (this *CY) GetHi() int32 {
	return *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + uintptr(4)))
}

func (this *CY) SetHi(v int32) {
	*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1[0])) + uintptr(4))) = v
}

func (this *CY) GetInt64() int64 {
	return *(*int64)(unsafe.Pointer(&this.union1[0]))
}

func (this *CY) SetInt64(v int64) {
	*(*int64)(unsafe.Pointer(&this.union1[0])) = v
}
