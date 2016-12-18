package win

//ref DWORD
//ref PERFLIBREQUEST
//ref PERF_MEM_ALLOC
//ref PERF_MEM_FREE
//ref LPVOID

type PERF_PROVIDER_CONTEXT struct {
	ContextSize     DWORD
	Reserved        DWORD
	ControlCallback uintptr // PERFLIBREQUEST
	MemAllocRoutine uintptr // PERF_MEM_ALLOC
	MemFreeRoutine  uintptr // PERF_MEM_FREE
	PMemContext     LPVOID
}
