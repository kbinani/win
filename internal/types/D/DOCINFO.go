package win

//ref LPCWSTR
//ref DWORD
type DOCINFO struct {
	CbSize       int32
	LpszDocName  LPCWSTR
	LpszOutput   LPCWSTR
	LpszDatatype LPCWSTR
	FwType       DWORD
}
