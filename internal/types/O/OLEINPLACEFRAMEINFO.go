package win

//ref UINT
//ref BOOL
//ref HWND
//ref HACCEL

type OLEINPLACEFRAMEINFO struct {
	Cb            UINT
	FMDIApp       BOOL
	HwndFrame     HWND
	Haccel        HACCEL
	CAccelEntries UINT
}
