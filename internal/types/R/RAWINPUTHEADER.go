package win

//ref HANDLE
type RAWINPUTHEADER struct {
	DwType  uint32
	DwSize  uint32
	HDevice HANDLE
	WParam  uintptr
}
