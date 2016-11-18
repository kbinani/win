package win

import "unsafe"

type VARIANT struct {
	union1 [24]byte
}

func (this *VARIANT) PRecInfo() *IRecordInfo {
	return (*IRecordInfo)(unsafe.Pointer(&this.union1[16]))
}
