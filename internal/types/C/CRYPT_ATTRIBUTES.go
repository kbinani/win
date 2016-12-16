package win

//ref DWORD
//ref PCRYPT_ATTRIBUTE

type CRYPT_ATTRIBUTES struct {
	CAttr  DWORD
	RgAttr PCRYPT_ATTRIBUTE
}
