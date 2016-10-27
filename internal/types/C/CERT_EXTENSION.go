package win

//ref CRYPT_OBJID_BLOB
//ref LPSTR
//ref BOOL
type CERT_EXTENSION struct {
	PszObjId  LPSTR
	FCritical BOOL
	Value     CRYPT_OBJID_BLOB
}
