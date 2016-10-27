package win

//ref HCURSOR
//ref POINT
type CURSORINFO struct {
	CbSize      uint32
	Flags       uint32
	HCursor     HCURSOR
	PtScreenPos POINT
}
