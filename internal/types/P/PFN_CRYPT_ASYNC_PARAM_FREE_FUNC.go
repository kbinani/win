package win

//ref LPSTR
//ref LPVOID

type PFN_CRYPT_ASYNC_PARAM_FREE_FUNC func(pszParamOid LPSTR, pvParam LPVOID)
