package win

//ref LRESULT
//ref WPARAM
//ref LPARAM
type HOOKPROC func(code int32, wParam WPARAM, lParam LPARAM) LRESULT
