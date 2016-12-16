package win

//ref DWORD
//ref LPSTR
//ref CRYPT_DATA_BLOB

type CERT_PHYSICAL_STORE_INFO struct {
	CbSize               DWORD
	PszOpenStoreProvider LPSTR
	DwOpenEncodingType   DWORD
	DwOpenFlags          DWORD
	OpenParameters       CRYPT_DATA_BLOB
	DwFlags              DWORD
	DwPriority           DWORD
}
