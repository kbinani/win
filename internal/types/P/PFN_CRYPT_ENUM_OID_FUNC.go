package win

type PFN_CRYPT_ENUM_OID_FUNC func(dwEncodingType DWORD, pszFuncName LPCSTR, pszOID LPCSTR, cValue DWORD, rgdwValueType /*const*/ *DWORD, rgpwszValueName /*const*/ *LPCWSTR, rgpbValueData /*const*/ **BYTE, rgcbValueData /*const*/ *DWORD, pvArg PVOID) BOOL
