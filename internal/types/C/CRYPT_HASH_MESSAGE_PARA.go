package win

//ref DWORD
//ref HCRYPTPROV_LEGACY
//ref CRYPT_ALGORITHM_IDENTIFIER
//ref PVOID

type CRYPT_HASH_MESSAGE_PARA struct {
	CbSize            DWORD
	DwMsgEncodingType DWORD
	HCryptProv        HCRYPTPROV_LEGACY
	HashAlgorithm     CRYPT_ALGORITHM_IDENTIFIER
	PvHashAuxInfo     PVOID
}
