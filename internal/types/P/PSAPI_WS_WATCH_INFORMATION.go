package win

//ref LPVOID
type PSAPI_WS_WATCH_INFORMATION struct {
	FaultingPc LPVOID
	FaultingVa LPVOID
}
