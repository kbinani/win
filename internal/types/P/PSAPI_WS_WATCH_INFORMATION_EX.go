package win

//ref PSAPI_WS_WATCH_INFORMATION
//ref ULONG_PTR
type PSAPI_WS_WATCH_INFORMATION_EX struct {
	BasicInfo        PSAPI_WS_WATCH_INFORMATION
	FaultingThreadId ULONG_PTR
	Flags            ULONG_PTR
}
