package win

//ref SIP_SUBJECTINFO
//ref DWORD
//ref SIP_INDIRECT_DATA
//ref BOOL

type PCryptSIPCreateIndirectData func(pSubjectInfo *SIP_SUBJECTINFO, pcbIndirectData *DWORD, pIndirectData *SIP_INDIRECT_DATA) BOOL
