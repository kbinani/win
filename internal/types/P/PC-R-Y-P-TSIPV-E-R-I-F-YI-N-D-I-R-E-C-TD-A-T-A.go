package win

//ref SIP_SUBJECTINFO
//ref SIP_INDIRECT_DATA
//ref BOOL

type PCryptSIPVerifyIndirectData func(pSubjectInfo *SIP_SUBJECTINFO, pIndirectData *SIP_INDIRECT_DATA) BOOL
