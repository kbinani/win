package win

//ref HWND
//ref BOOL
//ref LPARAM
type WNDENUMPROC func(hWnd HWND, lParam LPARAM) BOOL
