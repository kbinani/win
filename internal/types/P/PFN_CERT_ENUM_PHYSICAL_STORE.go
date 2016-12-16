package win

//ref DWORD
//ref LPCWSTR
//ref PCERT_PHYSICAL_STORE_INFO
//ref BOOL

type PFN_CERT_ENUM_PHYSICAL_STORE func(pvSystemStore /*const*/ uintptr, dwFlags DWORD, pwszStoreName LPCWSTR, pStoreInfo PCERT_PHYSICAL_STORE_INFO, pvReserved uintptr, pvArg uintptr) BOOL
