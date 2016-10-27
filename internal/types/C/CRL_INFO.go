package win

//ref CRYPT_ALGORITHM_IDENTIFIER
//ref CERT_NAME_BLOB
//ref FILETIME
//ref PCRL_ENTRY
//ref DWORD
//ref PCERT_EXTENSION
type CRL_INFO struct {
	DwVersion          DWORD
	SignatureAlgorithm CRYPT_ALGORITHM_IDENTIFIER
	Issuer             CERT_NAME_BLOB
	ThisUpdate         FILETIME
	NextUpdate         FILETIME
	CCRLEntry          DWORD
	RgCRLEntry         PCRL_ENTRY
	CExtension         DWORD
	RgExtension        PCERT_EXTENSION
}
