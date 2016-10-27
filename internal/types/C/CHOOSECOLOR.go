package win

//ref CCHOOKPROC
//ref DWORD
//ref HWND
//ref COLORREF
//ref LPARAM
//ref LPCWSTR
type CHOOSECOLOR struct {
	LStructSize    DWORD
	HwndOwner      HWND
	HInstance      HWND
	RgbResult      COLORREF
	LpCustColors   *COLORREF
	Flags          DWORD
	LCustData      LPARAM
	LpfnHook       uintptr // LPCCHOOKPROC
	LpTemplateName LPCWSTR
}
