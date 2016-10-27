package win

//ref RECT
type SCROLLBARINFO struct {
	CbSize        uint32
	RcScrollBar   RECT
	DxyLineButton int32
	XyThumbTop    int32
	XyThumbBottom int32
	Reserved      int32
	Rgstate       [CCHILDREN_SCROLLBAR + 1]uint32
}
