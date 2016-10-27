package win

//ref HWND
//ref LPWSTR
//ref HANDLE
//ref BOOL
type PROPENUMPROCEX func(hwnd HWND, lpszString LPWSTR, hData HANDLE, dwData uintptr) BOOL
