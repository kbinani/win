package win

//ref LPCWSTR
//ref DWORD
//ref BOOL

type PFN_CERT_ENUM_SYSTEM_STORE_LOCATION func(pwszStoreLocation LPCWSTR, dwFlags DWORD, pvReserved uintptr, pvArg uintptr) BOOL
