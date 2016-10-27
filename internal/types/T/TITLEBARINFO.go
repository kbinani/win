package win

//ref RECT
type TITLEBARINFO struct {
	CbSize     uint32
	RcTitleBar RECT
	Rgstate    [CCHILDREN_TITLEBAR + 1]uint32
}
