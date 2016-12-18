package win

//ref ETW_BUFFER_CONTEXT
//ref ULONG
//ref GUID
//ref PVOID

type EVENT_TRACE struct {
	Header           EVENT_TRACE_HEADER
	InstanceId       ULONG
	ParentInstanceId ULONG
	ParentGuid       GUID
	MofData          PVOID
	MofLength        ULONG
	union1           [pad4for64_8for32]byte
}

func (this *EVENT_TRACE) GetClientContext() ULONG {
	return *(*ULONG)(unsafe.Pointer(&this.union1[0]))
}

func (this *EVENT_TRACE) SetClientContext(v ULONG) {
	*(*ULONG)(unsafe.Pointer(&this.union1[0])) = v
}

func (this *EVENT_TRACE) GetBufferContext() ETW_BUFFER_CONTEXT {
	return *(*ETW_BUFFER_CONTEXT)(unsafe.Pointer(&this.union1[0]))
}

func (this *EVENT_TRACE) SetBufferContext(v ETW_BUFFER_CONTEXT) {
	*(*ETW_BUFFER_CONTEXT)(unsafe.Pointer(&this.union1[0])) = v
}
