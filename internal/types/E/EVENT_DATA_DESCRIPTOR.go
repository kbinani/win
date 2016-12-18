package win

//ref ULONGLONG
//ref ULONG
//ref UCHAR
//ref USHORT

type EVENT_DATA_DESCRIPTOR struct {
	Ptr    ULONGLONG
	Size   ULONG
	union1 ULONG
}

func (this *EVENT_DATA_DESCRIPTOR) GetReserved() ULONG {
	return this.union1
}

func (this *EVENT_DATA_DESCRIPTOR) SetReserved(v ULONG) {
	this.union1 = v
}

func (this *EVENT_DATA_DESCRIPTOR) GetType() UCHAR {
	return *(*UCHAR)(unsafe.Pointer(&this.union1))
}

func (this *EVENT_DATA_DESCRIPTOR) SetType(v UCHAR) {
	*(*UCHAR)(unsafe.Pointer(&this.union1)) = v
}

func (this *EVENT_DATA_DESCRIPTOR) GetReserved1() UCHAR {
	return *(*UCHAR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(1)))
}

func (this *EVENT_DATA_DESCRIPTOR) SetReserved1(v UCHAR) {
	*(*UCHAR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(1))) = v
}

func (this *EVENT_DATA_DESCRIPTOR) GetReserved2() USHORT {
	return *(*USHORT)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(2)))
}

func (this *EVENT_DATA_DESCRIPTOR) SetReserved2(v USHORT) {
	*(*USHORT)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(2))) = v
}
