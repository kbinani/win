package win

//ref HINSTANCE
//ref HICON
//ref HCURSOR
//ref HBRUSH
type WNDCLASS struct {
	Style         uint32
	LpfnWndProc   uintptr // WNDPROC
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  *uint16 // LPCWSTR
	LpszClassName *uint16 // LPCWSTR
}
