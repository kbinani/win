package win

//ref HWND
//ref WPARAM
//ref LPARAM
type DLGPROC func(hwndDlg HWND, uMsg uint32, wParam WPARAM, lParam LPARAM) int32
