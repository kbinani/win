package win

//ref DWORD
//ref UINT
//ref HKEY
//ref LPCWSTR
//ref MRUCMPPROC
type MRUINFO struct {
	CbSize      DWORD
	UMax        UINT
	FFlags      UINT
	HKey        HKEY
	LpszSubKey  LPCWSTR
	LpfnCompare uintptr // MRUCMPPROC
}
