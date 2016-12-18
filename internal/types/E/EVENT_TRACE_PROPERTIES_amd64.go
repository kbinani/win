package win

type EVENT_TRACE_PROPERTIES struct {
	Wnode               WNODE_HEADER
	BufferSize          ULONG
	MinimumBuffers      ULONG
	MaximumBuffers      ULONG
	MaximumFileSize     ULONG
	LogFileMode         ULONG
	FlushTimer          ULONG
	EnableFlags         ULONG
	union1              LONG
	NumberOfBuffers     ULONG
	FreeBuffers         ULONG
	EventsLost          ULONG
	BuffersWritten      ULONG
	LogBuffersLost      ULONG
	RealTimeBuffersLost ULONG
	LoggerThreadId      HANDLE
	LogFileNameOffset   ULONG
	LoggerNameOffset    ULONG
}
