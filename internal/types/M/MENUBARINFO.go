package win

//ref RECT
//ref HMENU
//ref HWND
type MENUBARINFO struct {
	CbSize          uint32
	RcBar           RECT
	HMenu           HMENU
	HwndMenu        HWND
	bitfieldedFlags uint32
	// BOOL fBarFocused:1;
	// BOOL fFocused:1;
}

func (i *MENUBARINFO) FBarFocused() bool {
	return (i.bitfieldedFlags & 1) == 1
}
func (i *MENUBARINFO) FFocused() bool {
	return (i.bitfieldedFlags & 2) == 2
}
