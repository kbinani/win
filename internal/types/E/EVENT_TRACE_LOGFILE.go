package win

//ref LPWSTR
//ref LONGLONG
//ref ULONG
//ref EVENT_TRACE
//ref TRACE_LOGFILE_HEADER
//ref PEVENT_TRACE_BUFFER_CALLBACK
//ref PVOID
//ref PEVENT_CALLBACK
//ref PEVENT_RECORD_CALLBACK

func (this *EVENT_TRACE_LOGFILE) GetLogFileMode() ULONG {
	return this.union1
}

func (this *EVENT_TRACE_LOGFILE) SetLogFileMode(v ULONG) {
	this.union1 = v
}

func (this *EVENT_TRACE_LOGFILE) GetProcessTraceMode() ULONG {
	return this.union1
}

func (this *EVENT_TRACE_LOGFILE) SetProcessTraceMode(v ULONG) {
	this.union1 = v
}

func (this *EVENT_TRACE_LOGFILE) GetEventCallback() uintptr {
	return this.union2
}

func (this *EVENT_TRACE_LOGFILE) SetEventCallback(v uintptr) {
	this.union2 = v
}

func (this *EVENT_TRACE_LOGFILE) GetEventRecordCallback() uintptr {
	return this.union2
}

func (this *EVENT_TRACE_LOGFILE) SetEventRecordCallback(v uintptr) {
	this.union2 = v
}
