package win

//ref HWND
type TIMERPROC func(hwnd HWND, uMsg uint32, idEvent uintptr, dwTime uint32)
