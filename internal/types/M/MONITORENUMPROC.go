package win

//ref HMONITOR
//ref HDC
//ref RECT
//ref BOOL
type MONITORENUMPROC func(hMonitor HMONITOR, hdcMonitor HDC, lprcMonitor *RECT, dwData uintptr) BOOL
