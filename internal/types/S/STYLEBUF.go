package win

//ref DWORD
//ref WCHAR
type STYLEBUF struct {
	DwStyle       DWORD
	SzDescription [STYLE_DESCRIPTION_SIZE]WCHAR
}
