package win

//ref DWORD
//ref SIP_INDIRECT_DATA

type MS_ADDINFO_FLAT struct {
	CbStruct      DWORD
	PIndirectData *SIP_INDIRECT_DATA
}
