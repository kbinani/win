package win

//ref ULONG
//ref GUID
//ref PEVENT_FILTER_DESCRIPTOR

type ENABLE_TRACE_PARAMETERS struct {
	Version          ULONG
	EnableProperty   ULONG
	ControlFlags     ULONG
	SourceId         GUID
	EnableFilterDesc PEVENT_FILTER_DESCRIPTOR
	FilterDescCount  ULONG
}
