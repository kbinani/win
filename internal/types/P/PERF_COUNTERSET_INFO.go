package win

//ref GUID
//ref ULONG

type PERF_COUNTERSET_INFO struct {
	CounterSetGuid GUID
	ProviderGuid   GUID
	NumCounters    ULONG
	InstanceType   ULONG
}
