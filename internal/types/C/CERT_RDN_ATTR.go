package win

//ref LPSTR
//ref DWORD
//ref CERT_RDN_VALUE_BLOB

type CERT_RDN_ATTR struct {
	PszObjId    LPSTR
	DwValueType DWORD
	Value       CERT_RDN_VALUE_BLOB
}
