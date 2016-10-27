package win

//ref LPCWSTR
//ref DWORD
//ref LPVOID
type REGISTERWORDENUMPROC func(lpszReading LPCWSTR, unnamed1 DWORD, lpszString LPCWSTR, unnamed3 LPVOID) int32
