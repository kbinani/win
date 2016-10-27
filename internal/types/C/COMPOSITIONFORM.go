package win

//ref DWORD
//ref POINT
//ref RECT
type COMPOSITIONFORM struct {
	DwStyle      DWORD
	PtCurrentPos POINT
	RcArea       RECT
}
