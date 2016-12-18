package win

//ref USHORT
//ref ULONG
//ref LARGE_INTEGER
//ref ULONGLONG
//ref ULONG64
//ref UCHAR

type EVENT_TRACE_HEADER struct {
	Size      USHORT
	union1    USHORT
	union2    ULONG
	ThreadId  ULONG
	ProcessId ULONG
	TimeStamp LARGE_INTEGER
	union3    GUID
	union4    ULONG64
}

func (this *EVENT_TRACE_HEADER) GetFieldTypeFlags() USHORT {
	return this.union1
}

func (this *EVENT_TRACE_HEADER) SetFieldTypeFlags(v USHORT) {
	this.union1 = v
}

func (this *EVENT_TRACE_HEADER) GetHeaderType() UCHAR {
	return *(*UCHAR)(unsafe.Pointer(&this.union1))
}

func (this *EVENT_TRACE_HEADER) SetHeaderType(v UCHAR) {
	*(*UCHAR)(unsafe.Pointer(&this.union1)) = v
}

func (this *EVENT_TRACE_HEADER) GetMarkerFlags() UCHAR {
	return *(*UCHAR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(1)))
}

func (this *EVENT_TRACE_HEADER) SetMarkerFlags(v UCHAR) {
	*(*UCHAR)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union1)) + uintptr(1))) = v
}

func (this *EVENT_TRACE_HEADER) GetVersion() ULONG {
	return this.union2
}

func (this *EVENT_TRACE_HEADER) SetVersion(v ULONG) {
	this.union2 = v
}

type EVENT_TRACE_HEADER_Class struct {
	Type    UCHAR
	Level   UCHAR
	Version USHORT
}

func (this *EVENT_TRACE_HEADER) GetClass() EVENT_TRACE_HEADER_Class {
	return *(*EVENT_TRACE_HEADER_Class)(unsafe.Pointer(&this.union2))
}

func (this *EVENT_TRACE_HEADER) SetClass(v EVENT_TRACE_HEADER_Class) {
	*(*EVENT_TRACE_HEADER_Class)(unsafe.Pointer(&this.union2)) = v
}

func (this *EVENT_TRACE_HEADER) GetGuid() GUID {
	return this.union3
}

func (this *EVENT_TRACE_HEADER) SetGuid(v GUID) {
	this.union3 = v
}

func (this *EVENT_TRACE_HEADER) GetGuidPtr() ULONGLONG {
	return *(*ULONGLONG)(unsafe.Pointer(&this.union3))
}

func (this *EVENT_TRACE_HEADER) SetGuidPtr(v ULONGLONG) {
	*(*ULONGLONG)(unsafe.Pointer(&this.union3)) = v
}

func (this *EVENT_TRACE_HEADER) GetKernelTime() ULONG {
	return *(*ULONG)(unsafe.Pointer(&this.union4))
}

func (this *EVENT_TRACE_HEADER) SetKernelTime(v ULONG) {
	*(*ULONG)(unsafe.Pointer(&this.union4)) = v
}

func (this *EVENT_TRACE_HEADER) GetUserTime() ULONG {
	return *(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union4)) + uintptr(4)))
}

func (this *EVENT_TRACE_HEADER) SetUserTime(v ULONG) {
	*(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union4)) + uintptr(4))) = v
}

func (this *EVENT_TRACE_HEADER) GetProcessorTime() ULONG64 {
	return this.union4
}

func (this *EVENT_TRACE_HEADER) SetProcessorTime(v ULONG64) {
	this.union4 = v
}

func (this *EVENT_TRACE_HEADER) GetClientContext() ULONG {
	return *(*ULONG)(unsafe.Pointer(&this.union4))
}

func (this *EVENT_TRACE_HEADER) SetClientContext(v ULONG) {
	*(*ULONG)(unsafe.Pointer(&this.union4)) = v
}

func (this *EVENT_TRACE_HEADER) GetFlags() ULONG {
	return *(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union4)) + uintptr(4)))
}

func (this *EVENT_TRACE_HEADER) SetFlags(v ULONG) {
	*(*ULONG)(unsafe.Pointer(uintptr(unsafe.Pointer(&this.union4)) + uintptr(4))) = v
}
