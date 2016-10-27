package win

//ref HDC
//ref BOOL
//ref LPARAM
type GRAYSTRINGPROC func(hdc HDC, lParam LPARAM, cchData int) BOOL
