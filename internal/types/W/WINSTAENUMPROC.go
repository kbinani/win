package win

//ref LPWSTR
//ref BOOL
//ref LPARAM
type WINSTAENUMPROC func(lpszWindowStation LPWSTR, lParam LPARAM) BOOL
