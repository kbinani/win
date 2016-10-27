package win

//ref HWND
//ref UINT
//ref WPARAM
//ref LPARAM
//ref UINT_PTR
type LPSETUPHOOKPROC func(unnamed0 HWND, unnamed1 UINT, unnamed2 WPARAM, unnamed3 LPARAM) UINT_PTR
