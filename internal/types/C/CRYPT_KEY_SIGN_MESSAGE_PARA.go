package win

//ref DWORD
//ref CRYPT_ALGORITHM_IDENTIFIER
//ref PVOID
//ref HCRYPTPROV
//ref NCRYPT_KEY_HANDLE

type CRYPT_KEY_SIGN_MESSAGE_PARA struct {
	CbSize                   DWORD
	DwMsgAndCertEncodingType DWORD
	union1                   uintptr
	DwKeySpec                DWORD
	HashAlgorithm            CRYPT_ALGORITHM_IDENTIFIER
	PvHashAuxInfo            PVOID
	PubKeyAlgorithm          CRYPT_ALGORITHM_IDENTIFIER
}

func (this *CRYPT_KEY_SIGN_MESSAGE_PARA) GetHCryptProv() HCRYPTPROV {
	return *(*HCRYPTPROV)(unsafe.Pointer(&this.union1))
}

func (this *CRYPT_KEY_SIGN_MESSAGE_PARA) SetHCryptProv(v HCRYPTPROV) {
	*(*HCRYPTPROV)(unsafe.Pointer(&this.union1)) = v
}

func (this *CRYPT_KEY_SIGN_MESSAGE_PARA) GetHNCryptKey() NCRYPT_KEY_HANDLE {
	return *(*NCRYPT_KEY_HANDLE)(unsafe.Pointer(&this.union1))
}

func (this *CRYPT_KEY_SIGN_MESSAGE_PARA) SetHNCryptKey(v NCRYPT_KEY_HANDLE) {
	*(*NCRYPT_KEY_HANDLE)(unsafe.Pointer(&this.union1)) = v
}
