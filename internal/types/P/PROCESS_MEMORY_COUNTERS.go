package win

//ref DWORD
//ref SIZE_T
type PROCESS_MEMORY_COUNTERS struct {
	Cb                         DWORD
	PageFaultCount             DWORD
	PeakWorkingSetSize         SIZE_T
	WorkingSetSize             SIZE_T
	QuotaPeakPagedPoolUsage    SIZE_T
	QuotaPagedPoolUsage        SIZE_T
	QuotaPeakNonPagedPoolUsage SIZE_T
	QuotaNonPagedPoolUsage     SIZE_T
	PagefileUsage              SIZE_T
	PeakPagefileUsage          SIZE_T
}
