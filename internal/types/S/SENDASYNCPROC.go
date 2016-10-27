package win

//ref HWND
//ref LRESULT
type SENDASYNCPROC func(hwnd HWND, uMsg uint32, dwData uintptr, lResult LRESULT)
