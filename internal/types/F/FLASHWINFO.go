package win

type FLASHWINFO struct {
	CbSize    uint32 // UINT
	Hwnd      HWND
	DwFlags   uint32
	UCount    uint32 // UINT
	DwTimeout uint32
}
