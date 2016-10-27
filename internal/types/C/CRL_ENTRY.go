package win

//ref CRYPT_INTEGER_BLOB
//ref FILETIME
//ref DWORD
//ref PCERT_EXTENSION
type CRL_ENTRY struct {
	SerialNumber   CRYPT_INTEGER_BLOB
	RevocationDate FILETIME
	CExtension     DWORD
	RgExtension    PCERT_EXTENSION
}
