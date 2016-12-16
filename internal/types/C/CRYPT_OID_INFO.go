package win

//ref DWORD
//ref LPCSTR
//ref LPCWSTR
//ref CRYPT_DATA_BLOB
//ref ALG_ID

type CRYPT_OID_INFO struct {
	CbSize    DWORD
	PszOID    LPCSTR
	PwszName  LPCWSTR
	DwGroupId DWORD
	union1    DWORD
	ExtraInfo CRYPT_DATA_BLOB
}

func (this *CRYPT_OID_INFO) GetDwValue() DWORD {
	return *(*DWORD)(unsafe.Pointer(&this.union1))
}

func (this *CRYPT_OID_INFO) SetDwValue(v DWORD) {
	*(*DWORD)(unsafe.Pointer(&this.union1)) = v
}

func (this *CRYPT_OID_INFO) GetAlgid() ALG_ID {
	return *(*ALG_ID)(unsafe.Pointer(&this.union1))
}

func (this *CRYPT_OID_INFO) SetAlgid(v ALG_ID) {
	*(*ALG_ID)(unsafe.Pointer(&this.union1)) = v
}

func (this *CRYPT_OID_INFO) GetDwLength() DWORD {
	return *(*DWORD)(unsafe.Pointer(&this.union1))
}

func (this *CRYPT_OID_INFO) SetDwLength(v DWORD) {
	*(*DWORD)(unsafe.Pointer(&this.union1)) = v
}
