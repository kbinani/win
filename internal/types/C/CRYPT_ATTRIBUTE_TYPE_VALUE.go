package win

//ref LPSTR
//ref CRYPT_OBJID_BLOB

type CRYPT_ATTRIBUTE_TYPE_VALUE struct {
	PszObjId LPSTR
	Value    CRYPT_OBJID_BLOB
}
