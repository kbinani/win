package win

type EVENT_TRACE_LOGFILE struct {
	LogFileName    LPWSTR
	LoggerName     LPWSTR
	CurrentTime    LONGLONG
	BuffersRead    ULONG
	union1         ULONG
	CurrentEvent   EVENT_TRACE
	LogfileHeader  TRACE_LOGFILE_HEADER
	BufferCallback uintptr
	BufferSize     ULONG
	Filled         ULONG
	EventsLost     ULONG
	union2         uintptr
	IsKernelTrace  ULONG
	Context        PVOID
}
