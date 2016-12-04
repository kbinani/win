package win

//ref DWORD
//ref CTL_USAGE
//ref CRYPT_DATA_BLOB
//ref CRYPT_INTEGER_BLOB
//ref FILETIME
//ref CRYPT_ALGORITHM_IDENTIFIER
//ref PCTL_ENTRY
//ref PCERT_EXTENSION

type CTL_INFO struct {
	DwVersion        DWORD
	SubjectUsage     CTL_USAGE
	ListIdentifier   CRYPT_DATA_BLOB
	SequenceNumber   CRYPT_INTEGER_BLOB
	ThisUpdate       FILETIME
	NextUpdate       FILETIME
	SubjectAlgorithm CRYPT_ALGORITHM_IDENTIFIER
	CCTLEntry        DWORD
	RgCTLEntry       PCTL_ENTRY
	CExtension       DWORD
	RgExtension      PCERT_EXTENSION
}
