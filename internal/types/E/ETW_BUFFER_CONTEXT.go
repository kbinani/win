package win

//ref USHORT
//ref UCHAR

type ETW_BUFFER_CONTEXT struct {
	union1   USHORT
	LoggerId USHORT
}

func (this *ETW_BUFFER_CONTEXT) GetProcessorNumber() UCHAR {
	return *(*UCHAR)(unsafe.Pointer(&this.union1))
}

func (this *ETW_BUFFER_CONTEXT) SetProcessorNumber(v UCHAR) {
	*(*UCHAR)(unsafe.Pointer(&this.union1)) = v
}

func (this *ETW_BUFFER_CONTEXT) GetAlignment() UCHAR {
	return *(*UCHAR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(1)))
}

func (this *ETW_BUFFER_CONTEXT) SetAlignment(v UCHAR) {
	*(*UCHAR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(1))) = v
}

func (this *ETW_BUFFER_CONTEXT) GetProcessorIndex() USHORT {
	return this.union1
}

func (this *ETW_BUFFER_CONTEXT) SetProcessorIndex(v USHORT) {
	this.union1 = v
}
