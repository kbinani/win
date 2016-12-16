package win

//ref CRYPT_HASH_BLOB
//ref DWORD
//ref PVOID
//ref BOOL

type PFN_CRYPT_ENUM_KEYID_PROP func(pKeyIdentifier /*const*/ *CRYPT_HASH_BLOB, dwFlags DWORD, pvReserved PVOID, pvArg PVOID, cProp DWORD, rgdwPropId *DWORD, rgpvData *PVOID, rgcbData *DWORD) BOOL
