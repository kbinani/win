package win

//ref DWORD
//ref CRYPT_DATA_BLOB
//ref HCERTSTORE

type CTL_VERIFY_USAGE_PARA struct {
	CbSize         DWORD
	ListIdentifier CRYPT_DATA_BLOB
	CCtlStore      DWORD
	RghCtlStore    *HCERTSTORE
	CSignerStore   DWORD
	RghSignerStore *HCERTSTORE
}
