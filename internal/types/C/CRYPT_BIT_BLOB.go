package win

//ref DWORD
//ref BYTE
//ref DWORD

type CRYPT_BIT_BLOB struct {
	CbData      DWORD
	PbData      *BYTE
	CUnusedBits DWORD
}
