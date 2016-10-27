package win

//ref DWORD
//ref SIZE_T
type ENUM_PAGE_FILE_INFORMATION struct {
	Cb         DWORD
	Reserved   DWORD
	TotalSize  SIZE_T
	TotalInUse SIZE_T
	PeakUsage  SIZE_T
}
