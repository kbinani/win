package win

//ref HDESK
//ref HWND
//ref LUID
type BSMINFO struct {
	CbSize uint32
	Hdesk  HDESK
	Hwnd   HWND
	Luid   LUID
}
