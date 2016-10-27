package win

//ref RECT
//ref ATOM
type WINDOWINFO struct {
	CbSize          uint32 // DWORD
	RcWindow        RECT
	RcClient        RECT
	DwStyle         uint32 // DWORD
	DwExStyle       uint32 // DWORD
	DwWindowStatus  uint32 // DWORD
	CxWindowBorders uint32 // UINT
	CyWindowBorders uint32 // UINT
	AtomWindowType  ATOM
	WCreatorVersion uint16 // WORD
}
