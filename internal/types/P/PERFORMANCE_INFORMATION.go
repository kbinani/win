package win

//ref DWORD
//ref SIZE_T
type PERFORMANCE_INFORMATION struct {
	Cb                DWORD
	CommitTotal       SIZE_T
	CommitLimit       SIZE_T
	CommitPeak        SIZE_T
	PhysicalTotal     SIZE_T
	PhysicalAvailable SIZE_T
	SystemCache       SIZE_T
	KernelTotal       SIZE_T
	KernelPaged       SIZE_T
	KernelNonpaged    SIZE_T
	PageSize          SIZE_T
	HandleCount       DWORD
	ProcessCount      DWORD
	ThreadCount       DWORD
}
