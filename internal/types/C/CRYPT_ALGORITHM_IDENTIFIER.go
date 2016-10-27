package win

//ref LPSTR
//ref CRYPT_OBJID_BLOB
type CRYPT_ALGORITHM_IDENTIFIER struct {
	PszObjId   LPSTR
	Parameters CRYPT_OBJID_BLOB
}
