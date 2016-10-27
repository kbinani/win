package win

//ref GUID
//ref ULONG
type PERF_COUNTERSET_INSTANCE struct {
	CounterSetGuid     GUID
	DwSize             ULONG
	InstanceId         ULONG
	InstanceNameOffset ULONG
	InstanceNameSize   ULONG
}
