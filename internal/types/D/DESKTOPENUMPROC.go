package win

//ref BOOL
//ref LPWSTR
//ref LPARAM
type DESKTOPENUMPROC func(lpszDesktop LPWSTR, lParam LPARAM) BOOL
