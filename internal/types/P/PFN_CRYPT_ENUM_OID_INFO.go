package win

//ref PCCRYPT_OID_INFO
//ref PVOID
//ref BOOL

type PFN_CRYPT_ENUM_OID_INFO func(pInfo PCCRYPT_OID_INFO, pvArg PVOID) BOOL
