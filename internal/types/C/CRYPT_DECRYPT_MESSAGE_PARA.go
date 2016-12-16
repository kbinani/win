package win

//ref DWORD
//ref HCERTSTORE

type CRYPT_DECRYPT_MESSAGE_PARA struct {
	CbSize                   DWORD
	DwMsgAndCertEncodingType DWORD
	CCertStore               DWORD
	RghCertStore             *HCERTSTORE
}
