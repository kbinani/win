package win

//ref DWORD
//ref PCERT_INFO
//ref CRYPT_ALGORITHM_IDENTIFIER
//ref PVOID
//ref PCRYPT_ATTRIBUTE
//ref NCRYPT_KEY_HANDLE
//ref HCRYPTPROV

type CMSG_SIGNER_ENCODE_INFO struct {
	CbSize        DWORD
	PCertInfo     PCERT_INFO
	union1        uintptr
	DwKeySpec     DWORD
	HashAlgorithm CRYPT_ALGORITHM_IDENTIFIER
	PvHashAuxInfo PVOID
	CAuthAttr     DWORD
	RgAuthAttr    PCRYPT_ATTRIBUTE
	CUnauthAttr   DWORD
	RgUnauthAttr  PCRYPT_ATTRIBUTE
}

func (this *CMSG_SIGNER_ENCODE_INFO) GetHCryptProv() HCRYPTPROV {
	return *(*HCRYPTPROV)(unsafe.Pointer(&this.union1))
}

func (this *CMSG_SIGNER_ENCODE_INFO) SetHCryptProv(v HCRYPTPROV) {
	*(*HCRYPTPROV)(unsafe.Pointer(&this.union1)) = v
}

func (this *CMSG_SIGNER_ENCODE_INFO) GetHNCryptKey() NCRYPT_KEY_HANDLE {
	return *(*NCRYPT_KEY_HANDLE)(unsafe.Pointer(&this.union1))
}

func (this *CMSG_SIGNER_ENCODE_INFO) SetHNCryptKey(v NCRYPT_KEY_HANDLE) {
	*(*NCRYPT_KEY_HANDLE)(unsafe.Pointer(&this.union1)) = v
}
