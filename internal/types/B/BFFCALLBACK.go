package win

//ref HWND
//ref UINT
//ref LPARAM

type BFFCALLBACK func(hwnd HWND, uMsg UINT, lParam LPARAM, lpData LPARAM) int32
