package win

//ref DWORD
//ref HCRYPTPROV_LEGACY

type CRYPT_KEY_VERIFY_MESSAGE_PARA struct {
	CbSize            DWORD
	DwMsgEncodingType DWORD
	HCryptProv        HCRYPTPROV_LEGACY
}
