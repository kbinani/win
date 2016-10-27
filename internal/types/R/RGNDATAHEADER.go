package win

//ref DWORD
//ref RECT
type RGNDATAHEADER struct {
	DwSize   DWORD
	IType    DWORD
	NCount   DWORD
	NRgnSize DWORD
	RcBound  RECT
}
