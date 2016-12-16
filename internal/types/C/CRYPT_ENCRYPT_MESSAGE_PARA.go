package win

//ref DWORD
//ref HCRYPTPROV_LEGACY
//ref CRYPT_ALGORITHM_IDENTIFIER
//ref PVOID

type CRYPT_ENCRYPT_MESSAGE_PARA struct {
	CbSize                     DWORD
	DwMsgEncodingType          DWORD
	HCryptProv                 HCRYPTPROV_LEGACY
	ContentEncryptionAlgorithm CRYPT_ALGORITHM_IDENTIFIER
	PvEncryptionAuxInfo        PVOID
	DwFlags                    DWORD
	DwInnerContentType         DWORD
}
