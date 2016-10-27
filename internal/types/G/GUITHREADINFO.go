package win

//ref HWND
//ref RECT
type GUITHREADINFO struct {
	CbSize        uint32
	Flags         uint32
	HwndActive    HWND
	HwndFocus     HWND
	HwndCapture   HWND
	HwndMenuOwner HWND
	HwndMoveSize  HWND
	HwndCaret     HWND
	RcCaret       RECT
}
