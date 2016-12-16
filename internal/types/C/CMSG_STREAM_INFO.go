package win

//ref DWORD
//ref PFN_CMSG_STREAM_OUTPUT
//ref PVOID

type CMSG_STREAM_INFO struct {
	CbContent       DWORD
	PfnStreamOutput uintptr // PFN_CMSG_STREAM_OUTPUT
	PvArg           PVOID
}
