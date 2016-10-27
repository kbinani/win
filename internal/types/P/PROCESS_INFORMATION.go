package win

//ref HANDLE
//ref DWORD
type PROCESS_INFORMATION struct {
	HProcess    HANDLE
	HThread     HANDLE
	DwProcessId DWORD
	DwThreadId  DWORD
}
