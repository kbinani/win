package win

//ref DWORD
//ref POINT
//ref RECT
type CANDIDATEFORM struct {
	DwIndex      DWORD
	DwStyle      DWORD
	PtCurrentPos POINT
	RcArea       RECT
}
