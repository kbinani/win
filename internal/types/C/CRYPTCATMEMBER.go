package win

//ref DWORD
//ref LPWSTR
//ref GUID
//ref SIP_INDIRECT_DATA
//ref HANDLE
//ref CRYPT_ATTR_BLOB

type CRYPTCATMEMBER struct {
	CbStruct             DWORD
	PwszReferenceTag     LPWSTR
	PwszFileName         LPWSTR
	GSubjectType         GUID
	FdwMemberFlags       DWORD
	PIndirectData        *SIP_INDIRECT_DATA
	DwCertVersion        DWORD
	DwReserved           DWORD
	HReserved            HANDLE
	SEncodedIndirectData CRYPT_ATTR_BLOB
	SEncodedMemberInfo   CRYPT_ATTR_BLOB
}
