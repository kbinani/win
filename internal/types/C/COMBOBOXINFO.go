package win

//ref HWND
//ref RECT
type COMBOBOXINFO struct {
	CbSize      uint32
	RcItem      RECT
	RcButton    RECT
	StateButton uint32
	HwndCombo   HWND
	HwndItem    HWND
	HwndList    HWND
}
