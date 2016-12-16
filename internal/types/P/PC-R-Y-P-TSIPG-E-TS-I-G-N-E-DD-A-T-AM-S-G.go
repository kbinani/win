package win

//ref SIP_SUBJECTINFO
//ref DWORD
//ref BYTE
//ref BOOL

type PCryptSIPGetSignedDataMsg func(pSubjectInfo *SIP_SUBJECTINFO, pdwEncodingType *DWORD, dwIndex DWORD, pcbSignedDataMsg *DWORD, pbSignedDataMsg *BYTE) BOOL
