package win

//import unsafe
type VARIANT struct {
	union1 [16]byte
}

func (this *VARIANT) PRecInfo() *IRecordInfo {
	return (*IRecordInfo)(unsafe.Pointer(&this.union1[12]))
}
