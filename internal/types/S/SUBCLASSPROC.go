package win

//ref HWND
//ref UINT
//ref WPARAM
//ref LPARAM
//ref UINT_PTR
//ref DWORD_PTR
//ref LRESULT
type SUBCLASSPROC func(hWnd HWND, uMsg UINT, wParam WPARAM, lParam LPARAM, uIdSubclass UINT_PTR, dwRefData DWORD_PTR) LRESULT
