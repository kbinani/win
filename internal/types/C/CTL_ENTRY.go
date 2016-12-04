package win

//ref CRYPT_DATA_BLOB
//ref DWORD
//ref PCRYPT_ATTRIBUTE

type CTL_ENTRY struct {
	SubjectIdentifier CRYPT_DATA_BLOB
	CAttribute        DWORD
	RgAttribute       PCRYPT_ATTRIBUTE
}
