package win

//ref USHORT
//ref ULONG
//ref LARGE_INTEGER
//ref GUID
//ref EVENT_DESCRIPTOR
//ref ULONG64

type EVENT_HEADER struct {
	Size            USHORT
	HeaderType      USHORT
	Flags           USHORT
	EventProperty   USHORT
	ThreadId        ULONG
	ProcessId       ULONG
	TimeStamp       LARGE_INTEGER
	ProviderId      GUID
	EventDescriptor EVENT_DESCRIPTOR
	union1          ULONG64
	ActivityId      GUID
}

func (this *EVENT_HEADER) GetKernelTime() ULONG {
	return *(*ULONG)(unsafe.Pointer(&this.union1))
}

func (this *EVENT_HEADER) SetKernelTime(v ULONG) {
	*(*ULONG)(unsafe.Pointer(&this.union1)) = v
}

func (this *EVENT_HEADER) GetUserTime() ULONG {
	return *(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4)))
}

func (this *EVENT_HEADER) SetUserTime(v ULONG) {
	*(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(4))) = v
}

func (this *EVENT_HEADER) GetProcessorTime() ULONG64 {
	return this.union1
}

func (this *EVENT_HEADER) SetProcessorTime(v ULONG64) {
	this.union1 = v
}
