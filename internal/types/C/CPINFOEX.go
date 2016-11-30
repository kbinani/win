package win

//ref UINT
//ref BYTE
//ref WCHAR
type CPINFOEX struct {
	MaxCharSize        UINT
	DefaultChar        [MAX_DEFAULTCHAR]BYTE
	LeadByte           [MAX_LEADBYTES]BYTE
	UnicodeDefaultChar WCHAR
	CodePage           UINT
	CodePageName       [MAX_PATH]WCHAR
}
