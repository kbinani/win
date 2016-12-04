package win

//ref HICON
//ref DWORD
//ref WCHAR

type SHFILEINFO struct {
	HIcon         HICON
	IIcon         int32
	DwAttributes  DWORD
	SzDisplayName [MAX_PATH]WCHAR
	SzTypeName    [80]WCHAR
}
