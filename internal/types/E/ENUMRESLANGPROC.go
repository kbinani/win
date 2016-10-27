package win

//ref HMODULE
//ref LPCWSTR
//ref WORD
//ref LONG_PTR
//ref BOOL
type ENUMRESLANGPROC func(hModule HMODULE, lpType LPCWSTR, lpName LPCWSTR, wLanguage WORD, lParam LONG_PTR) BOOL
