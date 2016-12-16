package win

//ref DWORD
//ref PCERT_SYSTEM_STORE_INFO
//ref BOOL

type PFN_CERT_ENUM_SYSTEM_STORE func(pvSystemStore /*const*/ uintptr, dwFlags DWORD, pStoreInfo PCERT_SYSTEM_STORE_INFO, pvReserved uintptr, pvArg uintptr) BOOL
