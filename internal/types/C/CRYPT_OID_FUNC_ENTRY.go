package win

//ref LPCSTR
//ref PVOID

type CRYPT_OID_FUNC_ENTRY struct {
	PszOID     LPCSTR
	PvFuncAddr PVOID
}
