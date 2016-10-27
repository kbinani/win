package win

//ref DWORD
//ref LPWSTR
//ref WORD
//ref LPBYTE
//ref HANDLE
type STARTUPINFO struct {
	Cb              DWORD
	LpReserved      LPWSTR
	LpDesktop       LPWSTR
	LpTitle         LPWSTR
	DwX             DWORD
	DwY             DWORD
	DwXSize         DWORD
	DwYSize         DWORD
	DwXCountChars   DWORD
	DwYCountChars   DWORD
	DwFillAttribute DWORD
	DwFlags         DWORD
	WShowWindow     WORD
	CbReserved2     WORD
	LpReserved2     LPBYTE
	HStdInput       HANDLE
	HStdOutput      HANDLE
	HStdError       HANDLE
}
