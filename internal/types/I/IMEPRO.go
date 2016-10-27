package win

//ref HWND
//ref DATETIME
type IMEPRO struct {
	HWnd          HWND
	InstDate      DATETIME
	WVersion      uint32 // UINT
	SzDescription [50]uint16
	SzName        [80]uint16
	SzOptions     [30]uint16
}
