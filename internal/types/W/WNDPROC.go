package win

//ref HWND
//ref UINT
//ref WPARAM
//ref LPARAM
//ref LRESULT
type WNDPROC func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) LRESULT
