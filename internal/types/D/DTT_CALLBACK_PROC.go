package win

//ref HDC
//ref LPWSTR
//ref LPRECT
//ref UINT
//ref LPARAM
type DTT_CALLBACK_PROC func(hdc HDC, pszText LPWSTR, cchText int32, prc LPRECT, dwFlags UINT, lParam LPARAM) int32
