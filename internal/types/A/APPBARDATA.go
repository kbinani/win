package win

//ref DWORD
//ref HWND
//ref UINT
//ref RECT
//ref LPARAM

type APPBARDATA struct {
	CbSize           DWORD
	HWnd             HWND
	UCallbackMessage UINT
	UEdge            UINT
	Rc               RECT
	LParam           LPARAM
}
