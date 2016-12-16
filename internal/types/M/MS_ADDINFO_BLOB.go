package win

//ref DWORD
//ref BYTE

type MS_ADDINFO_BLOB struct {
	CbStruct       DWORD
	CbMemObject    DWORD
	PbMemObject    *BYTE
	CbMemSignedMsg DWORD
	PbMemSignedMsg *BYTE
}
