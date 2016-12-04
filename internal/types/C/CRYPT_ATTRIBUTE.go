package win

//ref LPSTR
//ref DWORD
//ref PCRYPT_ATTR_BLOB

type CRYPT_ATTRIBUTE struct {
	PszObjId LPSTR
	CValue   DWORD
	RgValue  PCRYPT_ATTR_BLOB
}
