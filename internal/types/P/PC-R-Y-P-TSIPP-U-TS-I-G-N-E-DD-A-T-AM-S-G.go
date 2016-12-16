package win

//ref SIP_SUBJECTINFO
//ref DWORD
//ref BYTE
//ref BOOL

type PCryptSIPPutSignedDataMsg func(pSubjectInfo *SIP_SUBJECTINFO, dwEncodingType DWORD, pdwIndex *DWORD, cbSignedDataMsg DWORD, pbSignedDataMsg *BYTE) BOOL
