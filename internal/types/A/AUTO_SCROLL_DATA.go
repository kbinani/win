package win

//ref DWORD
//ref BOOL
//ref POINT
//ref DWORD

type AUTO_SCROLL_DATA struct {
	INextSample  int32
	DwLastScroll DWORD
	BFull        BOOL
	Pts          [NUM_POINTS]POINT
	DwTimes      [NUM_POINTS]DWORD
}
