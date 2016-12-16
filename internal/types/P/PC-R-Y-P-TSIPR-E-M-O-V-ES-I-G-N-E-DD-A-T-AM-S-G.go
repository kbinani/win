package win

//ref SIP_SUBJECTINFO
//ref DWORD
//ref BOOL

type PCryptSIPRemoveSignedDataMsg func(pSubjectInfo *SIP_SUBJECTINFO, dwIndex DWORD) BOOL
