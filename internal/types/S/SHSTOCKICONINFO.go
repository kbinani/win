package win

//ref DWORD
//ref HICON
//ref WCHAR

type SHSTOCKICONINFO struct {
	CbSize         DWORD
	HIcon          HICON
	ISysImageIndex int32
	IIcon          int32
	SzPath         [MAX_PATH]WCHAR
}
