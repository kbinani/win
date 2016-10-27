package win

//ref DWORD
type CANDIDATELIST struct {
	DwSize      DWORD
	DwStyle     DWORD
	DwCount     DWORD
	DwSelection DWORD
	DwPageStart DWORD
	DwPageSize  DWORD
	DwOffset    [1]DWORD
}
