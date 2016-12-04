package win

//ref DWORD
//ref CRYPT_INTEGER_BLOB
//ref CRYPT_ALGORITHM_IDENTIFIER
//ref CERT_NAME_BLOB
//ref FILETIME
//ref CRYPT_BIT_BLOB
//ref CERT_PUBLIC_KEY_INFO
//ref PCERT_EXTENSION

type CERT_INFO struct {
	DwVersion            DWORD
	SerialNumber         CRYPT_INTEGER_BLOB
	SignatureAlgorithm   CRYPT_ALGORITHM_IDENTIFIER
	Issuer               CERT_NAME_BLOB
	NotBefore            FILETIME
	NotAfter             FILETIME
	Subject              CERT_NAME_BLOB
	SubjectPublicKeyInfo CERT_PUBLIC_KEY_INFO
	IssuerUniqueId       CRYPT_BIT_BLOB
	SubjectUniqueId      CRYPT_BIT_BLOB
	CExtension           DWORD
	RgExtension          PCERT_EXTENSION
}
