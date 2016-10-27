package win

//ref HWND
//ref UINT
//ref WPARAM
//ref LPARAM
//ref LONG_PTR
//ref HRESULT
type TASKDIALOGCALLBACK func(hwnd HWND, msg UINT, wParam WPARAM, lParam LPARAM, lpRefData LONG_PTR) HRESULT
