package win

//ref HWND
//ref LPCWSTR
//ref HANDLE
//ref BOOL
type PROPENUMPROC func(hWnd HWND, lpszString LPCWSTR, hData HANDLE) BOOL
