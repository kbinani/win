package win

//ref DWORD
//ref HCRYPTPROV_LEGACY
//ref PFN_CRYPT_GET_SIGNER_CERTIFICATE
//ref PVOID

type CRYPT_VERIFY_MESSAGE_PARA struct {
	CbSize                   DWORD
	DwMsgAndCertEncodingType DWORD
	HCryptProv               HCRYPTPROV_LEGACY
	PfnGetSignerCertificate  PFN_CRYPT_GET_SIGNER_CERTIFICATE
	PvGetArg                 PVOID
}
